package postalservice

import "fmt"

type NorthPolePostService struct {
}

func (np NorthPolePostService) Send(origin, destination string) string, error {
	if origin == "" || destination == "" {
		return nil, errors.New("Invalid parameters")
	}

	fmt.Println("We will always deliver your letters before Christmas!")
	return "success", nil
}