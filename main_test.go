package main

import (
	"fmt"
	"testing"
)

func BenchmarkRunEventbuses(b *testing.B) {
	funcs := []struct {
		name string
		fn   func(a, b, c int)
	}{
		{"asaskevich/EventBus", RunAsaskevichEventbus},
		{"cskr/pubsub", RunCskrPubsub},
		{"olebedev/emitter", RunOlebedevEmitter},
	}

	cases := []struct {
		a int
		b int
		c int
	}{
		{1, 1, 1},
		{1, 1, 10},
		{2, 2, 2},
		{1, 10, 10},
		{10, 10, 10},
		{10, 100, 1000},
		{100, 10, 1000},
		{100, 100, 1000},
	}

	for _, c := range cases {
		for _, f := range funcs {
			name := fmt.Sprintf("%s %vx%vx%v", f.name, c.a, c.b, c.c)
			b.Run(name, func(b *testing.B) {
				f.fn(c.a, c.b, c.c)
			})
		}
	}
}
