package handler

import (
	"groupie-tracker/artists"
	"groupie-tracker/variables"
)

func formValues(search string, location string, memberNrs []string, albumDateFrom string, albumDateTo string, creationDateTo string, creationDateFrom string) {
	// Search-bar
	variables.NewPage.Filters.FormData.Search = search
	// Concert location
	variables.NewPage.Filters.FormData.ConcertLocation = location
	// Members
	/*
		Make every memberNr to an Int
	*/
	var memberNrsInt []int
	for i := 0; i < len(memberNrs); i++ {
		memberNrsInt = append(memberNrsInt, artists.ValuesToInt(memberNrs[i]))
	}
	variables.NewPage.Filters.FormData.CheckedMembers = memberNrsInt
	// First album dates
	/*
		Making sure we are not trying to change an empty string to an integer
		Empty string means no value from the form input
	*/
	if albumDateFrom != "" && albumDateTo != "" {
		variables.NewPage.Filters.FormData.FirstAlbum.From = artists.ValuesToInt(albumDateFrom)
		variables.NewPage.Filters.FormData.FirstAlbum.To = artists.ValuesToInt(albumDateTo)
	} else if albumDateFrom != "" {
		variables.NewPage.Filters.FormData.FirstAlbum.From = artists.ValuesToInt(albumDateFrom)
	} else if albumDateTo != "" {
		variables.NewPage.Filters.FormData.FirstAlbum.To = artists.ValuesToInt(albumDateTo)
	} else { // If user hasn't filtered by First album date, we set the dates to 0
		variables.NewPage.Filters.FormData.FirstAlbum.From = variables.NewPage.Filters.FirstAlbumMin
		variables.NewPage.Filters.FormData.FirstAlbum.To = variables.NewPage.Filters.FirstAlbumMax
	}
	// Creation date
	if creationDateFrom != "" && creationDateTo != "" {
		variables.NewPage.Filters.FormData.CreationDate.From = artists.ValuesToInt(creationDateFrom)
		variables.NewPage.Filters.FormData.CreationDate.To = artists.ValuesToInt(creationDateTo)
	} else {
		variables.NewPage.Filters.FormData.CreationDate.From = 1958
		variables.NewPage.Filters.FormData.CreationDate.To = 2015
	}
}
