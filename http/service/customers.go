package service

import (
	"errors"
	"time"
)

type OwnModel struct {
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-";sql:"index"`
}

type CustomerSchema struct {
	//OwnModel `swaggerignore:"true"`
	Users []UserSchema
}

type Customer interface {
	All() *CustomerSchema
	NewUser(user UserSchema) []UserSchema
	GetUsers() []UserSchema
	GetUser(email string) (UserSchema, error)
	FindCards(email string) ([]CardSchema, error)
}

func InitCustomer() *CustomerSchema {
	return &CustomerSchema{
		Users: []UserSchema{},
	}
}

func (r *CustomerSchema) All() *CustomerSchema {
	return r
}

func (r *CustomerSchema) NewUser(user UserSchema) []UserSchema {
	r.Users = append(r.Users, user)
	return r.Users
}

func (r *CustomerSchema) GetUsers() []UserSchema {
	return r.Users
}

func (r *CustomerSchema) GetUser(email string) (UserSchema, error) {
	for _, schema := range r.Users {
		if schema.Email == email {
			return schema, nil
		}
	}
	return UserSchema{}, errors.New("User with email(" + email + ") not found")
}

func (r *CustomerSchema) FindCards(email string) ([]CardSchema, error) {
	for _, user := range r.Users {
		if email == user.Email {
			return user.ListCards(), nil
		}
	}
	return nil, errors.New("Email (" + email + ") not found")
}
