package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const apiKey = "502b6b5f783eeed88240b5e28fc10fb5"

// pass channel into function
func fetchWeather(city string) interface{} {
	var data struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Error fetching weather info for %s: %s", city, err)
		return data
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Printf("Error decoding weather data for %s: %s", city, err)
		return data
	}

	return data
}

func main() {
	startNow := time.Now()

	cities := []string{"Toronto", "London", "Paris", "Tokyo"}

	for _, city := range cities {
		data := fetchWeather(city)
		fmt.Println(city, "Temp: ", data)
	}

	fmt.Println("The operation took: ", time.Since(startNow))
}
