package interrupt

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Cleaner type referes to a function with no inputs that returns an error
type Cleaner func() error

var cleaners []Cleaner
var active = false

// RegisterCleaner is responsible for regisreting a cleaner function. When a function is registered, the Signal watcher is started in a goroutine.
func RegisterCleaner(f Cleaner) {
	cleaners = append(cleaners, f)
	if !active {
		active = true
		go func() {
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
			<-ch
			// Prevent un-terminated ^C character in terminal
			fmt.Println()
			clean()
			os.Exit(1)
		}()
	}
}

// Clean invokes all registered cleanup functions
func clean() {
	fmt.Println("Cleaning")
	for _, f := range cleaners {
		_ = f()
	}
	cleaners = []Cleaner{}
}
