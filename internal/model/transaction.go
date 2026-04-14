package model

type TrxRequest struct {
	Amount  uint `json:"amount"`
	Success bool `json:"success"`
}

type TrxRecord struct {
	Id      uint `json:"id"`
	Amount  uint `json:"amount"`
	Success bool `json:"success"`
}
