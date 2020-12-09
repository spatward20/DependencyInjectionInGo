package postalservice

import (
	"errors"
	"fmt"
)

type NorthPolePostService struct {
}

func (np NorthPolePostService) Send(origin, destination string) (string, error) {
	if origin == "" || destination == "" {
		return "", errors.New("Invalid parameters")
	}

	fmt.Println("We will always deliver your letters before Christmas!")
	return "success", nil
}
