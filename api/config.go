package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-playground/validator"
)

const (
	DefaultDBName         = "attendance"
	DefaultDatabaseDriver = "mongodb"
	DefaultTableName      = "reports"
)

// var AttendanceKeys []string

var DB Databaser

func Bind[T interface{}](r *http.Request, requestBody T) (interface{}, error) {
	// var requestBody interface{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	if json.Unmarshal(body, &requestBody) != nil {
		return nil, err
	}
	return requestBody, nil
}

func Validate[T interface{}](d T, fields ...string) validator.ValidationErrors {
	v := validator.New()
	// if len(fields) > 0 {
	// _ = v.RegisterValidation(fields[0], func(fl validator.FieldLevel) bool {
	// 	return len(fl.Field().String()) > 6
	// })
	// }
	err := v.Struct(d)
	return err.(validator.ValidationErrors)
}

func HasDefaultKey(key string) bool {
	attendanceKeys := []string{
		"main-auditorium",
		"extension",
		"overflow",
	}
	hasValue := false
	for _, val := range attendanceKeys {
		if key == val {
			hasValue = true
		}
	}
	return hasValue
}
