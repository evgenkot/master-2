package main

import (
	"io"
	"net/http"
	"sync"
)

type EchoService struct {
	text string
	mu   sync.RWMutex
}

func (e *EchoService) HandleEcho(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	switch r.Method {
	case http.MethodPut:
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		e.mu.Lock()
		e.text = string(body)
		e.mu.Unlock()

		w.WriteHeader(http.StatusOK)

	case http.MethodGet:
		e.mu.RLock()
		defer e.mu.RUnlock()
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(e.text))

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	echoService := &EchoService{}
	http.HandleFunc("/echo", echoService.HandleEcho)
	http.ListenAndServe(":8080", nil)
}
