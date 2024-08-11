// Copyright (c) HopopOps
// SPDX-License-Identifier: MPL-2.0

package cloudbeaver

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	HostURL string = "http://localhost:8978/api/gql"

	operationAuthLogin string = "authLogin"
	queryAuthLogin     string = "\n    query authLogin($provider: ID!, $configuration: ID, $credentials: Object, $linkUser: Boolean, $customIncludeOriginDetails: Boolean!, $forceSessionsLogout: Boolean) {\n  authInfo: authLogin(\n    provider: $provider\n    configuration: $configuration\n    credentials: $credentials\n    linkUser: $linkUser\n    forceSessionsLogout: $forceSessionsLogout\n  ) {\n    redirectLink\n    authId\n    authStatus\n    userTokens {\n      ...AuthToken\n    }\n  }\n}\n    \n    fragment AuthToken on UserAuthToken {\n  authProvider\n  authConfiguration\n  loginTime\n  message\n  origin {\n    ...ObjectOriginInfo\n  }\n}\n    \n    fragment ObjectOriginInfo on ObjectOrigin {\n  type\n  subType\n  displayName\n  icon\n  details @include(if: $customIncludeOriginDetails) {\n    id\n    required\n    displayName\n    description\n    category\n    dataType\n    defaultValue\n    validValues\n    value\n    length\n    features\n    order\n  }\n}\n    "
)

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Auth       AuthGQL
	Cookies    *string
}

type AuthGQL struct {
	Query         string `json:"query"`
	OperationName string `json:"operationName"`
	Variables     Auth   `json:"variables"`
}

type AuthCredentials struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type Auth struct {
	Provider                   string          `json:"provider"`
	Credentials                AuthCredentials `json:"credentials"`
	LinkUser                   bool            `json:"linkUser"`
	CustomIncludeOriginDetails bool            `json:"customIncludeOriginDetails"`
	ForceSessionsLogout        bool            `json:"forceSessionsLogout"`
}

// AuthResponse -
type AuthResponse struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func (c *Client) doRequest(req *http.Request, authCookie *string) ([]byte, string, error) {
	if len(*c.Cookies) > 0 {
		req.Header.Set("Cookie", fmt.Sprintf("cb-session-id=%s", *c.Cookies))
	}

	if authCookie != nil {
		req.Header.Set("Cookie", *authCookie)
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, "", err
	}

	// fmt.Printf("%s\n", body)

	if res.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	cookie := ""
	cookies := res.Cookies()
	for _, c := range cookies {
		if c.Name == "cb-session-id" {
			cookie = c.Value
			break
		}
	}

	return body, cookie, err
}

// SignIn - Get a new token for user
func (c *Client) SignIn() (*string, error) {
	if c.Auth.Variables.Credentials.User == "" || c.Auth.Variables.Credentials.Password == "" {
		return nil, fmt.Errorf("define username and password")
	}
	rb, err := json.Marshal(c.Auth)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.HostURL, strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	_, cookies, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	return &cookies, nil
}

// NewClient -
func NewClient(host, username, password *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default CloudBeaver URL
		HostURL: HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	// If username or password not provided, return empty client
	if username == nil || password == nil {
		return &c, nil
	}

	c.Auth = AuthGQL{
		Query:         queryAuthLogin,
		OperationName: operationAuthLogin,
		Variables: Auth{
			Provider: "local",
			Credentials: AuthCredentials{
				User:     *username,
				Password: strings.ToUpper(fmt.Sprintf("%x", md5.Sum([]byte(*password)))),
			},
			LinkUser:                   false,
			CustomIncludeOriginDetails: true,
			ForceSessionsLogout:        false,
		},
	}

	initCookie := ""
	c.Cookies = &initCookie

	cookies, err := c.SignIn()
	if err != nil {
		return nil, err
	}

	c.Cookies = cookies

	return &c, nil
}
