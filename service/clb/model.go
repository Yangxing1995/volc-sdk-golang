package clb

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
	AccountID     string          `json:"AccountId"`
	PageNumber    int             `json:"PageNumber"`
	PageSize      int             `json:"PageSize"`
	TotalCount    int             `json:"TotalCount"`
	LoadBalancers []*LoadBalancer `json:"LoadBalancers"`
}

type LoadBalancer struct {
	ExclusiveClusterID           string    `json:"ExclusiveClusterId"`
	LoadBalancerID               string    `json:"LoadBalancerId"`
	LoadBalancerName             string    `json:"LoadBalancerName"`
	Status                       string    `json:"Status"`
	Description                  string    `json:"Description"`
	CreateTime                   time.Time `json:"CreateTime"`
	UpdateTime                   time.Time `json:"UpdateTime"`
	Type                         string    `json:"Type"`
	MasterZoneID                 string    `json:"MasterZoneId"`
	SlaveZoneID                  string    `json:"SlaveZoneId"`
	VpcID                        string    `json:"VpcId"`
	SubnetID                     string    `json:"SubnetId"`
	ModificationProtectionStatus string    `json:"ModificationProtectionStatus"`
	ModificationProtectionReason string    `json:"ModificationProtectionReason"`
	AddressIPVersion             string    `json:"AddressIpVersion"`
	EipID                        string    `json:"EipID"`
	Ipv6EipID                    string    `json:"Ipv6EipId"`
	EipAddress                   string    `json:"EipAddress"`
	EniIpv6Address               string    `json:"EniIpv6Address"`
	EniID                        string    `json:"EniID"`
	EniAddress                   string    `json:"EniAddress"`
	BusinessStatus               string    `json:"BusinessStatus"`
	LockReason                   string    `json:"LockReason"`
	OverdueTime                  string    `json:"OverdueTime"`
	DeletedTime                  string    `json:"DeletedTime"`
	LoadBalancerSpec             string    `json:"LoadBalancerSpec"`
	LoadBalancerBillingType      int       `json:"LoadBalancerBillingType"`
	ServiceManaged               bool      `json:"ServiceManaged"`
	ProjectName                  string    `json:"ProjectName"`
	Tags                         []*Tag    `json:"Tags"`
}

type Tag struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

type DescribeLoadBalancerAttributesRequest struct {
	LoadBalancerId string `json:"LoadBalancerId"`
}

type DescribeLoadBalancerAttributesResponse struct {
	ResponseMetadata *ResponseMetadata `json:"ResponseMetadata"`
}

type DescribeLoadBalancerAttributesResult struct {
	RequestID                    string                `json:"RequestId"`
	AccountID                    string                `json:"AccountId"`
	ExclusiveClusterID           string                `json:"ExclusiveClusterId"`
	LoadBalancerID               string                `json:"LoadBalancerId"`
	LoadBalancerName             string                `json:"LoadBalancerName"`
	Status                       string                `json:"Status"`
	Description                  string                `json:"Description"`
	CreateTime                   time.Time             `json:"CreateTime"`
	UpdateTime                   time.Time             `json:"UpdateTime"`
	Type                         string                `json:"Type"`
	MasterZoneID                 string                `json:"MasterZoneId"`
	SlaveZoneID                  string                `json:"SlaveZoneId"`
	VpcID                        string                `json:"VpcId"`
	SubnetID                     string                `json:"SubnetId"`
	ModificationProtectionStatus string                `json:"ModificationProtectionStatus"`
	ModificationProtectionReason string                `json:"ModificationProtectionReason"`
	AddressIPVersion             string                `json:"AddressIpVersion"`
	EipID                        string                `json:"EipID"`
	Ipv6EipID                    string                `json:"Ipv6EipId"`
	EipAddress                   string                `json:"EipAddress"`
	EniIpv6Address               string                `json:"EniIpv6Address"`
	EniID                        string                `json:"EniID"`
	EniAddress                   string                `json:"EniAddress"`
	BusinessStatus               string                `json:"BusinessStatus"`
	LockReason                   string                `json:"LockReason"`
	LogTopicID                   string                `json:"LogTopicId"`
	OverdueTime                  string                `json:"OverdueTime"`
	DeletedTime                  string                `json:"DeletedTime"`
	LoadBalancerSpec             string                `json:"LoadBalancerSpec"`
	LoadBalancerBillingType      int                   `json:"LoadBalancerBillingType"`
	Enabled                      bool                  `json:"Enabled"`
	ProjectName                  string                `json:"ProjectName"`
	ServiceManaged               bool                  `json:"ServiceManaged"`
	AccessLog                    *AccessLog            `json:"AccessLog"`
	Eip                          *Eip                  `json:"Eip"`
	Ipv6AddressBandwidth         *Ipv6AddressBandwidth `json:"Ipv6AddressBandwidth"`
	ServerGroups                 []*ServerGroup        `json:"ServerGroups"`
	Listeners                    []*Listener           `json:"Listeners"`
	Tags                         []*Tag                `json:"Tags"`
}

