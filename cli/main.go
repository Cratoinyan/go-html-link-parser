package main

import (
	"flag"
	"fmt"
	"gophercises/link"
	"os"
)

func main() {
	filePath := flag.String("file", "ex1.html", "html file that will be parsed")
	reader, err := os.Open(*filePath)
	if err != nil {
		fmt.Errorf(err.Error())
		panic(err)
	}

	defer reader.Close()

	links, err := link.Parse(reader)

	if err != nil {
		fmt.Errorf(err.Error())
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
