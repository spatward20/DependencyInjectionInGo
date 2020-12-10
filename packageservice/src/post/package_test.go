package postalservice

import (
	"testing"
)

func TestSendPackage(t *testing.T) {

	p := Package{
		DestinationCountry: "Canada",
		OriginCountry:      "Canada",
	}

	p.PostPackage()

}
