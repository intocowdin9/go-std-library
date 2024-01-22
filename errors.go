package main

import (
	"errors"
	"fmt"
)

var (
	validationError = errors.New("Validation Error")
	notFoundError   = errors.New("Not Found Error")
)

func getById(id string) error {

	if id == "" {
		return validationError
	}

	if id != "muhammad" {
		return notFoundError
	}

	return nil
}

func main() {

	err := getById("")

	if err != nil {
		if errors.Is(err, validationError) {
			fmt.Println("validation error")
		} else if errors.Is(err, notFoundError) {
			fmt.Println("not found error")
		} else {
			fmt.Println("unknow error")
		}
	}
}
