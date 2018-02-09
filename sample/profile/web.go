package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

func main() {
	// we need a webserver to get the pprof webserver
	go func() {
		log.Println(http.ListenAndServe("localhost:8080", nil))
	}()
	fmt.Println("hello world")
	var wg sync.WaitGroup
	wg.Add(1)
	go leakyFunction(wg)
	wg.Wait()
}

// go tool pprof -top http://localhost:8080/debug/pprof/goroutine
// go tool pprof -top http://localhost:8080/debug/pprof/heap
// go tool pprof -top http://localhost:8080/debug/pprof/threadcreate
// go tool pprof -top http://localhost:8080/debug/pprof/block
// go tool pprof -top http://localhost:8080/debug/pprof/mutex

// go tool pprof -top http://localhost:8080/debug/pprof/profile
// go tool pprof -top http://localhost:8080/debug/pprof/trace?seconds=5
// go tool pprof -top http://localhost:8080/debug/pprof/heap
// go tool pprof -png http://localhost:8080/debug/pprof/heap > out.png
func leakyFunction(wg sync.WaitGroup) {
	defer wg.Done()
	s := make([]string, 3)
	for i := 0; i < 10000000; i++ {
		s = append(s, "magical pandas")
		if (i % 100000) == 0 {
			time.Sleep(500 * time.Millisecond)
		}
	}
}
