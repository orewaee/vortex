package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func MustWriteBytes(writer http.ResponseWriter, data []byte, code int) {
	writer.WriteHeader(code)
	if _, err := writer.Write(data); err != nil {
		panic(err)
	}
}

func MustWriteString(writer http.ResponseWriter, data string, code int) {
	bytes := []byte(data)
	MustWriteBytes(writer, bytes, code)
}

func MustWriteJson(writer http.ResponseWriter, data interface{}, code int) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	writer.Header().Set("Content-Type", "application/json")
	MustWriteBytes(writer, bytes, code)
}

func MustReadJson[T interface{}](request *http.Request) *T {
	bytes, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}

	data := new(T)
	if err := json.Unmarshal(bytes, data); err != nil {
		panic(err)
	}

	return data
}
