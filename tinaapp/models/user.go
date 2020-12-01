package models

type User struct {
	Id 			int64	`json:"id"`
	Username 	string	`json:"username"`
	Password 	string	`json:"password"`
}

type CreateUser struct {
	Username 	string	`json:"username"`
	Password 	string	`json:"password"`
}