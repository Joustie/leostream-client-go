package leostream

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// GetPools - Returns list of pools (no auth required)
func (c *Client) GetPools() ([]NewPool, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/pools", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	pools := []NewPool{}
	err = json.Unmarshal(body, &pools)
	if err != nil {
		return nil, err
	}

	return pools, nil
}

// GetPool - Returns a specific pool
func (c *Client) GetPool(poolID string) (*PoolsStored, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/pools/%s", c.HostURL, poolID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	pool := PoolsStored{}
	err = json.Unmarshal(body, &pool)
	if err != nil {
		return nil, err
	}

	return &pool, nil
}

// CreatePool - Create new pool
func (c *Client) CreatePool(pool NewPool, authToken *string) (*PoolsStored, error) {

	rb, err := json.Marshal(pool)
	if err != nil {
		return nil, err
	}

	req1, err1 := http.NewRequest("POST", fmt.Sprintf("%s/pools", "http://localhost"), strings.NewReader(string(rb)))

	body1, err1 := c.doRequest(req1, authToken)
	_ = body1
	_ = err1

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/pools", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	pl := PoolsStored{}
	err = json.Unmarshal(body, &pl)
	if err != nil {
		return nil, err
	}

	return &pl, nil
}

// UpdatePool - Updates an pool
func (c *Client) UpdatePool(poolID string, pool NewPool, authToken *string) (*PoolsStored, error) {
	rb, err := json.Marshal(pool)
	if err != nil {
		return nil, err
	}

	req1, err1 := http.NewRequest("PUT", fmt.Sprintf("%s/pools/%s", "http://localhost", poolID), strings.NewReader(string(rb)))

	body1, err1 := c.doRequest(req1, authToken)
	_ = body1
	_ = err1

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/pools/%s", c.HostURL, poolID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	gw := PoolsStored{}
	err = json.Unmarshal(body, &gw)
	if err != nil {
		return nil, err
	}

	return &gw, nil
}

// DeletePool - Deletes an pool
func (c *Client) DeletePool(poolID string, authToken *string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/pools/%s", c.HostURL, poolID), nil)
	if err != nil {
		return err
	}

	req1, err1 := http.NewRequest("DELETE", fmt.Sprintf("%s/pools/%s", "http://localhost", poolID), nil)
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
