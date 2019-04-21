package response

type Vehicle struct {
    Name string `json:"name"`
    Direction string `json:"direction"`
    DirectionId int `json:"richtungsId"`
    BarrierFree bool `json:"barrierFree"`
    RealtimeSupported bool `json:"realtimeSupported"`
    TrafficJam bool `json:"trafficJam"`
    VehicleType string `json:"type"`
}
