package response

import (
    "testing"
    "encoding/json"

    "github.com/stretchr/testify/assert"
)

func TestCreatePoiCategoryGroup(t *testing.T) {
    poiCategoryGroup := PoiCategoryGroup{
        Name: "Name",
        Id: 155,
        Title: "Title",
    }
    assert.Equal(t, "Name", poiCategoryGroup.Name)
    assert.Equal(t, 155, poiCategoryGroup.Id)
    assert.Equal(t, "Title", poiCategoryGroup.Title)
}

func TestParsePoiCategoryGroupFromJson(t *testing.T) {
    jsonString := `{
        "name":"Name",
        "id":982,
        "title":"Title"
    }`
    var poiCategoryGroup PoiCategoryGroup
    err := json.Unmarshal([]byte(jsonString), &poiCategoryGroup)
    assert.Empty(t, err)
    assert.Equal(t, "Name", poiCategoryGroup.Name)
    assert.Equal(t, 982, poiCategoryGroup.Id)
    assert.Equal(t, "Title", poiCategoryGroup.Title)
}
