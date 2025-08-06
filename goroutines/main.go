package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var signals []string = []string{"debugLogs"}
var wg sync.WaitGroup
var mut sync.Mutex

func main() {
	start := time.Now()
	endpoints := []string{
		"https://google.com",
		"https://signworks.app/api/health",
		"https://httpbin.org/ip", "https://google.com",
		"https://signworks.app/api/health",
		"https://httpbin.org/ip", "https://google.com",
		"https://signworks.app/api/health",
		"https://httpbin.org/ip", "https://google.com",
		"https://signworks.app/api/health",
		"https://httpbin.org/ip", "https://google.com",
		"https://signworks.app/api/health",
		"https://httpbin.org/ip", "https://google.com",
		"https://signworks.app/api/health",
		"https://httpbin.org/ip", "https://google.com",
		"https://signworks.app/api/health",
		"https://httpbin.org/ip", "https://google.com",
		"https://signworks.app/api/health",
		"https://httpbin.org/ip", "https://google.com",
		"https://signworks.app/api/health",
		"https://httpbin.org/ip", "https://google.com",
		"https://signworks.app/api/health",
		"https://httpbin.org/ip", "https://google.com",
		"https://signworks.app/api/health",
		"https://httpbin.org/ip", "https://google.com",
		"https://signworks.app/api/health",
		"https://httpbin.org/ip",
	}
	for _, ep := range endpoints {
		go getStatusCode(ep)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
	elapsed := time.Since(start)
	fmt.Printf("Execution took %s\n", elapsed)
}

func getStatusCode(ep string) {
	defer wg.Done()
	res, err := http.Get(ep)
	if err != nil {
		fmt.Printf("Error %s in endpoint: %s", err, ep)
		mut.Lock()
		signals = append(signals, fmt.Sprintf("Error: %v in endpoint: %s", err, ep))
		mut.Unlock()
	} else {
		fmt.Printf("Status code is: %v for endpoint %s\n", res.StatusCode, ep)
		mut.Lock()
		signals = append(signals, fmt.Sprintf("Status code is: %v for endpoint %s\n", res.StatusCode, ep))
		mut.Unlock()
	}

}
