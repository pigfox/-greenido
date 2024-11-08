package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type CommandRequest struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
}

type CommandResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error,omitempty"`
}

func main() {
	commander := NewCommander()
	server := &http.Server{
		Addr:    ":8080",
		Handler: handleRequests(commander),
	}
	log.Fatal(server.ListenAndServe())
}

func handleRequests(cmdr Commander) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/execute", handleCommand(cmdr))
	return mux
}

func handleCommand(cmdr Commander) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CommandRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		var resp CommandResponse
		switch req.Type {
		case "ping":
			result, err := cmdr.Ping(req.Payload)
			if err != nil {
				resp = CommandResponse{Success: false, Error: err.Error()}
			} else {
				resp = CommandResponse{Success: true, Data: result}
			}
		case "sysinfo":
			result, err := cmdr.GetSystemInfo()
			if err != nil {
				resp = CommandResponse{Success: false, Error: err.Error()}
			} else {
				resp = CommandResponse{Success: true, Data: result}
			}
		default:
			resp = CommandResponse{Success: false, Error: "unknown command type"}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
