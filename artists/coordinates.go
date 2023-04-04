package artists

import (
	"fmt"
	"groupie-tracker/api"
	"groupie-tracker/variables"
	"strconv"
	"strings"
)

var (
	coordToken string = "pk.eyJ1Ijoid2lsbGVta3VuaW5nYXMiLCJhIjoiY2t2amxpdDlvMDVyajJvdWdxNXZsMDdqMiJ9.yq_MUbhZMIAl5bFPQ0tNdw"
	coordUrl   string = "https://api.mapbox.com/geocoding/v5/mapbox.places/[location].json?limit=1&access_token=" + coordToken
)

type featureCollection struct {
	Features []struct {
		Center []float64 `json:"center"`
	}
}

// Gets coordinates from API
func getCoordinates(location string) string {
	var box featureCollection
	api.ParseJSON(strings.Replace(coordUrl, "[location]", location+" ", -1), &box)
	return fmt.Sprintf("%f", box.Features[0].Center)
}

// Adds coordinates info to artist
func getCoordinatesFromAPI(a []string) {
	d := make(map[string]string)
	for _, b := range a {
		d[b] = getCoordinates(b)
	}
	for artistIndex, artist := range variables.Artists {
		var coordList []string
		for key := range artist.LocationDates {
			coordList = append(coordList, d[key])
		}
		variables.Artists[artistIndex].LocationCoord = coordList
		variables.Artists[artistIndex].LocationMap = linkFunc(coordList)
	}
}

// Creates link for map
func linkFunc(a []string) string {
	mapStyle := "streets-v11"
	beginning := "https://api.mapbox.com/styles/v1/mapbox/" + mapStyle + "/static/geojson(%7B%22type%22%3A%22MultiPoint%22%2C%22coordinates%22%3A["
	mapZoom := "1"
	var forLink string
	var lat int
	var long int
	var count int
	for _, b := range a {
		c := strings.Split(b, " ")
		lat += negNum(c[0][1:])
		long += negNum(c[1])
		count++
		forLink += strings.Replace(b, " ", "%2C", -1) + "%2C"
	}
	lat /= count
	long /= count
	mapCenter := strconv.Itoa(lat) + "," + strconv.Itoa(long)
	end := "]%7D)/" + mapCenter + "," + mapZoom + "/1000x900?access_token="
	return beginning + forLink[:len(forLink)-3] + end + coordToken
}

func negNum(a string) int {
	if string(a[1]) == "-" {
		d, _ := strconv.Atoi(strings.Split(a, ".")[0][1:])
		return -1 * d
	} else {
		d, _ := strconv.Atoi(strings.Split(a, ".")[0])
		return d
	}
}
