package response

type TrafficInfoAttribute struct {
    Status string `json:"status"`
    Station string `json:"station"`
    Location string `json:"location"`
    Reason string `json:"reason"`
    Towards string `json:"towards"`
    RelatedLines []string `json:"relatedLines"`
    RelatedStops []string `json:"relatedStops"`
}
