package response

type MonitorResponse struct {
    MonitorResponseData MonitorResponseData `json:"data"`
    ResponseMessage ResponseMessage `json:"message"`
}
