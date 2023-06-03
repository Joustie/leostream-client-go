package leostream

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// GetGateways - Returns list of gateways (no auth required)
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

// GetGateway - Returns a specific gateway
func (c *Client) GetGateway(gatewayID string) (*Gateway, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/gateways/%s", c.HostURL, gatewayID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	gateway := Gateway{}
	err = json.Unmarshal(body, &gateway)
	if err != nil {
		return nil, err
	}

	return &gateway, nil
}

// CreateGateway - Create new gateway
func (c *Client) CreateGateway(gateway NewGateway, authToken *string) (*GatewaysStored, error) {

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

	gw := GatewaysStored{}
	err = json.Unmarshal(body, &gw)
	if err != nil {
		return nil, err
	}

	return &gw, nil
}

// UpdateGateway - Updates an gateway
func (c *Client) UpdateGateway(gatewayID string, gateway NewGateway, authToken *string) (*GatewaysStored, error) {
	rb, err := json.Marshal(gateway)
	if err != nil {
		return nil, err
	}

	req1, err1 := http.NewRequest("PUT", fmt.Sprintf("%s/gateways/%s", "http://localhost", gatewayID), strings.NewReader(string(rb)))

	body1, err1 := c.doRequest(req1, authToken)
	_ = body1
	_ = err1
	
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/gateways/%s", c.HostURL, gatewayID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	gw := GatewaysStored{}
	err = json.Unmarshal(body, &gw)
	if err != nil {
		return nil, err
	}

	return &gw, nil
}

// DeleteGateway - Deletes an gateway
func (c *Client) DeleteGateway(gatewayID string, authToken *string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/gateways/%s", c.HostURL, gatewayID), nil)
	if err != nil {
		return err
	}

	req1, err1 := http.NewRequest("DELETE", fmt.Sprintf("%s/gateways/%s", "http://localhost", gatewayID), nil)
	body1, err1 := c.doRequest(req1, authToken)
	_ = body1
	_ = err1

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return err
	}

	if string(body) != "" {
		return errors.New(string(body))
	}

	return nil
}
