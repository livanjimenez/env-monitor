package mock

import "github.com/livanjimenez/env-monitor/data"

func FetchAirQualityData() []data.AirQualityData {
	CO := data.GenerateMockData().CO
	CO2 := data.GenerateMockData().CO2
	Humidity := data.GenerateMockData().Humidity
	PM10 := data.GenerateMockData().PM10
	PM25 := data.GenerateMockData().PM25
	O3 := data.GenerateMockData().O3
	Temperature := data.GenerateMockData().Temperature

	return []data.AirQualityData{
		{
			CO:          CO,
			CO2:         CO2,
			Humidity:    Humidity,
			PM10:        PM10,
			PM25:        PM25,
			O3:          O3,
			Temperature: Temperature,
		},
	}
}
