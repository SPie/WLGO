package response

import (
    "testing"
    "encoding/json"

    "github.com/stretchr/testify/assert"
)

func TestCreateTrafficInfoCategoryGroup(t *testing.T) {
    trafficInfoCategoryGroup := TrafficInfoCategoryGroup{
        Id: 425,
        Name: "Name",
    }
    assert.Equal(t, 425, trafficInfoCategoryGroup.Id)
    assert.Equal(t, "Name", trafficInfoCategoryGroup.Name)
}

func TestParseTrafficInfoCategoryGroupFromJson(t *testing.T) {
    jsonString := `{
        "id":466,
        "name":"Name"
    }`
    var trafficInfoCategoryGroup TrafficInfoCategoryGroup
    err := json.Unmarshal([]byte(jsonString), &trafficInfoCategoryGroup)
    assert.Empty(t, err)
    assert.Equal(t, 466, trafficInfoCategoryGroup.Id)
    assert.Equal(t, "Name", trafficInfoCategoryGroup.Name)
}
