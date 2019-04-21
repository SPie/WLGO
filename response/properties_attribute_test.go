package response

import (
    "testing"
     "encoding/json"
)

func TestCreatePropertiesAttribute(t *testing.T) {
    propertiesAttribute := PropertiesAttribute{RblNumber: 123}
    if propertiesAttribute.RblNumber != 123 {
        t.Fatal("RblNumber should be 123")
    }
}

func TestParsePropertiesAttributeFromJson(t *testing.T) {
    jsonString := `{"rbl":123}`
    var propertiesAttribute PropertiesAttribute
    err := json.Unmarshal([]byte(jsonString), &propertiesAttribute)
    if err != nil {
        t.Fatal("JSON should be parsed to PropertiesAttribute")
    }
    if propertiesAttribute.RblNumber != 123 {
        t.Fatal("RblNumber should be 123")
    }
}
