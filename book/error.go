package main

import (
	"errors"
	"fmt"
)

var validError = errors.New("valid error")
var validError2 = errors.New("valid error")

func says() error {
	return validError
}

func main() {
	err := says()
	if err == validError{
		fmt.Println(err)
	}
	if err == validError2{
		fmt.Println(err)
	}
}
