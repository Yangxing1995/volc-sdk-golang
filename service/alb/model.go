package alb

import "time"

type ErrorObj struct {
	CodeN   int64
	Code    string
	Message string
}

type ResponseMetadata struct {
	RequestId string
	Service   *string   `json:",omitempty"`
	Region    *string   `json:",omitempty"`
	Action    *string   `json:",omitempty"`
	Version   *string   `json:",omitempty"`
	Error     *ErrorObj `json:",omitempty"`
}

type BaseFindCond struct {
	PageNum  *int64 `json:",omitempty"`
	PageSize *int64 `json:",omitempty"`
}

type DescribeLoadBalancersRequest struct {
	PageNum  *int64 `json:",omitempty"`
	PageSize *int64 `json:",omitempty"`
}

type DescribeLoadBalancersResponse struct {
	ResponseMetadata *ResponseMetadata           `json:"ResponseMetadata"`
	Result           DescribeLoadBalancersResult `json:"Result"`
}

type DescribeLoadBalancersResult struct {
	RequestID     string          `json:"RequestId"`
	PageNumber    int             `json:"PageNumber"`
	PageSize      int             `json:"PageSize"`
	TotalCount    int             `json:"TotalCount"`
	LoadBalancers []*LoadBalancer `json:"LoadBalancers"`
}

type LoadBalancer struct {
	LoadBalancerID          string       `json:"LoadBalancerId"`
	LoadBalancerName        string       `json:"LoadBalancerName"`
	Status                  string       `json:"Status"`
	Description             string       `json:"Description"`
	CreateTime              time.Time    `json:"CreateTime"`
	UpdateTime              time.Time    `json:"UpdateTime"`
	Type                    string       `json:"Type"`
	VpcID                   string       `json:"VpcId"`
	SubnetID                string       `json:"SubnetId"`
	EipID                   string       `json:"EipID"`
	EipAddress              string       `json:"EipAddress"`
	EniID                   string       `json:"EniID"`
	EniAddress              string       `json:"EniAddress"`
	BusinessStatus          string       `json:"BusinessStatus"`
	LockReason              string       `json:"LockReason"`
	OverdueTime             string       `json:"OverdueTime"`
	DeletedTime             string       `json:"DeletedTime"`
	LoadBalancerBillingType int          `json:"LoadBalancerBillingType"`
	DNSName                 string       `json:"DNSName"`
	ZoneMappings            *ZoneMapping `json:"ZoneMappings"`
}

type ZoneMapping struct {
	ZoneID                string                 `json:"ZoneId"`
	SubnetID              string                 `json:"SubnetId"`
	LoadBalancerAddresses []*LoadBalancerAddress `json:"LoadBalancerAddresses"`
}

type LoadBalancerAddress struct {
	EniAddress string      `json:"EniAddress"`
	EniID      string      `json:"EniId"`
	EipAddress string      `json:"EipAddress"`
	EipID      string      `json:"EipId"`
	Eip        interface{} `json:"Eip"`
}

type DescribeLoadBalancerAttributesRequest struct {
	LoadBalancerId string `json:"LoadBalancerId"`
}

type DescribeLoadBalancerAttributesResponse struct {
	ResponseMetadata *ResponseMetadata `json:"ResponseMetadata"`
}

type DescribeLoadBalancerAttributesResult struct {
	RequestID               string         `json:"RequestId"`
	LoadBalancerID          string         `json:"LoadBalancerId"`
	LoadBalancerName        string         `json:"LoadBalancerName"`
	Status                  string         `json:"Status"`
	Description             string         `json:"Description"`
	CreateTime              time.Time      `json:"CreateTime"`
	UpdateTime              time.Time      `json:"UpdateTime"`
	Type                    string         `json:"Type"`
	VpcID                   string         `json:"VpcId"`
	SubnetID                string         `json:"SubnetId"`
	EipID                   string         `json:"EipID"`
	EipAddress              string         `json:"EipAddress"`
	EniID                   string         `json:"EniID"`
	EniAddress              string         `json:"EniAddress"`
	BusinessStatus          string         `json:"BusinessStatus"`
	LockReason              string         `json:"LockReason"`
	OverdueTime             string         `json:"OverdueTime"`
	DeletedTime             string         `json:"DeletedTime"`
	LoadBalancerBillingType int            `json:"LoadBalancerBillingType"`
	AccessLog               *AccessLog     `json:"AccessLog"`
	Eip                     *Eip           `json:"Eip"`
	Listeners               []*Listener    `json:"Listeners"`
	DNSName                 string         `json:"DNSName"`
	ZoneMappings            []*ZoneMapping `json:"ZoneMappings"`
}

