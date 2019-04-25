package response

import (
    "testing"
    "encoding/json"

    "github.com/stretchr/testify/assert"
)

func TestCreateTrafficInfoListData(t *testing.T) {
    trafficInfoListData := TrafficInfoListData{
        TrafficInfoCategoryGroups: []TrafficInfoCategoryGroup{TrafficInfoCategoryGroup{Id: 567}},
        TrafficInfoCategories: []TrafficInfoCategory{TrafficInfoCategory{Id: 828}},
        TrafficInfos: []TrafficInfo{TrafficInfo{Name: "Name"}},
    }
    assert.Equal(t, 567, trafficInfoListData.TrafficInfoCategoryGroups[0].Id)
    assert.Equal(t, 828, trafficInfoListData.TrafficInfoCategories[0].Id)
    assert.Equal(t, "Name", trafficInfoListData.TrafficInfos[0].Name)
}

func TestParseTrafficInfoListDataFromJson(t *testing.T) {
    jsonString := `{
        "trafficInfoCategoryGroups":[
            {
                "id":567
            }        
        ],
        "trafficInfoCategories":[
            {
                "id":828
            }
        ],
        "trafficInfos":[
            {
                "name":"Name"
            }
        ]
    }`
    var trafficInfoListData TrafficInfoListData
    err := json.Unmarshal([]byte(jsonString), &trafficInfoListData)
    assert.Empty(t, err)
    assert.Equal(t, 567, trafficInfoListData.TrafficInfoCategoryGroups[0].Id)
    assert.Equal(t, 828, trafficInfoListData.TrafficInfoCategories[0].Id)
    assert.Equal(t, "Name", trafficInfoListData.TrafficInfos[0].Name)
}
