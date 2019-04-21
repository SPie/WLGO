package response

type TrafficInfo struct {
    RefTrafficInfoCategoryId int `json:"refTrafficInfoCategoryId"`
    Name string `json:"name"`
    Priority string `json:"priority"`
    Owner string `json:"owner"`
    Title string `json:"title"`
    Description string `json:"description"`
    RelatedLines []string `json:"relatedLines"`
    RelatedStops []string `json:"relatedStops"`
    TrafficTime TrafficTime `json:"time"`
    TrafficInfoAttributes []TrafficInfoAttribute `json:"attributes"`
}
