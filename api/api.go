package api

import (
	"github.com/aldaprojects/tinaap/tinaapp/web"
	"log"
	"net/http"

	"github.com/aldaprojects/tinaap/internal/storage"
)

func Start()  {
	db := storage.ConnectToDB()
	defer db.Close()

	r := routes(web.NewUserHTTPService(db))

	log.Fatal(http.ListenAndServe(":3000", r))

}