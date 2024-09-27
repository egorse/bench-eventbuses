package main

import (
	"sync"

	"github.com/olebedev/emitter"
)

func RunOlebedevEmitter(a, b, c int) {
	topic := "main"
	message := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."
	e := emitter.New(1)

	// receivers
	wgRx := sync.WaitGroup{}
	for x := 0; x < b; x++ {
		wgRx.Add(1)
		ch := e.On(topic)
		go func(ch <-chan emitter.Event) {
			defer wgRx.Done()
			n := 0
			for event := range ch {
				text := event.String(0)
				n++
				if text != message {
					panic(text)
				}
				if n >= a*c {
					return
				}
			}

		}(ch)
	}

	// transmitters
	wg := sync.WaitGroup{}
	for x := 0; x < a; x++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for x := 0; x < c; x++ {
				<-e.Emit(topic, message)
			}
		}()
	}
	wg.Wait()
	wgRx.Wait()
}
