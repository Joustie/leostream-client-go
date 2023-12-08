package leostream

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// GetPoolassignments - Returns list of pool-assignments (no auth required)
func (c *Client) GetPoolAssignments(policyID string) ([]PoolAssignmentSummary, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/policies/%s/pool-assignments", c.HostURL, policyID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	poolassignmentsummary := []PoolAssignmentSummary{}
	err = json.Unmarshal(body, &poolassignmentsummary)
	if err != nil {
		return nil, err
	}

	return poolassignmentsummary, nil
}

// GetPolicy - Returns a specific policy
func (c *Client) GetPoolAssignment(policyID string, poolAssignmentID string) (*PoolAssignment, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/policies/%s/pool-assignments/%s", c.HostURL, policyID, poolAssignmentID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	poolassignment := PoolAssignment{}
	err = json.Unmarshal(body, &poolassignment)
	if err != nil {
		return nil, err
	}

	return &poolassignment, nil
}

// CreatePoolAssignment- Create new pool assignment
func (c *Client) CreatePoolAssignment(poolassignment PoolAssignment, policyID int, authToken *string) (*PoolAssignmentssStored, error) {

	rb, err := json.Marshal(poolassignment)
	if err != nil {
		return nil, err
	}

	// req1, err1 := http.NewRequest("POST", fmt.Sprintf("%s/policies/%s/pool-assignments", "http://localhost", policyID), strings.NewReader(string(rb)))

	// body1, err1 := c.doRequest(req1, authToken)
	// _ = body1
	// _ = err1

	println("policyID: ", policyID)
	fmt.Printf("%+v\n", poolassignment)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/policies/%s/pool-assignments", c.HostURL, policyID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	pl := PoolAssignmentssStored{}
	err = json.Unmarshal(body, &pl)
	if err != nil {
		return nil, err
	}

	return &pl, nil
}

// UpdatePoolAssignment - Updates a pool assignment
func (c *Client) UpdatePoolAssignment(poolassignementID string, poolassignment PoolAssignment, policyID string, authToken *string) (*PoolAssignmentssStored, error) {
	rb, err := json.Marshal(poolassignment)
	if err != nil {
		return nil, err
	}

	// req1, err1 := http.NewRequest("PUT", fmt.Sprintf("%s/policies/%s/pool-assignments/%s", "http://localhost", policyID, poolassignementID), strings.NewReader(string(rb)))

	// fmt.Printf("%+v\n", req1)
	// body1, err1 := c.doRequest(req1, authToken)
	// _ = body1
	// _ = err1

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/policies/%s/pool-assignments/%s", c.HostURL, policyID, poolassignementID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	gw := PoolAssignmentssStored{}
	err = json.Unmarshal(body, &gw)
	if err != nil {
		return nil, err
	}

	return &gw, nil
}

// DeletePoolAssignment- Deletes a pool assignment
func (c *Client) DeletePoolAssignment(poolassignementID string, policyID int, authToken *string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/policies/%s/pool-assignments/%s", c.HostURL, policyID, poolassignementID), nil)
	if err != nil {
		return err
	}

	req1, err1 := http.NewRequest("DELETE", fmt.Sprintf("%s/policies/%s/pool-assignments/%s", "http://localhost", policyID, poolassignementID), nil)
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
