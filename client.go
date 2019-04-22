package wlgo

import (
    "errors"
    "io/ioutil"

    wlgohttp "wlgo/http"
)

type Client interface {
    GetApiEndpoint() (string)
    GetSenderId() (string)
    GetHttpClient() (wlgohttp.Client)
    Request(action string, parameters map[string]string) ([]byte, error)
}

type WLClient struct {
    apiEndpoint string
    senderId string
    httpClient wlgohttp.Client
}

func NewClient(apiEndpoint string, senderId string, httpClient wlgohttp.Client) (WLClient, error) {
    if len(apiEndpoint) < 1 {
        return WLClient{}, errors.New("Empty API endpoint")
    }
    if len(senderId) < 1 {
        return WLClient{}, errors.New("Empty sender id")
    }

    return WLClient{apiEndpoint: apiEndpoint, senderId: senderId, httpClient: httpClient}, nil
}

func (wlClient WLClient) GetApiEndpoint() (string) {
    return wlClient.apiEndpoint
}

func (wlClient WLClient) GetSenderId() (string) {
    return wlClient.senderId
}

func (wlClient WLClient) GetHttpClient() (wlgohttp.Client) {
    return wlClient.httpClient
}

func (wlClient WLClient) Request(action string, parameters map[string]string) ([]byte, error) {
    if len(action) < 1 {
        return nil, errors.New("Empty action")
    }

    response, err := wlClient.GetHttpClient().Get(wlClient.buildURL(action, parameters))
    if err != nil {
        return nil, err
    }

    defer response.Body.Close()
    return ioutil.ReadAll(response.Body)
}

func (wlClient WLClient) buildURL(action string, parameters map[string]string) (string) {
    url := wlClient.GetApiEndpoint() + "/" + action + "?sender=" + wlClient.GetSenderId()
    for key, value := range parameters {
        url += "&" + key + "=" + value
    }

    return url
}
