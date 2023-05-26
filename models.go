package leostream

// Center -
type Center struct {
	ID           int                `json:"id"`
	Name         string             `json:"name"`
	os			 string             `json:"os"`
	flavor		 string             `json:"flavor"`
	online       int                `json:"online"`
	status       int                `json:"status"`
	status_label string             `json:"status_label"`
	center_type  string             `json:"type"`
	type_label   string             `json:"type_label"`
}
