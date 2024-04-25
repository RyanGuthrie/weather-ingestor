package points

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func New(lat, long float32) (*Instance, error) {
	instance := Instance{
		lat:  lat,
		long: long,
	}

	return instance.get()
}

type Instance struct {
	lat  float32
	long float32

	Context  []interface{} `json:"@context"`
	Id       string        `json:"id"`
	Type     string        `json:"type"`
	Geometry struct {
		Type        string    `json:"type"`
		Coordinates []float32 `json:"coordinates"`
	} `json:"geometry"`
	Properties struct {
		Id                  string `json:"@id"`
		Type                string `json:"@type"`
		Cwa                 string `json:"cwa"`
		ForecastOffice      string `json:"forecastOffice"`
		GridId              string `json:"gridId"`
		GridX               int    `json:"gridX"`
		GridY               int    `json:"gridY"`
		Forecast            string `json:"forecast"`
		ForecastHourly      string `json:"forecastHourly"`
		ForecastGridData    string `json:"forecastGridData"`
		ObservationStations string `json:"observationStations"`
		RelativeLocation    struct {
			Type     string `json:"type"`
			Geometry struct {
				Type        string    `json:"type"`
				Coordinates []float64 `json:"coordinates"`
			} `json:"geometry"`
			Properties struct {
				City     string `json:"city"`
				State    string `json:"state"`
				Distance struct {
					UnitCode string  `json:"unitCode"`
					Value    float64 `json:"value"`
				} `json:"distance"`
				Bearing struct {
					UnitCode string `json:"unitCode"`
					Value    int    `json:"value"`
				} `json:"bearing"`
			} `json:"properties"`
		} `json:"relativeLocation"`
		ForecastZone    string `json:"forecastZone"`
		County          string `json:"county"`
		FireWeatherZone string `json:"fireWeatherZone"`
		TimeZone        string `json:"timeZone"`
		RadarStation    string `json:"radarStation"`
	} `json:"properties"`
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
	escapedLatLong := url.QueryEscape(fmt.Sprintf("%.2f,%.2f", o.lat, o.long))
	return fmt.Sprintf("https://api.weather.gov/points/%s", escapedLatLong)
}
