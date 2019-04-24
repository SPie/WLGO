package wlgo

import (
    "bytes"
    "errors"
    "io/ioutil"
    "net/http"
    "testing"

    wlgohttp "wlgo/http"

    "github.com/stretchr/testify/assert"
)

type MockHttpClient struct {
    Response []byte
    Error error
    ExpectedUrl string
}

func (mockHttpClient MockHttpClient) Get(url string) (*http.Response, error) {
    if mockHttpClient.Error != nil {
        return nil, mockHttpClient.Error
    }

    if mockHttpClient.ExpectedUrl != "" && mockHttpClient.ExpectedUrl != url {
        return nil, errors.New("Invalid parameters")
    }

    response := http.Response{}
    if mockHttpClient.Response == nil {
        response.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(url)))
        return &response, nil
    }

    response.Body = ioutil.NopCloser(bytes.NewBuffer(mockHttpClient.Response))
    return &response, mockHttpClient.Error
}

func TestCreateNewWLClient(t *testing.T) {
    wlClient, err := NewClient("https://endpoint.api", "SenderId1234", wlgohttp.NewClient())
    assert.Empty(t, err)
    assert.Equal(t, "https://endpoint.api", wlClient.apiEndpoint)
    assert.Equal(t, "SenderId1234", wlClient.senderId)
    _, ok := wlClient.httpClient.(wlgohttp.Client)
    assert.True(t, ok)
}

func TestCreateNewWLClientWithEmptyApiEndpoint(t *testing.T) {
    _, err := NewClient("", "SenderId1234", wlgohttp.NewClient())
    assert.EqualError(t, err, "Empty API endpoint")
}

func TestCreateNewWLClientWithEmptySenderId(t *testing.T) {
    _, err := NewClient("https://endpoint.api", "", wlgohttp.NewClient())
    assert.EqualError(t, err, "Empty sender id")
}

func TestSuccessfulRequestWithParameters(t *testing.T) {
    wlClient, _ := NewClient("https://test.x", "SenderId1234", MockHttpClient{})
    responseString := "https://test.x/action?sender=SenderId1234&key=value"

    response, err := wlClient.Request("action", map[string][]string{"key": []string{"value"}})
    defer response.Close()

    assert.Empty(t, err)
    body, _ := ioutil.ReadAll(response)
    assert.Equal(t, []byte(responseString), body)
}

func TestSuccessfulRequestWithoutParameters(t *testing.T) {
    wlClient, _ := NewClient("https://test.x", "SenderId4321", MockHttpClient{})
    responseString := "https://test.x/action?sender=SenderId4321"

    response, err := wlClient.Request("action", nil)
    defer response.Close()

    assert.Empty(t, err)
    body, _ := ioutil.ReadAll(response)
    assert.Equal(t, []byte(responseString), body)
}

func TestRequestErrorWithEmptyAction(t *testing.T) {
    wlClient, _ := NewClient("https://test.x", "SenderId1234", MockHttpClient{})

    _, err := wlClient.Request("", nil)
    assert.EqualError(t, err,  "Empty action")
}

func TestRequestWithError(t *testing.T) {
    wlClient, _ := NewClient("https://test.x", "SenderId1234", MockHttpClient{Error: errors.New("Error")})

    _, err := wlClient.Request("action", nil)
    assert.EqualError(t, err, "Error")
}

func TestGetMonitorsWithoutStationNumbers(t *testing.T) {
    wlClient, _ := NewClient("https://test.x", "SenderId1234", MockHttpClient{})

    _, err := wlClient.GetMonitors([]string{}, []string{})
    assert.EqualError(t, err, "Empty station numbers")
}

func TestGetMonitorsWithInvalidFaultTypes(t *testing.T) {
    wlClient, _ := NewClient("https://test.x", "SenderId1234", MockHttpClient{})

    _, err := wlClient.GetMonitors([]string{"123"}, []string{"InvalidFault"})
    assert.EqualError(t, err, "Invalid fault types")
}

func TestSuccessfullyGetMonitorsWithOneStationNumberAndNoFaultTypes(t *testing.T) {
    response := `{
        "data":{
            "trafficInfo":{
                "name":"Name"
            }
        },
        "message":{
            "value":"Value" 
        }
    }`
    wlClient, _ := NewClient(
        "https://test.x",
        "SenderId1234",
        MockHttpClient{Response: []byte(response), ExpectedUrl: "https://test.x/monitor?sender=SenderId1234&rbl=123"},
    )

    monitorResponse, err := wlClient.GetMonitors([]string{"123"}, []string{})

    assert.Empty(t, err)
    assert.Equal(t, "Name", monitorResponse.MonitorResponseData.TrafficInfo.Name)
    assert.Equal(t, "Value", monitorResponse.ResponseMessage.Value)
}

func TestGetMonitorsWithError(t *testing.T) {
    wlClient, _ := NewClient("https://test.x", "SenderId1234", MockHttpClient{Error: errors.New("Error")})

    _, err := wlClient.GetMonitors([]string{"123"}, []string{})
    assert.EqualError(t, err, "Error")
}

func TestGetMonitorsWithFaultType(t *testing.T) {
    response := `{
        "data":{
            "trafficInfo":{
                "name":"Name"
            }
        },
        "message":{
            "value":"Value" 
        }
    }`
    wlClient, _ := NewClient(
        "https://test.x",
        "SenderId1234",
        MockHttpClient{Response: []byte(response), ExpectedUrl: "https://test.x/monitor?sender=SenderId1234&rbl=123&activateTrafficInfo=stoerungkurz"},
    )

    _, err := wlClient.GetMonitors([]string{"123"}, []string{"stoerungkurz"})

    assert.Empty(t, err)
}

func TestGetMonitorsWithMultipleStationNumbersAndFaultTypes(t *testing.T) {
     response := `{
        "data":{
            "trafficInfo":{
                "name":"Name"
            }
        },
        "message":{
            "value":"Value" 
        }
    }`
    wlClient, _ := NewClient(
        "https://test.x",
        "SenderId1234",
        MockHttpClient{
            Response: []byte(response),
            ExpectedUrl: "https://test.x/monitor?sender=SenderId1234&rbl=123&rbl=456&activateTrafficInfo=stoerungkurz&activateTrafficInfo=stoerunglang",
        },
    )

    _, err := wlClient.GetMonitors([]string{"123","456"}, []string{"stoerungkurz","stoerunglang"})

    assert.Empty(t, err)
}
