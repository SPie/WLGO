package response

import (
    "testing"
    "encoding/json"

    "github.com/stretchr/testify/assert"
)

func TestCreateDepartures(t *testing.T) {
    departures := Departures{Departures: []Departure{Departure{Vehicle: Vehicle{Name: "Name"}}}}
    assert.Equal(t, "Name", departures.Departures[0].Vehicle.Name)
}

func TestParseDeparturesFromJson(t *testing.T) {
    jsonString := `{
        "departure":[
            {
                "vehicle":{
                    "name":"Name"
                }
            }
        ]
    }`
    var departures Departures
    err := json.Unmarshal([]byte(jsonString), &departures)
    assert.Empty(t, err)
    assert.Equal(t, "Name", departures.Departures[0].Vehicle.Name)
}
