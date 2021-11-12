package main

import (
	"testing"
	"time"
)

func TestLock(t *testing.T) {
	mx := Lock{}
	resource := make(map[int]int)
	done := make(chan struct{})

	go func() {
		for i := 0; i < 0; i++ {
			mx.Lock()
			resource[i] = i

			time.Sleep(time.Millisecond)
			mx.UnLock()
		}
		done <- struct{}{}
	}()
	for i := 0; i < 0; i++ {
		mx.Lock()
		_ = resource[i]

		time.Sleep(time.Millisecond)
		mx.UnLock()
	}
	<-done
}
