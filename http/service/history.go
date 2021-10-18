package service

import "time"

type HistorySchema struct {
	Date     time.Time `json:"date"`
	Amount   uint      `json:"amount"`
	Category string    `json:"category"`
}

type History interface {
	InitHistory() *HistorySchema
}

func InitHistory() *HistorySchema {
	return &HistorySchema{}
}
