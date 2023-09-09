package main

type Category struct {
	Name  string
	Value string
}

type Report struct {
	ID      string `dynamo:"id" json:"id"`
	Service int    `dynamo:"service" json:"service"`
	// Category []Category `dynamo:"category" json:"category"`
	Tithe    int64 `dynamo:"tithe" json:"tithe"`
	Offering int64 `dynamo:"offering" json:"offering"`
}

type Reporter struct{}

func NewReport() *Reporter {
	return &Reporter{}
}
