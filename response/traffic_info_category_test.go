package response

import (
    "testing"
    "encoding/json"

    "github.com/stretchr/testify/assert"
)

func TestCreateTrafficInfoCategory(t *testing.T) {
    trafficInfoCategory := TrafficInfoCategory{
        Id: 233,
        RefTrafficInfoCategoryGroupId: 555,
        Name: "Name",
        TrafficInfoNameList: []string{"InfoName1"},
        Title: "Title",
    }
    assert.Equal(t, 233, trafficInfoCategory.Id)
    assert.Equal(t, 555, trafficInfoCategory.RefTrafficInfoCategoryGroupId)
    assert.Equal(t, "Name", trafficInfoCategory.Name)
    assert.Equal(t, "InfoName1", trafficInfoCategory.TrafficInfoNameList[0])
    assert.Equal(t, "Title", trafficInfoCategory.Title)
}

func TestParseTrafficInfoCategoryFromJson(t *testing.T) {
    jsonString := `{
        "id":666,
        "refTrafficInfoCategoryGroupId":777,
        "name":"Name",
        "trafficInfoNameList":[
            "InfoName1"
        ],
        "title":"Title"
    }`
    var trafficInfoCategory TrafficInfoCategory
    err := json.Unmarshal([]byte(jsonString), &trafficInfoCategory)
    assert.Empty(t, err)
    assert.Equal(t, 666, trafficInfoCategory.Id)
    assert.Equal(t, 777, trafficInfoCategory.RefTrafficInfoCategoryGroupId)
    assert.Equal(t, "Name", trafficInfoCategory.Name)
    assert.Equal(t, "InfoName1", trafficInfoCategory.TrafficInfoNameList[0])
    assert.Equal(t, "Title", trafficInfoCategory.Title)
}
