package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Mukunth-arya/golangapp/helpers"

	"github.com/gorilla/mux"
)

func main() {

	l := log.New(os.Stdout, "Product API", log.LstdFlags)
	router := mux.NewRouter()
	Getd := router.Methods(http.MethodGet).Subrouter()
	Getd.HandleFunc("/Products", helpers.GetMyAllData)

	Postd := router.Methods(http.MethodPost).Subrouter()
	Postd.HandleFunc("/Products/Applyit", helpers.CreateData)
	Postd.Use(helpers.MiddlewareData)

	Putd := router.Methods(http.MethodPut).Subrouter()
	Putd.HandleFunc("/Products/Modit/{id}", helpers.Satisfication)

	Deleted := router.Methods(http.MethodDelete).Subrouter()
	Deleted.HandleFunc("/Products/Delit/{id}", helpers.DeleteAData)

	DeleteAlld := router.Methods(http.MethodDelete).Subrouter()
	DeleteAlld.HandleFunc("/Products/Delallit", helpers.DeleteAllData)

	server := &http.Server{

		Addr:         ":9000",
		Handler:      router,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {

		err := server.ListenAndServe()
		if err != nil {
			log.Panic(err)
		}

	}()

	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt)
	signal.Notify(signals, os.Kill)

	sig := <-signals

	l.Println("Received an graceful shutdown", sig)
	ctx, err := context.WithTimeout(context.Background(), 10*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	server.Shutdown(ctx)
}
