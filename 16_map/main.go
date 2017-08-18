package main

import "fmt"

func main() {
	var details = map[int]string{101: "Jordan", 102: "Roger", 103: "Rafel"}
	fmt.Println(details)
	fmt.Println(details[102])            //Accessing value based on keys
	countries := make(map[string]string) //Shorthand and make
	countries["US"] = "United States"
	countries["IND"] = "India"
	countries["UK"] = "United Kingdom"
	fmt.Println(countries)
	capitals := map[string]string{ //Shorthand composite literal
		"India": "Delhi",
		"US":    "Washington"}
	capitals["UK"] = "Great Britian" //Adding an entry
	capitals["NZ"] = "Wellington"
	capitals["UK"] = "London"                          //Updating entry
	fmt.Println(capitals, " of length", len(capitals)) //len() for Length
	delete(capitals, "NZ")                             //deleting entry
	fmt.Println(capitals)
	//Comma OK
	if value, present := capitals["US"]; present {
		fmt.Println("Value", value, "is present")
		fmt.Println("Status:", present)
	}
	for country, capital := range countries {
		fmt.Println(country, "-->", capital)
	}
	//map inside map
	CountryInfo := map[string]map[string]int{
		"INDIA":                    {"IN": 91},
		"United States of America": {"US": 1}, //always give the , after last element
	}
	for key, val := range CountryInfo {
		fmt.Println(key, "--->", val)
	}
}
