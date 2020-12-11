package postalservice

import (
	"errors"
	"fmt"
	"strings"
)

type NorthPolePostService struct {
	Location string
}

func (npp NorthPolePostService) Send(origin, destination string) (string, error) {
	if origin == "" || destination == "" {
		return "", errors.New("Invalid parameters")
	}

	if !strings.Contains(destination, "North Pole") {
		return "", errors.New("Sorry, we only ship North Pole")
	}

	fmt.Println("Your letter to Santa will reach in time")
	return "success", nil
}
