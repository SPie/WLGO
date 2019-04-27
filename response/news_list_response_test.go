package response

import (
    "testing"
    "encoding/json"

    "github.com/stretchr/testify/assert"
)

func TestCreateNewsListResponse(t *testing.T) {
    newsListResponse := NewsListResponse{
        NewsListData: NewsListData{
            Pois: []Poi{Poi{Name: "Name"}},
        },
    }
    assert.Equal(t, "Name", newsListResponse.NewsListData.Pois[0].Name)
}

func TestParseNewsListResponseFromJson(t *testing.T) {
    jsonString := `{
        "data":{
            "pois":[
                {
                    "name":"Name"
                }
            ]
        }
    }`
    var newsListResponse NewsListResponse
    err := json.Unmarshal([]byte(jsonString), &newsListResponse)
    assert.Empty(t, err)
    assert.Equal(t, "Name", newsListResponse.NewsListData.Pois[0].Name)
}
