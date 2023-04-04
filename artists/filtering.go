package artists

import (
	"groupie-tracker/variables"
	"log"
	"strconv"
	"strings"
)

func FilterArtists(search string, cdf string, cdt string, fadf string, fadt string, mem []string, loc string) {
	var searchResults []string
	var filteredArtists []variables.Artist
	if search != "" {
		searchResults = fromSearch(search)
	}
	if ((search != "" && len(searchResults) != 0) || search == "") && len(cdf) > 0 { //if Creation Date From Exists
		copyArtist := false
		cdfn := ValuesToInt(cdf)
		cdtn := ValuesToInt(cdt)
		if cdtn >= cdfn { //if Creation Date From is not after Creation Date To
			for _, artist := range variables.Artists {
				if artist.CreationDate >= cdfn && artist.CreationDate <= cdtn { //if artists creation date is in between these dates
					copyArtist = inSearch(artist.Name, searchResults)
					if len(loc) > 0 { //if location is defined
						copyArtist = false
						for location := range artist.LocationDates {
							if location == loc {
								copyArtist = inSearch(artist.Name, searchResults)
							}
						}
					}
					if copyArtist && (len(fadf) > 0 || len(fadt) > 0) { //if First Album Dates are defined
						copyArtist = false
						fa := ValuesToInt(artist.FirstAlbum[6:10])
						if len(fadf) > 0 && len(fadt) > 0 { //if both are defined
							fadfn := ValuesToInt(fadf)
							fadtn := ValuesToInt(fadt)
							if fa >= fadfn && fa <= fadtn {
								copyArtist = inSearch(artist.Name, searchResults)
							}
						} else if len(fadf) > 0 { //if From is defined
							fadfn := ValuesToInt(fadf)
							if fa >= fadfn {
								copyArtist = inSearch(artist.Name, searchResults)
							}
						} else if len(fadt) > 0 { //if To is defined
							fadtn := ValuesToInt(fadt)
							if fa <= fadtn {
								copyArtist = inSearch(artist.Name, searchResults)
							}
						}
					}
					if copyArtist && len(mem) > 0 { //if number of Members is defined
						copyArtist = false
						memLen := len(artist.Members)
						if memLen > 0 {
							for _, nr := range mem {
								if ValuesToInt(nr) == memLen {
									copyArtist = inSearch(artist.Name, searchResults)
								}
							}
						}
					}
					if copyArtist { //if all filters apply to this artist, add artist to list
						filteredArtists = append(filteredArtists, artist)
					}
				}
			}
		}
		variables.NewPage.Artists = filteredArtists
	} else if (search != "" && len(searchResults) == 0) && len(cdf) > 0 {
		variables.NewPage.Artists = filteredArtists
	} else {
		variables.NewPage.Artists = variables.Artists
	}
}

func fromSearch(a string) []string {
	var b []string
	for _, artist := range variables.Artists {
		if a == artist.Name {
			return []string{a}
		} else {
			if strings.Contains(artist.Name, a) {
				b = append(b, artist.Name)
			}
			for _, member := range artist.Members {
				if strings.Contains(member, a) {
					b = append(b, artist.Name)
				}
			}
			if strings.Contains(strconv.Itoa(artist.CreationDate), a) {
				b = append(b, artist.Name)
			}
			if strings.Contains(artist.FirstAlbum, a) {
				b = append(b, artist.Name)
			}
			for location := range artist.LocationDates {
				if strings.Contains(location, a) {
					b = append(b, artist.Name)
				}
			}
		}
	}
	return b
}

func inSearch(name string, names []string) bool {
	if len(names) > 0 {
		for _, one := range names {
			if name == one {
				return true
			}
		}
		return false
	} else {
		return true
	}
}

func ValuesToInt(a string) int { //string to int and check for errors
	an, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)
	}
	return an
}
