package data

import (
	_ "embed"
	"encoding/csv"
	"errors"
	"io"
	"log"
	"math"
	"strconv"
	"strings"
)

type GeoPoint struct {
	CountryCode  string
	ProvinceCode string
	ProvinceName string
	City         string
	Lat          float64
	Lon          float64
}

func (g1 GeoPoint) Distance(g2 GeoPoint) float64 {
	radlat1 := float64(math.Pi * g1.Lat / 180)
	radlat2 := float64(math.Pi * g2.Lat / 180)

	theta := float64(g1.Lon - g2.Lon)
	radtheta := float64(math.Pi * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)
	if dist > 1 {
		dist = 1
	}
	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515
	return dist * 1.609344
}

//go:embed data.csv
var data string

var ProvinceCities []GeoPoint

func init() {
	r := csv.NewReader(strings.NewReader(data))
	if _, err := r.Read(); err != nil {
		log.Println(err)
		return
	}
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
			continue
		}
		countryCode := row[0]
		provinceCode := row[1]
		provinceName := row[2]
		city := row[3]
		lat, err := strconv.ParseFloat(row[4], 64)
		if err != nil {
			log.Println(err)
			continue
		}
		lon, err := strconv.ParseFloat(row[5], 64)
		if err != nil {
			log.Println(err)
			continue
		}
		ProvinceCities = append(ProvinceCities, GeoPoint{
			CountryCode:  countryCode,
			ProvinceCode: provinceCode,
			ProvinceName: provinceName,
			City:         city,
			Lat:          lat,
			Lon:          lon,
		})
	}
}

var (
	ErrNotFoundProvinceCities = errors.New("not found province city data")
)

func GetNearestProvinceCity(lat, lon float64) (*GeoPoint, error) {
	if len(ProvinceCities) == 0 {
		return nil, ErrNotFoundProvinceCities
	}
	g1 := GeoPoint{Lat: lat, Lon: lon}
	minDist := math.MaxFloat64
	var nearest GeoPoint
	for _, g2 := range ProvinceCities {
		dist := g1.Distance(g2)
		minDist = math.Min(minDist, dist)
		if minDist == dist {
			nearest = g2
		}
	}
	return &nearest, nil
}
