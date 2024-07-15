package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	
	"log"
	"io/ioutil"
	
)

func serveweatherBit(w http.ResponseWriter, r *http.Request){
	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w,"City name is required", http.StatusBadRequest)
		return
	}
	
	weatherData, err := weatherBitForecast(city)
	if err != nil {
		log.Printf("Error! Something went wrong: %v\n",err)
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(weatherData)
}


func weatherBitForecast(city string) (WeatherbitForecast, error){
	weatherbitAPIKey := "9fa3bd564e6b49c1a6ab91b1b03b8eae"
	apiURL := "https://api.weatherbit.io/v2.0/forecast/daily?city=%s&days=7&key=%s"
	url := fmt.Sprintf(apiURL, city, weatherbitAPIKey)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return WeatherbitForecast{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return WeatherbitForecast{}, fmt.Errorf("error reading response body: %v", err)
	}

	var forecast WeatherbitForecast
	if err := json.Unmarshal(body, &forecast); err != nil {
		return WeatherbitForecast{}, fmt.Errorf("error decoding response: %v", err)
	}

	return forecast, nil
}





func serveAPI3(w http.ResponseWriter, r *http.Request){
	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w,"City name is required", http.StatusBadRequest)
		return
	}
	
	weatherData, err := WeatherAPI3(city)
	if err != nil {
		log.Printf("Error! Something went wrong: %v\n",err)
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(weatherData)
}

func WeatherAPI3(city string) (WeatherAPI3Response, error) {
	weatherAPIKey := "c64f0240269f46a4a2d214353241805"

   	 apiURL := "http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no"
    	url := fmt.Sprintf(apiURL, weatherAPIKey, city)

    	log.Printf("Requesting URL: %s", url)

    	client := &http.Client{Timeout: 10 * time.Second}
    	resp, err := client.Get(url)
    	if err != nil {
       	 return WeatherAPI3Response{}, err
    	}
    	defer resp.Body.Close()

    	if resp.StatusCode != http.StatusOK {
      	  return WeatherAPI3Response{}, fmt.Errorf("received non-200 response status: %s", resp.Status)
    	}

   	 body, err := ioutil.ReadAll(resp.Body)
   	 if err != nil {
    	    return WeatherAPI3Response{}, fmt.Errorf("error reading response body: %v", err)
   	 }

    	log.Printf("Response body: %s", string(body))

    	var weatherAPIResponse WeatherAPI3Response
    	if err := json.Unmarshal(body, &weatherAPIResponse); err != nil {
    	    return WeatherAPI3Response{}, fmt.Errorf("error decoding response: %v", err)
   	 }
   	 return weatherAPIResponse, nil
}
