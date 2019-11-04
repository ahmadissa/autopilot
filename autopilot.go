//Package Autopilot wrapper for Autopilo API
package Autopilot

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

//Resquest  autopilot request object
type Resquest struct {
	Contact Contact `json:"contact"`
}

//Respond  autopilot respond object
type Respond struct {
	Contact    Contact `json:"contact"`
	HTTPStatus int
	Error      ATError
}

//Contact - autopilot contact object
type Contact struct {
	ContactID  string    `json:"contact_id"`
	FirstName  string    `json:"FirstName"`
	LastName   string    `json:"LastName"`
	Type       string    `json:"type"`
	Email      string    `json:"Email"`
	Phone      string    `json:"Phone"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	LeadSource string    `json:"LeadSource"`
	Status     string    `json:"Status"`
	Company    string    `json:"Company"`
	Lists      []string  `json:"lists"`
}

//ATError - autopilot error resp
type ATError struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

var apiKey = ""
var baseURL = "https://api2.autopilothq.com/v1/contact"

//do send data to auto pilot and return a contact object or error
func do(method string, url string, data interface{}) (res Respond, err error) {

	res = Respond{}
	client := &http.Client{}
	body, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return
	}
	req.Header.Add("autopilotapikey", apiKey)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		return
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode == 200 {
		contact := Contact{}
		err = json.Unmarshal(respBody, &contact)
		if err != nil {
			return
		}
		res.HTTPStatus = resp.StatusCode
		res.Contact = contact
		return
	}
	errResp := ATError{}
	err = json.Unmarshal(respBody, &errResp)
	if err != nil {
		return
	}
	err = errors.New(errResp.Error + ":" + errResp.Message)
	res.HTTPStatus = resp.StatusCode
	res.Error = errResp
	return
}

//Put add contact or update if exists
func Put(contact Contact) (res Respond, err error) {
	req := Resquest{Contact: contact}
	return do("POST", baseURL, req)
}

//Get contact by email or ID
func Get(emailOrID string) (res Respond, err error) {
	url := baseURL + "/" + emailOrID
	return do("GET", url, nil)
}

//Init set Autopilot api key
func Init(key string) {
	apiKey = key
}
