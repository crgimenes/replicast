package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

var targetPorts = []string{"8081", "8082", "8083"}

func replicateRequest(port string, data []byte) {
	url := fmt.Sprintf("http://localhost:%s", port)
	_, err := http.Post(url, "application/octet-stream", io.Reader(data))
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	var wg sync.WaitGroup
	for _, port := range targetPorts {
		wg.Add(1)
		go func(p string) {
			defer wg.Done()
			replicateRequest(p, body)
		}(port)
	}
	wg.Wait()
}

func main() {
	http.HandleFunc("/", handler)

	http.ListenAndServe(":8080", nil)
}
