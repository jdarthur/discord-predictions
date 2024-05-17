package main

import (
	"discord-predictions/src"
	"flag"
)

var genderFilename = "gender.html"
var ageFilename = "age.html"

func main() {

	filename := flag.String("f", "", "full name of the events.json filename we are going to read")
	flag.Parse()

	if filename == nil || *filename == "" {
		flag.PrintDefaults()
		return
	}

	// parse the main file for predicted gender
	gender, err := src.ParseFileForGender(*filename)
	if err != nil {
		panic(err)
	}

	src.Graph(gender, genderFilename)

	// parse the main file for predicted age
	age, err := src.ParseFileForAge(*filename)
	if err != nil {
		panic(err)
	}

	src.Graph(age, ageFilename)
}
