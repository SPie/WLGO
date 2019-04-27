package response

import (
    "strings"
    "time"
)

type Poi struct {
    RefPoiCategoryId int `json:"refPoiCategoryId"`
    Name string `json:"name"`
    Start int64 `json:"start"`
    End int64 `json:"end"`
    Title string `json:"title"`
    Subtitle string `json:"subtitle"`
    Description string `json:"description"`
    RelatedLines string `json:"relatedLines"`
    RelatedStops string `json:"relatedStops"`
    PoiAttributes PoiAttributes `json:"attributes"`
}

func (poi Poi) GetStartAsTime() (time.Time) {
    return time.Unix(poi.Start, 0)
}

func (poi Poi) GetEndAsTime() (time.Time) {
    return time.Unix(poi.End, 0)
}

func (poi Poi) GetRelatedLines() ([]string) {
    return strings.Split(poi.RelatedLines, ",")
}

func (poi Poi) GetRelatedStops() ([]string) {
    return strings.Split(poi.RelatedStops, ",")
}
