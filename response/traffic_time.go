package response

import "time"

type TrafficTime struct {
    Start string `json:"start"`
    End string `json:"end"`
    Resume string `json:"resume"`
}

func (trafficTime TrafficTime) GetStartAsTime() (time.Time, error) {
    return time.Parse(time.RFC3339, trafficTime.Start)
}

func (trafficTime TrafficTime) GetEndAsTime() (time.Time, error) {
    return time.Parse(time.RFC3339, trafficTime.End)
}

func (trafficTime TrafficTime) GetResumeAsTime() (time.Time, error) {
    return time.Parse(time.RFC3339, trafficTime.Resume)
}
