package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"TaskList/routes"
)

func main() {
	r := routes.Init()
	port := os.Getenv("PORT")

	// Interrupt handler.
	var errChan = make(chan error, 1)

	// Start gin server
	go func() {
		fmt.Println("run port:", port)
		errChan <- r.Run(":" + port)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	<-errChan
}
