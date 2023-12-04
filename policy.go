package leostream

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// GetPolicies - Returns list of policies (no auth required)
func (c *Client) GetPolicies() ([]Policy, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/policies", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	policies := []Policy{}
	err = json.Unmarshal(body, &policies)
	if err != nil {
		return nil, err
	}

	return policies, nil
}

// GetPolicy - Returns a specific policy
func (c *Client) GetPolicy(policyID string) (*Policy, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/policies/%s", c.HostURL, policyID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	policy := Policy{}
	err = json.Unmarshal(body, &policy)
	if err != nil {
		return nil, err
	}

	return &policy, nil
}

// CreatePolicy - Create new policy
func (c *Client) CreatePolicy(policy Policy, authToken *string) (*PoliciesStored, error) {

	rb, err := json.Marshal(policy)
	if err != nil {
		return nil, err
	}

	req1, err1 := http.NewRequest("POST", fmt.Sprintf("%s/policies", "http://localhost"), strings.NewReader(string(rb)))

	body1, err1 := c.doRequest(req1, authToken)
	_ = body1
	_ = err1

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/policies", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	pl := PoliciesStored{}
	err = json.Unmarshal(body, &pl)
	if err != nil {
		return nil, err
	}

	return &pl, nil
}

// UpdatePolicy - Updates an policy
func (c *Client) UpdatePolicy(policyID string, policy Policy, authToken *string) (*PoliciesStored, error) {
	rb, err := json.Marshal(policy)
	if err != nil {
		return nil, err
	}

	req1, err1 := http.NewRequest("PUT", fmt.Sprintf("%s/policies/%s", "http://localhost", policyID), strings.NewReader(string(rb)))

	body1, err1 := c.doRequest(req1, authToken)
	_ = body1
	_ = err1

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/policies/%s", c.HostURL, policyID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	gw := PoliciesStored{}
	err = json.Unmarshal(body, &gw)
	if err != nil {
		return nil, err
	}

	return &gw, nil
}

// DeletePolicy - Deletes an policy
func (c *Client) DeletePolicy(policyID string, authToken *string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/policies/%s", c.HostURL, policyID), nil)
	if err != nil {
		return err
	}

	req1, err1 := http.NewRequest("DELETE", fmt.Sprintf("%s/policies/%s", "http://localhost", policyID), nil)
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
