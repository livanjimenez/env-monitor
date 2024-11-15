package models

type AirQualityModel struct {
	ID          int     `json:"id"`
	PM25        float64 `json:"pm2_5"`
	PM10        float64 `json:"pm10"`
	CO2         float64 `json:"co2"`
	CO          float64 `json:"co"`
	O3          float64 `json:"o3"`
	Humidity    float64 `json:"humidity"`
	Temperature float64 `json:"temperature"`
}
