package site

import (
	"fmt"
	"strings"
	"regexp"
)

type Site struct {
	Url	string		// url
	UrlReg	string		// regex
}

func (website *Site) IsUrlValid() bool {
	valid, err := regexp.MatchString("^http[s]?://(www.)?[a-z0-9]{3,}.(com|net)(.br)?$", website.Url)

	if err != nil {
		return false
	} else {
		return valid
	}
}

func (website *Site) GetAllUrls(body []byte) []string {
	var sites []string

	r, err := regexp.Compile("href=\"" + website.UrlReg)

	if err != nil {
		fmt.Println("There was an error")
	} else {
		for _, url := range r.FindAll(body, -1) {
			sites = append(sites, strings.ReplaceAll(string(url), "href=\"", website.Url));
		}
	}

	return sites
}
