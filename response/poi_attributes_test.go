package response

import (
    "testing"
    "encoding/json"
    "time"

    "github.com/stretchr/testify/assert"
)

func TestCreatePoiAttributes(t *testing.T) {
    poiAttributes := PoiAttributes{
        Status: "Status",
        Station: "Station",
        Location: "Location",
        RelatedLines: "Line1,Line2",
        RelatedStops: "Stop1,Stop2",
        Towards: "Towards",
        AusVon: 1556317247282,
        AusBis: 1556317247282,
        Rbls: "123,343",
    }
    assert.Equal(t, "Status", poiAttributes.Status)
    assert.Equal(t, "Station", poiAttributes.Station)
    assert.Equal(t, "Location", poiAttributes.Location)
    assert.Equal(t, "Line1,Line2", poiAttributes.RelatedLines)
    assert.Equal(t, "Stop1,Stop2", poiAttributes.RelatedStops)
    assert.Equal(t, "Towards", poiAttributes.Towards)
    assert.Equal(t, int64(1556317247282), poiAttributes.AusVon)
    assert.Equal(t, int64(1556317247282), poiAttributes.AusBis)
    assert.Equal(t, "123,343", poiAttributes.Rbls)
}

func TestParsePoiAttributesFromJson(t *testing.T) {
    jsonString := `{
        "status":"Status",
        "station":"Station",
        "location":"Location",
        "relatedLines":"Line1,Line2",
        "relatedStops":"Stop1,Stop2",
        "towards":"Towards",
        "ausVon":1556317247282,
        "ausBis":1556317247282,
        "rbls":"123,343"
    }`
    var poiAttributes PoiAttributes
    err := json.Unmarshal([]byte(jsonString), &poiAttributes)
    assert.Empty(t, err)
    assert.Equal(t, "Status", poiAttributes.Status)
    assert.Equal(t, "Station", poiAttributes.Station)
    assert.Equal(t, "Location", poiAttributes.Location)
    assert.Equal(t, "Line1,Line2", poiAttributes.RelatedLines)
    assert.Equal(t, "Stop1,Stop2", poiAttributes.RelatedStops)
    assert.Equal(t, "Towards", poiAttributes.Towards)
    assert.Equal(t, int64(1556317247282), poiAttributes.AusVon)
    assert.Equal(t, int64(1556317247282), poiAttributes.AusBis)
    assert.Equal(t, "123,343", poiAttributes.Rbls)
}

func TestPoiAttributesGetRelatedLines(t *testing.T) {
    poiAttributes := PoiAttributes{RelatedLines: "Line1,Line2"}
    assert.Equal(t, []string{"Line1","Line2"}, poiAttributes.GetRelatedLines())
}

func TestPoiAttributesGetRelatedStops(t *testing.T) {
    poiAttributes := PoiAttributes{RelatedStops: "Stop1,Stop2"}
    assert.Equal(t, []string{"Stop1","Stop2"}, poiAttributes.GetRelatedStops())
}

func TestGetAusVonAsTime(t *testing.T) {
    poiAttributes := PoiAttributes{AusVon: 1556317247282}
    assert.Equal(t, time.Unix(1556317247282, 0), poiAttributes.GetAusVonAsTime())
}

func TestGetAusBisAsTime(t *testing.T) {
    poiAttributes := PoiAttributes{AusBis: 1556317247282}
    assert.Equal(t, time.Unix(1556317247282, 0), poiAttributes.GetAusBisAsTime())
}

func TestGetRbls(t *testing.T) {
    poiAttributes := PoiAttributes{Rbls: "123,574"}
    assert.Equal(t, []string{"123","574"}, poiAttributes.GetRbls())
}
