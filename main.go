package main

import (
	"log"
	"fmt"
	"os"
	"encoding/json"
	"net/http"
	"strings"
)


type apiConfigData struct {
	OpenWeatherMapApiKey string `json:"OpenWeatherMapKey`
}

type weatherData struct{
	Name string `json:"name"`
	Main struct{
		Kelvin float64 `json:"temp"`
	}`json:"main"`
}

func loadApiConfig(filename string) (apiConfigData, error){
	bytes, err := os.ReadFile(filename)

	if err != nil {
		return apiConfigData{}, err
	}

	var c apiConfigData

	err =json.Unmarshal(bytes, &c)
	if err != nil {
		return apiConfigData{}, err
	}
	return c, nil
}

func hello (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from go!\n"))

}

func query(city string)(weatherData, error){
	apiConfig, err := loadApiConfig(".apiConfig")
	if err != nil {
		return weatherData{}, err
	}

	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + apiConfig.OpenWeatherMapApiKey + "&q=" + city)
	if err != nil{
		return weatherData{}, err
	}

	defer resp.Body.Close()

	var d weatherData
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}
	return d, nil
}

func main (){
	http.HandleFunc("/hello", hello)

	http.HandleFunc("/weather/",
	func(w http.ResponseWriter, r *http.Request){
		city := strings.SplitN(r.URL.Path, "/", 3)[2]
		data,err := query(city)
		if err != nil{
			http.Error(w,err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type","application/json; charset=utf-8")
		json.NewEncoder(w).Encode(data)
	})


	port := ":3000"
	fmt.Printf("Server running on http://localhost%s\n", port)

	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatal("Server encountered an error: %v", err)
	}
}