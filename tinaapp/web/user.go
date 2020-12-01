package web

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/aldaprojects/tinaap/tinaapp/gateway"
	"github.com/aldaprojects/tinaap/tinaapp/models"
	"net/http"
)

type UserHTTPService struct {
	gtw gateway.UserGateway
}

func NewUserHTTPService(db *sql.DB) *UserHTTPService {
	return &UserHTTPService{
		gateway.NewUserGateway(db),
	}
}

func (uS *UserHTTPService) PostUser(w http.ResponseWriter, r *http.Request) {

	body := r.Body

	defer body.Close()
	var user models.CreateUser
	err := json.NewDecoder(body).Decode(&user)
	if err != nil {
		fmt.Fprint(w, "error")
		return
	}

	fmt.Println(user)

	uS.gtw.CreateUser()
}
func (uS *UserHTTPService) GetUser(w http.ResponseWriter, r *http.Request) {
	uS.gtw.GetUser()
}