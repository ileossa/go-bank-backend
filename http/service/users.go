package service

import (
	"errors"
	"time"
)

type UserSchema struct {
	Email          string    `json:"email" gorm:"primary_key"`
	Name           string    `json:"name"`
	Firstname      string    `json:"firstname"`
	Address        string    `json:"address"`
	Birthday       time.Time `json:"birthday"`
	Active         bool      `json:"active"`
	Card           []CardSchema
	MoneyOnAccount int64 `json:"money_on_account"`
}

type User interface {
	InitUser() *UserSchema
	ListCards() []CardSchema
	SetStatusAccount(b bool) bool
	newCard(card CardSchema) []CardSchema
	CardExist(number uint) bool
	Minus(amount int64) (bool, int64)
	HaveEnough(amount int64) bool
}

func InitUser() *UserSchema {
	return &UserSchema{
		Card: []CardSchema{},
	}
}

func (r *UserSchema) ListCards() []CardSchema {
	return r.Card
}

func (r *UserSchema) GetCard(number uint) (CardSchema, error) {
	for _, card := range r.Card {
		if card.Number == number {
			return card, nil
		}
	}
	return CardSchema{}, errors.New("Card not find")
}

func (r *UserSchema) SetStatusAccount(b bool) bool {
	r.Active = b
	return r.Active
}

func (us *UserSchema) newCard(card CardSchema) []CardSchema {
	us.Card = append(us.Card, card)
	return us.Card
}

func (r *UserSchema) CardExist(number uint) bool {
	if _, err := r.GetCard(number); nil == err {
		return true
	}
	return false
}

func (r *UserSchema) HaveEnough(amount int64) bool {
	if amount > 0 {
		return r.MoneyOnAccount-amount >= 0
	} else {
		return r.MoneyOnAccount-(amount*-1) >= 0
	}
}

func (r *UserSchema) Minus(amount int64) (bool, int64) {
	if false == r.HaveEnough(amount) {
		return false, r.MoneyOnAccount
	}
	r.MoneyOnAccount = r.MoneyOnAccount - amount
	return true, r.MoneyOnAccount
}
