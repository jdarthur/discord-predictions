package src

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

// Parseable is the interface that allows a struct to be
// parsed out from the main events.json file
type Parseable interface {
	UniqueKey() string    // this is a unique key that marks a line as relevant, e.g. "predicted_age"
	ParseInto() Parseable // this is the struct we will parse the JSON line into
}

// ParseFileForGender opens the given filename and parses out all of the
// Gender items it can find in the file, returning a ListOfGender
// that can be passed to the Graph function
func ParseFileForGender(filename string) (ListOfGender, error) {
	fmt.Println("Parsing predicted gender data: ")

	v, err := ParseFileForParseable(filename, Gender{})
	if err != nil {
		return nil, err
	}

	output := make([]*Gender, 0)
	for _, value := range v {
		output = append(output, value.(*Gender))
	}

	sort.Slice(output, func(i, j int) bool {
		return output[i].ModelVersion.Before(output[j].ModelVersion)
	})

	fmt.Printf("Got %d items\n", len(output))

	return output, nil
}

// ParseFileForAge opens the given filename and parses out all of the
// Age items it can find in the file, returning a ListOfAge
// that can be passed to the Graph function
func ParseFileForAge(filename string) (ListOfAge, error) {
	fmt.Println("Parsing predicted age data: ")

	v, err := ParseFileForParseable(filename, Age{})
	if err != nil {
		return nil, err
	}

	output := make([]*Age, 0)
	for _, value := range v {
		output = append(output, value.(*Age))
	}

	sort.Slice(output, func(i, j int) bool {
		return output[i].ModelVersion.Before(output[j].ModelVersion)
	})

	fmt.Printf("Got %d items\n", len(output))

	return output, nil

}

// ParseFileForParseable parses the events.json filename for a particular Parseable
// type, returning the list of []Parseable items that were found in the file
func ParseFileForParseable(filename string, p Parseable) ([]Parseable, error) {

	// we will save all of our parsed values in here
	output := make([]Parseable, 0)

	// open the file
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	// read all of the data
	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	// split the input on the newline character
	lines := bytes.Split(data, []byte("\n"))

	// we'll use this to break out of the loop when we are done
	found := false

	// find all of the items that match the Parseable.UniqueKey
	// that we are interested in
	for _, rawLine := range lines {

		// check if the line contains our unique key
		s := string(rawLine)
		if strings.Contains(s, p.UniqueKey()) {

			// found what we are looking for
			found = true

			// parse it into the provided struct
			v := p.ParseInto()
			err := json.Unmarshal(rawLine, &v)
			if err != nil {
				return nil, err
			}

			output = append(output, v)
		} else {

			// we have found all of the relevant info we are going to find,
			// so we can break out of the loop once we no longer have the
			// unique key the we were interested in in the line.
			if found {
				break
			}
		}
	}

	return output, nil
}
