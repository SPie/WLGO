package response

import (
    "testing"
    "encoding/json"

    "github.com/stretchr/testify/assert"
)

func TestCreateDeparture(t *testing.T) {
    departure := Departure{
        DepartureTime: DepartureTime{Countdown: 141},
        Vehicle: Vehicle{Direction: "Direction"},
    }
    assert.Equal(t, 141, departure.DepartureTime.Countdown)
    assert.Equal(t, "Direction", departure.Vehicle.Direction)
}

func TestParseDepartureFromJson(t *testing.T) {
    jsonString := `{
        "departureTime": {
            "countdown":164
        },
        "vehicle": {
            "direction":"Direction"
        }
    }`
    var departure Departure
    err := json.Unmarshal([]byte(jsonString), &departure)
    assert.Empty(t, err)
    assert.Equal(t, 164, departure.DepartureTime.Countdown)
    assert.Equal(t, "Direction", departure.Vehicle.Direction)
}
