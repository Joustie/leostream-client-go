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
