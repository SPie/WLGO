package response

import (
    "testing"
    "encoding/json"
)

func TestCreateGeometry(t *testing.T) {
    geometry := Geometry{GeometryType: "type", Coordinates: [2]float64{10.10, 11.11}}
    if geometry.GeometryType != "type" {
        t.Fatal("GeometryType should be 'type'")
    }
    if geometry.Coordinates.GetLatitude() != 10.10 {
        t.Fatal("Latitude of coordinates should be 10.10")
    }
    if geometry.Coordinates.GetLongitude() != 11.11 {
        t.Fatal("Longitude of coordinates shold be 11.11")
    }
}

func TestParseGeometryFromJson(t *testing.T) {
    jsonString := `{"type":"Test","coordinates":[10.10,11.11]}`
    var geometry Geometry
    err := json.Unmarshal([]byte(jsonString), &geometry)
    if err != nil {
        t.Fatalf("JSON should be parsed. Error: %s", err)
    }
    if geometry.GeometryType != "Test" {
        t.Fatal("GeometryType should be 'Test'")
    }
    if geometry.Coordinates.GetLatitude() != 10.10 {
        t.Fatal("Latitude of coordinates should be 10.10")
    }
    if geometry.Coordinates.GetLongitude() != 11.11 {
        t.Fatal("Longitude of coordinates shold be 11.11")
    }
}
