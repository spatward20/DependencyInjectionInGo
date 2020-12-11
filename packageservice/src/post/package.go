package post

import (
	"fmt"

	"postalservice"
)

type Package struct {
	Origin          string
	Destination     string
	DeliveryService postalservice.PostalService
}

func (p Package) PostPackage() (string, error) {

	result, err := p.DeliveryService.Send(p.Origin, p.Destination)

	if err != nil {
		return "", fmt.Errorf("%+v", err)
	}

	return result, nil
}
