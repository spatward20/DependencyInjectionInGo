package main

import (
	"fmt"
	"post"
)

func main() {

	p := post.Package{
		DestinationCountry: "Toronto, Canada",
		OriginCountry:      "Halifax, Canada",
	}
	result, err := p.PostPackage()

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(result)
}
