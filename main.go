package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"

	"github.com/Cjiq/ubd/data"
	"github.com/fatih/color"
)

type context struct {
	n int
	h bool
	e bool
}

var con *context = &context{}

var apiURL = "https://api.urbandictionary.com/v0/define?term="

func init() {
	flag.IntVar(&con.n, "n", 3, "Number of results. -n 0 returns all.")
	flag.BoolVar(&con.h, "h", false, "Show help")
	flag.BoolVar(&con.h, "help", false, "Show help")
	flag.BoolVar(&con.e, "e", false, "Show example or definition")
	flag.BoolVar(&con.e, "example", false, "Show example or definition")
	flag.Parse()
	if con.h {
		fmt.Println("UrbanDictionary.com v0.0.1")
		showExample()
		fmt.Println("Help: ")
		flag.PrintDefaults()
		os.Exit(0)
	}
}

func main() {
	args := flag.Args()
	if len(args) == 0 {
		showArgumentErr()
		os.Exit(1)
	}
	searchTerm := args[0]
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
		os.Exit(1)
	}

	sort.Slice(result.Definitions, func(i, j int) bool {
		return result.Definitions[i].Rating > result.Definitions[j].Rating
	})

	d := color.New(color.FgWhite, color.Bold)
	d.Printf("Term: %s\n\n", searchTerm)

	for i, def := range result.Definitions {
		d.Printf("Definition #%d: ", (i + 1))
		fmt.Printf("%s\n", trim(def.Text))
		if con.e {
			fmt.Printf("  Example: %s\n", trim(def.Example))
		}
		fmt.Println()
		if (i+1) == con.n && con.n != 0 {
			break
		}
	}
}

func showArgumentErr() {
	fmt.Println("Error: You have to enter a search term.")
	showExample()
}

func showExample() {
	fmt.Println("Usage: [-n..] ubd <search-term>")
	fmt.Printf("Example: ubd bird\n\n")
}

func trim(input string) string {
	return strings.Replace(input, "^M", "", -1)
}
