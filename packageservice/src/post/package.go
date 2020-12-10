package post

import (
	"fmt"

	"../postalservice"
)

type Package struct {
	Origin      string
	Destination string
}

func (p Package) PostPackage() (string, error) {

	postalService := postalservice.CanadaPostService{
		Location: "Toronto",
	}

	result, err := postalService.Send(p.Origin, p.Destination)

	if err != nil {
		return "", fmt.Errorf("%+v", err)
	}

	return result, nil
}
