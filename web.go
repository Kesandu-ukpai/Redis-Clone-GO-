package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

//go:embed index.html
var page []byte

// net/http runs each request in its own goroutine, and a Go map isn't safe
// for concurrent writes, so we lock around every command in this demo.
var mu sync.Mutex

// startWeb serves the demo terminal on http://localhost:8080
// It blocks, so call it last in main().
func startWeb(s *Store) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(page)
	})

	http.HandleFunc("/cmd", func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			Cmd string `json:"cmd"`
		}
		json.NewDecoder(r.Body).Decode(&body)

		mu.Lock()
		reply := run(s, body.Cmd)
		mu.Unlock()

		json.NewEncoder(w).Encode(map[string]string{"reply": reply})
	})

	fmt.Println("demo running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// run turns a typed line like "SET name kessy" into a store call.
// Throwaway: the real server replaces this with the RESP parser + dispatcher.
func run(s *Store, line string) string {
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return ""
	}

	switch strings.ToUpper(parts[0]) {
	case "PING":
		return "PONG"
	case "SET":
		if len(parts) != 3 {
			return "(error) wrong number of arguments for 'set'"
		}
		s.Set(parts[1], parts[2])
		return "OK"
	case "GET":
		if len(parts) != 2 {
			return "(error) wrong number of arguments for 'get'"
		}
		v, err := s.Get(parts[1])
		if err != nil {
			return "(nil)"
		}
		return `"` + v + `"`
	case "DEL":
		if len(parts) != 2 {
			return "(error) wrong number of arguments for 'del'"
		}
		if err := s.Delete(parts[1]); err != nil {
			return "(integer) 0"
		}
		return "(integer) 1"

		case "EXISTS":
    	if len(parts) != 2 {
        	return "(error) wrong number of arguments for 'exists'"
    	}
    	if s.Exist(parts[1]) {
        	return "(integer) 1"
    	}
    return "(integer) 0"
	}

	return "(error) unknown command '" + parts[0] + "'"

	
}
