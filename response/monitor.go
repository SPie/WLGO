package response

type Monitor struct {
    LocationStop LocationStop `json:"locationStop"`
    Lines []Line `json:"lines"`
    RefTrafficInfoNames []string `json:"refTrafficInfoNames"`
}
