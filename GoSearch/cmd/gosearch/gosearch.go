package main

import (
	"GO_thinknetica/GoSearch/pkg/crawler/spider"
	"fmt"
	"reflect"
)

func main() {
	// create slice for data
	// scan sites
	// iterate through result and append to data
	// write data to a file

	s := spider.New()
	result, err := s.Scan("https://go.dev/", 1)
	//result, err := s.Scan("https://www.programiz.com/golang/struct", 1)
	fmt.Println(reflect.TypeOf(result[0]))
	//fmt.Println(result)
	fmt.Printf("len=%d cap=%d %v\n", len(result), cap(result), result)
	fmt.Println(err)
}
