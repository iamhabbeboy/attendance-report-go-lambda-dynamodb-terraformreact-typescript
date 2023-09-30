package main

import (
	"encoding/json"
	"fmt"
)

type ReportRepository struct {
	report Report
}

func NewReportRepository() *ReportRepository {
	return &ReportRepository{
		report: *NewReport(),
	}
}

func (re *ReportRepository) Create(report Report) (interface{}, error) {
	resp := re.report.SetFirstTimer(report.FirstTimer).
		// SetMainAuditoriumAdult(report.Categories).
		SetMainAuditoriumBaby(report.Categories).
		SetCategory("traffic", report.Categories).
		// SetCategory("traffic", )
		Build()
	j, _ := json.Marshal(resp)
	fmt.Println(string(j))
	return nil, nil
	// return DB.Store(*resp)
}
