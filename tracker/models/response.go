package models

type Response struct {
	ISSPosition struct {
		Longitude string `json:"longitude"`
		Latitude  string `json:"latitude"`
	} `json:"iss_position"`
	Date    int    `json:"timestamp"`
	Message string `json:"message"`
}
