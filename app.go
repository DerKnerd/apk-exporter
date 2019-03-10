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
	flag.Parse()

	exp := exporter.New()

	prometheus.MustRegister(exp)
	listenAddress := flag.String("listen.address", ":9552", "Supply the address or port that the exporter should listen on")

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/metrics", http.StatusMovedPermanently)
	})

	log.Printf("Starting APK exporter on %q", *listenAddress)

	if err := http.ListenAndServe(*listenAddress, nil); err != nil {
		log.Fatalf("Cannot start APK exporter: %s", err)
	} else {
		log.Printf("Started APK exporter on %q", listenAddress)
	}
}
