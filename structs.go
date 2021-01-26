package openweathermap

type Weather struct {
	APIKey    string
	BaseUnits string
}

type WeatherResponse struct {
	GeoPos Coord  `json:"coord"`
	Main   Main   `json:"main"`
	Name   string `json:"name"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	FeelsLike float64 `json:"feels_like"`
	Pressure  float64 `json:"pressure"`
	SeaLevel  float64 `json:"sea_level"`
	GrndLevel float64 `json:"grnd_level"`
	Humidity  int     `json:"humidity"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}
