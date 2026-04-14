package model

type TrxRequest struct {
	Amount  uint `json:"amount"`
	Success bool `json:"success"`
}
