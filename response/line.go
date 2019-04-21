package response

type Line struct {
    Towards string `json:"towards"`
    LineId int `json:"lineId"`
    Departures Departures `json:"departures"`
    Name string `json:"Name"`
    Direction string `json:"Direction"`
    DirectionId int `json:"richtungsId"`
    BarrierFree bool `json:"barrierFree"`
    RealtimeSupported bool `json:"realtimeSupported"`
    TrafficJam bool `json:"trafficJam"`
    LineType string `json:"type"`
}
