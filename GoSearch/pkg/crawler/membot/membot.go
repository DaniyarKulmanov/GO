package membot

import (
	"GO_thinknetica/GoSearch/pkg/crawler"
)

// Service - имитация служба поискового робота.
type Service struct{}

// New - констрктор имитации службы поискового робота.
func New() *Service {
	s := Service{}
	return &s
}

// Scan возвращает заранее подготовленный набор данных
func (s *Service) Scan(url string, depth int) ([]crawler.Document, error) {

	data := []crawler.Document{
		{
			ID:    0,
			URL:   "https://go.dev/",
			Title: "GO_dev",
		},
		{
			ID:    1,
			URL:   "https://www.programiz.com/golang",
			Title: "Programiz",
		},
	}

	return data, nil
}
