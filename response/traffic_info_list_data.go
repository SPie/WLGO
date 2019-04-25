package response

type TrafficInfoListData struct {
    TrafficInfoCategoryGroups []TrafficInfoCategoryGroup `json:"trafficInfoCategoryGroups"`
    TrafficInfoCategories []TrafficInfoCategory `json:"trafficInfoCategories"`
    TrafficInfos []TrafficInfo `json:"trafficInfos"`
}
