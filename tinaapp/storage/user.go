package storage

import (
	"database/sql"
	"fmt"
)

type UserStorageService struct {
	DB *sql.DB
}

func (u UserStorageService) CreateUser() {
	fmt.Println("Create User MYSQL")
}

func (u UserStorageService) GetUser() {
	fmt.Println("Get User MYSQL")
}
