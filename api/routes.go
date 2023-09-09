package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	report ReportHandler
}

func NewRoutes() *Router {
	return &Router{
		report: *NewReportHandler(),
	}
}

func (re *Router) Endpoints() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		httpResponse := HttpResponse{}
		httpResponse.JSON(w, map[string]string{"status": "All is well here"}, 200)
	})
	r.HandleFunc("/api/reports", re.report.Index)
	r.HandleFunc("/api/export", re.report.Export).Methods("GET")
	return r
}
