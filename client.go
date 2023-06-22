package leostream

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// HostURL - Default Leostream URL
const HostURL string = "http://localhost"

// Client struct
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
	Auth       AuthStruct
}

// AuthStruct struct	
type AuthStruct struct {
	Username string `json:"user_login"`
	Password string `json:"password"`
}

// AuthResponse struct
type AuthResponse struct {
	UserID   int    `json:"user_id`
	Username string `json:"username`
	Token    string `json:"sid"`
}

// NewClient - Create new Leostream client
func NewClient(host, username, password *string) (*Client, error) {

	tr := &http.Transport{
		Proxy: http.ProxyFromEnvironment,TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second, Transport: tr},
		// Default Leostream URL
		HostURL: HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	// If username or password not provided, return empty client
	if username == nil || password == nil {
		return &c, nil
	}

	c.Auth = AuthStruct{
		Username: *username,
		Password: *password,
	}

	ar, err := c.SignIn()
	if err != nil {
		return nil, err
	}

	c.Token = ar.Token

	return &c, nil
}

// perform a http request with an auth token
func (c *Client) doRequest(req *http.Request, authToken *string) ([]byte, error) {
	token := c.Token

	if authToken != nil {
		token = *authToken
	}

	req.Header.Set("Content-Type", "application/json ")
	req.Header.Set("Authorization", "Bearer "+token)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// If status code is not (200 or 201), return error
	statusOK := res.StatusCode >= 200 && res.StatusCode < 300
	if !statusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
