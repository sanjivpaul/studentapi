package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sanjivpaul/studentapi/internal/config"
	"github.com/sanjivpaul/studentapi/internal/http/handlers/student"
	"github.com/sanjivpaul/studentapi/internal/storage/sqlite"
)

func main(){
	// fmt.Println("welcome to student api")

	// 1. load config
	cfg := config.MustLoad()

	// 2. database setup
	_, err := sqlite.New(cfg)
	if err != nil{
		log.Fatal(err)
	}

	slog.Info("storage initalized", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))

	// 3. setup router

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to student api"))
	})

	router.HandleFunc("POST /api/student/new", student.New())

	// 4. setup server

	server := http.Server{
		Addr: cfg.Addr,
		Handler: router,
	}

	slog.Info("server is running", slog.String("address", cfg.Addr))
	// fmt.Printf("Server is running %s", cfg.HTTPServer.Addr)

	// creating channel for go routines of size 1
	done := make(chan os.Signal, 1) 

	// call done channel when these event are trigger
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)


	// implement gracefull shutdown server
	// wrap start server in go routines
	// note: go routines are like useEffects() in react
	// go routines runs concurrently so we have to implement a logic to stop that
	go func ()  {
		err := server.ListenAndServe()
		if err != nil{
			log.Fatalf("failed to start server: %s", err.Error())
		}
	}()

	// listen for done channel
	// if done channel get some signal then it break the code 
	// code will exit from here and go to next task
	<-done

	// now here we will write logic for shutdown
	slog.Info("shutting down the server.")

	// create a timer
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)

	defer cancel()

	// long form
	err = server.Shutdown(ctx)
	if err != nil{
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}

	// short form
	// err := server.Shutdown(ctx); err != nil {
	// 	slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	// }

	slog.Info("server shutdown successfully")


}