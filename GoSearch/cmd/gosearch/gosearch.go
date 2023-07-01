package main

import (
	"GO_thinknetica/GoSearch/pkg/crawler"
	"GO_thinknetica/GoSearch/pkg/crawler/spider"
	"fmt"
)

func main() {
	// create slice for data
	var data []crawler.Document
	sts := [2]string{
		"https://go.dev/",
		"https://www.programiz.com/golang/",
	}
	s := spider.New()

	// scan sites
	for i := range sts {
		result, err := s.Scan(sts[i], 1)

		if err != nil {
			fmt.Println(err)
			continue
		}
		// result append to data
		data = append(data, result...)
		//print section
		//fmt.Println(result)
	}
	// write data to a file
	fmt.Println(data)
}
