// Package main does a simple print out based on the input in file. I'm using this purely for the purposes of testing github-actions
package main

import (
	"fmt"
	"log"
)

func printAnimal(animal string) (string, error) {
	if animal != "" {
		return fmt.Sprintf("Your animal is: %s", animal), nil
	}
	return "", fmt.Errorf("why didn't you give me an animal")
}

func main() {
	animal := "Cat"

	res, err := printAnimal(animal)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(res)
}
