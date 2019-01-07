package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/Cjiq/ubd/data"
	"github.com/fatih/color"
)

var apiURL = "https://api.urbandictionary.com/v0/define?term="

func main() {
	if len(os.Args) <= 1 {
		showArgumentErr()
		os.Exit(1)
	}
	searchTerm := os.Args[1]
	if searchTerm == "" {
		showArgumentErr()
		os.Exit(1)
	}
	resp, err := http.Get(apiURL + url.QueryEscape(searchTerm))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	var result data.Result
	err = json.Unmarshal(res, &result)
	if err != nil {
		fmt.Printf("Failed to parse json: %s", err)
	}
	d := color.New(color.FgWhite, color.Bold)
	d.Printf("Term: %s\n\n", searchTerm)
	for i, def := range result.Definitions {
		trim := strings.Replace(def.Text, "^M", "", -1)
		d.Printf("Definition #%d: ", (i + 1))
		fmt.Printf("%s\n\n", trim)
	}
}

func showArgumentErr() {
	fmt.Println("Error: You have to enter a search term.")
	fmt.Println("Usage: ubd <search-term>")
	fmt.Printf("Example: ubd mansplain\n\n")
}
