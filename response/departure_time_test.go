package response

import (
    "testing"
    "time"
    "encoding/json"

    "github.com/stretchr/testify/assert"
)

func TestCreateDepartureTime(t *testing.T) {
    departureTime := DepartureTime{
        TimePlanned: "2019-04-14T02:11:02.000+0200",
        TimeReal: "2019-04-15T02:11:02.000+0200",
        Countdown: 526,
    }
    assert.Equal(t, "2019-04-14T02:11:02.000+0200", departureTime.TimePlanned)
    assert.Equal(t, "2019-04-15T02:11:02.000+0200", departureTime.TimeReal)
    assert.Equal(t, 526, departureTime.Countdown)
}

func TestParseDepartureTimeFromJson(t *testing.T) {
    jsonString := `{
        "timePlanned": "2019-04-14T02:11:02.000+0200",
        "timeReal": "2019-04-15T02:11:02.000+0200",
        "countdown": 774
    }`
    var departureTime DepartureTime
    err := json.Unmarshal([]byte(jsonString), &departureTime)
    assert.Empty(t, err)
    assert.Equal(t, "2019-04-14T02:11:02.000+0200", departureTime.TimePlanned)
    assert.Equal(t, "2019-04-15T02:11:02.000+0200", departureTime.TimeReal)
    assert.Equal(t, 774, departureTime.Countdown)
}

func TestGetTimePlannedAsTime(t *testing.T) {
    departureTime := DepartureTime{TimePlanned: "2019-04-15T02:11:02.000+0200"}
    timePlanned, _ := time.Parse(time.RFC3339, "2019-04-15T02:11:02.000+0200")
    actualTimePlanned, _ := departureTime.GetTimePlannedAsTime()
    assert.Equal(t, timePlanned, actualTimePlanned)
}

func TestGetTimeRealAsTime(t *testing.T) {
    departureTime := DepartureTime{TimePlanned: "2019-04-14T02:11:02.000+0200"}
    timeReal, _ := time.Parse(time.RFC3339, "2019-04-14T02:11:02.000+0200")
    actualTimeReal, _ := departureTime.GetTimeRealAsTime()
    assert.Equal(t, timeReal, actualTimeReal)
}
