package app

import (
	"github.com/fvukojevic/bookstore_items-api/clients/elasticsearch"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	elasticsearch.Init()

	mapUrls()

	src := &http.Server{
		Handler:      router,
		Addr:         "localhost:8082",
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
	}

	if err := src.ListenAndServe(); err != nil {
		panic(err)
	}
}
