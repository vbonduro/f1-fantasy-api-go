// f1fantasy Provides APIs to access data from the f1 fantasy website.
package f1fantasy

import (
	"encoding/json"

	"github.com/vbonduro/f1-fantasy-api-go/internal/client"
)

type Api struct {
}

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
