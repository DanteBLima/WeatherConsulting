package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"strconv"
	"log"
	"io/ioutil"
	
)



func serveCurrent(w http.ResponseWriter, r *http.Request){
	lat, lon, err := getLatLong(r)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	weather, err := CurrentWeather(lat, lon)
	if err != nil{
		log.Printf("Error! Something went wrong: %v\n",err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(weather)
	

	
}

func CurrentWeather(lat, lon float64) (WeatherResponse, error) {
	

	apiULR := "https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current_weather=true&timezone=America/Sao_Paulo"

	url:= fmt.Sprintf(apiULR, lat, lon)
	

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return WeatherResponse{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
    	if err != nil {
        return WeatherResponse{}, fmt.Errorf("error reading response body: %v", err)
    }

	var weatherResponse WeatherResponse
	if err := json.Unmarshal(body, &weatherResponse); err != nil {
		return WeatherResponse{}, err
	}
	return weatherResponse, nil
	
	
	
}




func serveDynamic(w http.ResponseWriter, r *http.Request){
	
	lat,lon, err := getLatLong(r)
	if err != nil {
		http.Error(w,err.Error(), http.StatusBadRequest)
		return
	}
	
	dynamicWeather, err := DynamicWeather(lat,lon)
	if err !=nil {
		log.Printf("Error! Something went wrong: %v\n",err)
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dynamicWeather)
	
}

func DynamicWeather(lat, lon float64) (DynamicForecast, error) {
	
	apiURL := "https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&daily=weather_code,temperature_2m_max,temperature_2m_min,apparent_temperature_max,apparent_temperature_min&timezone=America/Sao_Paulo"
	
	url:= fmt.Sprintf(apiURL, lat, lon)
	
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
	return DynamicForecast{}, err
	}
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
    	if err != nil {
        return DynamicForecast{}, fmt.Errorf("error reading response body: %v", err)
    }
	
	
	var dynamicResponse DynamicForecast 
	
	if err := json.Unmarshal(body, &dynamicResponse); err != nil {
		return DynamicForecast{}, err
	}
	return dynamicResponse, nil
}


func serveStartEndDate(w http.ResponseWriter, r *http.Request){

	lat,lon, err := getLatLong(r)
	if err != nil {
		http.Error(w,err.Error(), http.StatusBadRequest)
		return
	}

	startDate := r.URL.Query().Get("start")
	endDate := r.URL.Query().Get("end")
	if startDate == "" || endDate == "" {
		http.Error(w, "Start and End dates required for this type of consulting", http.StatusBadRequest)
		return
	}

	startEnd, err := StartEndDateWeather (lat, lon, startDate, endDate)
	if err != nil {
		log.Printf("Error! Something went wrong: %v\n",err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(startEnd)

}

func StartEndDateWeather(lat, lon float64, startDate, endDate string) (StartEndDate, error){
	
	apiURL := "https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&start_date=%s&end_date=%s&daily=temperature_2m_max,temperature_2m_min,weather_code,apparent_temperature_max,apparent_temperature_min&timezone=America/Sao_Paulo"
	url := fmt.Sprintf(apiURL, lat, lon, startDate, endDate)
	
	client := &http.Client{Timeout: 10 * time.Second}
    	resp, err := client.Get(url)
    	if err != nil {
        	return StartEndDate{}, err
    }
    defer resp.Body.Close()
    
     body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return StartEndDate{}, fmt.Errorf("error reading response body: %v", err)
    }
    
    var startEndDateResponse StartEndDate
    if err := json.Unmarshal(body, &startEndDateResponse); err != nil {
        return StartEndDate{}, fmt.Errorf("error decoding response: %v", err)
    }
    return startEndDateResponse, nil
	


}

func getLatLong(r *http.Request) (float64, float64, error){
	latStr := r.URL.Query().Get("lat")
	lonStr := r.URL.Query().Get("lon")
	
	if latStr == "" || lonStr == ""{
		return 0,0,fmt.Errorf("Lacking latitude and/or longitude")
	}
	
	
	lat, err := strconv.ParseFloat(latStr,64)	
	if err != nil {
		return 0,0,fmt.Errorf("Invalid Latitude")
	}
	
	lon, err := strconv.ParseFloat(lonStr,64)	
	if err != nil {
		return 0,0,fmt.Errorf("Invalid Longitude")
	}
	
	return lat,lon, nil
}


















