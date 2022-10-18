package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gotokatsuya/open-weather-sdk-go/openweather"
	"github.com/gotokatsuya/open-weather-sdk-go/openweather/data"
)

func main() {
	cli, err := openweather.NewClient(os.Getenv("OPEN_WEATHER_API_KEY"), http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		return
	}
	res, httpRes, err := cli.OneCall(context.Background(), &openweather.OneCallRequest{
		Lat: data.ProvinceCities[0].Lat,
		Lon: data.ProvinceCities[0].Lon,
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
