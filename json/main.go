package main

import (
	"bytes"
	"encoding/json"
)

func main() {
	// You can edit this code!
	// Click here and start typing.

	var blob []map[string]interface{}
	var jsonBlob = []byte(`[
      {"name": "Adam", "age": "10"},
      {"name": "Eve",    "age": "100"}
    ]`)
	json.NewDecoder(bytes.NewReader(jsonBlob)).Decode(&blob)
}
