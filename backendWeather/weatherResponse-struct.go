package main

type WeatherResponse struct {
    Latitude      float64 `json:"latitude"`
    Longitude     float64 `json:"longitude"`
    CurrentWeather struct {
        Temperature float64 `json:"temperature"`
        Weathercode int     `json:"weathercode"`
        Time        string  `json:"time"`
    } `json:"current_weather"`
}

type DynamicForecast struct {
    Latitude  float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
    Daily struct {
        Time                   []string  `json:"time"`
        WeatherCode            []int     `json:"weather_code"`
        Temperature2mMax       []float64 `json:"temperature_2m_max"`
        Temperature2mMin       []float64 `json:"temperature_2m_min"`
        ApparentTemperatureMax []float64 `json:"apparent_temperature_max"`
        ApparentTemperatureMin []float64 `json:"apparent_temperature_min"`
    } `json:"daily"`
}

type StartEndDate struct {
    Latitude  float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
    Daily struct {
        Time                   []string  `json:"time"`
        WeatherCode            []int     `json:"weather_code"`
        Temperature2mMax       []float64 `json:"temperature_2m_max"`
        Temperature2mMin       []float64 `json:"temperature_2m_min"`
        ApparentTemperatureMax []float64 `json:"apparent_temperature_max"`
        ApparentTemperatureMin []float64 `json:"apparent_temperature_min"`
    } `json:"daily"`
}

type VisualCrossingDynamic struct {
    Datetime     string  `json:"datetime"`
    TempMax      float64 `json:"tempmax"`
    TempMin      float64 `json:"tempmin"`
    FeelsLikeMax float64 `json:"feelslikemax"`
    FeelsLikeMin float64 `json:"feelslikemin"`
}

type VisualCrossingDay struct {
    Datetime     string                  `json:"datetime"`
    TempMax      float64                 `json:"tempmax"`
    TempMin      float64                 `json:"tempmin"`
    FeelsLikeMax float64                 `json:"feelslikemax"`
    FeelsLikeMin float64                 `json:"feelslikemin"`
    Hours        []VisualCrossingHours   `json:"hours"`
}

type VisualCrossingResponse struct {
    Address   string                  `json:"address"`
    Latitude  float64                 `json:"latitude"`
    Longitude float64                 `json:"longitude"`
    Days      []VisualCrossingDynamic `json:"days"`
}

type VisualCrossingHours struct {
    Datetime   string  `json:"datetime"`
    Temp       float64 `json:"temp"`
    Feelslike  float64 `json:"feelslike"`
    Conditions string  `json:"conditions"`
    Precipprob float64 `json:"precipprob"`
}

type VisualResponseToday struct {
    Address   string              `json:"address"`
    Latitude  float64             `json:"latitude"`
    Longitude float64             `json:"longitude"`
    Days      []VisualCrossingDay `json:"days"`
}

type VisualResponseData struct {
    Datetime     string  `json:"datetime"`
    Temp         float64 `json:"temp"`
    TempMax      float64 `json:"tempmax"`
    TempMin      float64 `json:"tempmin"`
    FeelsLike    float64 `json:"feelslike"`
    FeelsLikeMax float64 `json:"feelslikemax"`
    FeelsLikeMin float64 `json:"feelslikemin"`
    Conditions   string  `json:"conditions"`
    Description  string  `json:"description"`
    Precip       float64     `json:"precip"`
}

type VisualResponseDate struct {
    Address   string               `json:"address"`
    Latitude  float64              `json:"latitude"`
    Longitude float64              `json:"longitude"`
    Days      []VisualResponseData `json:"days"`
}


type WeatherAPI3Response struct {
    Location struct {
        Name           string `json: "name"`
        Region         string `json: "region"`
        Country        string `json: "country"`
        Lat            float64 `json: "lat"`
        Lon            float64 `json: "lon"`
        Localtime      string `json:  "localtime"`
    } `json:"location"`
	Current struct {
        Temp_c            float64 `json: "temp_c"`
        Temp_f            float64 `json: "temp_f"`
        IsDay            int     `json: "is_day"`
        Condition struct {
        	Text string `json: "text"`
        	Code int `json: "code"`
        }`json:"condition"`
        Feelslike_c float64 `json: "feelslike_c"`
        Feelslike_f float64 `json: "feelslike_f"`
        
	}
}


type WeatherbitForecast struct {
	Data []WeatherbitDay `json: "data"`
}

type WeatherbitDay struct {
	Datetime  string  `json: "datetime"`
	App_max_temp float64 `json: "app_max_temp"`
	App_min_temp float64 `json: "app_min_temp"`
	Max_temp float64 `json: "max_temp"`
	Min_temp float64 `json: "min_temp"`
	Precip  float64 `json: "precip"`
	Temp float64 `json: "temp"`
	Weather WeatherDescription `json: "weather"`
}

type WeatherDescription struct {
	Description string `json: "description"`
}



















