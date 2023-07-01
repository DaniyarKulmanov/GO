package main

import (
	"GO_thinknetica/GoSearch/pkg/crawler"
	"GO_thinknetica/GoSearch/pkg/crawler/spider"
	"fmt"
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
}
