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

	// read flag
	var search string

	flag.StringVar(&search, "s", "", "search links")
	flag.Parse()

	// filter content based on search
	var dataFilter []crawler.Document
	for i := range data {
		if strings.Contains(data[i].URL, search) {
			dataFilter = append(dataFilter, data[i])
		}
	}

	if search != "" {
		fmt.Println(dataFilter)
	} else {
		fmt.Println(data)
	}

}
