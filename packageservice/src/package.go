package main

import (
	"errors"
	"fmt"
)

type Package struct {
	recipientCountry string
	senderCountry    string
}

func (p Package) PostPackage() error {

	postalService := postalservice.CanadaPostService {}
	if p.recipientCountry != "Canada" {
		return errors.New("Sorry, We don't ship outside Canada")
	}

	result, err := postalService.Send()

	if err != nil {
		return fmt.Errorf("Effor: %+v", err)
	}

	return nil
}
