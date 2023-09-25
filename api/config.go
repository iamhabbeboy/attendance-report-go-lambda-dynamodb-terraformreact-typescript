package main

const (
	DefaultDBName         = "attendance"
	DefaultDatabaseDriver = "mongodb"
	DefaultTableName      = "reports"
)

// var AttendanceKeys []string

var DB Databaser

// func JsonDecode() (interface{}, error){
// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&requestBody); err != nil {
// 		String()
// 		return nil,
// 	}
// }

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
