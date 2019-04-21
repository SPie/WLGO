package response

import (
    "testing"
    "encoding/json"

    "github.com/stretchr/testify/assert"
)

func TestCreateLocationStop(t *testing.T) {
    locationStop := LocationStop{
        LocationStopType: "Type",
        Geometry: Geometry{GeometryType: "GeometryType",},
        Properties: Properties{Name: "Name"},
    }
    assert.Equal(t, "Type", locationStop.LocationStopType)
    assert.Equal(t, "GeometryType", locationStop.Geometry.GeometryType)
    assert.Equal(t, "Name", locationStop.Properties.Name)
}

func TestParseLocationStopFromJson(t *testing.T) {
    jsonString := `{
        "type":"Type",
        "geometry":{
            "type":"GeometryType"
        },
        "properties":{
            "name":"Name"
        }
    }`
    var locationStop LocationStop
    err := json.Unmarshal([]byte(jsonString), &locationStop)
    assert.Empty(t, err)
    assert.Equal(t, "Type", locationStop.LocationStopType)
    assert.Equal(t, "GeometryType", locationStop.Geometry.GeometryType)
    assert.Equal(t, "Name", locationStop.Properties.Name)
}
