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

type BaseFindCond struct {
	PageNum  *int64 `json:",omitempty"`
	PageSize *int64 `json:",omitempty"`
}

type DescribeLoadBalancersRequest struct {
	BaseFindCond
}

type DescribeLoadBalancersResponse struct {
	ResponseMetadata *ResponseMetadata           `json:"ResponseMetadata"`
	Result           DescribeLoadBalancersResult `json:"Result"`
}

type DescribeLoadBalancersResult struct {
	RequestID     string `json:"RequestId"`
	AccountID     string `json:"AccountId"`
	PageNumber    int    `json:"PageNumber"`
	PageSize      int    `json:"PageSize"`
	TotalCount    int    `json:"TotalCount"`
	LoadBalancers []struct {
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
		Tags                         []struct {
			Key   string `json:"Key"`
			Value string `json:"Value"`
		} `json:"Tags"`
	} `json:"LoadBalancers"`
}

type DescribeLoadBalancerAttributesRequest struct {
	LoadBalancerId string `json:"LoadBalancerId"`
}

type DescribeLoadBalancerAttributesResponse struct {
	ResponseMetadata *ResponseMetadata `json:"ResponseMetadata"`
}

type DescribeLoadBalancerAttributesResult struct {
	RequestID                    string    `json:"RequestId"`
	AccountID                    string    `json:"AccountId"`
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
	LogTopicID                   string    `json:"LogTopicId"`
	OverdueTime                  string    `json:"OverdueTime"`
	DeletedTime                  string    `json:"DeletedTime"`
	LoadBalancerSpec             string    `json:"LoadBalancerSpec"`
	LoadBalancerBillingType      int       `json:"LoadBalancerBillingType"`
	Enabled                      bool      `json:"Enabled"`
	ProjectName                  string    `json:"ProjectName"`
	ServiceManaged               bool      `json:"ServiceManaged"`
	AccessLog                    struct {
		Enabled    bool   `json:"Enabled"`
		BucketName string `json:"BucketName"`
	} `json:"AccessLog"`
	Eip struct {
		Isp                     string   `json:"ISP"`
		Bandwidth               int      `json:"Bandwidth"`
		EipBillingType          int      `json:"EipBillingType"`
		EipAddress              string   `json:"EipAddress"`
		BandwidthPackageID      string   `json:"BandwidthPackageId"`
		SecurityProtectionTypes []string `json:"SecurityProtectionTypes"`
	} `json:"Eip"`
	Ipv6AddressBandwidth struct {
		Isp                string `json:"ISP"`
		Bandwidth          int    `json:"Bandwidth"`
		EipBillingType     int    `json:"EipBillingType"`
		NetworkType        string `json:"NetworkType"`
		BandwidthPackageID string `json:"BandwidthPackageId"`
	} `json:"Ipv6AddressBandwidth"`
	ServerGroups []struct {
		ServerGroupID   string `json:"ServerGroupId"`
		ServerGroupName string `json:"ServerGroupName"`
	} `json:"ServerGroups"`
	Listeners []struct {
		ListenerID   string `json:"ListenerId"`
		ListenerName string `json:"ListenerName"`
	} `json:"Listeners"`
	Tags []struct {
		Key   string `json:"Key"`
		Value string `json:"Value"`
	} `json:"Tags"`
}

type DescribeListenersRequest struct {
	BaseFindCond
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
		CreateTime             time.Time `json:"CreateTime"`
		UpdateTime             time.Time `json:"UpdateTime"`
		ListenerID             string    `json:"ListenerId"`
		ListenerName           string    `json:"ListenerName"`
		ACLStatus              string    `json:"AclStatus"`
		ACLType                string    `json:"AclType"`
		Bandwidth              int       `json:"Bandwidth"`
		ACLIds                 []string  `json:"AclIds"`
		Enabled                string    `json:"Enabled"`
		Protocol               string    `json:"Protocol"`
		Scheduler              string    `json:"Scheduler"`
		ProxyProtocolType      string    `json:"ProxyProtocolType"`
		PersistenceType        string    `json:"PersistenceType"`
		PersistenceTimeut      int       `json:"PersistenceTimeut"`
		ConnectionDrainEnabled string    `json:"ConnectionDrainEnabled"`
		ConnectionDrainTimeout int       `json:"ConnectionDrainTimeout"`
		Port                   int       `json:"Port"`
		Status                 string    `json:"Status"`
		ServerGroupID          string    `json:"ServerGroupId"`
		HealthCheck            struct {
			Enabled            string `json:"Enabled"`
			Interval           int    `json:"Interval"`
			HealthyThreshold   int    `json:"HealthyThreshold"`
			UnHealthyThreshold int    `json:"UnHealthyThreshold"`
			Timeout            int    `json:"Timeout"`
		} `json:"HealthCheck"`
	} `json:"Listeners"`
}

type ModifyListenerAttributesRequest struct {
	ListenerId    string  `json:"ListenerId"`
	CertificateId *string `json:"CertificateId"`
}

type ModifyListenerAttributesResponse struct {
	ResponseMetadata *ResponseMetadata `json:"ResponseMetadata"`
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
