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

	l := log.New(os.Stdout, "SmartDetectReviewsystem", log.LstdFlags)
	router := mux.NewRouter()
	router.HandleFunc("/Getdata", helpers.GetMyAllData).Methods("GET")
	router.HandleFunc("/Insertone", helpers.CreateData).Methods("POST")
	router.HandleFunc("/Updateone/{id}", helpers.Satisfication).Methods("PUT")
	router.HandleFunc("/Deleteone/{id}", helpers.DeleteAData).Methods("DELETE")
	router.HandleFunc("/Deleteall", helpers.DeleteAllData).Methods("DELETE")

	server := &http.Server{

		Addr:         ":9000",
		Handler:      router,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		server.ListenAndServe()

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
