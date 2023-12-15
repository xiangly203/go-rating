package test

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"strconv"
	"testing"
)

func TestConsumer(t *testing.T) {
	// Create a Pulsar client
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://localhost:6650",
	})
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	// Start a separate goroutine for Prometheus metrics
	// In this case, Prometheus metrics can be accessed via http://localhost:2112/metrics
	go func() {
		prometheusPort := 2113
		log.Printf("Starting Prometheus metrics at http://localhost:%v/metrics\n", prometheusPort)
		http.Handle("/metrics", promhttp.Handler())
		err = http.ListenAndServe(":"+strconv.Itoa(prometheusPort), nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Create a consumer
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "topic-1",
		SubscriptionName: "sub-1",
		Type:             pulsar.Shared,
	})
	if err != nil {
		log.Fatal(err)
	}

	defer consumer.Close()

	ctx := context.Background()

	// Write your business logic here
	// In this case, you build a simple Web server. You can consume messages by requesting http://localhost:8083/consume
	webPort := 8083
	http.HandleFunc("/consume", func(w http.ResponseWriter, r *http.Request) {
		msg, err := consumer.Receive(ctx)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Printf("Received message msgId: %v -- content: '%s'\n", msg.ID(), string(msg.Payload()))
			fmt.Fprintf(w, "Received message msgId: %v -- content: '%s'\n", msg.ID(), string(msg.Payload()))
			consumer.Ack(msg)
		}
	})

	err = http.ListenAndServe(":"+strconv.Itoa(webPort), nil)
	if err != nil {
		log.Fatal(err)
	}
}
