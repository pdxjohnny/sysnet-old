package api

import (
	"net/http"
	"net/http/httptest"
)

func TestConnectStream() {
	ts := httptest.NewServer(http.HandlerFunc(ConnectStream))
	defer ts.Close()

	res, err := http.Get(ts.URL)
}
