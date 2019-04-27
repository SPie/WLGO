package response

type NewsListData struct {
    PoiCategoryGroups []PoiCategoryGroup `json:"poiCategoryGroups"`
    PoiCategories []PoiCategory `json:"poiCategories"`
    Pois []Poi `json:"pois"`
}
