package main

import "time"

type DBModel struct {
	ID                   int
	CreatedAt, UpdatedAt time.Time
}

type User struct {
	DBModel
	name  string
	email string
	s     struct {
		anonStructField string
	}
}

func handleUser() {
	/*u := User{
		DBModel: DBModel{
			ID: 1,
		},
		name: "ivan",
	}*/

}