type AccessLog struct {
	Enabled    string `json:"Enabled"`
	BucketName string `json:"BucketName"`
}

type Eip struct {
	Isp            string `json:"ISP"`
	Bandwidth      int    `json:"Bandwidth"`
	EipBillingType int    `json:"EipBillingType"`
	EipAddress     string `json:"EipAddress"`
}
type Listener struct {
	ListenerID   string `json:"ListenerId"`
	ListenerName string `json:"ListenerName"`
}

type DescribeListenersRequest struct {
	PageNum        *int64  `json:",omitempty"`
	PageSize       *int64  `json:",omitempty"`
	LoadBalancerId *string `json:"LoadBalancerId"`
}

type DescribeListenersResponse struct {
	ResponseMetadata *ResponseMetadata       `json:"ResponseMetadata"`
	Result           DescribeListenersResult `json:"Result"`
}

type DescribeListenersResult struct {
	RequestID  string `json:"RequestId"`
	PageNumber int    `json:"PageNumber"`
	PageSize   int    `json:"PageSize"`
	TotalCount int    `json:"TotalCount"`
	Listeners  []struct {
		CreateTime    time.Time `json:"CreateTime"`
		UpdateTime    time.Time `json:"UpdateTime"`
		ListenerID    string    `json:"ListenerId"`
		ListenerName  string    `json:"ListenerName"`
		Enabled       string    `json:"Enabled"`
		Protocol      string    `json:"Protocol"`
		Port          int       `json:"Port"`
		Status        string    `json:"Status"`
		ServerGroupID string    `json:"ServerGroupId"`
		ServerGroups  []struct {
			ServerGroupID   string `json:"ServerGroupId"`
			ServerGroupName string `json:"ServerGroupName"`
		} `json:"ServerGroups"`
		Description   string   `json:"Description,omitempty"`
		ACLStatus     string   `json:"AclStatus,omitempty"`
		ACLType       string   `json:"AclType,omitempty"`
		ACLIds        []string `json:"AclIds,omitempty"`
		CertificateID string   `json:"CertificateId,omitempty"`
	}
}

type Listeners struct {
	CreateTime    time.Time      `json:"CreateTime"`
	UpdateTime    time.Time      `json:"UpdateTime"`
	ListenerID    string         `json:"ListenerId"`
	ListenerName  string         `json:"ListenerName"`
	Enabled       string         `json:"Enabled"`
	Protocol      string         `json:"Protocol"`
	Port          int            `json:"Port"`
	Status        string         `json:"Status"`
	ServerGroupID string         `json:"ServerGroupId"`
	ServerGroups  []*ServerGroup `json:"ServerGroups"`
	Description   string         `json:"Description,omitempty"`
	ACLStatus     string         `json:"AclStatus,omitempty"`
	ACLType       string         `json:"AclType,omitempty"`
	ACLIds        []string       `json:"AclIds,omitempty"`
	CertificateID string         `json:"CertificateId,omitempty"`
}

type ServerGroup struct {
	ServerGroupID   string `json:"ServerGroupId"`
	ServerGroupName string `json:"ServerGroupName"`
}

type ModifyListenerAttributesRequest struct {
	ListenerId    string  `json:"ListenerId"`
	CertificateId *string `json:"CertificateId"`
}

type ModifyListenerAttributesResponse struct {
	ResponseMetadata *ResponseMetadata `json:"ResponseMetadata"`
}

type DescribeRulesRequest struct {
	ListenerId string `json:"ListenerId"`
}

type DescribeRulesResponse struct {
	ResponseMetadata *ResponseMetadata   `json:"ResponseMetadata"`
	Result           DescribeRulesResult `json:"Result"`
}

type DescribeRulesResult struct {
	RequestID string   `json:"RequestId"`
	Rules     []*Rules `json:"Rules"`
}

type Rules struct {
	RuleID        string `json:"RuleId"`
	Domain        string `json:"Domain"`
	URL           string `json:"Url"`
	ServerGroupID string `json:"ServerGroupId"`
	Description   string `json:"Description"`
}

type UploadCertificateRequest struct {
	CertificateName *string `json:"CertificateName"`
	CertificateType string  `json:"CertificateType,default:Server"`
	PublicKey       string  `json:"PublicKey"`
	PrivateKey      string  `json:"PrivateKey"`
	Description     *string `json:"Description"`
	ProjectName     *string `json:"ProjectName"`
}

type UploadCertificateResponse struct {
	ResponseMetadata *ResponseMetadata       `json:"ResponseMetadata"`
	Result           UploadCertificateResult `json:"Result"`
}

type UploadCertificateResult struct {
	RequestId     string `json:"RequestId"`
	CertificateId string `json:"CertificateId"`
}
