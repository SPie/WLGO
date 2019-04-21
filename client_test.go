package wlgo

import (
    "fmt"
    "testing"
    "net/http"
    "net/http/httptest"

    "github.com/stretchr/testify/assert"
)

func mockServerForRequestUrlTest() (*httptest.Server) {
    return mockServer(
        func(responseWriter http.ResponseWriter, request *http.Request) {
            fmt.Fprint(responseWriter, request.URL.String())
        },
    )
}

func mockServer(handlerFunction func (responseWriter http.ResponseWriter, request *http.Request)) (*httptest.Server) {
    return httptest.NewServer(http.HandlerFunc(handlerFunction))
}

func TestCreateNewWLClient(t *testing.T) {
    wlClient, err := NewClient("https://endpoint.api", "SenderId1234")
    assert.Empty(t, err)
    assert.Equal(t, "https://endpoint.api", wlClient.GetApiEndpoint())
    assert.Equal(t, "SenderId1234", wlClient.GetSenderId())
}

func TestCreateNewWLClientWithEmptyApiEndpoint(t *testing.T) {
    _, err := NewClient("", "SenderId1234")
    assert.EqualError(t, err, "Empty API endpoint")
}

func TestCreateNewWLClientWithEmptySenderId(t *testing.T) {
    _, err := NewClient("https://endpoint.api", "")
    assert.EqualError(t, err, "Empty sender id")
}

func TestSuccessfulRequestRequestWithParameters(t *testing.T) {
    server := mockServerForRequestUrlTest()
    defer server.Close()

    wlClient, _ := NewClient(server.URL, "SenderId1234")
    responseString := "/action?sender=SenderId1234&key=value"

    response, err := wlClient.Request("action", map[string]string{"key": "value"})
    assert.Empty(t, err)
    assert.Equal(t, []byte(responseString), response)
}

func TestSuccessfulRequestWithoutParameters(t *testing.T) {
    server := mockServerForRequestUrlTest()
    defer server.Close()

    wlClient, _ := NewClient(server.URL, "SenderId4321")
    responseString := "/action?sender=SenderId4321"

    response, err := wlClient.Request("action", nil)
    assert.Empty(t, err)
    assert.Equal(t, []byte(responseString), response)
}

func TestRequestErrorWithEmptyAction(t *testing.T) {
    server := mockServerForRequestUrlTest()
    defer server.Close()

    wlClient, _ := NewClient(server.URL, "SenderId1234")

    _, err := wlClient.Request("", nil)
    assert.EqualError(t, err,  "Empty action")
}
