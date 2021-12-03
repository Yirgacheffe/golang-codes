package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
)

const (
	exitCodeErr       = 1
	exitCodeInterrupt = 2
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	defer func() {
		signal.Stop(signals)
		cancel()
	}()

	go func() {
		// wait signal, cancel context
		select {
		case <-signals:
			cancel()
		case <-ctx.Done():
		}
		<-signals
		os.Exit(exitCodeInterrupt)

	}()

	// avoid main function testing easy...
	if err := run(ctx, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitCodeErr)
	}
}

func run(ctx context.Context, args []string) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			fmt.Println("work as default ...")
		}
	}
}
