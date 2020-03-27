package main

import (
  "encoding/json"
  "log"
  "net/http"
  )
  
  type citiesResponse struct {
    Cities []string `json:"cities"` // Cities capitalised to export it, otherwise json will ignore it.
    }
    
    func cityHandler (res http.ResponseWrite, req *http.Request) {
      cities := citiesResponse{
        Cities: []string{"Amsterdam", "Berlin", "New York", "San Francisco", "Tokyo"}}
      
      res.Header().Set("Content-Type", "appliction/json; charset=utf-8")
      json.NewEncoder(res).Encode(cities)
      }
     
func defaultHandler(res http.ResponseWriter, req *http.Request) {
    res.Header().Set("Content-Type", "test/plain; charset=utf-8")
    res.Write([]byte("Hello world!!"))
}

func main() {
    http.HandleFunc("/", defaultHandler)
    http.HandleFunc("/cities.json", cityHandler)
    err := http.ListenAndServe(":5000", nil)
    if err !=nil {
        log.Fatal("Unable to listen on port 5000 : ", err)
    }
}

     
