package wlgo

import (
    "fmt"
    "bytes"
    "errors"
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "net/url"
    "testing"

    wlgoHttp "wlgo/http"

    "github.com/stretchr/testify/assert"
)

type MockHttpClient struct {
    Response         []byte
    Error            error
    ExpectedUrl      string
    ExpectedUrlParts []string
}

func (mockHttpClient MockHttpClient) Get(url string) (*http.Response, error) {
    if mockHttpClient.Error != nil {
    	return nil, mockHttpClient.Error
    }

    response := http.Response{}
    if mockHttpClient.Response == nil {
    	response.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(url)))
    	return &response, nil
    }

    response.Body = ioutil.NopCloser(bytes.NewBuffer(mockHttpClient.Response))
    return &response, mockHttpClient.Error
}

func newMockServer(response string, validationFunction func(url *url.URL)) (*httptest.Server) {
    return httptest.NewServer(
	http.HandlerFunc(
	    func(responseWriter http.ResponseWriter, request *http.Request) {
		validationFunction(request.URL)
		fmt.Fprint(responseWriter, response)
	    },
	),
    )
}

func sliceContainsSliceValues(expectedValuesSlice []string, actualValuesSlice []string) bool {
    if len(expectedValuesSlice) != len(actualValuesSlice) {
        return false
    }

    for _, expectedValue := range expectedValuesSlice {
	if !sliceContainsValue(actualValuesSlice, expectedValue) {
	    return false
	}
    }

    return true
}

func sliceContainsValue(slice []string, expectedValue string) bool {
    for _, value := range slice {
	if value == expectedValue {
	    return true
	}
    }
    return false
}

func TestCreateNewWLClient(t *testing.T) {
    wlClient, err := NewClient("https://endpoint.api", "SenderId1234", wlgoHttp.NewClient())
    assert.Empty(t, err)
    assert.Equal(t, "https://endpoint.api", wlClient.apiEndpoint)
    assert.Equal(t, "SenderId1234", wlClient.senderId)
    _, ok := wlClient.httpClient.(wlgoHttp.Client)
    assert.True(t, ok)
}

func TestCreateNewWLClientWithEmptyApiEndpoint(t *testing.T) {
    _, err := NewClient("", "SenderId1234", wlgoHttp.NewClient())
    assert.EqualError(t, err, "Empty API endpoint")
}

func TestCreateNewWLClientWithEmptySenderId(t *testing.T) {
    _, err := NewClient("https://endpoint.api", "", wlgoHttp.NewClient())
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
    assert.EqualError(t, err, "Empty action")
}

func TestRequestWithError(t *testing.T) {
    wlClient, _ := NewClient("https://test.x", "SenderId1234", MockHttpClient{Error: errors.New("Error")})

    _, err := wlClient.Request("action", nil)
    assert.EqualError(t, err, "Error")
}

func TestGetMonitorsWithoutStationNumbers(t *testing.T) {
    endpoint, senderId := "https://test.x", "SenderId1234"
    wlClient, _ := NewClient(endpoint, senderId, MockHttpClient{})

    _, err := wlClient.GetMonitors([]string{}, []string{})
    assert.EqualError(t, err, "Empty station numbers")
}

func TestGetMonitorsWithInvalidFaultTypes(t *testing.T) {
    wlClient, _ := NewClient("https://test.x", "SenderId1234", MockHttpClient{})

    _, err := wlClient.GetMonitors([]string{"123"}, []string{"InvalidFault"})
    assert.EqualError(t, err, "Invalid fault types")
}

func TestSuccessfullyGetMonitorsWithOneStationNumberAndNoFaultTypes(t *testing.T) {
    mockServer := newMockServer(
	`{
	    "data":{
	        "trafficInfo":{
		    "name":"Name"
	        }
	    },
	    "message":{
		"value":"Value" 
	    }
        }`,
	func (url *url.URL) {
	    queryParameters := url.Query()
	    if queryParameters.Get("rbl") != "123" || queryParameters.Get("activateTrafficInfo") != "" {
		t.Error("Invalid parameters")
	    }
	},
    )
    defer mockServer.Close()

    wlClient, _ := NewClient(
	mockServer.URL,
	"SenderId1234",
	wlgoHttp.NewClient(),
    )

    monitorResponse, _ := wlClient.GetMonitors([]string{"123"}, []string{})

    assert.Equal(t, "Name", monitorResponse.MonitorResponseData.TrafficInfo.Name)
    assert.Equal(t, "Value", monitorResponse.ResponseMessage.Value)
}

