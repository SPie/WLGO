package response

type Properties struct {
    Name string `json:"name"`
    Title string `json:"title"`
    Municipality string `json:"municipality"`
    MunicipalityId int `json:"municipalityId"`
    PropertiesType string `json:"type"`
    CoordinateName string `json:"coordName"`
    Gate string `json:"gate"`
    Attributes PropertiesAttribute `json:"attributes"`
}
