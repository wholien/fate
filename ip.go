package main

import (
	"encoding/json"
	"net/http"
)

type Ip_response_success struct {
	as          string
	city        string
	country     string
	countryCode string
	isp         string
	lat         string
	lon         string
	org         string
	query       string
	region      string
	regionName  string
	status      string	
	timezone    string
	zip         string
}

func Query(url string) (Ip_response_success, error) {
	resp, err := http.Get(url)
	if err != nil {
		return Ip_response_success{}, err
	}

	ip_response := Ip_response_success{}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&ip_response)
	if err != nil {
		return Ip_response_success{}, err
	}
	return ip_response, nil
}
