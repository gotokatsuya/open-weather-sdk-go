package openweather

import (
	"context"
	"net/http"
)

const (
	ExcludeCurrent  = "current"
	ExcludeMinutely = "minutely"
	ExcludeHourly   = "hourly"
	ExcludeDaily    = "daily"
	ExcludeAlerts   = "alerts"
)

const (
	UnitStandard = "standard"
	UnitMetric   = "metric"
	UnitImperial = "imperial"
)

const (
	LangJa = "ja"
)

// OneCallRequest type
type OneCallRequest struct {
	Lat     float64  `url:"lat"`
	Lon     float64  `url:"lon"`
	Exclude []string `url:"exclude,omitempty" del:","`
	Units   string   `url:"units,omitempty"`
	Lang    string   `url:"lang,omitempty"`
}

// OneCallResponse type
type OneCallResponse struct {
	ErrorResponse
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	Timezone       string  `json:"timezone"`
	TimezoneOffset int     `json:"timezone_offset"`
	Current        struct {
		Dt         int64   `json:"dt"`
		Sunrise    int64   `json:"sunrise"`
		Sunset     int64   `json:"sunset"`
		Temp       float64 `json:"temp"`
		FeelsLike  float64 `json:"feels_like"`
		Pressure   int     `json:"pressure"`
		Humidity   int     `json:"humidity"`
		DewPoint   float64 `json:"dew_point"`
		Uvi        float64 `json:"uvi"`
		Clouds     int     `json:"clouds"`
		Visibility int     `json:"visibility"`
		WindSpeed  float64 `json:"wind_speed"`
		WindDeg    int     `json:"wind_deg"`
		WindGust   float64 `json:"wind_gust"`
		Weather    []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
	} `json:"current"`
	Minutely []struct {
		Dt            int64 `json:"dt"`
		Precipitation int   `json:"precipitation"`
	} `json:"minutely"`
	Hourly []struct {
		Dt         int64   `json:"dt"`
		Temp       float64 `json:"temp"`
		FeelsLike  float64 `json:"feels_like"`
		Pressure   int     `json:"pressure"`
		Humidity   int     `json:"humidity"`
		DewPoint   float64 `json:"dew_point"`
		Uvi        float64 `json:"uvi"`
		Clouds     int     `json:"clouds"`
		Visibility int     `json:"visibility"`
		WindSpeed  float64 `json:"wind_speed"`
		WindDeg    int     `json:"wind_deg"`
		WindGust   float64 `json:"wind_gust"`
		Weather    []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Pop float64 `json:"pop"`
	} `json:"hourly"`
	Daily []struct {
		Dt        int64   `json:"dt"`
		Sunrise   int64   `json:"sunrise"`
		Sunset    int64   `json:"sunset"`
		Moonrise  int64   `json:"moonrise"`
		Moonset   int64   `json:"moonset"`
		MoonPhase float64 `json:"moon_phase"`
		Temp      struct {
			Day   float64 `json:"day"`
			Min   float64 `json:"min"`
			Max   float64 `json:"max"`
			Night float64 `json:"night"`
			Eve   float64 `json:"eve"`
			Morn  float64 `json:"morn"`
		} `json:"temp"`
		FeelsLike struct {
			Day   float64 `json:"day"`
			Night float64 `json:"night"`
			Eve   float64 `json:"eve"`
			Morn  float64 `json:"morn"`
		} `json:"feels_like"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		DewPoint  float64 `json:"dew_point"`
		WindSpeed float64 `json:"wind_speed"`
		WindDeg   int     `json:"wind_deg"`
		WindGust  float64 `json:"wind_gust"`
		Weather   []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Clouds int     `json:"clouds"`
		Pop    float64 `json:"pop"`
		Rain   float64 `json:"rain"`
		Uvi    float64 `json:"uvi"`
	} `json:"daily"`
	Alerts []struct {
		SenderName  string   `json:"sender_name"`
		Event       string   `json:"event"`
		Start       int64    `json:"start"`
		End         int64    `json:"end"`
		Description string   `json:"description"`
		Tags        []string `json:"tags"`
	} `json:"alerts"`
}

// OneCall method
func (c *Client) OneCall(ctx context.Context, req *OneCallRequest) (*OneCallResponse, *http.Response, error) {
	path := "/data/3.0/onecall"
	httpReq, err := c.NewRequest(http.MethodGet, path, req)
	if err != nil {
		return nil, nil, err
	}
	resp := new(OneCallResponse)
	httpResp, err := c.Do(ctx, httpReq, resp)
	if err != nil {
		return nil, httpResp, err
	}
	return resp, httpResp, nil
}
