package grid_point

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
	"weather-ingestor/domain/temperature"
	"weather-ingestor/domain/uom"
	"weather-ingestor/domain/weather_time"
)

type gridPoint struct {
	gridX int
	gridY int
}

func (g gridPoint) AsEscaped() string {
	return url.QueryEscape(fmt.Sprintf("%d,%d", g.gridX, g.gridY))
}

// queryValues are the values used to query an Instance of a GridPoint and not part of the actual hydrated response
type queryValues struct {
	gridID string
	point  gridPoint
}

type Instance struct {
	queryValues

	Context  []any  `json:"@context"`
	Id       string `json:"id"`
	Type     string `json:"type"`
	Geometry struct {
		Type        string        `json:"type"`
		Coordinates [][][]float64 `json:"coordinates"`
	} `json:"geometry"`
	Properties struct {
		Id         string                   `json:"@id"`
		Type       string                   `json:"@type"`
		UpdateTime time.Time                `json:"updateTime"`
		ValidTimes weather_time.WeatherTime `json:"validTimes"`
		Elevation  struct {
			UnitCode string  `json:"unitCode"`
			Value    float64 `json:"value"`
		} `json:"elevation"`
		ForecastOffice string                   `json:"forecastOffice"`
		GridId         string                   `json:"gridId"`
		GridX          string                   `json:"gridX"`
		GridY          string                   `json:"gridY"`
		Temperature    temperature.Temperatures `json:"temperature"`
		Dewpoint       struct {
			Uom    uom.UOM `json:"uom"`
			Values []struct {
				ValidTime weather_time.WeatherTime `json:"validTime"`
				Value     float64                  `json:"value"`
			} `json:"values"`
		} `json:"dewpoint"`
		MaxTemperature   temperature.Temperatures `json:"maxTemperature"`
		MinTemperature   temperature.Temperatures `json:"minTemperature"`
		RelativeHumidity struct {
			Uom    uom.UOM `json:"uom"`
			Values []struct {
				ValidTime weather_time.WeatherTime `json:"validTime"`
				Value     int                      `json:"value"`
			} `json:"values"`
		} `json:"relativeHumidity"`
		ApparentTemperature     temperature.Temperatures `json:"apparentTemperature"`
		WetBulbGlobeTemperature temperature.Temperatures `json:"wetBulbGlobeTemperature"`
		HeatIndex               struct {
			Uom    uom.UOM `json:"uom"`
			Values []struct {
				ValidTime weather_time.WeatherTime `json:"validTime"`
				Value     *float64                 `json:"value"`
			} `json:"values"`
		} `json:"heatIndex"`
		WindChill temperature.Temperatures `json:"windChill"`
		SkyCover  struct {
			Uom    uom.UOM `json:"uom"`
			Values []struct {
				ValidTime weather_time.WeatherTime `json:"validTime"`
				Value     int                      `json:"value"`
			} `json:"values"`
		} `json:"skyCover"`
		WindDirection struct {
			Uom    uom.UOM `json:"uom"`
			Values []struct {
				ValidTime weather_time.WeatherTime `json:"validTime"`
				Value     int                      `json:"value"`
			} `json:"values"`
		} `json:"windDirection"`
		WindSpeed struct {
			Uom    uom.UOM `json:"uom"`
			Values []struct {
				ValidTime weather_time.WeatherTime `json:"validTime"`
				Value     float64                  `json:"value"`
			} `json:"values"`
		} `json:"windSpeed"`
		WindGust struct {
			Uom    uom.UOM `json:"uom"`
			Values []struct {
				ValidTime weather_time.WeatherTime `json:"validTime"`
				Value     float64                  `json:"value"`
			} `json:"values"`
		} `json:"windGust"`
		Weather struct {
			Values []struct {
				ValidTime weather_time.WeatherTime `json:"validTime"`
				Value     []struct {
					Coverage   *string `json:"coverage"`
					Weather    *string `json:"weather"`
					Intensity  *string `json:"intensity"`
					Visibility struct {
						UnitCode string `json:"unitCode"`
						Value    any    `json:"value"`
					} `json:"visibility"`
					Attributes []any `json:"attributes"`
				} `json:"value"`
			} `json:"values"`
		} `json:"weather"`
		Hazards struct {
			Values []any `json:"values"`
		} `json:"hazards"`
		ProbabilityOfPrecipitation struct {
			Uom    uom.UOM `json:"uom"`
			Values []struct {
				ValidTime weather_time.WeatherTime `json:"validTime"`
				Value     float64                  `json:"value"`
			} `json:"values"`
		} `json:"probabilityOfPrecipitation"`
		QuantitativePrecipitation struct {
			Uom    uom.UOM `json:"uom"`
			Values []struct {
				ValidTime weather_time.WeatherTime `json:"validTime"`
				Value     float64                  `json:"value"`
			} `json:"values"`
		} `json:"quantitativePrecipitation"`
		IceAccumulation struct {
			Uom    uom.UOM `json:"uom"`
			Values []struct {
				ValidTime weather_time.WeatherTime `json:"validTime"`
				Value     float64                  `json:"value"`
			} `json:"values"`
		} `json:"iceAccumulation"`
		SnowfallAmount struct {
			Uom    uom.UOM `json:"uom"`
			Values []struct {
				ValidTime weather_time.WeatherTime `json:"validTime"`
				Value     float64                  `json:"value"`
			} `json:"values"`
		} `json:"snowfallAmount"`
		SnowLevel struct {
			Uom    uom.UOM `json:"uom"`
			Values []struct {
				ValidTime weather_time.WeatherTime `json:"validTime"`
				Value     float64                  `json:"value"`
			} `json:"values"`
		} `json:"snowLevel"`
		CeilingHeight struct {
			Uom    uom.UOM `json:"uom"`
			Values []struct {
				ValidTime weather_time.WeatherTime `json:"validTime"`
				Value     float64                  `json:"value"`
			} `json:"values"`
		} `json:"ceilingHeight"`
		Visibility struct {
			Uom    uom.UOM `json:"uom"`
			Values []struct {
				ValidTime weather_time.WeatherTime `json:"validTime"`
				Value     float64                  `json:"value"`
			} `json:"values"`
		} `json:"visibility"`
		TransportWindSpeed struct {
			Uom    uom.UOM `json:"uom"`
			Values []struct {
				ValidTime weather_time.WeatherTime `json:"validTime"`
				Value     float64                  `json:"value"`
			} `json:"values"`
		} `json:"transportWindSpeed"`
		TransportWindDirection struct {
			Uom    uom.UOM `json:"uom"`
			Values []struct {
				ValidTime weather_time.WeatherTime `json:"validTime"`
				Value     float64                  `json:"value"`
			} `json:"values"`
		} `json:"transportWindDirection"`
		MixingHeight struct {
			Uom    uom.UOM `json:"uom"`
			Values []struct {
				ValidTime weather_time.WeatherTime `json:"validTime"`
				Value     float64                  `json:"value"`
			} `json:"values"`
		} `json:"mixingHeight"`
		HainesIndex struct {
			Values []struct {
				ValidTime weather_time.WeatherTime `json:"validTime"`
				Value     float64                  `json:"value"`
			} `json:"values"`
		} `json:"hainesIndex"`
		LightningActivityLevel struct {
			Values []struct {
				ValidTime weather_time.WeatherTime `json:"validTime"`
				Value     float64                  `json:"value"`
			} `json:"values"`
		} `json:"lightningActivityLevel"`
		TwentyFootWindSpeed struct {
			Values []any `json:"values"`
		} `json:"twentyFootWindSpeed"`
		TwentyFootWindDirection struct {
			Values []any `json:"values"`
		} `json:"twentyFootWindDirection"`
		WaveHeight struct {
			Values []any `json:"values"`
		} `json:"waveHeight"`
		WavePeriod struct {
			Values []any `json:"values"`
		} `json:"wavePeriod"`
		WaveDirection struct {
			Values []any `json:"values"`
		} `json:"waveDirection"`
		PrimarySwellHeight struct {
			Values []any `json:"values"`
		} `json:"primarySwellHeight"`
		PrimarySwellDirection struct {
			Values []any `json:"values"`
		} `json:"primarySwellDirection"`
		SecondarySwellHeight struct {
			Values []any `json:"values"`
		} `json:"secondarySwellHeight"`
		SecondarySwellDirection struct {
			Values []any `json:"values"`
		} `json:"secondarySwellDirection"`
		WavePeriod2 struct {
			Values []any `json:"values"`
		} `json:"wavePeriod2"`
		WindWaveHeight struct {
			Values []any `json:"values"`
		} `json:"windWaveHeight"`
		DispersionIndex struct {
			Values []any `json:"values"`
		} `json:"dispersionIndex"`
		Pressure struct {
			Values []any `json:"values"`
		} `json:"pressure"`
		ProbabilityOfTropicalStormWinds struct {
			Values []any `json:"values"`
		} `json:"probabilityOfTropicalStormWinds"`
		ProbabilityOfHurricaneWinds struct {
			Values []any `json:"values"`
		} `json:"probabilityOfHurricaneWinds"`
		PotentialOf15MphWinds struct {
			Values []any `json:"values"`
		} `json:"potentialOf15mphWinds"`
		PotentialOf25MphWinds struct {
			Values []any `json:"values"`
		} `json:"potentialOf25mphWinds"`
		PotentialOf35MphWinds struct {
			Values []any `json:"values"`
		} `json:"potentialOf35mphWinds"`
		PotentialOf45MphWinds struct {
			Values []any `json:"values"`
		} `json:"potentialOf45mphWinds"`
		PotentialOf20MphWindGusts struct {
			Values []any `json:"values"`
		} `json:"potentialOf20mphWindGusts"`
		PotentialOf30MphWindGusts struct {
			Values []any `json:"values"`
		} `json:"potentialOf30mphWindGusts"`
		PotentialOf40MphWindGusts struct {
			Values []any `json:"values"`
		} `json:"potentialOf40mphWindGusts"`
		PotentialOf50MphWindGusts struct {
			Values []any `json:"values"`
		} `json:"potentialOf50mphWindGusts"`
		PotentialOf60MphWindGusts struct {
			Values []any `json:"values"`
		} `json:"potentialOf60mphWindGusts"`
		GrasslandFireDangerIndex struct {
			Values []any `json:"values"`
		} `json:"grasslandFireDangerIndex"`
		ProbabilityOfThunder struct {
			Values []struct {
				ValidTime weather_time.WeatherTime `json:"validTime"`
				Value     float64                  `json:"value"`
			} `json:"values"`
		} `json:"probabilityOfThunder"`
		DavisStabilityIndex struct {
			Values []any `json:"values"`
		} `json:"davisStabilityIndex"`
		AtmosphericDispersionIndex struct {
			Values []any `json:"values"`
		} `json:"atmosphericDispersionIndex"`
		LowVisibilityOccurrenceRiskIndex struct {
			Values []any `json:"values"`
		} `json:"lowVisibilityOccurrenceRiskIndex"`
		Stability struct {
			Values []any `json:"values"`
		} `json:"stability"`
		RedFlagThreatIndex struct {
			Values []any `json:"values"`
		} `json:"redFlagThreatIndex"`
	} `json:"properties"`
}

func New(gridID string, gridX, gridY int) (*Instance, error) {
	instance := Instance{
		queryValues: queryValues{
			gridID: gridID,
			point: gridPoint{
				gridX: gridX,
				gridY: gridY,
			},
		},
	}

	return instance.get()
}

func (o *Instance) get() (*Instance, error) {
	response, err := http.Get(o.url())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	if err = decoder.Decode(&o); err != nil {
		return nil, err
	}

	return o, nil
}

func (o *Instance) url() string {
	return fmt.Sprintf("https://api.weather.gov/gridpoints/%s/%s", o.gridID, o.point.AsEscaped())
}
