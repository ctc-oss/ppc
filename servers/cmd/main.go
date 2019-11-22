package main

import (
	"fmt"
	"github.com/alexandrevicenzi/go-sse"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/jw3/ppc/servers"
	"github.com/xujiajun/gorouter"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func mqttOpts(cfg *servers.ServerConfig) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(cfg.BrokerURI)
	opts.SetClientID(cfg.ClientID)
	opts.SetKeepAlive(2 * time.Second)
	opts.SetPingTimeout(1 * time.Second)
	return opts
}

func req2chan(r *http.Request) string {
	e := strings.TrimLeft(r.URL.Path, "/v1/events/")
	return fmt.Sprintf("%s/%s", "/E", e)
}

func main() {
	cfg := servers.NewServerConfiguration()
	log.Printf("mqtt @ %s", cfg.BrokerURI)

	c := mqtt.NewClient(mqttOpts(cfg))
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// channel for shuttling mqtt events
	events := make(chan [2]string)

	if token := c.Subscribe(cfg.EventPrefix, 0, func(client mqtt.Client, msg mqtt.Message) {
		t := strings.TrimLeft(msg.Topic(), cfg.EventPrefix)
		events <- [2]string{t, string(msg.Payload())}
		println(t)
	}); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	so := sse.Options{ChannelNameFunc: req2chan}
	s := sse.NewServer(&so)
	defer s.Shutdown()

	mux := gorouter.New()
	http.HandleFunc("/v1/health", handleHealth)
	mux.GET("/v1/events/:t", func(w http.ResponseWriter, r *http.Request) {
		s.ServeHTTP(w, r)
	})

	// POST /v1/devices/{DEVICE_ID}/{FUNCTION}
	mux.POST("/v1/devices/:device/:function", func(w http.ResponseWriter, r *http.Request) {
		d := gorouter.GetParam(r, "device")
		f := gorouter.GetParam(r, "function")
		t := fmt.Sprintf("%s%s/%s", cfg.FunctionPrefix, d, f)

		println(fmt.Sprintf("function called %s(?)", t))

		c.Publish(t, 0, false, r.Body)
		w.WriteHeader(http.StatusOK)
	})
	mux.POST("/v1/devices/events", func(w http.ResponseWriter, r *http.Request) {
		t := r.FormValue("name")
		d := r.FormValue("data")

		t = fmt.Sprintf("%s%s", cfg.EventPrefix, t)
		println(fmt.Sprintf("event published to %s", t))

		c.Publish(t, 0, false, d)
		events <- [2]string{t, d}

		w.WriteHeader(http.StatusOK)
	})
	http.Handle("/", mux)

	go func() {
		for {
			e := <-events
			s.SendMessage(e[0], sse.NewMessage("", e[1], e[0]))
		}
	}()

	log.Fatal(http.ListenAndServe(":9000", nil))
}

func handleHealth(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusOK)
}
