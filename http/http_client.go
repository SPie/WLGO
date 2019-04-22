package http

import (
    "net/http"
)

type Client interface {
    Get(url string) (*http.Response, error)
}

type HttpClient struct {}

func NewClient() (HttpClient) {
    return HttpClient{}
}

func (httpClient HttpClient) Get(url string) (*http.Response, error) {
    return http.Get(url)
}
