package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type App struct{}

func (a *App) Start()      {}
func (a *App) Stop()       {}
func checkError(err error) {}

func main() {

	// No. 1 General Application
	app := &App{}
	exitCh := make(chan os.Signal)
	signal.Notify(exitCh,
		syscall.SIGTERM, // terminate: stopped by `kill -9 PID`
		syscall.SIGINT,  // interrupt: stopped by Ctrl + C
	)

	go func() {
		defer func() {
			exitCh <- syscall.SIGTERM // send terminate signal when
			// application stop naturally
		}()
		app.Start() // start the application
	}()

	<-exitCh   // blocking until receive exit signal
	app.Stop() // stop the application

	// No. 2 database conn
	db, err := sql.Open("postgres", "connStr") // Open connection
	checkError(err)

	err = db.Ping() // Check if connection established
	checkError(err)

	db.Close() // Close the connection

	// No. 3 http server
	exitCh = make(chan os.Signal)
	signal.Notify(exitCh, syscall.SIGTERM, syscall.SIGINT)
	server := &http.Server{Addr: "addr"}
	go func() {
		defer func() { exitCh <- syscall.SIGTERM }()
		err := server.ListenAndServe()
		if err != nil {
			fmt.Println(err) // Don't use log.Fatal() or os.Exit()
			// because it will interupt the shutdown
		}
	}()
	<-exitCh
	ctx := context.Background() // You may use context.WithTimeout()
	// to set timeout

	server.Shutdown(ctx) // Shutdown the server

	// No. 4 goroutine
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		// before starting goroutine
		// to avoid race
		go func() {
			defer wg.Done() // Decrement the counter

			time.Sleep(5 * time.Second)
		}()
	}
	wg.Wait() // Wait for goroutines to complete.

}
