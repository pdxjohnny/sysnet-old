package stream

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPStreamer(t *testing.T) {
	httpstreamer := NewHTTPStreamer()
	ts := httptest.NewServer(http.HandlerFunc(httpstreamer.HTTPStream))
	defer ts.Close()

	httpstreamer.Connect = func(id string, w http.ResponseWriter, r *http.Request, f http.Flusher) error {
		fmt.Println("Client connected", id)
		fmt.Fprintf(w, "Some data\n")
		f.Flush()
		fmt.Println("Sent data")
		return nil
	}

	fmt.Println(ts.URL)
	res, err := http.Get(ts.URL)
	if err != nil {
		panic(err)
	}

	done := make(chan bool)
	go OnReadString(res.Body, func(text string) {
		fmt.Println("Output:", text)
		done <- true
	})
	<-done
	res.Body.Close()
	close(done)
}
