package Autopilot

import (
	"os"
	"testing"
)

//TestInit configuration and init autopilot module
func TestInit(t *testing.T) {
	autoPilotAPIKey := os.Getenv("autoPilotAPIKey")
	if autoPilotAPIKey == "" {
		t.Error("Environment variable 'autoPilotAPIKey' was not set")
	}
	Init(autoPilotAPIKey)
}

func TestAddContact(t *testing.T) {
	contact := Contact{
		Email:     "issa.ahmd@gmail.com",
		FirstName: "Ahmad",
		LastName:  "Issa",
	}
	res, err := Put(contact)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if res.Contact.ContactID == "" {
		t.Error("ContactID is empty")
	}
}

func TestGetContact(t *testing.T) {

	res, err := Get("issa.ahmd@gmail.com")
	if err != nil {
		t.Error(err.Error())
		return
	}
	if res.Contact.ContactID == "" {
		t.Error("ContactID is empty")
		return
	}
	if res.Contact.FirstName != "Ahmad" {
		t.Error("FirstName didnt match")
		return
	}
	if res.Contact.LastName != "Issa" {
		t.Error("LastName didnt match")
		return
	}
}

func TestNotFound(t *testing.T) {

	_, err := Get("not_found@gmail.com")
	if err == nil {
		t.Error("contact should not be found")
		return
	}
}

func TestUpdate(t *testing.T) {

	contact := Contact{
		Email:     "issa.ahmd@gmail.com",
		FirstName: "Ahmad",
		LastName:  "We Have Changed your name",
	}
	res, err := Put(contact)
	if err != nil {
		t.Error(err.Error())
		return
	}
	res, err = Get("issa.ahmd@gmail.com")
	if err != nil {
		t.Error(err.Error())
		return
	}
	if res.Contact.LastName != "We Have Changed your name" {
		t.Error("LastName should be 'We Have Changed your name'")
	}

}
