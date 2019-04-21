package response

import (
    "testing"
    "encoding/json"
    "time"

    "github.com/stretchr/testify/assert"
)

func TestCreateResponseMessage(t *testing.T) {
    responseMessage := ResponseMessage{
        Value: "Value",
        MessageCode: 1,
        ServerTime: "2019-04-14T02:11:02.000+0200",
    }
    assert.Equal(t, "Value", responseMessage.Value)
    assert.Equal(t, 1, responseMessage.MessageCode)
    assert.Equal(t, "2019-04-14T02:11:02.000+0200", responseMessage.ServerTime)
}

func TestParseResponseMessageFromJson(t *testing.T) {
    jsonString := `{
        "value":"Value",
        "messageCode":2,
        "serverTime":"2019-04-14T02:11:02.000+0200"
    }`
    var responseMessage ResponseMessage
    err := json.Unmarshal([]byte(jsonString), &responseMessage)
    assert.Empty(t, err)
    assert.Equal(t, "Value", responseMessage.Value)
    assert.Equal(t, 2, responseMessage.MessageCode)
    assert.Equal(t, "2019-04-14T02:11:02.000+0200", responseMessage.ServerTime)
}

func TestGetServerTimeAsTime(t *testing.T) {
    responseMessage := ResponseMessage{ServerTime: "2019-04-14T02:11:02.000+0200"}
    serverTime, _ := time.Parse(time.RFC3339, "2019-04-14T02:11:02.000+0200")
    actualServerTime, _ := responseMessage.GetServerTimeAsTime()
    assert.Equal(t, serverTime, actualServerTime)
}
