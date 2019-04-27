package response

import (
    "testing"
    "encoding/json"

    "github.com/stretchr/testify/assert"
)

func TestCreatePoiCategory(t *testing.T) {
    poiCategory := PoiCategory{
        Name: "Name",
        Title: "Title",
        Id: 838,
        RefPoiCategoryGroupId: 639,
    }
    assert.Equal(t, "Name", poiCategory.Name)
    assert.Equal(t, "Title", poiCategory.Title)
    assert.Equal(t, 838, poiCategory.Id)
    assert.Equal(t, 639, poiCategory.RefPoiCategoryGroupId)
}

func TestParsePoiCategoryFromJson(t *testing.T) {
    jsonString := `{
        "name":"Name",
        "title":"Title",
        "id":382,
        "refPoiCategoryGroupId":282
    }`
    var poiCategory PoiCategory
    err := json.Unmarshal([]byte(jsonString), &poiCategory)
    assert.Empty(t, err)
    assert.Equal(t, "Name", poiCategory.Name)
    assert.Equal(t, "Title", poiCategory.Title)
    assert.Equal(t, 382, poiCategory.Id)
    assert.Equal(t, 282, poiCategory.RefPoiCategoryGroupId)
}
