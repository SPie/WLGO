package response

type MonitorResponseData struct {
    Monitors []Monitor `json:"monitors"`
    TrafficInfoCategoryGroups []TrafficInfoCategoryGroup `json:"trafficInfoCategoryGroups"`
    TrafficInfoCategories []TrafficInfoCategory `json:"trafficInfoCategories"`
    TrafficInfo TrafficInfo `json:"trafficInfo"`
}
