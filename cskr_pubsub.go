package main

import (
	"sync"

	"github.com/cskr/pubsub/v2"
)

func RunCskrPubsub(a, b, c int) {
	topic := "main"
	message := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."
	ps := pubsub.New[string, string](1)

	// receivers
	wgRx := sync.WaitGroup{}
	for x := 0; x < b; x++ {
		wgRx.Add(1)
		sub := ps.Sub(topic)
		go func(sub chan string) {
			defer wgRx.Done()
			n := 0
			for text := range sub {
				n++
				if text != message {
					panic(text)
				}
				if n >= a*c {
					return
				}
			}
		}(sub)
	}

	// transmitters
	wg := sync.WaitGroup{}
	for x := 0; x < a; x++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := 0; n < c; n++ {
				ps.Pub(message, topic)
			}
		}()
	}
	wg.Wait()
	wgRx.Wait()
}
