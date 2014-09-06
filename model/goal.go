package model

type Goal struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Tallies int64  `json:"tallies"`
}
