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
	Center_type  string             `json:"center_type"`
	Type_label   string             `json:"type_label"`
}
