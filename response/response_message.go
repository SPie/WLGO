package response

import "time"

type ResponseMessage struct {
    Value string `json:"value"`
    MessageCode int `json:"messageCode"`
    ServerTime string `json:"serverTime"`
}

func (responseMessage ResponseMessage) GetServerTimeAsTime() (time.Time, error) {
    return time.Parse(time.RFC3339, responseMessage.ServerTime)
}
