package main

import (
	"GO_thinknetica/03-algorithms/GoSearch/pkg/crawler"
	"GO_thinknetica/03-algorithms/GoSearch/pkg/crawler/spider"
	"GO_thinknetica/03-algorithms/GoSearch/pkg/index"
	"GO_thinknetica/03-algorithms/GoSearch/pkg/search"
	"flag"
	"fmt"
)

func main() {
	var word string
	var depth int
	var docs []crawler.Document
	sts := [2]string{
		"https://go.dev/",
		"https://www.programiz.com/golang/",
	}
	s := spider.New()
	flag.StringVar(&word, "s", "", "word links")
	flag.IntVar(&depth, "d", 1, "depth size")
	flag.Parse()

	for i := range sts {
		result, err := s.Scan(sts[i], depth)
		if err != nil {
			fmt.Println(err)
			continue
		}
		for i := range result {
			index.Create(result[i])
		}
	}
	if word != "" {
		docs = search.Binary(index.MapIndex[word], index.Documents)
		index.Documents = docs
	}
	fmt.Println(index.Documents)
}
