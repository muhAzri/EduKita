package utils

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"io"
)

func IOReaderToBytes(r io.Reader) ([]byte, error) {
	b, err := io.ReadAll(r)
	return b, err
}

func TypeToBytes[T any](t T) ([]byte, error) {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	err := enc.Encode(t)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func AnyToType[T any](v any) (T, error) {
	var out T
	b, err := json.Marshal(v)
	if err != nil {
		return out, err
	}
	err = json.Unmarshal(b, &out)
	if err != nil {
		return out, err
	}
	return out, err
}
