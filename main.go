package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

/*

	goPsdmpMe - it's a simple tool that uses https://psbdmp.ws API
	to search for pastebins' pastes by keyword.
	No API key is needed.

*/

var (
	keyword *string
)

type Response struct {
	Search string `json:"search"'`
	Count  int    `json:"count"`
	Data   []data `json:"data"`
}

type data struct {
	ID     string  `json:"id"`
	Tags   string  `json:"tags"`
	Length float64 `json:"length"`
	Time   string  `json:"time"`
	Text   string  `json:"text"`
}

func parseFlags() bool {

	keyword = flag.String("k", "", "Keyword to use in search query")
	flag.Parse()

	if *keyword == "" {
		fmt.Printf("No keyword is specified\n")
		flag.PrintDefaults()
		return false
	}

	return true
}

func main() {

	if !parseFlags() {
		os.Exit(1)
	}

	var response = new(Response)
	searchUrl := "https://psbdmp.ws/api/v3/search/" + *keyword
	mainUrl := "https://pastebin.com/"

	// Set up http client and make request.
	httpClient := &http.Client{}
	req, _ := http.NewRequest("GET", searchUrl, nil)
	req.Header.Set("Content-Type", "text/html; charset=UTF-8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:92.0) Gecko/20100101 Firefox/92.0")
	time.Sleep(time.Second) // Sleep for the sake of the website.
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	// Read response body and unmarshal it into Response structure
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(body, &response)

	fmt.Printf("[+] Found %d pastes!\n", response.Count)

	for _, r := range response.Data {
		fmt.Printf(mainUrl + r.ID + "\n")
	}

}
