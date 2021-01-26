package openweathermap

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

const (
	BaseURL = "http://api.openweathermap.org/data/2.5/weather?%s"
)

type Units string

const (
	UnitsStandard Units = "standard"
	UnitsMetric   Units = "metric"
	UnitsImperial Units = "imperial"
)

func NewWeather(token string) *Weather {
	return &Weather{
		APIKey: token,
	}
}

func (wt *Weather) WeatherByName(location string, units Units) (response *WeatherResponse, err error) {

	client := http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf(BaseURL, fmt.Sprintf("q=%s&appid=%s&units=%s", url.QueryEscape(location), wt.APIKey, units)), nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			return nil, err
		} else {
			return response, nil
		}
	} else {
		return nil, errors.New("unexpected error")
	}

}

func (wt *Weather) WeatherByCityID(cityID int, units Units) (response *WeatherResponse, err error) {

	client := http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf(BaseURL, fmt.Sprintf("id=%d&appid=%s&units=%s", cityID, wt.APIKey, units)), nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			return nil, err
		} else {
			return response, nil
		}
	} else {
		return nil, errors.New("unexpected error")
	}

}

func (wt *Weather) WeatherByGeographicCoordinates(lat, lon float64, units Units) (response *WeatherResponse, err error) {

	client := http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf(BaseURL, fmt.Sprintf("lat=%f&lon=%f&appid=%s&units=%s", lat, lon, wt.APIKey, units)), nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			return nil, err
		} else {
			return response, nil
		}
	} else {
		return nil, errors.New("unexpected error")
	}

}
