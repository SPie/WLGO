package response

import (
    "testing"
    "encoding/json"

    "github.com/stretchr/testify/assert"
)

func TestCreateMonitorResponseData(t *testing.T) {
    monitorResponseData := MonitorResponseData{
        Monitors: []Monitor{
            Monitor{RefTrafficInfoNames: []string{"Info1"}},
        },
        TrafficInfoCategoryGroups: []TrafficInfoCategoryGroup{TrafficInfoCategoryGroup{Id: 992}},
        TrafficInfoCategories: []TrafficInfoCategory{TrafficInfoCategory{Id: 444}},
        TrafficInfo: TrafficInfo{Name: "Name"},
    }
    assert.Equal(t, "Info1", monitorResponseData.Monitors[0].RefTrafficInfoNames[0])
    assert.Equal(t, 992, monitorResponseData.TrafficInfoCategoryGroups[0].Id)
    assert.Equal(t, 444, monitorResponseData.TrafficInfoCategories[0].Id)
    assert.Equal(t, "Name", monitorResponseData.TrafficInfo.Name)
}

func TestParseMonitorResponseDataFromJson(t *testing.T) {
    jsonString := `{
        "monitors":[
            {
                "locationStop":{
                    "type":"Type"
                }
            }
        ],
        "trafficInfoCategoryGroups":[
            {
                "id":992
            }
        ],
        "trafficInfoCategories":[
            {
                "id":444
            }
        ],
        "trafficInfo":{
            "name":"Name"
        }
    }`
    var monitorResponseData MonitorResponseData
    err := json.Unmarshal([]byte(jsonString), &monitorResponseData)
    assert.Empty(t, err)
    assert.Equal(t, "Type", monitorResponseData.Monitors[0].LocationStop.LocationStopType)
    assert.Equal(t, 992, monitorResponseData.TrafficInfoCategoryGroups[0].Id)
    assert.Equal(t, 444, monitorResponseData.TrafficInfoCategories[0].Id)
    assert.Equal(t, "Name", monitorResponseData.TrafficInfo.Name)
}
