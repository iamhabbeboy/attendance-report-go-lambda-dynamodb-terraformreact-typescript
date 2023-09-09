package main

// func Seeder() {
// 	services := []Report{{
// 		ID:      uuid.New().String(),
// 		Service: 1,
// 		Category: []Category{{
// 			Name:  "children",
// 			Value: "20",
// 		}, {
// 			Name:  "main-auditorium",
// 			Value: "100",
// 		}, {
// 			Name:  "extension",
// 			Value: "20",
// 		}},
// 	}, {
// 		ID:      uuid.New().String(),
// 		Service: 2,
// 		Category: []Category{{
// 			Name:  "children",
// 			Value: "60",
// 		}, {
// 			Name:  "main-auditorium",
// 			Value: "200",
// 		}, {
// 			Name:  "extension",
// 			Value: "20",
// 		}},
// 	}}
// 	database := NewDatabaseDriver(DefaultDatabaseDriver)
// 	for _, val := range services {
// 		err := database.Store(val)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println("Migrated successfuly")
// 	}

// }
