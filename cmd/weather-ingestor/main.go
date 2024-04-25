package main

import (
	"fmt"
	"weather-ingestor/domain/forecast_office"
	"weather-ingestor/domain/grid_point"
	"weather-ingestor/domain/points"
)

// And the forecast: https://api.weather.gov/gridpoints/BOU/64,73/forecast

// Stations are located via https://api.weather.gov/stations?state=CO

func main() {
	point, err := points.New(40.0150, -105.2705)
	if err != nil {
		panic(err)
	}

	office, err := forecast_office.Instances.InstanceFromString(point.Properties.GridId)
	if err != nil {
		panic(err)
	}
	fmt.Println(office.ForecastOfficeName)

	gridPoint, err := grid_point.New(point.Properties.GridId, point.Properties.GridX, point.Properties.GridY)
	if err != nil {
		panic(err)
	}

	fmt.Println(gridPoint.Properties.Temperature.Uom)
}
