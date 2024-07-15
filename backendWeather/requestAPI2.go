package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	
	"log"
	"io/ioutil"
	
)

func serveVisualCurr(w http.ResponseWriter, r *http.Request){
	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w,"City name is required", http.StatusBadRequest)
		return
	}
	
	weatherData, err := VisualCrossingToday(city)
	if err != nil {
		log.Printf("Error! Something went wrong: %v\n",err)
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(weatherData)
}


func VisualCrossingToday(city string) (VisualResponseToday, error){
	var visualCrossingAPIKey = "396GEQ95Z2VM4XJNMJFPG6CB2"

	apiURL := "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s/today?unitGroup=metric&key=%s"
	url := fmt.Sprintf(apiURL, city, visualCrossingAPIKey)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return VisualResponseToday{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return VisualResponseToday{}, fmt.Errorf("error reading response body: %v", err)
	}

	var weatherResponse VisualResponseToday
	if err := json.Unmarshal(body, &weatherResponse); err != nil {
		return VisualResponseToday{}, err
	}
	return weatherResponse, nil
}



func serveVisual(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "City name is required", http.StatusBadRequest)
		return
	}

	weatherData, err := VisualCrossingLast7Days(city)
	if err != nil {
		log.Printf("Error! Something went wrong: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weatherData)
}


	

func VisualCrossingLast7Days(city string) (VisualCrossingResponse, error) {

	var visualCrossingAPIKey = "396GEQ95Z2VM4XJNMJFPG6CB2"

	apiURL := "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s/last7days?unitGroup=metric&key=%s"
	url := fmt.Sprintf(apiURL, city, visualCrossingAPIKey)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return VisualCrossingResponse{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return VisualCrossingResponse{}, fmt.Errorf("error reading response body: %v", err)
	}

	var weatherResponse VisualCrossingResponse
	if err := json.Unmarshal(body, &weatherResponse); err != nil {
		return VisualCrossingResponse{}, err
	}
	return weatherResponse, nil
}


func serveVisualDate(w http.ResponseWriter, r *http.Request) {
    city := r.URL.Query().Get("city")
    if city == "" {
        http.Error(w, "City name is required", http.StatusBadRequest)
        return
    }

    startDate := r.URL.Query().Get("start")
    endDate := r.URL.Query().Get("end")
    if startDate == "" || endDate == "" {
        http.Error(w, "Start and End dates required for this type of consulting", http.StatusBadRequest)
        return
    }

    weatherData, err := VisualDate(city, startDate, endDate)
    if err != nil {
        log.Printf("Error! Something went wrong: %v\n", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(weatherData)
}


func VisualDate(city, startDate, endDate string) (VisualResponseDate, error) {
    visualCrossingAPIKey := "GHTZ5EZNQPBAE4SGDU64ZAXFW"
    //"4EAB7ZYJJW26EK3C54KKYXYTT"

    apiURL := "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s/%s/%s?key=%s"
    url := fmt.Sprintf(apiURL, city, startDate, endDate, visualCrossingAPIKey)

    log.Printf("Requesting URL: %s", url)

    client := &http.Client{Timeout: 10 * time.Second}
    resp, err := client.Get(url)
    if err != nil {
        return VisualResponseDate{}, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return VisualResponseDate{}, fmt.Errorf("received non-200 response status: %s", resp.Status)
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return VisualResponseDate{}, fmt.Errorf("error reading response body: %v", err)
    }

    log.Printf("Response body: %s", string(body))

    var startEndDateResponse VisualResponseDate
    if err := json.Unmarshal(body, &startEndDateResponse); err != nil {
        return VisualResponseDate{}, fmt.Errorf("error decoding response: %v", err)
    }
    return startEndDateResponse, nil
}
















