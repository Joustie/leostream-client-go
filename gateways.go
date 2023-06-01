package leostream

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetGateways - Returns list of coffees (no auth required)
func (c *Client) GetGateways() ([]Gateway, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/gateways", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	gateways := []Gateway{}
	err = json.Unmarshal(body, &gateways)
	if err != nil {
		return nil, err
	}

	return gateways, nil
}

// GetGateway - Returns specific gateway (no auth required)
func (c *Client) GetGateway(gatewayID string) ([]Gateway, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/gateways/%s", c.HostURL, gatewayID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	gateways := []Gateway{}
	err = json.Unmarshal(body, &gateways)
	if err != nil {
		return nil, err
	}

	return gateways, nil
}

// CreateGateway - Create new gateway
func (c *Client) CreateGateway(gateway Gateway, authToken *string) (*Gateway, error) {
	rb, err := json.Marshal(gateway)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/gateways", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	newGateway := Gateway{}
	err = json.Unmarshal(body, &newGateway)
	if err != nil {
		return nil, err
	}

	return &newGateway, nil
}
