package leostream

import "testing"

/*
The following function tests the creation of a new Leostream client with no username or password.
*/
func TestNewClientNoAuth(t *testing.T) {
	c, err := NewClient(nil, nil, nil)
	if err != nil {
		t.Errorf("Error creating new client: %s", err)
	}
	if c == nil {
		t.Errorf("Client is nil")
	}
}

/*
The following function tests the creation of a new Leostream client with a random username and an empty password.
*/
func TestNewClientNoPassword(t *testing.T) {
	username := "test"
	c, err := NewClient(nil, &username, nil)
	if err != nil {
		t.Errorf("Error creating new client: %s", err)
	}
	if c == nil {
		t.Errorf("Client is nil")
	}
}

/*
The following function tests the creation of a new Leostream client with a empty username and an random password.
*/
func TestNewClientNoUsername(t *testing.T) {
	password := "test"
	c, err := NewClient(nil, nil, &password)
	if err != nil {
		t.Errorf("Error creating new client: %s", err)
	}
	if c == nil {
		t.Errorf("Client is nil")
	}
}

/*
The following function tests the creation of a new Leostream client and do a request with no username or password.
*/
func TestDoRequestNoAuth(t *testing.T) {
	c, err := NewClient(nil, nil, nil)
	if err != nil {
		t.Errorf("Error creating new client: %s", err)
	}
	if c == nil {
		t.Errorf("Client is nil")
	}
	if c.doRequest(nil, nil) == c {
		t.Errorf("Client is nil")
	}

	// _, err = c.doRequest(nil, nil)
	// if err == nil {
	// 	t.Errorf("Expected error")
	// }
}
