package response

import (
    "strings"
    "time"
)

type PoiAttributes struct {
    Status string `json:"status"`
    Station string `json:"station"`
    Location string `json:"location"`
    RelatedLines string `json:"relatedLines"`
    RelatedStops string `json:"relatedStops"`
    Towards string `json:"towards"`
    AusVon int64 `json:"ausVon"`
    AusBis int64 `json:"ausBis"`
    Rbls string `json:"rbls"`
}

func (poiAttributes PoiAttributes) GetRelatedLines() ([]string) {
    return strings.Split(poiAttributes.RelatedLines, ",")
}

func (poiAttributes PoiAttributes) GetRelatedStops() ([]string) {
    return strings.Split(poiAttributes.RelatedStops, ",")
}

func (poiAttributes PoiAttributes) GetAusVonAsTime() (time.Time) {
    return time.Unix(poiAttributes.AusVon, 0)
}

func (poiAttributes PoiAttributes) GetAusBisAsTime() (time.Time) {
    return time.Unix(poiAttributes.AusBis, 0)
}

func (poiAttributes PoiAttributes) GetRbls() ([]string) {
    return strings.Split(poiAttributes.Rbls, ",")
}
