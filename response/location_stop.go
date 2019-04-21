package response

type LocationStop struct {
    LocationStopType string `json:"type"`
    Geometry Geometry `json:"geometry"`
    Properties Properties `json:"properties"`
}
