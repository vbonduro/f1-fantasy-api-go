// Package client implementation for interacting with the F1 Fantasy website.
package client

import (
	"bytes"
	"encoding/json"
)

// BASE_URL for requests to F1 Fantasy Server.
const BASE_URL = "https://fantasy-api.formula1.com/f1/2022"

type loginRequest struct {
	DistributionChannel string `json:"DistributionChannel"`
	User                string `json:"Login"`
	Password            string `json:"Password"`
}

type subscriberData struct {
	FirstName   string `json:"FirstName"`
	LastName    string `json:"LastName"`
	HomeCountry string `json:"HomeCountry"`
	Id          int    `json:"Id"`
	Email       string `json:"Email"`
	Login       string `json:"Login"`
}

type subscriptionData struct {
	Status string `json:"subscriptionStatus"`
	Token  string `json:"subscriptionToken"`
}

type loginResponse struct {
	SessionId           string           `json:"SessionId"`
	PasswordIsTemporary bool             `json:"PasswordIsTemporary"`
	Subscriber          subscriberData   `json:"Subscriber"`
	Country             string           `json:"Country"`
	Subscription        subscriptionData `json:"data"`
}

// Login will create an authenticated session with the provided user name and password.
func Login(user string, password string) (*Session, error) {
	const URI = "https://api.formula1.com/v2/account/subscriber/authenticate/by-password"
	const DISTRIBUTION_CHANNEL = "d861e38f-05ea-4063-8776-a7e2b6d885a4"

	login := loginRequest{
		DistributionChannel: DISTRIBUTION_CHANNEL,
		User:                user,
		Password:            password,
	}

	payload, err := json.Marshal(&login)
	if err != nil {
		return nil, err
	}

	request := Request{
		Endpoint: URI,
		Payload:  bytes.NewReader(payload),
		Header:   MakeHeader(),
	}

	_, body, err := request.Post()
	if err != nil {
		return nil, err
	}

	var response loginResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	session, err := NewSession(response.Subscriber.Id, response.Subscription.Token)
	if err != nil {
		return nil, err
	}

	return session, nil
}
