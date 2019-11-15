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
	"time"
)

func main() {
	cfg := servers.NewServerConfiguration()
	log.Printf("mqtt @ %s", cfg.BrokerURI)

	opts := mqtt.NewClientOptions().AddBroker(cfg.BrokerURI).SetClientID(cfg.ClientID)
	opts.SetKeepAlive(2 * time.Second)
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// channel for shuttling mqtt events
	events := make(chan [2]string)

	if token := c.Subscribe("#", 0, func(client mqtt.Client, msg mqtt.Message) {
		events <- [2]string{msg.Topic(), string(msg.Payload())}
		println(msg.Topic())
	}); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	s := sse.NewServer(nil)
	defer s.Shutdown()
	http.Handle("/v1/events", s)

	mux := gorouter.New()
	http.HandleFunc("/v1/health", handleHealth)

	// POST /v1/devices/{DEVICE_ID}/{FUNCTION}
	mux.POST("/v1/devices/:device/:function", func(w http.ResponseWriter, r *http.Request) {
		d := gorouter.GetParam(r, "device")
		f := gorouter.GetParam(r, "function")
		t := fmt.Sprintf("/F/%s/%s", d, f)
		c.Publish(t, 0, false, r.Body)
		w.WriteHeader(http.StatusOK)
	})
	http.Handle("/v1/devices", mux)

	go func() {
		for {
			e := <-events
			s.SendMessage("/v1/events", sse.NewMessage("", e[1], e[0]))
		}
	}()

	log.Fatal(http.ListenAndServe(":9000", nil))
}

func handleHealth(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusOK)
}
