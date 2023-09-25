package main

import (
	"encoding/json"
	"io"
	"log"
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

func (re *ReportHandler) create(w http.ResponseWriter, r *http.Request) error {
	var requestData Report
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print(err.Error())
		String(w, "Error reading request body", http.StatusBadRequest)
		return nil
	}
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		log.Print(err.Error())
		JSON(w, map[string]string{"message": "Error reading request body"}, http.StatusBadRequest)
		return nil
	}
	res, _ := re.repo.Create(requestData)
	JSON(w, map[string]interface{}{"data": res}, http.StatusBadRequest)
	return nil
}
