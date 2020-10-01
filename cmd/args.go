package cmd

import (
	"flag"
)

type (
	Args struct {
		Keyword    string `json:"keyword"`
		PageNumber int    `json:"page_number"`
	}
)

func Parse() (args Args) {
	flag.StringVar(&args.Keyword, "keyword", "", "Keyword that is used to search contents")
	flag.IntVar(&args.PageNumber, "page-number", 1, "Page number of search result that will be accessed")
	flag.Parse()

	return
}
