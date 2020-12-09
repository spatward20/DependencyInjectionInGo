package post

import (
	"errors"
	"fmt"
	"postalservice"
	"strings"
)

type Package struct {
	OriginCountry      string
	DestinationCountry string
}

func (p Package) PostPackage() (string, error) {

	postalService := postalservice.CanadaPostService{}

	if !strings.Contains(p.DestinationCountry, "Canada") {
		return "", errors.New("Sorry, We don't ship outside Canada")
	}

	result, err := postalService.Send(p.OriginCountry, p.DestinationCountry)

	if err != nil {
		return "", fmt.Errorf("Effor: %+v", err)
	}

	return result, nil
}
