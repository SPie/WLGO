package response

import (
    "testing"
    "encoding/json"

    "github.com/stretchr/testify/assert"
)

func TestCreateProperties(t *testing.T) {
    properties := Properties{
        Name: "Name",
        Title: "Title",
        Municipality: "Municipality",
        MunicipalityId: 123,
        PropertiesType: "Type",
        CoordinateName: "CoordinateName",
        Gate: "Gate",
        Attributes: PropertiesAttribute{RblNumber: 234},
    }
    assert.Equal(t, "Name", properties.Name)
    assert.Equal(t, "Title", properties.Title)
    assert.Equal(t, "Municipality", properties.Municipality)
    assert.Equal(t, 123, properties.MunicipalityId)
    assert.Equal(t, "Type", properties.PropertiesType)
    assert.Equal(t, "CoordinateName", properties.CoordinateName)
    assert.Equal(t, "Gate", properties.Gate)
    assert.Equal(t, 234, properties.Attributes.RblNumber)
}

func TestParsePropertiesFromJson(t *testing.T) {
    jsonString := `{
        "name":"Name",
        "title":"Title",
        "municipality":"Municipality",
        "municipalityId":152,
        "type":"Type",
        "coordName":"CoordinateName",
        "gate":"Gate",
        "attributes": {
            "rbl": 522
        }
    }`
    var properties Properties
    err := json.Unmarshal([]byte(jsonString), &properties)
    assert.Empty(t, err)
    assert.Equal(t, "Name", properties.Name)
    assert.Equal(t, "Title", properties.Title)
    assert.Equal(t, "Municipality", properties.Municipality)
    assert.Equal(t, 152, properties.MunicipalityId)
    assert.Equal(t, "Type", properties.PropertiesType)
    assert.Equal(t, "CoordinateName", properties.CoordinateName)
    assert.Equal(t, "Gate", properties.Gate)
    assert.Equal(t, 522, properties.Attributes.RblNumber)}
