package wlgo

import (
    "encoding/json"
    "errors"
    "io"

    wlgoHttp "github.com/spie/wlgo/http"
    wlgoResponse "github.com/spie/wlgo/response"
)

const (
    ACTION_MONITOR string = "monitor"
    ACTION_TRAFFIC_INFO_LIST string = "trafficInfoList"
    ACTION_NEWS_LIST string = "newsList"

    FAULT_TYPE_FAULT_SHORT string = "stoerungkurz"
    FAULT_TYPE_FAULT_LONG string = "stoerunglang"
    FAULT_TYPE_ELEVATOR_INFO string = "aufzugsinfo"
)

type Client interface {
    GetMonitor(stationNumbers []string, faultTypes []string) (wlgoResponse.MonitorResponse, error)
    GetTrafficInfoList(relatedLines []string, relatedStops []string, names []string) (wlgoResponse.TrafficInfoListResponse, error)
    GetNewsList(relatedLines []string, relatedStops []string, names []string) (wlgoResponse.NewsListResponse, error)
    Request(action string, parameters map[string]string) (io.ReadCloser, error)
}

type WLClient struct {
    apiEndpoint string
    senderId string
    httpClient wlgoHttp.Client
}

func NewClient(apiEndpoint string, senderId string, httpClient wlgoHttp.Client) (WLClient, error) {
    if len(apiEndpoint) < 1 {
        return WLClient{}, errors.New("Empty API endpoint")
    }
    if len(senderId) < 1 {
        return WLClient{}, errors.New("Empty sender id")
    }

    return WLClient{apiEndpoint: apiEndpoint, senderId: senderId, httpClient: httpClient}, nil
}

func (wlClient WLClient) GetMonitors(stationNumbers []string, faultTypes []string) (wlgoResponse.MonitorResponse, error) {
    if isEmpty(stationNumbers) {
        return wlgoResponse.MonitorResponse{}, errors.New("Empty station numbers")
    }
    if !isEmpty(faultTypes) && !areFaultTypesValid(faultTypes) {
        return wlgoResponse.MonitorResponse{}, errors.New("Invalid fault types")
    }

    response, err := wlClient.Request(ACTION_MONITOR, map[string][]string{"rbl": stationNumbers, "activateTrafficInfo": faultTypes})
    if err != nil {
        return wlgoResponse.MonitorResponse{}, err
    }

    var monitorResponse wlgoResponse.MonitorResponse
    json.NewDecoder(response).Decode(&monitorResponse)

    return monitorResponse, nil
}

func areFaultTypesValid(faultTypes []string) (bool) {
    for _, faultType := range faultTypes {
        if !(faultType == FAULT_TYPE_FAULT_SHORT || faultType == FAULT_TYPE_FAULT_LONG || faultType == FAULT_TYPE_ELEVATOR_INFO) {
            return false
        }
    }
    return true
}

func (wlClient WLClient) GetTrafficInfoList(relatedLines []string, relatedStops []string, names []string) (wlgoResponse.TrafficInfoListResponse, error) {
    response, err := wlClient.Request(
        ACTION_TRAFFIC_INFO_LIST,
        map[string][]string{"relatedLine": relatedLines, "relatedStop": relatedStops, "name": names},
    )
    if err != nil {
        return wlgoResponse.TrafficInfoListResponse{}, err
    }

    var trafficInfoListResponse wlgoResponse.TrafficInfoListResponse
    json.NewDecoder(response).Decode(&trafficInfoListResponse)

    return trafficInfoListResponse, nil
}

func (wlClient WLClient) GetNewsList(relatedLines []string, relatedStops []string, names []string) (wlgoResponse.NewsListResponse, error) {
    response, err := wlClient.Request(
        ACTION_NEWS_LIST, 
        map[string][]string{"relatedLine": relatedLines, "relatedStop": relatedStops, "name": names},
    )
    if err != nil {
        return wlgoResponse.NewsListResponse{}, err
    }

    var newsListResponse wlgoResponse.NewsListResponse
    json.NewDecoder(response).Decode(&newsListResponse)

    return newsListResponse, nil
}

func (wlClient WLClient) Request(action string, parameters map[string][]string) (io.ReadCloser, error) {
    if len(action) < 1 {
        return nil, errors.New("Empty action")
    }

    response, err := wlClient.httpClient.Get(wlClient.buildURL(action, parameters))
    if err != nil {
        return nil, err
    }

    return response.Body, nil
}

func (wlClient WLClient) buildURL(action string, parameters map[string][]string) (string) {
    url := wlClient.apiEndpoint + "/" + action + "?sender=" + wlClient.senderId
    for key, values := range parameters {
        for _, value := range values {
            url += "&" + key + "=" + value
        }
    }

    return url
}

func isEmpty(values []string) (bool) {
    return len(values) < 1
}
