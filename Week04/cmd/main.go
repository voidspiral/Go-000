package main

import (
	"errors"
	"os"
	"os/signal"
	"school/internal/di"
	"syscall"

	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	app := di.InitApp()

	// grpc server
	g.Go(func() error {
		fmt.Println("grpc")
		go func() {
			<-ctx.Done()
			fmt.Println("grpc ctx done")
			app.Stop()
		}()
		return app.Start()
	})

	// signal
	g.Go(func() error {
		exitSignals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT} // SIGTERM is POSIX specific
		sig := make(chan os.Signal, len(exitSignals))
		signal.Notify(sig, exitSignals...)
		for {
			fmt.Println("signal listen")
			select {
			case <-ctx.Done():
				fmt.Println("signal ctx done")
				return ctx.Err()
			case s := <-sig:
				return errors.New("Close by singal " + s.String())
			}
		}
	})

	// inject error
	/*
		g.Go(func() error {
			fmt.Println("inject")
			time.Sleep(time.Second)
			fmt.Println("inject finish")
			return errors.New("inject error")
		})
	*/

	err := g.Wait() // first error return
	fmt.Println(err)
}
