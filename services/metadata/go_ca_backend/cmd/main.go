package main

import (
	"fmt"
	"github.com/EventStore/EventStore-Client-Go/client"
	"github.com/dasch-swiss/dsp-meta-svc/services/metadata/go_ca_backend/api/handler"
	"github.com/dasch-swiss/dsp-meta-svc/services/metadata/go_ca_backend/api/middleware"
	addressRepository "github.com/dasch-swiss/dsp-meta-svc/services/metadata/go_ca_backend/infrastructure/repository/address"
	"github.com/gorilla/context"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/negroni"
	"net/http"
	"strconv"
	"time"

	addressService "github.com/dasch-swiss/dsp-meta-svc/services/metadata/go_ca_backend/service/address"
	"github.com/dasch-swiss/dsp-meta-svc/shared/go/pkg/metric"
	"github.com/gorilla/mux"
	"log"
	"os"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)

	config, err := client.ParseConnectionString("esdb://localhost:2113?tls=false")
	if err != nil {
		log.Fatal("Unexpected configuration error: ", err.Error())
	}

	client, err := client.NewClient(config)
	if err != nil {
		log.Fatal("Unexpected failure setting up test connection: ", err.Error())
	}
	err = client.Connect()
	if err != nil {
		log.Fatal("Unexpected failure connecting: ", err.Error())
	}

	ar := addressRepository.NewAddressRepository(client)
	as := addressService.NewService(ar)

	metricService, err := metric.NewPrometheusService()
	if err != nil {
		log.Fatal(err.Error())
	}

	r := mux.NewRouter()
	n := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.HandlerFunc(middleware.Metrics(metricService)),
		negroni.NewLogger(),
	)

	handler.HandleAddressRoutes(r, *n, as)
	http.Handle("/", r)
	http.Handle("/metrics", promhttp.Handler())
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + strconv.Itoa(8080),
		Handler:  context.ClearHandler(http.DefaultServeMux),
		ErrorLog: logger,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
