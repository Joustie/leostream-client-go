package leostream


// Centers -
type CenterSummary struct {
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

// Center -
type Center struct {
	ID           int64  `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Active	   int64  `json:"active,omitempty"`
	Center_definition CenterDefinition `json:"center_definition,omitempty"`
	Center_info  CenterInfo `json:"center_info,omitempty"`
	Images []AwsImages `json:"images,omitempty"`
}

type CenterDefinition struct {
	Allow_rogue int64  `json:"allow_rogue,omitempty"`
	Flavor string `json:"flavor,omitempty"`
	Name string `json:"name,omitempty"`
	Offer_vms int64  `json:"off_vms,omitempty"`
	Poll_interval int64  `json:"poll_interval,omitempty"`
	Type string `json:"type,omitempty"`
	Type_label string `json:"type_label,omitempty"`
	Vc_name string `json:"vc_name,omitempty"`
	Vc_password string `json:"vc_password,omitempty"`
	Vc_datacenter string `json:"vc_datacenter,omitempty"`
}

type CenterInfo struct {
	Aws_sec_groups []Aws_sec_groups `json:"aws_sec_groups,omitempty"`
	Aws_sizes []string `json:"aws_sizes,omitempty"`
	Aws_subnets []string `json:"aws_subnets,omitempty"`
	Os string `json:"os,omitempty"`
	Os_version string `json:"os_version,omitempty"`
	Type string `json:"type,omitempty"`
}

type Aws_sec_groups struct {
	GDesc string `json:"gDesc,omitempty"`
	GId string `json:"gId,omitempty"`
	GName string `json:"gName,omitempty"`
	VpcId string `json:"vpcId,omitempty"`
}

type AwsImages struct {
	ID int64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
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

type Provision struct {
	Provision_on_off             int64            `json:"provision_on_off"`
	Provision_tenant_id          int64            `json:"provision_tenant_id"`
	Provision_threshold          int64            `json:"provision_threshold"`
	Provision_vm_id              int64            `json:"provision_vm_id"`
	Provision_vm_name            string           `json:"provision_vm_name"`
	Provision_vm_name_next_value int64            `json:"provision_vm_name_next_value"`
	Provision_vm_display_name    string           `json:"provision_vm_display_name"`
	Provision_url                string           `json:"provision_url"`
	Provision_server_id          int64            `json:"provision_server_id"`
	Provision_max                int64            `json:"provision_max"`
	Provision_limits_enforce     int64            `json:"provision_limits_enforce"`
	Mark_deleteable    			 int64            `json:"mark_deletable"`
	Mark_unavailable   			 int64            `json:"mark_unavailable"`
	Time_limits        			 *[]PoolTimeLimits `json:"time_limits,omitempty"`
	Center             			 *PoolAwsCenter    `json:"center,omitempty"`
}

type PoolAwsCenter struct {
	ID               	int64  	`json:"id,omitempty"`
	Name             	string 	`json:"name,omitempty"`
	Type             	string 	`json:"type,omitempty"`
	Aws_size         	string 	`json:"aws_size,omitempty"`
	Aws_t2_unlimited 	int64  	`json:"aws_t2_unlimited,omitempty"`
	Aws_iam_name     	string 	`json:"aws_iam_name,omitempty"`
	Aws_sub_net      	string 	`json:"aws_sub_net,omitempty"`
	Aws_sec_group    	string 	`json:"aws_sec_group,omitempty"`
	Aws_vpc_id       	string 	`json:"aws_vpc_id,omitempty"`
	Vc_resource_pool_id int64  	`json:"vc_resource_pool_id,omitempty"`
	Vc_spec_file_id	 	int64  	`json:"vc_spec_file_id,omitempty"`
}

// Pool summary
type PoolSummary struct {
	ID           int64  `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Assigned_vm  int64  `json:"assigned_vm,omitempty"`
	Available_vm int64  `json:"available_vm,omitempty"`
	Unavailable_vm int64  `json:"unavailable_vm,omitempty"`
	Total_vm    int64  `json:"total_vm,omitempty"`
	Parent_pool_id int64  `json:"parent_pool_id,omitempty"`
	Total_agent_running int64  `json:"total_agent_running,omitempty"`
	Total_connected int64  `json:"total_connected,omitempty"`
	Total_logged_in int64  `json:"total_logged_in,omitempty"`
	Total_vm_running int64  `json:"total_vm_running,omitempty"`
	Total_vm_stopped int64  `json:"total_vm_stopped,omitempty"`
	Total_vm_suspended int64  `json:"total_vm_suspended,omitempty"`
}


// Pool -
type Pool struct {
	ID                         int64          `json:"id,omitempty"`
	Name                       string         `json:"name"`
	Display_name               string         `json:"display_name,omitempty"`
	Notes                      string         `json:"notes,omitempty"`
	Running_desktops_threshold int64          `json:"running_desktops_threshold,omitempty"`
	Pool_definition            *PoolDefinition `json:"pool_definition,omitempty"`
	Provision	               *Provision 	  `json:"provision,omitempty"`
	PoolStats				  *PoolStats     `json:"pool_stats,omitempty"`
	Pool_domain                *PoolDomain    `json:"pool_domain,omitempty"`
	Pool_log                   *PoolLog       `json:"pool_log,omitempty"`
	Vms_list				   *[]Vms       `json:"vms_list,omitempty"`
}

 
type Vms struct {
	Display_name string `json:"display_name,omitempty"`
	Hda_status string `json:"hda_status,omitempty"`
	Id int64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Out_of_service int64 `json:"out_of_service,omitempty"`
	Status int64 `json:"status,omitempty"`
	User_id int64 `json:"user_id,omitempty"`	
}

type PoolDefinition struct {
	Restrict_by             string           `json:"restrict_by,omitempty"`
	Vm_tags                 []string         `json:"vm_tags,omitempty"`
	Vm_tags_join            string           `json:"vm_tags_join,omitempty"`
	Pool_attribute_join     string           `json:"pool_attribute_join,omitempty"`
	Adhoc_vms               []int64          `json:"adhoc_vms,omitempty"`
	Adhoc_vc_hosts          []int64          `json:"adhoc_vc_hosts,omitempty"`
	Adhoc_vc_clusters       []int64          `json:"adhoc_vc_clusters,omitempty"`
	Adhoc_vc_resource_pools []int64          `json:"adhoc_vc_resource_pools,omitempty"`
	Never_rogue             int64            `json:"never_rogue"`
	Server_ids              []int64          `json:"server_ids,omitempty"`
    Use_vmotion             int64            `json:"use_vmotion"`
	Vm_server_distribution  int64            `json:"vm_server_distribution,omitempty"`
	Attributes              []PoolAttributes `json:"attributes,omitempty"`
}

type PoolAttributes struct {
	Vm_table_field     string `json:"vm_table_field,omitempty"`
	Ad_attribute_field string `json:"ad_attribute_field"`
	Vm_gpu_field       string `json:"vm_gpu_field"`
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

// Policy -
type Policy struct {
	ID                         int64          `json:"id,omitempty"`
	Name                       string         `json:"name"`
	Notes                      string         `json:"notes,omitempty"`
	
}

type PoliciesStored struct {
	Stored_data Policy `json:"stored_data,omitempty"`
}


// PoolAssignmentSummary -	
type PoolAssignmentSummary struct {
	ID           int64  `json:"id,omitempty"`
	Pool_name         string `json:"pool_name,omitempty"`
	Pool_id   		int64  `json:"pool_id,omitempty"`
}


// PoolAssignment -
type PoolAssignment struct {
	ID           int64  `json:"id,omitempty"`
	Pool_id   		int64  `json:"pool_id,omitempty"`
	Policy_id  		int64  `json:"policy_id,omitempty"`
	Offer_filter	string `json:"offer_filter,omitempty"`
	Offer_filter_json	*PoolAssignmentFilters `json:"offer_filter_json,omitempty"`
	Plan_protocol_id	int64  `json:"plan_protocol_id,omitempty"`
	Plan_power_control_id	int64  `json:"plan_power_control_id,omitempty"`
	Plan_release_id	int64  `json:"plan_release_id,omitempty"`
	Offer_quantity	int64  `json:"offer_quantity,omitempty"`
	Display_mode	string `json:"display_mode,omitempty"`
	Start_if_stopped int64  `json:"start_if_stopped"`
}

type PoolAssignmentFilters struct {
	Filters []Filters `json:"filters,omitempty"`
	Join string `json:"join,omitempty"`
}

type Filters struct {
	Offer_filter_attribute	string `json:"offer_filter_attribute,omitempty"`
	Offer_filter_condition	string `json:"offer_filter_condition,omitempty"`
	Offer_filter_value	string `json:"offer_filter_value,omitempty"`
}

type PoolAssignmentssStored struct {
	Stored_data PoolAssignment `json:"stored_data,omitempty"`
}
