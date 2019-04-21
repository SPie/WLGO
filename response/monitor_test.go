package response

import (
    "testing"
    "encoding/json"

    "github.com/stretchr/testify/assert"
)

func TestCreateMonitor(t *testing.T) {
    monitor := Monitor{
        LocationStop: LocationStop{LocationStopType: "Type"},
        Lines: []Line{Line{Towards: "Towards"}},
        RefTrafficInfoNames: []string{"Info1"},
    }
    assert.Equal(t, "Type", monitor.LocationStop.LocationStopType)
    assert.Equal(t, "Towards", monitor.Lines[0].Towards)
    assert.Equal(t, "Info1", monitor.RefTrafficInfoNames[0])
}

func TestParseMonitorFromJson(t *testing.T) {
    jsonString := `{
        "locationStop":{
            "type": "Type"
        },
        "lines":[
            {
                "towards":"Towards"
            }
        ],
        "refTrafficInfoNames":[
            "Info1"
        ]
    }`
    var monitor Monitor
    err := json.Unmarshal([]byte(jsonString), &monitor)
    assert.Empty(t, err)
    assert.Equal(t, "Type", monitor.LocationStop.LocationStopType)
    assert.Equal(t, "Towards", monitor.Lines[0].Towards)
    assert.Equal(t, "Info1", monitor.RefTrafficInfoNames[0])
}
