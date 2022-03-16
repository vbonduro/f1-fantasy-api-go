// Client implementation for interacting with the F1 Fantasy website.
package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"
)

// Base URL for requests to F1 Fantasy Server.
const BASE_URL = "https://fantasy-api.formula1.com/f1/2022"

// Encodes pertinent information for an authenticated session.
type Session struct {
	Token  string
	UserId int
	Cookie string
	Expiry time.Time
}

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

func makeHeader() map[string][]string {
	return map[string][]string{
		"authority":          {"api.formula1.com"},
		"pragma":             {"no-cache"},
		"cache-control":      {"no-cache"},
		"sec-ch-ua":          {`" Not A;Brand";v="99", "Chromium";v="98", "Google Chrome";v="98"`},
		"dnt":                {"1"},
		"sec-ch-ua-mobile":   {"?0"},
		"user-agent":         {"RaceControl"},
		"Content-Type":       {"application/json"},
		"accept":             {"application/json, text/javascript, */*; q=0.01"},
		"apikey":             {"fCUCjWrKPu9ylJwRAv8BpGLEgiAuThx7"},
		"sec-ch-ua-platform": {`"macOS"`},
		"origin":             {"https://account.formula1.com"},
		"sec-fetch-site":     {"same-site"},
		"sec-fetch-mode":     {"cors"},
		"sec-fetch-dest":     {"empty"},
		"referer":            {"https://account.formula1.com/"},
		"accept-language":    {"en-US,en;q=0.9"},
	}
}

func generateXF1Cookie(token string) string {
	tokenStr := fmt.Sprintf(`{"data":{"subscriptionToken":"%s"}}`, token)
	return base64.StdEncoding.EncodeToString([]byte(tokenStr))
}

func generateCookie(session Session) (string, error) {
	cookie := map[string]interface{}{
		"formula-1-session": map[string]interface{}{"authenticated": map[string]interface{}{}},
		"notice_behavior":   "implied,eu",
		// Measures the user's bandwidth and throttles webchat functionality
		"talkative_qos_bandwidth": 5.33,
		// EU consent string, contains generated time
		"euconsent-v2":       "CPVNvewPVNvewAvACDENCECgAAAAAAAAAAAAAAAAAAAA.YAAAAAAAAAAA",
		"notice_preferences": "0:",
		// TrustArc consent id
		"TAconsentID": "d3ad6bbe-9dbe-449a-abb9-589173f0b16f",
		//"TAconsentID","",
		"notice_gdpr_prefs":    "0::implied,eu",
		"notice_poptime":       1620322920000,
		"cmapi_gtm_bl":         "ga-ms-ua-ta-asp-bzi-sp-awct-cts-csm-img-flc-fls-mpm-mpr-m6d-tc-tdc",
		"cmapi_cookie_privacy": "permit 1 required",
		"register": map[string]interface{}{
			"event":                   "register",
			"eventCategory":           "account registration",
			"eventAction":             "Marketing Consent",
			"userID":                  0, // Loaded in from "account/subscriber/authenticate/by-password" body.Subscriber.Id
			"userType":                0,
			"subscriptionSource":      "web",
			"countryOfRegisteredUser": "USA",
			"eventLabel":              "Unchecked",
			"actionType":              "success"},
		"login": map[string]interface{}{"event": "login", "componentId": "component_login_page", "actionType": "success"},
		// Login token
		"login-session": map[string]interface{}{
			"data": map[string]interface{}{"subscriptionToken": ""},
		},
		"user-metadata": map[string]interface{}{"subscriptionSource": "", "userRegistrationLevel": "full", "subscribedProduct": "", "subscriptionExpiry": "99/99/9999"},
	}
	cookie["register"].(map[string]interface{})["userID"] = session.UserId
	cookie["login-session"].(map[string]interface{})["data"].(map[string]interface{})["subscriptionToken"] = session.Token

	cookieStr := ""
	for key, value := range cookie {
		switch value.(type) {
		case map[string]interface{}:
			data, err := json.Marshal(value)
			if err != nil {
				return "", err
			}
			cookieStr += url.Values{key: {string(data)}}.Encode() + "; "
		case float64:
			cookieStr += key + "=" + fmt.Sprintf("%.2f", value.(float64)) + "; "
		case int:
			cookieStr += key + "=" + fmt.Sprintf("%d", value.(int)) + "; "
		default:
			cookieStr += key + "=" + value.(string) + "; "
		}
	}
	cookieStr = cookieStr[:len(cookieStr)-2]
	return cookieStr, nil
}

func getCookie(session Session) (string, error) {
	const PAYLOAD = `{"user":{"date_of_birth":null,"email":null,"first_name":null,"last_name":null,"password":null}}`
	request := Request{
		Endpoint: BASE_URL + "/sessions",
		Header:   makeHeader(),
		Payload:  strings.NewReader(PAYLOAD),
	}

	request.Header["X-F1-COOKIE-DATA"] = []string{generateXF1Cookie(session.Token)}

	cookie, err := generateCookie(session)
	if err != nil {
		return "", err
	}
	request.Header["Cookie"] = []string{cookie}

	header, _, err := request.Post()
	if err != nil {
		return "", err
	}

	responseCookie := strings.Split(strings.Split(header["Set-Cookie"][0], ";")[0], "=")[1]
	return responseCookie, nil
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
		Header:   makeHeader(),
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

	const DEFAULT_EXPIRY_HOURS = 4
	session := Session{
		Token:  response.Subscription.Token,
		UserId: response.Subscriber.Id,
		Expiry: time.Now().Local().Add(time.Hour * time.Duration(DEFAULT_EXPIRY_HOURS)),
	}

	session.Cookie, err = getCookie(session)
	if err != nil {
		return nil, err
	}

	return &session, nil
}
