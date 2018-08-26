package concurrency

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//Loop runs a function on every time interval
//usage: go Loop(time.Second(15), doSomething())
func Loop(interval time.Duration, f func()) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	tk := time.NewTicker(interval)
	defer tk.Stop()

	for {
		select {
		case <-tk.C:
			f()
		case <-sigChan:
			log.Println("shutdown received, stopping loop")
			return
		}
	}
}

//LoopSync runs a function on sync with an existing time.Ticker
//usage: go LoopSync(tk, doSomething())
func LoopSync(tk *time.Ticker, f func()) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-tk.C:
			f()
		case <-sigChan:
			log.Println("shutdown received, stopping loop")
			return
		}
	}
}

//RunSyncronously runs a collection of functions synchronously
func RunSyncronously(tk *time.Ticker, fs []func()) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-tk.C:
			for _, f := range fs {
				go f()
			}
		case <-sigChan:
			log.Println("shutdown received, stopping loop")
			return
		}
	}
}
