package models

type Week struct {
	Id		int64	`json:"id"`
	IdUser	int64	`json:"id_user"`
	InitDay	int64	`json:"init_day"`
	EndDay	int64	`json:"end_day"`
}
