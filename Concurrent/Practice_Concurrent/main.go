package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const apiKey = "502b6b5f783eeed88240b5e28fc10fb5"

// pass channel into function
func fetchWeather(city string, ch chan<- string, wg *sync.WaitGroup) interface{} {
	var data struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}

	defer wg.Done()

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
	//pass info into channel
	ch <- fmt.Sprintf("%s Temp: %v", city, data)

	return data
}

func main() {
	startNow := time.Now()

	cities := []string{"Toronto", "London", "Paris", "Tokyo"}
	//1. make a channel and wait group
	ch := make(chan string)
	var wg sync.WaitGroup

	for _, city := range cities {
		//add 1 wg for each goroutine
		wg.Add(1)
		//initiate goroutine
		go fetchWeather(city, ch, &wg)
	}

	//initiate goroutine
	go func() {
		//wait for all wait group then close channel
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		//print all result from channel
		fmt.Println(result)
	}

	fmt.Println("The operation took: ", time.Since(startNow))
}
