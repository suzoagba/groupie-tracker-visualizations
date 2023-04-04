package api

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/variables"
	"io"
	"log"
	"net/http"
)

func GetFromApi() {
	fmt.Println("Please wait, loading information from APIs:")

	fmt.Println("- Getting sub-links from API")
	ParseJSON(variables.ApiLink, &variables.ApiSubLinks) // get info from API

	fmt.Println("- Getting artists info from API")
	ParseJSON(variables.ApiSubLinks.Artists, &variables.Artists) // get artists from API

	fmt.Println("- Getting relations from info")
	ParseJSON(variables.ApiSubLinks.Relation, &variables.Relations) // get relations from API
}

// Parse the API JSON and put it in v
func ParseJSON(link string, v interface{}) {
	rawJSON := pullData(link)
	err := json.Unmarshal([]byte(rawJSON), v)
	if err != nil {
		log.Fatal(err)
	}
}

// Pull file from link
func pullData(link string) string {
	var content []byte
	res, err := http.Get(link)
	if err == nil {
		content, err = io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
	return string(content)
}
