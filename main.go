package main

import (
	"fmt"

	"github.com/livanjimenez/env-monitor/data"
)

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

// goal: 
// fetch air quality data and print it to the console
// collect data from iot device/ESP32 with sensors; use emulator for now
// send data to a server and update db
// fetch data from db and display it in a web app / dashboard to end users

// thoughts:
// write in go the backend of an air quality monitoring system
// write embedded code for esp32 in C or C++? not sure yet

func main() {
	airQualityData := FetchAirQualityData()
	for i := 0; i < 10; i++ {
		fmt.Println(airQualityData)
	}
}