type AccessLog struct {
	Enabled    bool   `json:"Enabled"`
	BucketName string `json:"BucketName"`
}

type Eip struct {
	Isp                     string   `json:"ISP"`
	Bandwidth               int      `json:"Bandwidth"`
	EipBillingType          int      `json:"EipBillingType"`
	EipAddress              string   `json:"EipAddress"`
	BandwidthPackageID      string   `json:"BandwidthPackageId"`
	SecurityProtectionTypes []string `json:"SecurityProtectionTypes"`
}

type Listener struct {
	ListenerID   string `json:"ListenerId"`
	ListenerName string `json:"ListenerName"`
}

type ServerGroup struct {
	ServerGroupID   string `json:"ServerGroupId"`
	ServerGroupName string `json:"ServerGroupName"`
}

type Ipv6AddressBandwidth struct {
	Isp                string `json:"ISP"`
	Bandwidth          int    `json:"Bandwidth"`
	EipBillingType     int    `json:"EipBillingType"`
	NetworkType        string `json:"NetworkType"`
	BandwidthPackageID string `json:"BandwidthPackageId"`
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
	RequestID  string            `json:"RequestId"`
	PageNumber int               `json:"PageNumber"`
	PageSize   int               `json:"PageSize"`
	TotalCount int               `json:"TotalCount"`
	Listeners  []*ListenerDetail `json:"Listeners"`
}

type ListenerDetail struct {
	CreateTime             time.Time    `json:"CreateTime"`
	UpdateTime             time.Time    `json:"UpdateTime"`
	ListenerID             string       `json:"ListenerId"`
	ListenerName           string       `json:"ListenerName"`
	ACLStatus              string       `json:"AclStatus"`
	ACLType                string       `json:"AclType"`
	Bandwidth              int          `json:"Bandwidth"`
	ACLIds                 []string     `json:"AclIds"`
	Enabled                string       `json:"Enabled"`
	Protocol               string       `json:"Protocol"`
	Scheduler              string       `json:"Scheduler"`
	ProxyProtocolType      string       `json:"ProxyProtocolType"`
	PersistenceType        string       `json:"PersistenceType"`
	PersistenceTimeut      int          `json:"PersistenceTimeut"`
	ConnectionDrainEnabled string       `json:"ConnectionDrainEnabled"`
	ConnectionDrainTimeout int          `json:"ConnectionDrainTimeout"`
	Port                   int          `json:"Port"`
	Status                 string       `json:"Status"`
	ServerGroupID          string       `json:"ServerGroupId"`
	HealthCheck            *HealthCheck `json:"HealthCheck"`
}

type HealthCheck struct {
	Enabled            string `json:"Enabled"`
	Interval           int    `json:"Interval"`
	HealthyThreshold   int    `json:"HealthyThreshold"`
	UnHealthyThreshold int    `json:"UnHealthyThreshold"`
	Timeout            int    `json:"Timeout"`
}

type ModifyListenerAttributesRequest struct {
	ListenerId    string  `json:"ListenerId"`
	CertificateId *string `json:"CertificateId"`
}

type ModifyListenerAttributesResponse struct {
	ResponseMetadata *ResponseMetadata `json:"ResponseMetadata"`
}

type DescribeRulesRequest struct {
	ListenerId *string `json:",omitempty"` // required
}

type DescribeRulesResponse struct {
	ResponseMetadata *ResponseMetadata   `json:"ResponseMetadata"`
	Result           DescribeRulesResult `json:"Result"`
}

type DescribeRulesResult struct {
	RequestID string  `json:"RequestId"`
	Rules     []*Rule `json:"Rules"`
}

type Rule struct {
	RuleID        string `json:"RuleId"`
	Domain        string `json:"Domain"`
	URL           string `json:"Url"`
	ServerGroupID string `json:"ServerGroupId"`
	Description   string `json:"Description"`
}

type UploadCertificateRequest struct {
	CertificateName *string `json:"CertificateName"`
	PublicKey       *string `json:"PublicKey"`
	PrivateKey      *string `json:"PrivateKey"`
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
