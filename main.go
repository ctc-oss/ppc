package main

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/xujiajun/gorouter"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	brokerHost, ok := os.LookupEnv("MQTT_HOST")
	if !ok {
		brokerHost = "localhost"
	}
	brokerUri := fmt.Sprintf("tcp://%s:1883", brokerHost)

	opts := mqtt.NewClientOptions().AddBroker(brokerUri).SetClientID("notpc")
	opts.SetKeepAlive(2 * time.Second)
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	mux := gorouter.New()
	mux.GET("/health", handleHealth)

	// POST /v1/devices/{DEVICE_ID}/{FUNCTION}
	mux.POST("/v1/devices/:device/:function", func(w http.ResponseWriter, r *http.Request) {
		d := gorouter.GetParam(r, "device")
		f := gorouter.GetParam(r, "function")
		t := fmt.Sprintf("/F/%s/%s", d, f)
		c.Publish(t, 0, false, r.Body)
		w.WriteHeader(http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(":9000", mux))
}

func handleHealth(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusOK)
}
