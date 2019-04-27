package response

import (
    "testing"
    "encoding/json"

    "github.com/stretchr/testify/assert"
)

func TestCreateNewsListData(t *testing.T) {
    newsListData := NewsListData{
        PoiCategoryGroups: []PoiCategoryGroup{
            PoiCategoryGroup{Id: 522},
        },
        PoiCategories: []PoiCategory{
            PoiCategory{Id: 173},
        },
        Pois: []Poi{
            Poi{Name: "Name"},
        },
    }
    assert.Equal(t, 522, newsListData.PoiCategoryGroups[0].Id)
    assert.Equal(t, 173, newsListData.PoiCategories[0].Id)
    assert.Equal(t, "Name", newsListData.Pois[0].Name)
}

func TestParseNewsListDataFromJson(t *testing.T) {
    jsonString := `{
        "poiCategoryGroups":[
            {
                "id":522
            }
        ],
        "poiCategories":[
            {
                "id":173
            }
        ],
        "pois":[
            {
                "name":"Name"
            }
        ]
    }`
    var newsListData NewsListData
    err := json.Unmarshal([]byte(jsonString), &newsListData)
    assert.Empty(t, err)
    assert.Equal(t, 522, newsListData.PoiCategoryGroups[0].Id)
    assert.Equal(t, 173, newsListData.PoiCategories[0].Id)
    assert.Equal(t, "Name", newsListData.Pois[0].Name)
}
