package http

import (
    "testing"
    "net/http"
    "net/http/httptest"

    "github.com/stretchr/testify/assert"
)

func mockServer(status int) (*httptest.Server) {
    return httptest.NewServer(http.HandlerFunc(
        func(responseWriter http.ResponseWriter, request *http.Request) {
            responseWriter.WriteHeader(status)
        },
    ))
}

func TestGet(t *testing.T) {
    server := mockServer(200)
    defer server.Close()

    client := NewClient()
    response, err := client.Get(server.URL)
    assert.Empty(t, err)
    assert.Equal(t, 200, response.StatusCode)
}

func TestGetWitherror(t *testing.T) {
    client := NewClient()
    _, err := client.Get("http://invalid-url.x")
    assert.NotEmpty(t, err)
}
