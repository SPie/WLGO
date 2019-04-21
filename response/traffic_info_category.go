package response

type TrafficInfoCategory struct {
    Id int `json:"id"`
    RefTrafficInfoCategoryGroupId int `json:"refTrafficInfoCategoryGroupId"`
    Name string `json:"name"`
    TrafficInfoNameList []string `json:"trafficInfoNameList"`
    Title string `json:"title"`
}
