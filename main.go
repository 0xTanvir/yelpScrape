package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gocarina/gocsv"
)

const baseURL = "https://www.yelp.com/search/snippet?find_desc=%s&find_loc=%s&start=%d"

// https://www.yelp.com/search/snippet?find_desc=Restaurants&find_loc=San Francisco, CA&start=10
func main() {
	var q string
	flag.StringVar(&q, "q", "", "query you searching for")
	var l string
	flag.StringVar(&l, "l", "", "location you searching for")
	flag.Parse()

	tr, rpp, err := getTotalResult(q, l)
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Println("TotalResult: ", tr, "Result per page: ", rpp)

	allResult := []SaveResult{}

	q = url.QueryEscape(q)
	l = url.QueryEscape(l)

	for i := 0; i < tr; i = i + rpp {

		fmt.Println("visiting: ", fmt.Sprintf(baseURL, q, l, i))
		req, err := http.NewRequest("GET", fmt.Sprintf(baseURL, q, l, i), nil)
		if err != nil {
			log.Println("http request creation error: ", err.Error())
		}
		req.Header.Add("Accept", "application/json")

		// Create request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("client request creation error: ", err.Error())
		}
		defer resp.Body.Close()

		// Read response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("response body reading error: ", err.Error())
		}

		var pl PageList
		// Unmarshal body to techlandbdRspns struct
		err = json.Unmarshal(body, &pl)
		if err != nil {
			log.Println("response body json unmarshalling error: ", err.Error())
		}

		for _, rs := range pl.Searchpageprops.Maincontentcomponentslistprops {
			if !rs.Searchresultbusiness.Isad && rs.Searchresultbusiness.Businessurl != "" {
				mu, phone, adrs := getWebSite("https://www.yelp.com" + rs.Searchresultbusiness.Businessurl)
				sr := SaveResult{
					Name:    rs.Searchresultbusiness.Name,
					Yelpurl: "https://www.yelp.com" + rs.Searchresultbusiness.Businessurl,
					Mainurl: mu,
					Phone:   rs.Searchresultbusiness.Phone,
					Address: rs.Searchresultbusiness.Formattedaddress,
				}

				if phone != "" {
					sr.Phone = phone
				}
				// sr.Address = adrs
				if adrs != "" {
					sr.Address = adrs
				}

				allResult = append(allResult, sr)
				
			}
		}
	}

	csvFile, err := os.OpenFile("csvFile.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()
	gocsv.MarshalFile(&allResult, csvFile)

}

func getTotalResult(q, l string) (int, int, error) {
	q = url.QueryEscape(q)
	l = url.QueryEscape(l)

	req, err := http.NewRequest("GET", fmt.Sprintf(baseURL, q, l, 0), nil)
	if err != nil {
		log.Println("http request creation error: ", err.Error())
		return 0, 0, err
	}
	req.Header.Add("Accept", "application/json")

	// Create request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("client request creation error: ", err.Error())
		return 0, 0, err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("response body reading error: ", err.Error())
		return 0, 0, err
	}

	var nr NumberOfResult
	// Unmarshal body to techlandbdRspns struct
	err = json.Unmarshal(body, &nr)
	if err != nil {
		log.Println("response body json unmarshalling error: ", err.Error())
		return 0, 0, err
	}

	for _, listProps := range nr.Searchpageprops.Maincontentcomponentslistprops {
		if listProps.Props.Totalresults != 0 {
			return listProps.Props.Totalresults, listProps.Props.Resultsperpage, nil
		}
	}

	return 0, 0, errors.New("No number of result found, something wrong on request")
}
