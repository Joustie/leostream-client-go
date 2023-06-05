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
	Provision       PoolProvision        `json:"provision,omitempty"`
}

type PoolProvision struct {
	Provision_on_off types.Int64 `json:"provision_on_off"`
	Provision_tenant_id types.Int64 `json:"provision_tenant_id"`
	Provision_threshold types.Int64 `json:"provision_threshold"`
	Provision_vm_id types.Int64 `json:"provision_vm_id"`
	Provision_vm_name types.String `json:"provision_vm_name"`
	Provision_vm_name_next_value types.Int64 `json:"provision_vm_name_next_value"`
	Provision_vm_display_name types.String `json:"provision_vm_display_name"`
	Provision_url types.String `json:"provision_url"`
	Provision_server_id types.Int64 `json:"provision_server_id"`
	Provision_max types.Int64 `json:"provision_max"`
	Provision_limits_enforce types.Int64 `json:"provision_limits_enforce"`
	Mark_deleteable types.Int64 `json:"mark_deleteable"`
	Mark_unavailable types.Int64 `json:"mark_unavailable"`
	Time_limits []PoolTimeLimits `json:"time_limits"`
	Center Center `json:"center"`
}

// Pool -
type Pool struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Display_name  	 types.String `json:"display_name"`
	Notes            types.String `json:"notes"`
	Running_desktops_threshold		 types.Int64 `json:"running_desktops_threshold"`
	Pool_definition PoolDefinition `json:"pool_definition"`
	Provision PoolProvision `json:"provision"`
	Domain PoolDomain `json:"domain"`
	Log PoolLog `json:"log"`
	Pool_stats PoolStats `json:"pool_stats"`
	Vc_spec_file_id       types.Int64  `json:"vc_spec_file_id "`
	Vc_resource_pool_id       types.Int64  `json:"vc_resource_pool_id "`
	Vc_datastore_id       types.Int64  `json:"vc_datastore_id "`
	Read_only       types.Int64  `json:"read_only "`
	Pool_type            types.String `json:"pool_type"`
	Is_root       types.Int64  `json:"is_root "`
	Created	 types.String `json:"created"`
	Updated 	 types.String `json:"updated"`
}

type PoolDefinition struct {
	Parent_pool_id       types.Int64  `json:"parent_pool_id "`
	Restrict_by 	 types.String  `json:"restrict_by "`
	Vm_tags	[]types.String `json:"vm_tags"`
	Vm_tags_join 	 types.String  `json:"vm_tags_join "`
	Pool_attribute_join 	 types.String  `json:"pool_attribute_join "`
	Adhoc_vms 	 []types.Int64 `json:"adhoc_vms "`
	Adhoc_vc_hosts 	 []types.Int64 `json:"adhoc_vc_hosts "`
	Adhoc_vc_clusters 	 []types.Int64 `json:"adhoc_vc_clusters "`
	Adhoc_vc_resource_pools 	 []types.Int64 `json:"adhoc_vc_resource_pools "`
	Never_rogue 	 types.Int64  `json:"never_rogue "`
	Server_ids 	 []types.Int64 `json:"server_ids "`
	Use_vmotion 	 types.Int64  `json:"use_vmotion "`
	Vm_server_distribution types.Int64 `json:"vm_server_distribution "`
	Attributes PoolAttributes `json:"attributes"`
}

type PoolAttributes struct {
	Vm_table_field 	 types.String  `json:"vm_table_field "`
	Ad_attribute_field 	 types.String  `json:"ad_attribute_field "`
	Vm_gpu_field 	 types.String  `json:"vm_gpu_field "`
	Text_to_match	 types.String  `json:"text_to_match "`
	Condition_type	 types.String  `json:"condition_type "`
}

type PoolTimeLimits struct {
	Provision_time_day types.Int64 `json:"provision_time_day"`
	Provision_time_start types.String `json:"provision_time_start"`
	Provision_time_stop types.String `json:"provision_time_stop"`
	Provision_time_threshold types.Int64 `json:"provision_time_threshold"`
	Provision_time_max_size types.Int64 `json:"provision_time_max_size"`
}

type PoolDomain struct {
	Domain_join types.Int64 `json:"domain_join"`
	Domain_join_ou types.String `json:"domain_join_ou"`
	Domain_join_hostname_as_vm_name types.Int64 `json:"domain_join_hostname_as_vm_name"`
	Remote_authentication PoolRemoteAuthentication `json:"remote_authentication"`
	Pool_join_ad_groups []PoolJoinAdGroups `json:"pool_join_ad_groups"`
}

type PoolJoinAdGroups struct {
	Group_dn types.String `json:"group_dn"`
}

type PoolRemoteAuthentication struct {
	Domain_join_remote_authentication_id types.Int64 `json:"domain_join_remote_authentication_id"`
}

type PoolLog struct {
	Log_information_threshold types.Int64 `json:"log_information_threshold"`
	Log_warning_threshold types.Int64 `json:"log_warning_threshold"`
	Log_error_threshold types.Int64 `json:"log_error_threshold"`
	Retain_history PoolRetainHistory `json:"retain_history"`
}

type PoolRetainHistory struct {
	Pool_history_age types.Int64 `json:"pool_history_age"`
	Pool_history_interval types.Int64 `json:"pool_history_interval"`
}

type PoolStats struct {
	Counts_updated types.String	`json:"counts_updated"`
	Total_vm types.Int64	`json:"total_vm"`
	Total_agent_running types.Int64	`json:"total_agent_running"`
	Total_vm_running types.Int64	`json:"total_vm_running"`
	Total_vm_stopped types.Int64	`json:"total_vm_stopped"`
	Total_vm_suspended types.Int64	`json:"total_vm_suspended"`
	Total_logged_in types.Int64	`json:"total_logged_in"`
	Total_connected types.Int64	`json:"total_connected"`
	Assigned_vm types.Int64	`json:"assigned_vm"`
	Available_vm types.Int64	`json:"available_vm"`
	Unavailable_vm types.Int64	`json:"unavailable_vm"`
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
