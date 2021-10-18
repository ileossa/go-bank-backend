package service

import (
	"testing"
)

func TestAdd(t *testing.T) {
	feed := New()
	//feed.Add(UserSchema{
	//	Name:      "Montcourt",
	//	Firstname: "Damien",
	//	Address:   "12 rue belle-rue",
	//	Birthday:  time.Now(),
	//})
	feed.Add(CustomerSchema{})

	if len(feed.Users) == 0 {
		t.Error("Users was not added")
	}
}

func TestGetAll(t *testing.T) {

	feed := New()
	feed.Add(CustomerSchema{})
	res := feed.GetCustomers()
	if len(res) != 1 {
		t.Error("Users was not added")
	}
}
