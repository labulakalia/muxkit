package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r1 := r.PathPrefix("/path")

	r1.Subrouter().HandleFunc("/test", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Query())
		err := r.ParseForm()
		fmt.Println(err)
		fmt.Println(r.PostForm,r.Form)
		fmt.Fprintln(rw, "test")
	})
	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}
	srv.RegisterOnShutdown(func() {
		log.Println("start clean data")
	})
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Println(err)
		}
	}()
	ctx, done := signal.NotifyContext(context.Background(), os.Interrupt)
	defer done()

	<-ctx.Done()
	srv.Shutdown(context.Background())
	log.Println("success exit")
}
