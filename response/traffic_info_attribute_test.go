package response

import (
    "testing"
    "encoding/json"

    "github.com/stretchr/testify/assert"
)

func TestCreateTrafficInfoAttribute(t *testing.T) {
    trafficInfoAttribute := TrafficInfoAttribute{
        Status: "Status",
        Station: "Station",
        Location: "Location",
        Reason: "Reason",
        Towards: "Towards",
        RelatedLines: []string{"Line1"},
        RelatedStops: []string{"Stop1"},
    }
    assert.Equal(t, "Status", trafficInfoAttribute.Status)
    assert.Equal(t, "Station", trafficInfoAttribute.Station)
    assert.Equal(t, "Location", trafficInfoAttribute.Location)
    assert.Equal(t, "Reason", trafficInfoAttribute.Reason)
    assert.Equal(t, "Towards", trafficInfoAttribute.Towards)
    assert.Equal(t, "Line1", trafficInfoAttribute.RelatedLines[0])
    assert.Equal(t, "Stop1", trafficInfoAttribute.RelatedStops[0])
}

func TestParseTrafficInfoAttributeFromJson(t *testing.T) {
    jsonString := `{
        "status":"Status",
        "station":"Station",
        "location":"Location",
        "reason":"Reason",
        "towards":"Towards",
        "relatedLines":[
            "Line1"
        ],
        "relatedStops":[
            "Stop1"
        ]
    }`
    var trafficInfoAttribute TrafficInfoAttribute
    err := json.Unmarshal([]byte(jsonString), &trafficInfoAttribute)
    assert.Empty(t, err)
    assert.Equal(t, "Status", trafficInfoAttribute.Status)
    assert.Equal(t, "Station", trafficInfoAttribute.Station)
    assert.Equal(t, "Location", trafficInfoAttribute.Location)
    assert.Equal(t, "Reason", trafficInfoAttribute.Reason)
    assert.Equal(t, "Towards", trafficInfoAttribute.Towards)
    assert.Equal(t, "Line1", trafficInfoAttribute.RelatedLines[0])
    assert.Equal(t, "Stop1", trafficInfoAttribute.RelatedStops[0])
}
