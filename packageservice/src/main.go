package main

import (
	"postalservice"
)

func main() {

	canadaPost := postalservice.CanadaPostService{}

	p := Package{
		recipientCountry: "US",
		senderCountry:    "Canada",
		postalService:    canadaPost,
	}
	p.PostPackage()

	// NorthPolePostService := postalservice.NorthPolePostService{}
	// p2 := Package{
	// 	recipientCountry: "US",
	// 	senderCountry:    "Canada",
	// 	postalService:    NorthPolePostService,
	// }
	// p2.PostPackage()
}
