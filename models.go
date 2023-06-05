package leostream

// Center -
type Center struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Os           string `json:"os"`
	Flavor       string `json:"flavor"`
	Online       int    `json:"online"`
	Status       int    `json:"status"`
	Status_label string `json:"status_label"`
	Center_type  string `json:"type"`
	Type_label   string `json:"type_label"`
}

// Gateway -
type Gateway struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Address          string `json:"address"`
	Address_private  string `json:"address_private"`
	Load_balancer_id int    `json:"load_balancer_id"`
	Use_src_ip       int    `json:"use_src_ip"`
	Notes            string `json:"notes"`
}

type GatewayStored struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Address          string `json:"address"`
	Address_private  string `json:"address_private"`
	Load_balancer_id int    `json:"load_balancer_id"`
	Use_src_ip       int    `json:"use_src_ip"`
	Notes            string `json:"notes"`
	Created          string `json:"created"`
	Updated          string `json:"updated"`
	Online           int    `json:"online"`
}

type GatewaysStored struct {
	Stored_data GatewayStored `json:"stored_data,omitempty"`
}

type NewGateway struct {
	Name             string `json:"name"`
	Address          string `json:"address"`
	Address_private  string `json:"address_private"`
	Load_balancer_id int    `json:"load_balancer_id"`
	Use_src_ip       int    `json:"use_src_ip"`
	Notes            string `json:"notes"`
}

type NewPool struct {
	Display_name     string           `json:"display_name"`
	Name            string           `json:"name"`
	Pool_definition []PoolDefinition `json:"pool_definition,omitempty"`
	Provision       Provision        `json:"provision,omitempty"`
}

// OrderItem -
type PoolDefinition struct {
	Server_ids    []int  `json:"server_ids"`
	Never_rogue   int    `json:"never_rogue"`
	Restricted_by string `json:"restricted_by"`
}

type Provision struct {
	On_off int `json:"on_off"`
}

// Pool -
type Pool struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Address          string `json:"address"`
	Address_private  string `json:"address_private"`
	Load_balancer_id int    `json:"load_balancer_id"`
	Use_src_ip       int    `json:"use_src_ip"`
	Notes            string `json:"notes"`
}

type PoolsStored struct {
	Stored_data PoolStored `json:"stored_data,omitempty"`
}

type PoolStored struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Address          string `json:"address"`
	Address_private  string `json:"address_private"`
	Load_balancer_id int    `json:"load_balancer_id"`
	Use_src_ip       int    `json:"use_src_ip"`
	Notes            string `json:"notes"`
}
