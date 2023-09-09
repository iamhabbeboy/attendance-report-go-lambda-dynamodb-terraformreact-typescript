package main

import (
	"fmt"
	"net/http"
)

type ReportHandler struct{}

func NewReportHandler() *ReportHandler {
	return &ReportHandler{}
}

func (re *ReportHandler) Index(w http.ResponseWriter, r *http.Request) {
	var result interface{}
	switch r.Method {
	case GetMethod:
		result = findOneOrAll(nil)
	case PostMethod:
		result = create()
	default:
	}
	fmt.Println(result)
	JSON(w, map[string]string{"status": "Ok"})
}

func (re *ReportHandler) Export(w http.ResponseWriter, r *http.Request) {
}

func findOneOrAll(payload interface{}) error {
	return nil
}

func create() error {
	return nil
}

func Get(ID int) {
}
