package variables

// API links
type apiCategories struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

var (
	ApiLink     = "https://groupietrackers.herokuapp.com/api"
	ApiSubLinks apiCategories
)

// Artist information
type Artist struct {
	ID            int      `json:"id"`
	Image         string   `json:"image"` // Image URL
	Name          string   `json:"name"`
	Members       []string `json:"members"`
	CreationDate  int      `json:"creationDate"`
	FirstAlbum    string   `json:"firstAlbum"`
	LocationDates map[string][]string
	LocationCoord []string
	LocationMap   string
}

// Artist ID and Location/Dates
type relation struct {
	Index []struct {
		ID             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

// Filter for webpage
type Filter struct {
	Members         []int
	CreationDateMin int
	CreationDateMax int
	CreationDates   []int
	FirstAlbumMin   int
	FirstAlbumMax   int
	FirstAlbumDates []int
	Locations       []string
	FormData        struct {
		Search          string
		ConcertLocation string
		CheckedMembers  []int
		FirstAlbum      struct {
			From int
			To   int
		}
		CreationDate struct {
			From int
			To   int
		}
	}
}

var (
	Artists    []Artist
	Relations  relation
	ForFilters Filter
	NewPage    ForPage
)

type ForPage struct {
	Filters Filter
	Artists []Artist
	Search  []Artist
}
