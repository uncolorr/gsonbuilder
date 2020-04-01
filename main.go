package main

import (
	"fmt"
	"github.com/uncolorr/gsonbuilder/builder"
	"log"
)

func main() {
	var gsonClassBuilder builder.GsonClassBuilder
	res, err := gsonClassBuilder.Parse("{\"foo\":[{\"m\":3}], \"m\":5, \"arr\":[4,5]}")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

}
