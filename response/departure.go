package response

type Departure struct {
    DepartureTime DepartureTime `json:"departureTime"`
    Vehicle Vehicle `json:"vehicle"`
}
