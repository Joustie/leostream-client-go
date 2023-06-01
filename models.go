package leostream

// Center -
type Center struct {
	ID           int                `json:"id"`
	Name         string             `json:"name"`
	Os			 string             `json:"os"`
	Flavor		 string             `json:"flavor"`
	Online       int                `json:"online"`
	Status       int                `json:"status"`
	Status_label string             `json:"status_label"`
	Center_type  string             `json:"type"`
	Type_label   string             `json:"type_label"`
}

// Gateway -
type Gateway struct {
	ID           int                `json:"id"`
	Name         string             `json:"name"`
}
