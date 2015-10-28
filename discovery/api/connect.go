package api

import (
	"fmt"
	"log"
	"net/http"
)

// ConnectStream takes a conenction and tells us when it disconnects
func ConnectStream(w http.ResponseWriter, r *http.Request) {
	f, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Cannot stream", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	closed, ok := w.(http.CloseNotifier)
	if !ok {
		http.Error(w, "Cannot stream", http.StatusInternalServerError)
		return
	}

	for {
		select {
		case <-closed.CloseNotify():
			log.Println("Done: closed connection")
			return
		case msg := <-messages:
			fmt.Fprintf(w, "data: %s\n\n", msg)
			f.Flush()
		}
	}
}
