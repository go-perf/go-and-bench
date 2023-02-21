package goandbench

import (
	"runtime"
	"sync"
	"testing"
)

// adapted from go/src/runtime/proc_test.go
func BenchmarkPingPongHog(b *testing.B) {
	if b.N == 0 {
		return
	}
	defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(1))

	var wg sync.WaitGroup
	wg.Add(3)

	// Create a CPU hog
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				wg.Done()
				return
			default:
			}
		}
	}()

	// Ping-pong b.N times
	ping, pong := make(chan bool), make(chan bool)
	go func() {
		for j := 0; j < b.N; j++ {
			pong <- <-ping
		}
		close(done)
		wg.Done()
	}()

	go func() {
		for i := 0; i < b.N; i++ {
			ping <- <-pong
		}
		wg.Done()
	}()

	b.ResetTimer()

	ping <- true // Start ping-pong
	<-done
	b.StopTimer()
	<-ping // Let last ponger exit

	wg.Wait()
}
