package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/wholien/go-yelp/yelp"
)

func main() {
	ip, err := Query("http://ip-api.com/json")
	fmt.Printf("%#v\n", ip)
	// get config keys
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
	
	// make phone search query
	number := os.Args[1]
	phoneOptions := yelp.PhoneOptions{Phone: number}
	results, err := client.PhoneSearch(phoneOptions)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Printf("\nFound a total of %v results for number %v.\n", results.Total, number)
	fmt.Println("-----------------------------")
	for i := 0; i < len(results.Businesses); i++ {
		fmt.Printf("%v\t\t%v\n", results.Businesses[i].Name, results.Businesses[i].Rating)
	}
}
