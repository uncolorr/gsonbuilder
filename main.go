package main

import (
	"fmt"
	"github.com/uncolorr/gsonbuilder/builder"
	"log"
)

func main() {

	var gsonClassBuilder builder.GsonClassBuilder
	res, err := gsonClassBuilder.Parse("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

}
