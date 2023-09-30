package main

import (
	"log"
	"net/http"
)

type ReportHandler struct {
	repo ReportRepository
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
	data, err := Bind(r, requestData)
	if err != nil {
		log.Print(err.Error())
		String(w, "Error reading request body", http.StatusBadRequest)
		return nil
	}
	report := data.(Report)
	// var err1 []string
	// if len(Validate(report)) > 0 {
	// 	for _, e := range Validate(report) {
	// 		err1 = append(err1, e.Kind().String())
	// 	}
	// }
	// fmt.Println(err1)

	res, _ := re.repo.Create(report)
	JSON(w, map[string]interface{}{"data": res}, http.StatusBadRequest)
	return nil
}
