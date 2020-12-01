package gateway

import (
	"database/sql"
	"github.com/aldaprojects/tinaap/tinaapp/storage"
)

type UserGateway interface {
	CreateUser()
	GetUser()
}

func NewUserGateway(db *sql.DB) UserGateway {
	return &storage.UserStorageService{DB: db}
}