func TestGetMonitorsWithError(t *testing.T) {
    wlClient, _ := NewClient("https://test.x", "SenderId1234", MockHttpClient{Error: errors.New("Error")})
    _, err := wlClient.GetMonitors([]string{"123"}, []string{})
    assert.EqualError(t, err, "Error")
}

func TestGetMonitorsWithFaultType(t *testing.T) {
    mockServer := newMockServer(
	`{}`,
	func (url *url.URL) {
	    queryParameters := url.Query()
	    if queryParameters.Get("rbl") != "123" || queryParameters.Get("activateTrafficInfo") != "stoerungkurz" {
		t.Error("InvalidParameters")
	    }
	},
    )
    wlClient, _ := NewClient(
	mockServer.URL,
	"SenderId1234",
	wlgoHttp.NewClient(),
    )

    _, err := wlClient.GetMonitors([]string{"123"}, []string{"stoerungkurz"})
    assert.Empty(t, err)
}

func TestGetMonitorsWithMultipleStationNumbersAndFaultTypes(t *testing.T) {
    mockServer := newMockServer(
	`{}`,
	func (url *url.URL) {
	    if !sliceContainsSliceValues([]string{"123", "456"}, url.Query()["rbl"]) || 
	    !sliceContainsSliceValues([]string{"stoerungkurz", "stoerunglang"}, url.Query()["activateTrafficInfo"]) {
		t.Error("Invalid parameters")
	    }
	},
    )
    wlClient, _ := NewClient(
	mockServer.URL,
	"SenderId1234",
	wlgoHttp.NewClient(),
    )

    _, err := wlClient.GetMonitors([]string{"123", "456"}, []string{"stoerungkurz", "stoerunglang"})
    assert.Empty(t, err)
}

func TestGetTrafficInfoListWithoutParameters(t *testing.T) {
    mockServer := newMockServer(
	`{
	    "data":{
		"trafficInfos":[
		    {
			"name":"Name" 
		    }
		]
	    }
	}`,
	func (url *url.URL) {},
    )
    wlClient, _ := NewClient(
	mockServer.URL,
	"SenderId1234",
	wlgoHttp.NewClient(),
    )

    trafficInfoListResponse, err := wlClient.GetTrafficInfoList([]string{}, []string{}, []string{})
    assert.Empty(t, err)
    assert.Equal(t, "Name", trafficInfoListResponse.TrafficInfoListData.TrafficInfos[0].Name)
}

func TestGetTrafficInfoListWithSingleParameters(t *testing.T) {
    mockServer := newMockServer(
	`{
	    "data":{
		"trafficInfos":[
		    {
			"name":"Name" 
		    }
		]
	    }
	}`,
	func (url *url.URL) {
	    if !sliceContainsSliceValues([]string{"Line1"}, url.Query()["relatedLine"]) || 
	    !sliceContainsSliceValues([]string{"Stop1"}, url.Query()["relatedStop"]) || 
	    !sliceContainsSliceValues([]string{"Name"}, url.Query()["name"]) {
		t.Error("Invalid parameters")
	    }
	},
    )
    wlClient, _ := NewClient(
	mockServer.URL,
	"SenderId1234",
	wlgoHttp.NewClient(),
    )

    trafficInfoListResponse, err := wlClient.GetTrafficInfoList([]string{"Line1"}, []string{"Stop1"}, []string{"Name"})
    assert.Empty(t, err)
    assert.Equal(t, "Name", trafficInfoListResponse.TrafficInfoListData.TrafficInfos[0].Name)
}

