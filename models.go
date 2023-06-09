package leostream


// Center -
type Center struct {
	ID           int64  `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Os           string `json:"os,omitempty"`
	Flavor       string `json:"flavor,omitempty"`
	Online       int64  `json:"online,omitempty"`
	Status       int64  `json:"status,omitempty"`
	Status_label string `json:"status_label,omitempty"`
	Center_type  string `json:"type,omitempty"`
	Type_label   string `json:"type_label,omitempty"`
}

// Gateway -
type Gateway struct {
	ID               int64  `json:"id"`
	Name             string `json:"name"`
	Address          string `json:"address"`
	Address_private  string `json:"address_private"`
	Load_balancer_id int64  `json:"load_balancer_id"`
	Use_src_ip       int64  `json:"use_src_ip"`
	Notes            string `json:"notes"`
}

type GatewayStored struct {
	ID               int64  `json:"id"`
	Name             string `json:"name"`
	Address          string `json:"address"`
	Address_private  string `json:"address_private"`
	Load_balancer_id int64  `json:"load_balancer_id"`
	Use_src_ip       int64  `json:"use_src_ip"`
	Notes            string `json:"notes"`
	Created          string `json:"created"`
	Updated          string `json:"updated"`
	Online           int64  `json:"online"`
}

type GatewaysStored struct {
	Stored_data GatewayStored `json:"stored_data,omitempty"`
}

type NewGateway struct {
	Name             string `json:"name"`
	Address          string `json:"address"`
	Address_private  string `json:"address_private"`
	Load_balancer_id int64  `json:"load_balancer_id"`
	Use_src_ip       int64  `json:"use_src_ip"`
	Notes            string `json:"notes"`
}

type PoolProvision struct {
	On_off             int64            `json:"provision_on_off,omitempty"`
	Tenant_id          int64            `json:"provision_tenant_id,omitempty"`
	Threshold          int64            `json:"provision_threshold,omitempty"`
	Vm_id              int64            `json:"provision_vm_id,omitempty"`
	Vm_name            string           `json:"provision_vm_name,omitempty"`
	Vm_name_next_value int64            `json:"provision_vm_name_next_value,omitempty"`
	Vm_display_name    string           `json:"provision_vm_display_name,omitempty"`
	Url                string           `json:"provision_url,omitempty"`
	Server_id          int64            `json:"provision_server_id,omitempty"`
	Max                int64            `json:"provision_max,omitempty"`
	Limits_enforce     int64            `json:"provision_limits_enforce,omitempty"`
	Mark_deleteable    int64            `json:"mark_deleteable,omitempty"`
	Mark_unavailable   int64            `json:"mark_unavailable,omitempty"`
	Time_limits        []PoolTimeLimits `json:"time_limits,omitempty"`
	Center             PoolAwsCenter    `json:"center,omitempty"`
}

type PoolAwsCenter struct {
	ID               int64  `json:"id,omitempty"`
	Name             string `json:"name,omitempty"`
	Type             string `json:"type,omitempty"`
	Aws_size         string `json:"aws_size,omitempty"`
	aws_t2_unlimited int64  `json:"aws_t2_unlimited,omitempty"`
	aws_iam_name     string `json:"aws_iam_name,omitempty"`
	aws_sub_net      string `json:"aws_sub_net,omitempty"`
	aws_sec_group    string `json:"aws_sec_group,omitempty"`
	aws_vpc_id       string `json:"aws_vpc_id,omitempty"`
}

// Pool -
type Pool struct {
	ID                         int64          `json:"id,omitempty"`
	Name                       string         `json:"name"`
	Display_name               string         `json:"display_name,omitempty"`
	Notes                      string         `json:"notes,omitempty"`
	Running_desktops_threshold int64          `json:"running_desktops_threshold,omitempty"`
	Pool_definition            PoolDefinition `json:"pool_definition,omitempty"`
}

type PoolDefinition struct {
	Restrict_by             string           `json:"restrict_by"`
	Vm_tags                 []string         `json:"vm_tags,omitempty"`
	Vm_tags_join            string           `json:"vm_tags_join,omitempty"`
	Pool_attribute_join     string           `json:"pool_attribute_join,omitempty"`
	Adhoc_vms               []int64          `json:"adhoc_vms,omitempty"`
	Adhoc_vc_hosts          []int64          `json:"adhoc_vc_hosts,omitempty"`
	Adhoc_vc_clusters       []int64          `json:"adhoc_vc_clusters,omitempty"`
	Adhoc_vc_resource_pools []int64          `json:"adhoc_vc_resource_pools,omitempty"`
	Never_rogue             int64            `json:"never_rogue,omitempty"`
	Server_ids              []int64          `json:"server_ids,omitempty"`
	Use_vmotion             int64            `json:"use_vmotion,omitempty"`
	Vm_server_distribution  int64            `json:"vm_server_distribution,omitempty"`
	Attributes              []PoolAttributes `json:"attributes,omitempty"`
}

type PoolAttributes struct {
	Vm_table_field     string `json:"vm_table_field,omitempty"`
	Ad_attribute_field string `json:"ad_attribute_field,omitempty"`
	Vm_gpu_field       string `json:"vm_gpu_field,omitempty"`
	Text_to_match      string `json:"text_to_match,omitempty"`
	Condition_type     string `json:"condition_type,omitempty"`
}

type PoolTimeLimits struct {
	Provision_time_day       int64  `json:"provision_time_day,omitempty"`
	Provision_time_start     string `json:"provision_time_start,omitempty"`
	Provision_time_stop      string `json:"provision_time_stop,omitempty"`
	Provision_time_threshold int64  `json:"provision_time_threshold,omitempty"`
	Provision_time_max_size  int64  `json:"provision_time_max_size,omitempty"`
}

type PoolDomain struct {
	Domain_join                     int64                    `json:"domain_join,omitempty"`
	Domain_join_ou                  string                   `json:"domain_join_ou,omitempty"`
	Domain_join_hostname_as_vm_name int64                    `json:"domain_join_hostname_as_vm_name,omitempty"`
	Remote_authentication           PoolRemoteAuthentication `json:"remote_authentication,omitempty"`
	Pool_join_ad_groups             []PoolJoinAdGroups       `json:"pool_join_ad_groups,omitempty"`
}

type PoolJoinAdGroups struct {
	Group_dn string `json:"group_dn,omitempty"`
}

type PoolRemoteAuthentication struct {
	Domain_join_remote_authentication_id int64 `json:"domain_join_remote_authentication_id,omitempty"`
}

type PoolLog struct {
	Log_information_threshold int64             `json:"log_information_threshold,omitempty"`
	Log_warning_threshold     int64             `json:"log_warning_threshold,omitempty"`
	Log_error_threshold       int64             `json:"log_error_threshold,omitempty"`
	Retain_history            PoolRetainHistory `json:"retain_history,omitempty"`
}

type PoolRetainHistory struct {
	Pool_history_age      int64 `json:"pool_history_age,omitempty"`
	Pool_history_interval int64 `json:"pool_history_interval,omitempty"`
}

type PoolStats struct {
	Counts_updated      string `json:"counts_updated,omitempty"`
	Total_vm            int64  `json:"total_vm,omitempty"`
	Total_agent_running int64  `json:"total_agent_running,omitempty"`
	Total_vm_running    int64  `json:"total_vm_running,omitempty"`
	Total_vm_stopped    int64  `json:"total_vm_stopped,omitempty"`
	Total_vm_suspended  int64  `json:"total_vm_suspended,omitempty"`
	Total_logged_in     int64  `json:"total_logged_in,omitempty"`
	Total_connected     int64  `json:"total_connected,omitempty"`
	Assigned_vm         int64  `json:"assigned_vm,omitempty"`
	Available_vm        int64  `json:"available_vm,omitempty"`
	Unavailable_vm      int64  `json:"unavailable_vm,omitempty"`
}

type PoolsStored struct {
	Stored_data Pool `json:"stored_data,omitempty"`
}
