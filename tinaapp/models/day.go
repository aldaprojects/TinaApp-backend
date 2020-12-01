package models

type Day struct {
	Id			int64	`json:"id"`
	IdWeek 		int64	`json:"id_week"`
	Day			int64	`json:"day"`
	Consumption	int64	`json:"consumption"`
}
