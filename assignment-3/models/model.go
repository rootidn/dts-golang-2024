package models

type Data struct {
	Status Status `json:"status"`
}

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}
