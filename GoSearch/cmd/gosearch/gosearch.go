package main

import (
	"GO_thinknetica/GoSearch/pkg/crawler"
	"GO_thinknetica/GoSearch/pkg/crawler/spider"
	"flag"
	"fmt"
	"strings"
)

func main() {
	var data []crawler.Document
	sts := [2]string{
		"https://go.dev/",
		"https://www.programiz.com/golang/",
	}
	s := spider.New()

	for i := range sts {
		result, err := s.Scan(sts[i], 1)

		if err != nil {
			fmt.Println(err)
			continue
		}
		data = append(data, result...)
	}

	fmt.Println(data)

	// read flag
	var search string

	flag.StringVar(&search, "s", "documents", "search links")
	flag.Parse()
	fmt.Println(search)

	// filter content based on search
	var dataFilter []crawler.Document
	fmt.Println(strings.Contains(data[1].URL, search))

	for i := range data {
		if strings.Contains(data[i].URL, search) {
			dataFilter = append(dataFilter, data[i])
		}
	}
	fmt.Println(dataFilter)
}
