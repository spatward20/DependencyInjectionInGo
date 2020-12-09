package main

import "fmt"

type Package struct {
	recipientCountry string
	senderCountry    string
	postalService    postalservice.PostalService
}

func (p Package) PostPackage() error {
	result, err := p.postalService.Send()
	if err != nil {
		return fmt.Errorf("Effor: %+v", err)
	}
	return nil
}
