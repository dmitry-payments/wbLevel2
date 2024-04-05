package main

import (
	"fmt"
)

func or(channels ...<-chan struct{}) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		select {
		case <-done:
			return
		case <-merge(channels...):
			return
		}
	}()
	return done
}

func merge(channels ...<-chan struct{}) <-chan struct{} {
	done := make(chan struct{})
	for _, ch := range channels {
		go func(ch <-chan struct{}) {
			defer close(done)
			select {
			case <-done:
				return
			case <-ch:
				return
			}
		}(ch)
	}
	return done
}

func main() {
	done1 := make(chan struct{})
	done2 := make(chan struct{})

	go func() {
		defer close(done1)
		defer close(done2)
	}()

	done := or(done1, done2)

	<-done

	fmt.Println("At least one of the channels is closed.")
}
