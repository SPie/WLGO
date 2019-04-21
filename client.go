package wlgo

import (
    "errors"
    "io/ioutil"
    "net/http"
)

type Client interface {
    GetApiEndpoint() (string)
    GetSenderId() (string)
    Request(action string, parameters map[string]string) ([]byte, error)
}

type WLClient struct {
    apiEndpoint string
    senderId string
}

func NewClient(apiEndpoint string, senderId string) (WLClient, error) {
    if len(apiEndpoint) < 1 {
        return WLClient{}, errors.New("Empty API endpoint")
    }
    if len(senderId) < 1 {
        return WLClient{}, errors.New("Empty sender id")
    }

    return WLClient{apiEndpoint: apiEndpoint, senderId: senderId}, nil
}

func (wlClient WLClient) GetApiEndpoint() (string) {
    return wlClient.apiEndpoint
}

func (wlClient WLClient) GetSenderId() (string) {
    return wlClient.senderId
}

func (wlClient WLClient) Request(action string, parameters map[string]string) ([]byte, error) {
    if len(action) < 1 {
        return nil, errors.New("Empty action")
    }

    response, _ := http.Get(wlClient.buildURL(action, parameters))

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
