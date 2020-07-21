package test

import (
	"bytes"
	"encoding/json"
)

// ParseResponseBody parse json data
func ParseResponseBody(recBody *bytes.Buffer, res interface{}) interface{} {
	return json.Unmarshal(recBody.Bytes(), res)
}
