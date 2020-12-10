package postalservice

import (
	"errors"
	"fmt"
	"strings"

	"../postalservice"
)

type Package struct {
	OriginCountry      string
	DestinationCountry string
	DeliveryService    postalservice.PostalService
}

func (p Package) PostPackage() (string, error) {

	if !strings.Contains(p.DestinationCountry, "Canada") {
		return "", errors.New("Sorry, We don't ship outside Canada")
	}

	result, err := p.DeliveryService.Send(p.OriginCountry, p.DestinationCountry)

	if err != nil {
		return "", fmt.Errorf("Effor: %+v", err)
	}

	return result, nil
}
