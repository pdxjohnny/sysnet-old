package service

import "encoding/json"

func DecodeJSON() {
	w.Header().Set("Content-Type", "application/json")
	var result map[string]string
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&result)
	if err != nil {
		panic(err)
	}
}
