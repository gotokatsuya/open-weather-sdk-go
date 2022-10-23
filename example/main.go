package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gotokatsuya/open-weather-sdk-go/openweather"
)

func main() {
	cli, err := openweather.NewClient(os.Getenv("OPEN_WEATHER_API_KEY"), http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		return
	}
	res, httpRes, err := cli.OneCall(context.Background(), &openweather.OneCallRequest{
		Lat: 35.68944,
		Lon: 139.69167,
		Exclude: []string{
			openweather.ExcludeMinutely,
			openweather.ExcludeHourly,
			openweather.ExcludeAlerts,
		},
		Lang: openweather.LangJa,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	if httpRes.StatusCode != http.StatusOK {
		fmt.Println(res.ErrorCode)
		fmt.Println(res.ErrorMessage)
		return
	}
	fmt.Println(httpRes.StatusCode)
	fmt.Println(res)
}
