package client

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"
)

// Session encodes pertinent information for an authenticated session.
type Session struct {
	Cookie string
	Expiry time.Time
}

// NewSession establishes a playOn session and stores the resultant cookie for access to authorized APIs.
func NewSession(userId int, token string) (*Session, error) {
	const PAYLOAD = `{"user":{"date_of_birth":null,"email":null,"first_name":null,"last_name":null,"password":null}}`
	request := Request{
		Endpoint: BASE_URL + "/sessions",
		Header:   MakeHeader(),
		Payload:  strings.NewReader(PAYLOAD),
	}

	request.Header["X-F1-COOKIE-DATA"] = []string{generateXF1Cookie(token)}

	cookie, err := generateCookie(userId, token)
	if err != nil {
		return nil, err
	}
	request.Header["Cookie"] = []string{cookie}

	header, _, err := request.Post()
	if err != nil {
		return nil, err
	}

	playOnCookie := strings.Split(header["Set-Cookie"][0], ";")[0]
	cookie += "; " + playOnCookie

	const DEFAULT_EXPIRY_HOURS = 4
	return &Session{
		Cookie: cookie,
		Expiry: time.Now().Local().Add(time.Hour * time.Duration(DEFAULT_EXPIRY_HOURS)),
	}, nil
}

func generateXF1Cookie(token string) string {
	tokenStr := fmt.Sprintf(`{"data":{"subscriptionToken":"%s"}}`, token)
	return base64.StdEncoding.EncodeToString([]byte(tokenStr))
}

func generateCookie(userId int, token string) (string, error) {
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
	cookie["register"].(map[string]interface{})["userID"] = userId
	cookie["login-session"].(map[string]interface{})["data"].(map[string]interface{})["subscriptionToken"] = token

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
