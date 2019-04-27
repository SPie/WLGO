package response

type PoiCategory struct {
    Name string `json:"name"`
    Title string `json:"title"`
    Id int `json:"id"`
    RefPoiCategoryGroupId int `json:"refPoiCategoryGroupId"`
}
