package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `[
		{
			"first_name":"AKM",
			"last_name":"KH",
			"hair_color":"Black",
			"has_dog":"false",
		},
		{
			"first_name":"KK",
			"last_name":"KH",
			"hair_color":"Black",
			"has_dog":"true",
		},
	]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		log.Println("Log unmarshalled error", err)

	}
	log.Printf("unmarshalled %v", unmarshalled)

}
