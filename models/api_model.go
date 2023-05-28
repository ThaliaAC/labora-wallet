package models

import "time"

type Api_Request_To_Truora struct {
	National_id    string `json:"national_id"`
	Country        string `json:"country"`
	Type           string `json:"type"`
	UserAuthorized bool   `json:"userAuthorized"`
}

type Person struct {
	National_id string `json:"national_id"`
	Country     string `json:"country"`
}

type TruoraPostResponse struct {
	Check struct {
		CheckID string `json:"check_id"`
	} `json:"check"`
}

type TruoraGetResponse struct {
	Check struct {
		CheckID      string    `json:"check_id"`
		Country      string    `json:"country"`
		CreationDate time.Time `json:"creation_date"`
		NameScore    int       `json:"name_score"`
		IDScore      int       `json:"id_score"`
		Score        int       `json:"score"`
	}
}
