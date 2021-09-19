package main

import (
	"errors"
	"log"
)

func main() {

	result, err := divied(100.0, 0.0)

	if err != nil {

		log.Println(err)
		return
	}
	log.Println("result of division is: ", result)

}

func divied(x, y float32) (float32, error) {

	var result float32
	if y == 0 {
		return result, errors.New("cannot divided by 0")
	}

	result = x / y

	return result, nil

}
