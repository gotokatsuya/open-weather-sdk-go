package data

import (
	_ "embed"
	"reflect"
	"testing"
)

func TestGetNearestProvinceCity(t *testing.T) {
	type args struct {
		lat float64
		lon float64
	}
	tests := []struct {
		name    string
		args    args
		want    *GeoPoint
		wantErr bool
	}{
		{
			name: "Tokyo",
			args: args{
				lat: 35.66174295035835,
				lon: 139.72087293033627,
			},
			want: &GeoPoint{
				CountryCode:  "JP",
				ProvinceCode: "JP-13",
				ProvinceName: "東京都",
				City:         "新宿区",
				Lat:          35.68944,
				Lon:          139.69167,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetNearestProvinceCity(tt.args.lat, tt.args.lon)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNearestProvinceCity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNearestProvinceCity() = %v, want %v", got, tt.want)
			}
		})
	}
}