func TestGetTrafficInfoListWithMultipleParameters(t *testing.T) {
    mockServer := newMockServer(
	`{
	    "data":{
		"trafficInfos":[
		    {
			"name":"Name" 
		    }
		]
	    }
	}`,
	func (url *url.URL) {
	    if !sliceContainsSliceValues([]string{"Line1", "Line2"}, url.Query()["relatedLine"]) || 
	    !sliceContainsSliceValues([]string{"Stop1", "Stop2"}, url.Query()["relatedStop"]) || 
	    !sliceContainsSliceValues([]string{"Name1", "Name2"}, url.Query()["name"]) {
		t.Error("Invalid parameters")
	    }
	},
    )
    wlClient, _ := NewClient(
	mockServer.URL,
	"SenderId1234",
	wlgoHttp.NewClient(),
    )

    trafficInfoListResponse, err := wlClient.GetTrafficInfoList([]string{"Line1", "Line2"}, []string{"Stop1", "Stop2"}, []string{"Name1", "Name2"})
    assert.Empty(t, err)
    assert.Equal(t, "Name", trafficInfoListResponse.TrafficInfoListData.TrafficInfos[0].Name)
}

func TestGetTrafficInfoListWithError(t *testing.T) {
    wlClient, _ := NewClient(
	"https://test.x",
	"SenderId1234",
	MockHttpClient{
	    Error: errors.New("Error"),
	},
    )

    _, err := wlClient.GetTrafficInfoList([]string{}, []string{}, []string{})
    assert.EqualError(t, err, "Error")
}

func TestGetNewsListWithoutParameters(t *testing.T) {
    mockServer := newMockServer(
	`{
	    "data":{
		"pois":[
		    {
			"name":"Name"
		    }
		]
	    }
	}`,
	func (url *url.URL) {},
    )
    wlClient, _ := NewClient(
	mockServer.URL,
	"SenderId1234",
	wlgoHttp.NewClient(),
    )

    newsListResponse, err := wlClient.GetNewsList([]string{}, []string{}, []string{})
    assert.Empty(t, err)
    assert.Equal(t, "Name", newsListResponse.NewsListData.Pois[0].Name)
}

func TestGetNewsListWithSingleParameters(t *testing.T) {
    mockServer := newMockServer(
	`{}`,
	func (url *url.URL) {
	    if !sliceContainsSliceValues([]string{"Line1"}, url.Query()["relatedLine"]) ||
	    !sliceContainsSliceValues([]string{"Stop1"}, url.Query()["relatedStop"]) ||
	    !sliceContainsSliceValues([]string{"Name1"}, url.Query()["name"]) {
		t.Error("Invalid parameters")
	    }
	},
    )
    wlClient, _ := NewClient(
    	mockServer.URL,
	"SenderId1234",
	wlgoHttp.NewClient(),
    )

    _, err := wlClient.GetNewsList([]string{"Line1"}, []string{"Stop1"}, []string{"Name1"})
    assert.Empty(t, err)
}

func TestGetNewsListWithMultipleParameters(t *testing.T) {
    mockServer := newMockServer(
	`{}`,
	func (url *url.URL) {
	    if !sliceContainsSliceValues([]string{"Line1", "Line2"}, url.Query()["relatedLine"]) ||
	    !sliceContainsSliceValues([]string{"Stop1", "Stop2"}, url.Query()["relatedStop"]) ||
	    !sliceContainsSliceValues([]string{"Name1", "Name2"}, url.Query()["name"]) {
		t.Error("Invalid parameters")
	    }
	},
    )
    wlClient, _ := NewClient(
    	mockServer.URL,
    	"SenderId1234",
    	wlgoHttp.NewClient(),
    )

    _, err := wlClient.GetNewsList([]string{"Line1", "Line2"}, []string{"Stop1", "Stop2"}, []string{"Name1", "Name2"})
    assert.Empty(t, err)
}

func TestGetNewsListWithError(t *testing.T) {
    wlClient, _ := NewClient(
	"https://test.x",
    	"SenderId1234",
	MockHttpClient{Error: errors.New("Error")},
    )

    _, err := wlClient.GetNewsList([]string{}, []string{}, []string{})
    assert.EqualError(t, err, "Error")
}
