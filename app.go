package main

import (
	"./exporter"
	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func main() {
	flag.String("listen.address", ":9552", "Provides the listen address")
	flag.Parse()
	listenAddress := flag.Lookup("listen.address").Value.String()

	exp := exporter.New()

	prometheus.MustRegister(exp)

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/metrics", http.StatusMovedPermanently)
	})

	log.Printf("Starting APK exporter on %q", listenAddress)

	if err := http.ListenAndServe(listenAddress, nil); err != nil {
		log.Fatalf("Cannot start APK exporter: %s", err)
	} else {
		log.Printf("Started APK exporter on %q", listenAddress)
	}
}
