package main

import (
	"fmt"
	"github.com/uncolorr/gsonbuilder/builder"
	"log"
)

func main() {
	var gsonClassBuilder builder.GsonClassBuilder
	jsonData := "[{”foo”:[{”m”:3}], \"m\":5, \"arr\":[4,5]}]"
	res, err := gsonClassBuilder.Parse(jsonData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

}
