package models

type Tina struct {
	Id         int64   `json:"id"`
	IdUser     int64   `json:"id_user"`
	MacAddress string  `json:"mac_address"`
	Name       string  `json:"name"`
	Height     float64 `json:"height"`
}
