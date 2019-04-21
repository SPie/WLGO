package response

import (
    "testing"
    "encoding/json"

    "github.com/stretchr/testify/assert"
)

func TestCreateMonitorResponse(t *testing.T) {
    monitorResponse := MonitorResponse{
        MonitorResponseData: MonitorResponseData{TrafficInfo: TrafficInfo{Name: "Name"}},
        ResponseMessage: ResponseMessage{Value: "Value"},
    }
    assert.Equal(t, "Name", monitorResponse.MonitorResponseData.TrafficInfo.Name)
    assert.Equal(t, "Value", monitorResponse.ResponseMessage.Value)
}

func TestParseMonitorResponseFromJson(t *testing.T) {
    jsonString := `{
        "data":{
            "trafficInfo":{
                "name":"Name"
            }
        },
        "message":{
            "value":"Value"
        }
    }`
    var monitorResponse MonitorResponse
    err := json.Unmarshal([]byte(jsonString), &monitorResponse)
    assert.Empty(t, err)
    assert.Equal(t, "Name", monitorResponse.MonitorResponseData.TrafficInfo.Name)
    assert.Equal(t, "Value", monitorResponse.ResponseMessage.Value)
}
