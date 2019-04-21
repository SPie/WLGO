package response

type Geometry struct {
    GeometryType string `json:"type"`
    Coordinates Coordinates `json:"coordinates"`
}
