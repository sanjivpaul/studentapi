package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sanjivpaul/studentapi/internal/config"
)

func main(){
	// fmt.Println("welcome to student api")

	// 1. load config
	cfg := config.MustLoad()
	// 2. database setup
	// 3. setup router

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to student api"))
	})

	// 4. setup server

	server := http.Server{
		Addr: cfg.Addr,
		Handler: router,
	}

	fmt.Printf("Server is running %s", cfg.HTTPServer.Addr)

	err := server.ListenAndServe()

	if err != nil{
		log.Fatalf("failed to start server: %s", err.Error())
	}



}