package response

import (
    "testing"
    "encoding/json"

    "github.com/stretchr/testify/assert"
)

func TestCreateVehicle(t *testing.T) {
    vehicle := Vehicle{
        Name: "Name",
        Direction: "Direction",
        DirectionId: 142,
        BarrierFree: true,
        RealtimeSupported: true,
        TrafficJam: true,
        VehicleType: "Vehicle",
    }
    assert.Equal(t, "Name", vehicle.Name)
    assert.Equal(t, "Direction", vehicle.Direction)
    assert.Equal(t, 142, vehicle.DirectionId)
    assert.True(t, vehicle.BarrierFree)
    assert.True(t, vehicle.RealtimeSupported)
    assert.True(t, vehicle.TrafficJam)
    assert.Equal(t, "Vehicle", vehicle.VehicleType)
}

func TestParseVehicleFromJson(t *testing.T) {
    jsonString := `{
        "name":"Name",
        "direction":"Direction",
        "richtungsId":145,
        "barrierFree":true,
        "realtimeSupported":true,
        "trafficJam":true,
        "type":"Vehicle"
    }`
    var vehicle Vehicle
    err := json.Unmarshal([]byte(jsonString), &vehicle)
    assert.Empty(t, err)
    assert.Equal(t, "Name", vehicle.Name)
    assert.Equal(t, "Direction", vehicle.Direction)
    assert.Equal(t, 145, vehicle.DirectionId)
    assert.True(t, vehicle.BarrierFree)
    assert.True(t, vehicle.RealtimeSupported)
    assert.True(t, vehicle.TrafficJam)
    assert.Equal(t, "Vehicle", vehicle.VehicleType)
}
