package service

import (
	"encoding/json"
	"net/http"
)

// DecodeJSON creates an interface and decodes into it
func DecodeJSON(r *http.Request) (interface{}, error) {
	var result interface{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&result)
	if err != nil {
		return nil, err
	}
	return err, nil
}
