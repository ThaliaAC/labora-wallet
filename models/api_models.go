package models

type Api_Request struct {
	NationalID     string `json:"national_id"`
	Country        string `json:"country"`
	EntityType     string `json:"type"`
	UserAuthorized bool   `json:"userAutorized"`
}
