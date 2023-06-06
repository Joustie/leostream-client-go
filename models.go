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
	Display_name    string           `json:"display_name"`
	Name            string           `json:"name"`
	Pool_definition PoolDefinition 	 `json:"pool_definition,omitempty"`
	Provision       PoolProvision    `json:"provision,omitempty"`
}

type PoolProvision struct {
	On_off             int              `json:"provision_on_off"`
	Tenant_id          int              `json:"provision_tenant_id,omitempty"`
	Threshold          int              `json:"provision_threshold,omitempty"`
	Vm_id              int              `json:"provision_vm_id,omitempty"`
	Vm_name            string           `json:"provision_vm_name,omitempty"`
	Vm_name_next_value int              `json:"provision_vm_name_next_value,omitempty"`
	Vm_display_name    string           `json:"provision_vm_display_name,omitempty"`
	Url                string           `json:"provision_url,omitempty"`
	Server_id          int              `json:"provision_server_id,omitempty"`
	Max                int              `json:"provision_max,omitempty"`
	Limits_enforce     int              `json:"provision_limits_enforce,omitempty"`
	Mark_deleteable    int              `json:"mark_deleteable,omitempty"`
	Mark_unavailable   int              `json:"mark_unavailable,omitempty"`
	Time_limits        []PoolTimeLimits `json:"time_limits,omitempty"`
	Center             Center           `json:"center,omitempty"`
}

// Pool -
type Pool struct {
	ID                         int            `json:"id"`
	Name                       string         `json:"name"`
	Display_name               string         `json:"display_name"`
	Notes                      string         `json:"notes"`
	Running_desktops_threshold int            `json:"running_desktops_threshold"`
	Pool_definition            PoolDefinition `json:"pool_definition"`
	Provision                  PoolProvision  `json:"provision"`
	Domain                     PoolDomain     `json:"domain"`
	Log                        PoolLog        `json:"log"`
	Pool_stats                 PoolStats      `json:"pool_stats"`
	Vc_spec_file_id            int            `json:"vc_spec_file_id"`
	Vc_resource_pool_id        int            `json:"vc_resource_pool_id"`
	Vc_datastore_id            int            `json:"vc_datastore_id"`
	Read_only                  int            `json:"read_only"`
	Pool_type                  string         `json:"pool_type"`
	Is_root                    int            `json:"is_root "`
	Created                    string         `json:"created"`
	Updated                    string         `json:"updated"`
}

type PoolDefinition struct {
	Parent_pool_id          int            `json:"parent_pool_id,omitempty"`
	Restrict_by             string         `json:"restrict_by"`
	Vm_tags                 []string       `json:"vm_tags,omitempty"`
	Vm_tags_join            string         `json:"vm_tags_join,omitempty"`
	Pool_attribute_join     string         `json:"pool_attribute_join,omitempty"`
	Adhoc_vms               []int          `json:"adhoc_vms,omitempty"`
	Adhoc_vc_hosts          []int          `json:"adhoc_vc_hosts,omitempty"`
	Adhoc_vc_clusters       []int          `json:"adhoc_vc_clusters,omitempty"`
	Adhoc_vc_resource_pools []int          `json:"adhoc_vc_resource_pools,omitempty"`
	Never_rogue             int            `json:"never_rogue,omitempty"`
	Server_ids              []int          `json:"server_ids,omitempty"`
	Use_vmotion             int            `json:"use_vmotion,omitempty"`
	Vm_server_distribution  int            `json:"vm_server_distribution,omitempty"`
	Attributes              PoolAttributes `json:"attributes,omitempty"`
}

type PoolAttributes struct {
	Vm_table_field     string `json:"vm_table_field,omitempty"`
	Ad_attribute_field string `json:"ad_attribute_field,omitempty"`
	Vm_gpu_field       string `json:"vm_gpu_field,omitempty"`
	Text_to_match      string `json:"text_to_match,omitempty"`
	Condition_type     string `json:"condition_type,omitempty"`
}

type PoolTimeLimits struct {
	Provision_time_day       int    `json:"provision_time_day,omitempty"`
	Provision_time_start     string `json:"provision_time_start,omitempty"`
	Provision_time_stop      string `json:"provision_time_stop,omitempty"`
	Provision_time_threshold int    `json:"provision_time_threshold,omitempty"`
	Provision_time_max_size  int    `json:"provision_time_max_size,omitempty"`
}

type PoolDomain struct {
	Domain_join                     int                      `json:"domain_join,omitempty"`
	Domain_join_ou                  string                   `json:"domain_join_ou,omitempty"`
	Domain_join_hostname_as_vm_name int                      `json:"domain_join_hostname_as_vm_name,omitempty"`
	Remote_authentication           PoolRemoteAuthentication `json:"remote_authentication,omitempty"`
	Pool_join_ad_groups             []PoolJoinAdGroups       `json:"pool_join_ad_groups,omitempty"`
}

type PoolJoinAdGroups struct {
	Group_dn string `json:"group_dn,omitempty"`
}

type PoolRemoteAuthentication struct {
	Domain_join_remote_authentication_id int `json:"domain_join_remote_authentication_id,omitempty"`
}

type PoolLog struct {
	Log_information_threshold int               `json:"log_information_threshold,omitempty"`
	Log_warning_threshold     int               `json:"log_warning_threshold,omitempty"`
	Log_error_threshold       int               `json:"log_error_threshold,omitempty"`
	Retain_history            PoolRetainHistory `json:"retain_history,omitempty"`
}

type PoolRetainHistory struct {
	Pool_history_age      int `json:"pool_history_age,omitempty"`
	Pool_history_interval int `json:"pool_history_interval,omitempty"`
}

type PoolStats struct {
	Counts_updated      string `json:"counts_updated,omitempty"`
	Total_vm            int    `json:"total_vm,omitempty"`
	Total_agent_running int    `json:"total_agent_running,omitempty"`
	Total_vm_running    int    `json:"total_vm_running,omitempty"`
	Total_vm_stopped    int    `json:"total_vm_stopped,omitempty"`
	Total_vm_suspended  int    `json:"total_vm_suspended,omitempty"`
	Total_logged_in     int    `json:"total_logged_in,omitempty"`
	Total_connected     int    `json:"total_connected,omitempty"`
	Assigned_vm         int    `json:"assigned_vm,omitempty"`
	Available_vm        int    `json:"available_vm,omitempty"`
	Unavailable_vm      int    `json:"unavailable_vm,omitempty"`
}

type PoolsStored struct {
	Stored_data Pool `json:"stored_data,omitempty"`
}
