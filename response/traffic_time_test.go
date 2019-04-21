package response

import (
    "testing"
    "time"
    "encoding/json"

    "github.com/stretchr/testify/assert"
)

func TestCreateTrafficTime(t *testing.T) {
    trafficTime := TrafficTime{
        Start: "2019-04-14T02:11:02.000+0200",
        End: "2019-04-15T02:11:02.000+0200",
        Resume: "2019-04-16T02:11:02.000+0200",
    }
    assert.Equal(t, "2019-04-14T02:11:02.000+0200", trafficTime.Start)
    assert.Equal(t, "2019-04-15T02:11:02.000+0200", trafficTime.End)
    assert.Equal(t, "2019-04-16T02:11:02.000+0200", trafficTime.Resume)
}

func TestParseTrafficTimeFromJson(t *testing.T) {
    jsonString := `{
        "start":"2019-04-14T02:11:02.000+0200",
        "end":"2019-04-15T02:11:02.000+0200",
        "resume":"2019-04-16T02:11:02.000+0200"
    }`
    var trafficTime TrafficTime
    err := json.Unmarshal([]byte(jsonString), &trafficTime)
    assert.Empty(t, err)
    assert.Equal(t, "2019-04-14T02:11:02.000+0200", trafficTime.Start)
    assert.Equal(t, "2019-04-15T02:11:02.000+0200", trafficTime.End)
    assert.Equal(t, "2019-04-16T02:11:02.000+0200", trafficTime.Resume)
}

func TestGetStartAsTime(t *testing.T) {
    trafficTime := TrafficTime{Start: "2019-04-14T02:11:02.000+0200"}
    start, _ := time.Parse(time.RFC3339, "2019-04-14T02:11:02.000+0200")
    actualStart, _ := trafficTime.GetStartAsTime()
    assert.Equal(t, start, actualStart)
}

func TestGetEndAsTime(t *testing.T) {
    trafficTime := TrafficTime{End: "2019-04-15T02:11:02.000+0200"}
    end, _ := time.Parse(time.RFC3339, "2019-04-15T02:11:02.000+0200")
    actualEnd, _ := trafficTime.GetEndAsTime()
    assert.Equal(t, end, actualEnd)
}

func TestGetResumeAsTime(t *testing.T) {
    trafficTime := TrafficTime{Resume: "2019-04-16T02:11:02.000+0200"}
    resume, _ := time.Parse(time.RFC3339, "2019-04-16T02:11:02.000+0200")
    actualResume, _ := trafficTime.GetResumeAsTime()
    assert.Equal(t, resume, actualResume)
}
