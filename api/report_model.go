package main

import (
	"log"
	"strconv"
)

type Category struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Report struct {
	ID         string     `json:"id,omitempty" dynamo:"id" bson:"id,omitempty"`
	Service    string     ` json:"service,omitempty" dynamo:"service" bson:"service,omitempty"`
	Categories []Category `json:"categories,omitempty" dynamo:"categories" bson:"categories,omitempty"`
	Offering   int64      `json:"offering,omitempty" dynamo:"offering" bson:"offering,omitempty"`
	FirstTimer int        `json:"first_timer,omitempty" dynamo:"first_timer" bson:"first_timer,omitempty"`
}

func NewReport() *Report {
	return &Report{}
}

func (r *Report) SetService(service string) *Report {
	r.Service = service
	return r
}

func (r *Report) SetFirstTimer(ftimer int) *Report {
	r.FirstTimer = ftimer
	return r
}

func (r *Report) SetCategories(categories []Category) *Report {
	r.Categories = categories
	return r
}

func (r *Report) SetMainAuditoriumAdult(categories []Category) *Report {
	for _, val := range categories {
		if val.Name == "main_auditorium_adult" {
			r.Categories = append(r.Categories, categories...)
		}
	}

	return r
}

func (r *Report) SetMainAuditoriumBaby(categories []Category) *Report {
	for _, val := range categories {
		if val.Name == "main_auditorium_baby" {
			r.Categories = append(r.Categories, categories...)
		}
	}

	return r
}

func (r *Report) SetCategory(key string, val string) *Report {
	for _, val := range r.Categories {
		if val.Name == key {
			log.Fatal("category already exist")
		}
	}
	r.Categories = append(r.Categories, Category{
		Name:  key,
		Value: val,
	})

	return r
}

func (r *Report) GetCategory(key string) int {
	if len(r.Categories) == 0 {
		return 0
	}
	for _, val := range r.Categories {
		if val.Name == key {
			v, _ := strconv.Atoi(val.Value)
			return v
		}
	}
	return 0
}

func (r *Report) GetMainAuditoriumBaby() int {
	if len(r.Categories) == 0 {
		return 0
	}
	for _, val := range r.Categories {
		if val.Name == "main_auditorium_baby" {
			v, _ := strconv.Atoi(val.Value)
			return v
		}
	}
	return 0
}

func (r *Report) GetMainAuditoriumAdult() int {
	if len(r.Categories) == 0 {
		return 0
	}
	for _, val := range r.Categories {
		if val.Name == "main_auditorium_adult" {
			v, _ := strconv.Atoi(val.Value)
			return v
		}
	}
	return 0
}

func (r *Report) GetFirstTimer() int {
	return r.FirstTimer
}

func (r *Report) SetOffering(offering int64) *Report {
	r.Offering = offering
	return r
}

func (r *Report) GetOffering() int64 {
	return r.Offering
}

func (r *Report) Build() *Report {
	return r
}
