package main

import (
	"fmt"
	"post"
)

func main() {

	p := post.Package{
		Destination: "Toronto, Canada",
		Origin:      "Halifax, Canada",
	}

	result, err := p.PostPackage()

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(result)
}
