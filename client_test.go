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

type TestHttpClient struct {
    Error error
}

func (testHttpClient TestHttpClient) Get(url string) (*http.Response, error) {
    if testHttpClient.Error != nil {
        return nil, testHttpClient.Error
    }

    response := http.Response{
        Body: ioutil.NopCloser(bytes.NewBuffer([]byte(url))),
    }
    return &response, testHttpClient.Error
}

func TestCreateNewWLClient(t *testing.T) {
    wlClient, err := NewClient("https://endpoint.api", "SenderId1234", wlgohttp.NewClient())
    assert.Empty(t, err)
    assert.Equal(t, "https://endpoint.api", wlClient.GetApiEndpoint())
    assert.Equal(t, "SenderId1234", wlClient.GetSenderId())
    _, ok := wlClient.GetHttpClient().(wlgohttp.Client)
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

func TestSuccessfulRequestRequestWithParameters(t *testing.T) {
    wlClient, _ := NewClient("https://test.x", "SenderId1234", TestHttpClient{})
    responseString := "https://test.x/action?sender=SenderId1234&key=value"

    response, err := wlClient.Request("action", map[string]string{"key": "value"})
    assert.Empty(t, err)
    assert.Equal(t, []byte(responseString), response)
}

func TestSuccessfulRequestWithoutParameters(t *testing.T) {
    wlClient, _ := NewClient("https://test.x", "SenderId4321", TestHttpClient{})
    responseString := "https://test.x/action?sender=SenderId4321"

    response, err := wlClient.Request("action", nil)
    assert.Empty(t, err)
    assert.Equal(t, []byte(responseString), response)
}

func TestRequestErrorWithEmptyAction(t *testing.T) {
    wlClient, _ := NewClient("https://test.x", "SenderId1234", TestHttpClient{})

    _, err := wlClient.Request("", nil)
    assert.EqualError(t, err,  "Empty action")
}

func TestRequestWithError(t *testing.T) {
    wlClient, _ := NewClient("https://test.x", "SenderId1234", TestHttpClient{Error: errors.New("Error")})

    _, err := wlClient.Request("action", nil)
    assert.EqualError(t, err, "Error")
}
