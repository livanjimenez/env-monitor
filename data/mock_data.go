package data

import (
	"math/rand"
	"time"
)

type AirQualityData struct {
	PM25        float64 `json:"pm2_5"`       // PM2.5 in micrograms per cubic meter
	PM10        float64 `json:"pm10"`        // PM10 in micrograms per cubic meter
	CO2         float64 `json:"co2"`         // CO2 in parts per million (ppm)
	CO          float64 `json:"co"`          // CO in parts per million (ppm)
	O3          float64 `json:"o3"`          // Ozone (O3) in parts per billion (ppb)
	Humidity    float64 `json:"humidity"`    // Humidity in percentage
	Temperature float64 `json:"temperature"` // Temperature in Celsius
}

func GenerateMockData() AirQualityData {
	rand.Seed(time.Now().UnixNano())

	return AirQualityData{
		PM25:        rand.Float64()*100 + 10,   // Random value between 10 and 110
		PM10:        rand.Float64()*150 + 20,   // Random value between 20 and 170
		CO2:         rand.Float64()*500 + 400,  // Random value between 400 and 900 ppm
		CO:          rand.Float64()*10 + 0.1,   // Random value between 0.1 and 10.1 ppm
		O3:          rand.Float64()*100 + 10,   // Random value between 10 and 110 ppb
		Humidity:    rand.Float64()*50 + 30,    // Random value between 30% and 80%
		Temperature: rand.Float64()*15 + 15,    // Random value between 15 and 30 °C
	}
}
