# Wiener Linien Client in Go

## Create new client

To create a new Wiener Linien client use the `github.com/spie/wlgo.NewClient()` function. The function requires three parameters:
1. `apiEndpoint`: The endpoint for the Wiener Linien API. It should be http://www.wienerlinien.at/ogd_realtime
2. `senderId`: This is the api key provided by Wiener Linien
3. `httpClient`: A http client of type `github.com/spie/wlgo/http.Client`

## Methods

The client has three methods:
1. `GetMonitor`: This function has two parameters and returns a `github.com/spie/wlgo/response.MonitorResponse` type.
  1. `stationNumbers []string`: This parameter is required.
  2. `faultTypes []string`: This parameter is optional. The posible values are  'stoerungkurz', 'stoerunglang' and 'aufzugsinfo'.
2. `GetTrafficInfoList`: This function returns a `github.com/spie/wlgo/response.TrafficInfoListResponse` type and takes three optional parameters.
  1. `relatedLines []string`
  2. `relatedStops []string`
  3. `names []string`
3. `GetNewsList`: This function returns a `github.com/spie/wlgo/response.NewsListResponse` type. It takes the same parameters as `GetTrafficInfoList` does.
