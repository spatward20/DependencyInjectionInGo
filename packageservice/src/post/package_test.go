package post

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestSendPackageSuccess(t *testing.T) {

	deliveryService := postalservice.CanadaPostService{
		Location: "Toronto",
	}

	p := Package{
		Destination:     "Canada",
		Origin:          "Canada",
		DeliveryService: deliveryService,
	}

	result, err := p.PostPackage()

	assert.Equal(t, err, nil)
	assert.Equal(t, result, "success")
}
