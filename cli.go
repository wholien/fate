package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/wholien/go-yelp/yelp"
	"github.com/guregu/null"
)

func main() {
	ip, err := Query("http://ip-api.com/json")
	if err != nil {
		fmt.Printf("%#v\n", err)
	}
	fmt.Printf("%#v\n", ip)

	//// get config keys
	var o yelp.AuthOptions
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &o)
	if err != nil {
		panic(err)
	}
	
	// new yelp client
	client := yelp.New(&o, nil)

	locationOptions := yelp.LocationOptions{
		ip.Zip,
		&yelp.CoordinateOptions{
			Latitude: null.FloatFrom(ip.Lat),
			Longitude: null.FloatFrom(ip.Lon),
		},
	}

	generalOptions := yelp.GeneralOptions{
		Term: "food",
		RadiusFilter: null.FloatFrom(2000),
	}

	searchOptions := yelp.SearchOptions{
		GeneralOptions: &generalOptions,
		LocationOptions: &locationOptions,
	}
	
	// make phone search query
	results, err := client.DoSearch(searchOptions)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Printf("\nFound a total of %v results.\n", results.Total)
	fmt.Println("-----------------------------")
	for i := 0; i < len(results.Businesses); i++ {
		fmt.Printf("%v\t\t%v\n", results.Businesses[i].Name, results.Businesses[i].Rating)
	}
}
