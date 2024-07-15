package main

import (
	
	"net/http"
	"time"
	
)


func dealingCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Access-Control-Allow-Origin","*")
		w.Header().Set("Access-Control-Allow-Methods","GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers","Content-Type, Authorization")
		next.ServeHTTP(w,r)
	})
}








var ValidPaths = []string{"/weather"}

func isPathV(path string) bool{

	for i := 0; i<=len(ValidPaths); i++{
		if (path == ValidPaths[i]){
			return true
		}
		
	}
	return false
}


type Weather int

func (Weather) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	
	if (!isPathV(r.URL.Path)) {
		w.WriteHeader(404)
		w.Write([]byte("{\"code\": 404, \"error\": not found}"))
        return
	}


	weatherType := r.URL.Query().Get("type")
	if handler, ok := forecastType[weatherType]; ok {
		handler(w,r)
	} else{
	 http.Error(w,"Invalid weather type", http.StatusBadRequest)
}


}

func NewServer(w Weather) *http.Server {

	handler := dealingCors(w)

	return &http.Server{

		Addr:         "192.168.15.39:8080",
		Handler:      handler,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
}




