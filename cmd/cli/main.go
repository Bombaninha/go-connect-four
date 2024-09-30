package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Bombaninha/go-connect-four/pkg/game"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	go func() {
		<-sigs // Wait for the signal
		fmt.Println("\nFinishing game...")
		os.Exit(0) // Exit without error
	}()

	grid := game.NewGrid(6, 7)
	game := game.NewGame(*grid, 4, 2)
	game.Play()
}
