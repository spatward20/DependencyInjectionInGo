package postalservice

import (
	"testing"

	"../postalservice"
	"gopkg.in/go-playground/assert.v1"
)

func TestSendPackageSuccess(t *testing.T) {

	deliveryService := postalservice.CanadaPostService{}

	p := Package{
		DestinationCountry: "Canada",
		OriginCountry:      "Canada",
		DeliveryService:    deliveryService,
	}

	result, err := p.PostPackage()

	assert.Equal(t, err, nil)
	assert.Equal(t, result, "success")
}
