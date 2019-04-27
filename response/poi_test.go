package response

import (
    "testing"
    "encoding/json"
    "time"

    "github.com/stretchr/testify/assert"
)

func TestCreatePoi(t *testing.T) {
    poi := Poi{
        RefPoiCategoryId: 183,
        Name: "Name",
        Start: 1556317247282,
        End: 1556317247282,
        Title: "Title",
        Subtitle: "Subtitle",
        Description: "Description",
        RelatedLines: "Line1,Line2",
        RelatedStops: "Stop1,Stop2",
        PoiAttributes: PoiAttributes{Status: "Status"},
    }
    assert.Equal(t, 183, poi.RefPoiCategoryId)
    assert.Equal(t, "Name", poi.Name)
    assert.Equal(t, int64(1556317247282), poi.Start)
    assert.Equal(t, int64(1556317247282), poi.End)
    assert.Equal(t, "Title", poi.Title)
    assert.Equal(t, "Subtitle", poi.Subtitle)
    assert.Equal(t, "Description", poi.Description)
    assert.Equal(t, "Line1,Line2", poi.RelatedLines)
    assert.Equal(t, "Stop1,Stop2", poi.RelatedStops)
    assert.Equal(t, "Status", poi.PoiAttributes.Status)
}

func TestParsePoiFromJson(t *testing.T) {
    jsonString := `{
        "refPoiCategoryId":253,
        "name":"Name",
        "start": 1556317247282,
        "end": 1556317247282,
        "title":"Title",
        "subtitle":"Subtitle",
        "description":"Description",
        "relatedLines":"Line1,Line2",
        "relatedStops":"Stop1,Stop2",
        "attributes":{
            "status":"Status"
        }
    }`
    var poi Poi
    err := json.Unmarshal([]byte(jsonString), &poi)
    assert.Empty(t, err)
    assert.Equal(t, 253, poi.RefPoiCategoryId)
    assert.Equal(t, "Name", poi.Name)
    assert.Equal(t, int64(1556317247282), poi.Start)
    assert.Equal(t, int64(1556317247282), poi.End)
    assert.Equal(t, "Title", poi.Title)
    assert.Equal(t, "Subtitle", poi.Subtitle)
    assert.Equal(t, "Description", poi.Description)
    assert.Equal(t, "Line1,Line2", poi.RelatedLines)
    assert.Equal(t, "Stop1,Stop2", poi.RelatedStops)
    assert.Equal(t, "Status", poi.PoiAttributes.Status)
}

func TestPoiGetStartAsTime(t *testing.T) {
    poi := Poi{Start: 1556317247282}
    assert.Equal(t, time.Unix(1556317247282, 0), poi.GetStartAsTime())
}

func TestPoiGetEndAsTime(t *testing.T) {
    poi := Poi{End: 1556317247282}
    assert.Equal(t, time.Unix(1556317247282, 0), poi.GetEndAsTime())
}

func TestGetRelatedLines(t *testing.T) {
    poi := Poi{RelatedLines: "Line1,Line2"}
    assert.Equal(t, []string{"Line1","Line2"}, poi.GetRelatedLines())
}

func TestGetRelatedStops(t *testing.T) {
    poi := Poi{RelatedStops: "Stop1,Stop2"}
    assert.Equal(t, []string{"Stop1","Stop2"}, poi.GetRelatedStops())
}
