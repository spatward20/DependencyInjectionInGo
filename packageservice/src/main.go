package main

import (
	"fmt"
	"post"
	"postalservice"
)

func main() {

	deliveryService := postalservice.NorthPolePostService{
		Location: "North Branch",
	}

	p := post.Package{
		Destination:     "North Pole",
		Origin:          "Halifax, Canada",
		DeliveryService: deliveryService,
	}

	result, err := p.PostPackage()

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(result)
}
