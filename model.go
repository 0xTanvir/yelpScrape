package main

// type PageList struct {
// 	Searchpageprops struct {
// 		Searchmapprops struct {
// 			Hovercarddata map[string]interface{} `json:"hovercardData"`
// 		} `json:"searchMapProps"`
// 	} `json:"searchPageProps"`
// }

// type Result struct {
// 	Rating        float64  `json:"rating"`
// 	Photopageurl  string   `json:"photoPageUrl"`
// 	Name          string   `json:"name"`
// 	Neighborhoods []string `json:"neighborhoods"`
// 	Businessurl   string   `json:"businessUrl"`
// 	Isad          bool     `json:"isAd"`
// 	Phone         string   `json:"phone"`
// 	Photosrcset   string   `json:"photoSrcSet"`
// 	Photosrc      string   `json:"photoSrc"`
// 	Numreviews    int      `json:"numReviews"`
// 	Addresslines  []string `json:"addressLines"`
// }

type NumberOfResult struct {
	Searchpageprops struct {
		Maincontentcomponentslistprops []struct {
			Props struct {
				Totalresults   int `json:"totalResults,omitempty"`
				Resultsperpage int `json:"resultsPerPage,omitempty"`
				Startresult    int `json:"startResult,omitempty"`
			} `json:"props,omitempty"`
		} `json:"mainContentComponentsListProps"`
	} `json:"searchPageProps"`
}

type PageList struct {
	Searchpageprops struct {
		Maincontentcomponentslistprops []struct {
			Searchresultbusiness struct {
				Name             string `json:"name"`
				Businessurl      string `json:"businessUrl"`
				Isad             bool   `json:"isAd"`
				Phone            string `json:"phone"`
				Formattedaddress string `json:"formattedAddress"`
			} `json:"searchResultBusiness,omitempty"`
		} `json:"mainContentComponentsListProps"`
	} `json:"searchPageProps"`
}

type SaveResult struct {
	Name    string `json:"name" csv:"name"`
	Yelpurl string `json:"yelpurl" csv:"yelpurl"`
	Mainurl string `json:"mainurl" csv:"mainurl"`
	Phone   string `json:"phone" csv:"phone"`
	Address string `json:"address" csv:"address"`
}
