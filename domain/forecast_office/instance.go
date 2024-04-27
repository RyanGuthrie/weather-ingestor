package forecast_office

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type queryValues struct {
	ForecastOfficeName Name
}

type Instance struct {
	queryValues

	Context struct {
		Version string `json:"@version"`
		Vocab   string `json:"@vocab"`
	} `json:"@context"`
	Type    string `json:"@type"`
	Id      string `json:"@id"`
	Id1     string `json:"id"`
	Name    string `json:"name"`
	Address struct {
		Type            string `json:"@type"`
		StreetAddress   string `json:"streetAddress"`
		AddressLocality string `json:"addressLocality"`
		AddressRegion   string `json:"addressRegion"`
		PostalCode      string `json:"postalCode"`
	} `json:"address"`
	Telephone                   string   `json:"telephone"`
	FaxNumber                   string   `json:"faxNumber"`
	Email                       string   `json:"email"`
	SameAs                      string   `json:"sameAs"`
	NwsRegion                   string   `json:"nwsRegion"`
	ParentOrganization          string   `json:"parentOrganization"`
	ResponsibleCounties         []string `json:"responsibleCounties"`
	ResponsibleForecastZones    []string `json:"responsibleForecastZones"`
	ResponsibleFireZones        []string `json:"responsibleFireZones"`
	ApprovedObservationStations []string `json:"approvedObservationStations"`
}

func New(name Name) (*Instance, error) {
	instance := Instance{
		queryValues: queryValues{
			ForecastOfficeName: name,
		},
	}

	return instance.get()
}

func NewLazy(name Name) (*Instance, error) {
	instance := Instance{
		queryValues: queryValues{
			ForecastOfficeName: name,
		},
	}

	return &instance, nil
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
	sprintf := fmt.Sprintf("https://api.weather.gov/offices/%s", o.queryValues.ForecastOfficeName)
	return sprintf
}
