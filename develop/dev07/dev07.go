package main

import (
	"fmt"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan interface{})

	go func () {
		defer close(orDone)

		switch len(channels) {
		case 2:
			select {
			case <- channels[0]:
			case <- channels[1]:
			}
		default:
			select {
			case <- channels[0]:
			case <- channels[1]:
			case <- channels[2]:
			case <- or(append(channels[3:], orDone)...):
			}
		}
	}()

	return orDone
}

func main(){
	sig := func(after time.Duration) <- chan interface{} {
		c := make(chan interface{})
		
		go func() {
			defer close(c)
			time.Sleep(after)
			fmt.Println("closing chan")
		}()
		
		return c
	}

	start := time.Now()

	<-or(
		sig(2*time.Second),
		sig(5*time.Second),
		sig(3*time.Second),
		sig(4*time.Second),
		sig(6*time.Second),
	)

	fmt.Printf("Done after %v\n", time.Since(start))
}