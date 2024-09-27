package main

import (
	"sync"

	"github.com/asaskevich/EventBus"
)

func RunAsaskevichEventbus(a, b, c int) {
	channel := "main"
	message := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."
	bus := EventBus.New()

	// receivers
	for x := 0; x < b; x++ {
		bus.Subscribe(channel, func(s string) {
			if s != message {
				panic(s)
			}
		})
	}

	// transmitters
	wg := sync.WaitGroup{}
	for x := 0; x < a; x++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for x := 0; x < c; x++ {
				bus.Publish(channel, message)
			}
		}()
	}
	wg.Wait()
}
