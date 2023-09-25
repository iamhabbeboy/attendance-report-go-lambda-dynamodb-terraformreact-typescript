package main

import "fmt"

type ReportRepository struct {
	report Report
}

func NewReportRepository() *ReportRepository {
	return &ReportRepository{
		report: *NewReport(),
	}
}

func (re *ReportRepository) Create(report Report) (*Report, error) {
	fmt.Println(report.Categories)
	resp := re.report.SetFirstTimer(report.FirstTimer).
		SetMainAuditoriumAdult(report.Categories).
		SetMainAuditoriumBaby(report.Categories).
		Build()

	// r.SetCategory("traffic", "10")
	// r.SetCategory("info_desk", "5")
	// .Build()
	// db := DB.Store(*r)
	return resp, nil
}
