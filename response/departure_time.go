package response

import "time"

type DepartureTime struct {
    TimePlanned string `json:"timePlanned"`
    TimeReal string `json:"timeReal"`
    Countdown int `json:"countdown"`
}

func (departureTime DepartureTime) GetTimePlannedAsTime() (time.Time, error) {
    return time.Parse(time.RFC3339, departureTime.TimePlanned)
}

func (departureTime DepartureTime) GetTimeRealAsTime() (time.Time, error) {
    return time.Parse(time.RFC3339, departureTime.TimeReal)
}
