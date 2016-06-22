package main

import (
	"encoding/json"
	"net/http"
)

type Ip_response_success struct {
	As          string
	City        string
	Country     string
	CountryCode string
	Isp         string
	Lat         float64
	Lon         float64
	Org         string
	Query       string
	Region      string
	RegionName  string
	Status      string	
	Timezone    string
	Zip         string
}

func Query(url string) (Ip_response_success, error) {
	resp, err := http.Get(url)
	if err != nil {
		return Ip_response_success{}, err
	}

	var ip_response Ip_response_success
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&ip_response)
	if err != nil {
		return Ip_response_success{}, err
	}
	return ip_response, nil
}
