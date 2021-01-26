package openweathermap_test

import (
	"os"
	"testing"

	"github.com/nstkshkn/go-openweathermap"
)

var weather *openweathermap.Weather

func Test_Init(t *testing.T) {
	weather = openweathermap.NewWeather(os.Getenv("WEATHER_TOKEN"))
}

func TestWeatherByName(t *testing.T) {
	_, err := weather.WeatherByName("Moscow", openweathermap.UnitsMetric)
	if err != nil {
		t.Errorf("TestWeatherByName error: %s\n", err)
	}
}

func TestWeatherByCityID(t *testing.T) {
	_, err := weather.WeatherByCityID(2172797, openweathermap.UnitsImperial)
	if err != nil {
		t.Errorf("TestWeatherByCityID error: %s\n", err)
	}
}

func TestWeatherByGeographicCoordinates(t *testing.T) {
	_, err := weather.WeatherByGeographicCoordinates(53.2, 45, openweathermap.UnitsStandard)
	if err != nil {
		t.Errorf("TestWeatherByGeographicCoordinates error: %s\n", err)
	}
}
