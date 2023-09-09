package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ReportHandler struct {
	repo         ReportRepository
	httpResponse HttpResponse
}

func NewReportHandler() *ReportHandler {
	return &ReportHandler{
		repo: *NewReportRepository(),
	}
}

func (re *ReportHandler) Index(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	switch r.Method {
	case http.MethodGet:
		re.findOneOrAll(nil)
	case http.MethodPost:
		re.create(w, r)
	default:
	}
	// fmt.Println(result)
	// JSON(w, result, 200)
}

func (re *ReportHandler) Export(w http.ResponseWriter, r *http.Request) {
}

func (re *ReportHandler) findOneOrAll(payload interface{}) error {
	return nil
}

func (re *ReportHandler) create(w http.ResponseWriter, r *http.Request) (map[string]string, error) {
	var requestBody Report
	httpResponse := HttpResponse{}
	fmt.Println("===")
	fmt.Println(r.Body)
	fmt.Println("===x===")
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestBody); err != nil {
		httpResponse.String(w, "Invalid Request", http.StatusBadRequest)
		return nil, nil
	}
	return map[string]string{"status": "ok"}, nil
}
