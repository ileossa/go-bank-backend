package service

import "time"

type CardSchema struct {
	Number   uint      `json:"number"`
	validity time.Time `json:"validity"`
	crypto   uint      `json:"crypto"`
	History  []HistorySchema
}

type Card interface {
	InitCard() *CardSchema
	NewPayment(newEntry HistorySchema) []HistorySchema
	ListPayment() []HistorySchema
	getHistory(numberOfElements uint) []HistorySchema
}

func InitCard() *CardSchema {
	return &CardSchema{
		History: []HistorySchema{},
	}
}

func (cs *CardSchema) NewPayment(newEntry HistorySchema) []HistorySchema {
	cs.History = append(cs.History, newEntry)
	return cs.History
}

func (cs *CardSchema) ListPayment() []HistorySchema {
	return cs.History
}

func (cs *CardSchema) getHistory(numberOfElements uint) []HistorySchema {
	out := make([]HistorySchema, numberOfElements)
	copy(out, cs.History)
	return out
}
