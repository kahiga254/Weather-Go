# Weather API Server in Go

## Overview
This is a simple weather API server built with Go. It fetches weather data from OpenWeatherMap based on a city name provided in the request.

## Features
- Fetches weather information using OpenWeatherMap API
- Serves responses in JSON format
- Handles API configuration from a file
- Provides error handling for missing or invalid requests

## Prerequisites
Ensure you have the following installed:
- Go (1.16 or later)
- OpenWeatherMap API Key

## Installation
1. Clone the repository:
   ```sh
   git clone https://github.com/kahiga254/go-weather-api.git
   cd go-weather-api
   ```
2. Create an `.apiConfig` file in the project root and add your OpenWeatherMap API key:
   ```json
   {
       "OpenWeatherMapKey": "YOUR_API_KEY_HERE"
   }
   ```
3. Install dependencies:
   ```sh
   go mod tidy
   ```

## Usage
1. Run the server:
   ```sh
   go run main.go
   ```
2. Open your browser or use cURL to get weather data:
   ```sh
   curl http://localhost:3000/weather/Nairobi
   ```

## API Endpoints
### GET `/weather/{city}`
Fetches weather data for the given city.
#### Example Request:
```sh
curl http://localhost:3000/weather/London
```
#### Example Response:
```json
{
    "name": "London",
    "main": {
        "kelvin": 289.92
    }
}
```

## Code Explanation
### `loadApiConfig(filename string) (apiConfigData, error)`
Reads the API key from a file and returns it.

### `query(city string) (weatherData, error)`
- Reads API key using `loadApiConfig`.
- Sends a request to OpenWeatherMap API.
- Parses JSON response into `weatherData` struct.

### `http.HandleFunc("/weather/", handlerFunction)`
- Extracts the city name from the request URL.
- Calls `query(city)` to fetch weather data.
- Encodes the response as JSON and sends it to the client.

## Error Handling
- If the API key is missing, the server returns an error.
- If OpenWeatherMap API request fails, an internal server error is returned.
- If the requested city is invalid, the OpenWeatherMap API will return an error message.

## License
This project is open-source and available under the MIT License.

## Autho
Adams kahiga -[ [Your GitHu](https://github.com/kahiga254)

