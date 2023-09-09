package main

type Category struct {
	Name  string
	Value string
}

const (
	FirstService = iota
	SecondService
)

type Report struct {
	ID       string `dynamo:"id" json:"id"`
	Service  string `dynamo:"service" json:"service"`
	Category []Category
}

type Reporter struct{}

func NewReport() *Reporter {
	return &Reporter{}
}
