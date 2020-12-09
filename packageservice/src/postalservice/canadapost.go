package postalservice

import (
	"errors"
	"fmt"
)

type CanadaPostService struct {
	Location string
}

func (cp CanadaPostService) Send(origin, destination string) (string, error) {
	if origin == "" || destination == "" {
		return "", errors.New("Invalid parameters")
	}
	fmt.Println("Sending package from " + origin + " to " + destination)
	fmt.Println("Don't worry, your package will get there eventually! Thank you for using Canada Post!")
	return "success", nil
}
