package openweathermap_test

import (
	"testing"

	openweathermap "github.com/nstkshkn/go-openweathermap"
)

var weather *openweathermap.Weather

func Test_Init(t *testing.T) {
	weather = openweathermap.NewWeather("Token")
}

func TestWeatherByName(t *testing.T) {
	_, err := weather.WeatherByName("Moscow")
	if err != nil {
		t.Errorf("TestWeatherByName error: %s\n", err)
	}
}

func TestWeatherByCityID(t *testing.T) {
	_, err := weather.WeatherByCityID(2172797)
	if err != nil {
		t.Errorf("TestWeatherByCityID error: %s\n", err)
	}
}

func TestWeatherByGeographicCoordinates(t *testing.T) {
	_, err := weather.WeatherByGeographicCoordinates(53.2, 45)
	if err != nil {
		t.Errorf("TestWeatherByGeographicCoordinates error: %s\n", err)
	}
}
