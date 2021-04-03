package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

func getWebSite(address string) (string, string, string) {
	website := ""
	phone := ""
	adrs := ""
	// Instantiate default collector
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting: ", r.URL)
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// fmt.Println("before Link", link)
		if strings.HasPrefix(link, "/biz_redir?url") {
			website = e.Text
		}
	})

	//  p>p.css-e81eai
	c.OnHTML("body", func(e *colly.HTMLElement) {
		adrs = e.ChildText(`#wrap > div.main-content-wrap.main-content-wrap--full > yelp-react-root > div > div.margin-t3__373c0__1l90z.margin-b6__373c0__2Azj6.border-color--default__373c0__3-ifU > div > div > div.margin-b6__373c0__2Azj6.border-color--default__373c0__3-ifU > div > div.stickySidebar--fullHeight__373c0__1szWY.arrange-unit__373c0__o3tjT.arrange-unit-grid-column--4__373c0__33Wpc.padding-l2__373c0__1Dr82.border-color--default__373c0__3-ifU > div > div > section.margin-b3__373c0__q1DuY.border-color--default__373c0__3-ifU > div > div:nth-child(3) > div > div.arrange-unit__373c0__o3tjT.arrange-unit-fill__373c0__3Sfw1.border-color--default__373c0__3-ifU > p`)
		if strings.Contains(adrs,"View Service Area"){
			adrs = ""
		}else{
			adrs = strings.Replace(adrs, "Get Directions", "", -1)
		}
		phone = e.ChildText(`#wrap > div.main-content-wrap.main-content-wrap--full > yelp-react-root > div > div.margin-t3__373c0__1l90z.margin-b6__373c0__2Azj6.border-color--default__373c0__3-ifU > div > div > div.margin-b6__373c0__2Azj6.border-color--default__373c0__3-ifU > div > div.stickySidebar--fullHeight__373c0__1szWY.arrange-unit__373c0__o3tjT.arrange-unit-grid-column--4__373c0__33Wpc.padding-l2__373c0__1Dr82.border-color--default__373c0__3-ifU > div > div > section.margin-b3__373c0__q1DuY.border-color--default__373c0__3-ifU > div > div:nth-child(2) > div > div.arrange-unit__373c0__o3tjT.arrange-unit-fill__373c0__3Sfw1.border-color--default__373c0__3-ifU > p.css-1h1j0y3`)
		
	})

	
	c.Visit(address)

	return website, phone, adrs
}
