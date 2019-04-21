package response

import (
    "testing"
    "encoding/json"

    "github.com/stretchr/testify/assert"
)

func TestCreateTrafficInfo(t *testing.T) {
    trafficInfo := TrafficInfo{
        RefTrafficInfoCategoryId: 333,
        Name: "Name",
        Priority: "Priority",
        Owner: "Owner",
        Title: "Title",
        Description: "Description",
        RelatedLines: []string{"Line1"},
        RelatedStops: []string{"Stop1"},
        TrafficTime: TrafficTime{Start: "2019-04-14T02:11:02.000+0200"},
        TrafficInfoAttributes: []TrafficInfoAttribute{TrafficInfoAttribute{Status: "Status"}},
    }
    assert.Equal(t, 333, trafficInfo.RefTrafficInfoCategoryId)
    assert.Equal(t, "Name", trafficInfo.Name)
    assert.Equal(t, "Priority", trafficInfo.Priority)
    assert.Equal(t, "Owner", trafficInfo.Owner)
    assert.Equal(t, "Title", trafficInfo.Title)
    assert.Equal(t, "Description", trafficInfo.Description)
    assert.Equal(t, "Line1", trafficInfo.RelatedLines[0])
    assert.Equal(t, "Stop1", trafficInfo.RelatedStops[0])
    assert.Equal(t, "2019-04-14T02:11:02.000+0200", trafficInfo.TrafficTime.Start)
    assert.Equal(t, "Status", trafficInfo.TrafficInfoAttributes[0].Status)
}

func TestParseTrafficInfoFromJson(t *testing.T) {
    jsonString := `{
        "refTrafficInfoCategoryId":444,
        "name":"Name",
        "priority":"Priority",
        "owner":"Owner",
        "title":"Title",
        "description":"Description",
        "relatedLines":[
            "Line1"
        ],
        "relatedStops":[
            "Stop1"
        ],
        "time":{
            "start":"2019-04-14T02:11:02.000+0200"
        },
        "attributes":[
            {
                "status":"Status"
            }
        ]
    }`
    var trafficInfo TrafficInfo
    err := json.Unmarshal([]byte(jsonString), &trafficInfo)
    assert.Empty(t, err)
    assert.Equal(t, 444, trafficInfo.RefTrafficInfoCategoryId)
    assert.Equal(t, "Name", trafficInfo.Name)
    assert.Equal(t, "Priority", trafficInfo.Priority)
    assert.Equal(t, "Owner", trafficInfo.Owner)
    assert.Equal(t, "Title", trafficInfo.Title)
    assert.Equal(t, "Description", trafficInfo.Description)
    assert.Equal(t, "Line1", trafficInfo.RelatedLines[0])
    assert.Equal(t, "Stop1", trafficInfo.RelatedStops[0])
    assert.Equal(t, "2019-04-14T02:11:02.000+0200", trafficInfo.TrafficTime.Start)
    assert.Equal(t, "Status", trafficInfo.TrafficInfoAttributes[0].Status)
}
