package stream

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
)

// HTTPStreamer keeps track of when clients connect and disconenct
type HTTPStreamer struct {
	IDLength   int
	Connect    func(string, http.ResponseWriter, *http.Request, http.Flusher) error
	Disconnect func(string)
}

// NewHTTPStreamer creates a new HTTPStreamer and sets defaults
func NewHTTPStreamer() *HTTPStreamer {
	return &HTTPStreamer{
		IDLength:   12,
		Connect:    nil,
		Disconnect: nil,
	}
}

// HTTPStream takes a conenction and tells us when it disconnects
func (httpstreamer *HTTPStreamer) HTTPStream(w http.ResponseWriter, r *http.Request) {
	// Make sure we can stream to the client
	f, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Cannot stream", http.StatusInternalServerError)
		return
	}
	// Make sure we can check if the client has closed
	closed, ok := w.(http.CloseNotifier)
	if !ok {
		http.Error(w, "Cannot stream", http.StatusInternalServerError)
		return
	}
	// Set the stream headers
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	// Create an id and say that the client connected
	connectionID := randomID(httpstreamer.IDLength)
	if httpstreamer.Connect != nil {
		err := httpstreamer.Connect(connectionID, w, r, f)
		if err != nil {
			return
		}
	}
	for {
		fmt.Println("Waiting for close")
		select {
		case <-closed.CloseNotify():
			fmt.Println("Client closed connection")
			if httpstreamer.Disconnect != nil {
				httpstreamer.Disconnect(connectionID)
			}
			return
		}
	}
}

func randomID(length int) string {
	id := ""
	rb := make([]byte, length)
	_, err := rand.Read(rb)
	if err != nil {
		return id
	}
	return base64.URLEncoding.EncodeToString(rb)
}
