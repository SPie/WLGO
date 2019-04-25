package response

import (
    "testing"
    "encoding/json"

    "github.com/stretchr/testify/assert"
)

func TestCreateTrafficInfoListResponse(t *testing.T) {
    trafficInfoListResponse := TrafficInfoListResponse{
        TrafficInfoListData: TrafficInfoListData{
            TrafficInfos: []TrafficInfo{TrafficInfo{Name: "Name"}},
        },
    }
    assert.Equal(t, "Name", trafficInfoListResponse.TrafficInfoListData.TrafficInfos[0].Name)
}

func TestParseTrafficInfoListResponseFromJson(t *testing.T) {
    jsonString := `{
        "data":{
            "trafficInfos":[
                {
                    "name":"Name"
                }
            ]
        }
    }`
    var trafficInfoListResponse TrafficInfoListResponse
    err := json.Unmarshal([]byte(jsonString), &trafficInfoListResponse)
    assert.Empty(t, err)
    assert.Equal(t, "Name", trafficInfoListResponse.TrafficInfoListData.TrafficInfos[0].Name)
}
