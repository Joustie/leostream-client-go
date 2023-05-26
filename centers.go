package leostream

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetCenters - Returns list of coffees (no auth required)
func (c *Client) GetCenters() ([]Center, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/centers", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	centers := []Centers{}
	err = json.Unmarshal(body, &centers)
	if err != nil {
		return nil, err
	}

	return centers, nil
}

// GetCenter - Returns specific center (no auth required)
func (c *Client) GetCenter(centerID string) ([]Center, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/centers/%s", c.HostURL, centerID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	centers := []Center{}
	err = json.Unmarshal(body, &centers)
	if err != nil {
		return nil, err
	}

	return centers, nil
}

// CreateCenter - Create new center
func (c *Client) CreateCenter(center Center, authToken *string) (*Center, error) {
	rb, err := json.Marshal(center)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/centers", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	newCenter := Center{}
	err = json.Unmarshal(body, &newCenter)
	if err != nil {
		return nil, err
	}

	return &newCenter, nil
}
