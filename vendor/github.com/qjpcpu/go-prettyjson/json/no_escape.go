package json

import (
	"bufio"
	"bytes"
	"encoding/json"
)

func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)
	encoder := json.NewEncoder(writer)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(v); err != nil {
		return nil, err
	}
	writer.Flush()
	data := buf.Bytes()
	// Encode(v) will terminate each value with a newline.
	// e.WriteByte('\n')
	// What a fool!
	return data[:len(data)-1], nil
}

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
