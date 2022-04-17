// Package f1fantasy provides APIs to access data from the f1 fantasy website.
package f1fantasy

import (
	"encoding/json"
	"time"

	"github.com/vbonduro/f1-fantasy-api-go/internal/client"
)

// Api is used to access the public API for the f1 fantasy website.
type Api struct {
}

// AuthenticatedApi is used to access both public and authenticated endpoints from the f1 fantasy website.
type AuthenticatedApi struct {
	Api
	session *client.Session
}

// NewApi creates an Api instance that can be used to access the public f1 fantasy Apis.
func NewApi() *Api {
	return &Api{}
}

// NewAuthenticatedApi creates an authenticated API instance that allows access to APIs requiring user login.
// This API also can be used with public APIs.
func NewAuthenticatedApi(user string, password string) (*AuthenticatedApi, error) {
	session, err := client.Login(user, password)
	if err != nil {
		return nil, err
	}
	return &AuthenticatedApi{session: session}, nil
}

// Expired checks whether or not the session is still valid. If it isn't the client must create a new authenticated API instance.
func (api *AuthenticatedApi) Expired() bool {
	anHourAgo := time.Now().Add(time.Duration(-1) * time.Hour)
	return api.session.Expiry.Before(anHourAgo)
}

func (*Api) getAndDecode(page string, data interface{}) error {
	request := client.Request{
		Endpoint: client.BASE_URL + page,
		Header:   client.MakeHeader(),
	}
	_, body, err := request.Get()
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, data)
	if err != nil {
		return err
	}
	return nil
}

func (api *AuthenticatedApi) getAndDecode(page string, data interface{}) error {
	request := client.Request{
		Endpoint: client.BASE_URL + page,
		Header:   client.MakeHeader(),
	}
	request.Header["Cookie"] = []string{api.session.Cookie}
	_, body, err := request.Get()
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, data)
	if err != nil {
		return err
	}
	return nil
}

func (api *AuthenticatedApi) get(page string) ([]byte, error) {
	request := client.Request{
		Endpoint: client.BASE_URL + page,
		Header:   client.MakeHeader(),
	}
	request.Header["Cookie"] = []string{api.session.Cookie}
	_, body, err := request.Get()
	if err != nil {
		return nil, err
	}
	return body, nil
}
