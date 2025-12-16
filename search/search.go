package search

import (
	"fmt"

	"github.com/go-rod/rod"
)

func SearchPeople(page *rod.Page, keyword string) []string {
	searchURL := "https://www.linkedin.com/search/results/people/?keywords=" + keyword

	page.MustNavigate(searchURL)
	page.MustWaitLoad()

	fmt.Println("🔍 Searching for:", keyword)

	unique := make(map[string]bool)
	var profiles []string

	links := page.MustElements(`a[href*="/in/"]`)
	for _, link := range links {
		href, err := link.Attribute("href")
		if err != nil || href == nil {
			continue
		}

		url := *href
		if !unique[url] {
			unique[url] = true
			profiles = append(profiles, url)
		}
	}

	return profiles
}
