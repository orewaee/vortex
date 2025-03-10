package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

// MustWriteBytes writes a slice of bytes to the writer and sets the status code
func MustWriteBytes(writer http.ResponseWriter, data []byte, code int) {
	writer.WriteHeader(code)
	if _, err := writer.Write(data); err != nil {
		panic(err)
	}
}

// MustWriteString writes the string to the writer and sets the status code
func MustWriteString(writer http.ResponseWriter, data string, code int) {
	bytes := []byte(data)
	MustWriteBytes(writer, bytes, code)
}

// MustWriteJson writes data in json format to the writer and sets the status code
func MustWriteJson(writer http.ResponseWriter, data interface{}, code int) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	writer.Header().Set("Content-Type", "application/json")
	MustWriteBytes(writer, bytes, code)
}

// MustReadJson reads data from the request body in json format
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

func IntQueryParam(request *http.Request, key string) int {
	raw := request.URL.Query().Get(key)
	if raw == "" {
		return 0
	}

	param, err := strconv.Atoi(raw)
	if err != nil {
		return 0
	}

	return param
}
