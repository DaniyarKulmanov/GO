package main

import (
	"GO_thinknetica/GoSearch/pkg/crawler"
	"GO_thinknetica/GoSearch/pkg/crawler/spider"
	"flag"
	"fmt"
	"strings"
)

func main() {
	var fullData []crawler.Document
	var searchData []crawler.Document
	var search string
	sts := [2]string{
		"https://go.dev/",
		"https://www.programiz.com/golang/",
	}
	s := spider.New()
	flag.StringVar(&search, "s", "", "search links")
	flag.Parse()

	for i := range sts {
		result, err := s.Scan(sts[i], 1)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fullData = append(fullData, result...)
	}
	if search != "" {
		searchData = append(searchData, searchString(fullData, search)...)
	}
	fmt.Println(fullData)
	fmt.Println(searchData)
}

func searchString(document []crawler.Document, search string) []crawler.Document {
	var searchData []crawler.Document
	for i := range document {
		if strings.Contains(document[i].Title, search) {
			searchData = append(searchData, document[i])
		}
	}
	return searchData
}
