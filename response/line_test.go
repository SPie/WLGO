package response

import (
    "testing"
     "encoding/json"

    "github.com/stretchr/testify/assert"
)

func TestCreateLine(t *testing.T) {
    line := Line{
        Towards: "Towards",
        LineId: 124,
        Departures: Departures{
            Departures: []Departure{Departure{Vehicle: Vehicle{Name: "Name"}}},
        },
        Name: "Name",
        Direction: "Direction",
        DirectionId: 152,
        BarrierFree: true,
        RealtimeSupported: true,
        TrafficJam: true,
        LineType: "Type",
    }
    assert.Equal(t, "Towards", line.Towards)
    assert.Equal(t, 124, line.LineId)
    assert.Equal(t, "Name", line.Departures.Departures[0].Vehicle.Name)
    assert.Equal(t, "Name", line.Name)
    assert.Equal(t, "Direction", line.Direction)
    assert.Equal(t, 152, line.DirectionId)
    assert.True(t, line.BarrierFree)
    assert.True(t, line.RealtimeSupported)
    assert.True(t, line.TrafficJam)
    assert.Equal(t, "Type", line.LineType)
}

func TestParseLineFromJson(t *testing.T) {
    jsonString := `{
        "towards":"Towards",
        "lineId":126,
        "departures":{
            "departure":[
                {
                    "vehicle":{
                        "name":"Name"
                    }
                }
            ]
        },
        "name":"Name",
        "direction":"Direction",
        "richtungsId":525,
        "barrierFree":true,
        "realtimeSupported":true,
        "trafficJam":true,
        "type":"Type"
    }`
    var line Line
    err := json.Unmarshal([]byte(jsonString), &line)
    assert.Empty(t, err)
    assert.Equal(t, "Towards", line.Towards)
    assert.Equal(t, 126, line.LineId)
    assert.Equal(t, "Name", line.Departures.Departures[0].Vehicle.Name)
    assert.Equal(t, "Name", line.Name)
    assert.Equal(t, "Direction", line.Direction)
    assert.Equal(t, 525, line.DirectionId)
    assert.True(t, line.BarrierFree)
    assert.True(t, line.RealtimeSupported)
    assert.True(t, line.TrafficJam)
    assert.Equal(t, "Type", line.LineType)
}
