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

func NewWeather(token string) *Weather {
	return &Weather{
		APIKey: token,
	}
}

func (wt *Weather) WeatherByName(location string) (response *WeatherResponse, err error) {

	client := http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf(BaseURL, fmt.Sprintf("q=%s&appid=%s", url.QueryEscape(location), wt.APIKey)), nil)
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

func (wt *Weather) WeatherByCityID(cityID int) (response *WeatherResponse, err error) {

	client := http.Client{}
	rqst, err := http.NewRequest("GET", fmt.Sprintf(BaseURL, fmt.Sprintf("id=%d&appid=%s", cityID, wt.APIKey)), nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(rqst)
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

func (wt *Weather) WeatherByGeographicCoordinates(lat, lon float64) (response *WeatherResponse, err error) {

	client := http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf(BaseURL, fmt.Sprintf("lat=%f&lon=%f&appid=%s", lat, lon, wt.APIKey)), nil)
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
