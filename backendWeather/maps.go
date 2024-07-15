package main

import "net/http"

var forecastType = map[string]func(http.ResponseWriter, *http.Request) {
	
	"current": serveCurrent,
	"dynamic": serveDynamic,
	"historical": serveStartEndDate,
	
	"current2": serveVisualCurr,
	"dynamic2": serveVisual,
	"historical2": serveVisualDate,
	
	"current3": serveAPI3,
	"dynamic3": serveweatherBit,

}




