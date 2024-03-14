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
	DNSName                 string         `json:"DNSName"`
	ZoneMappings            []*ZoneMapping `json:"ZoneMappings"`
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
	RequestID               string            `json:"RequestId"`
	LoadBalancerID          string            `json:"LoadBalancerId"`
	LoadBalancerName        string            `json:"LoadBalancerName"`
	Status                  string            `json:"Status"`
	Description             string            `json:"Description"`
	CreateTime              time.Time         `json:"CreateTime"`
	UpdateTime              time.Time         `json:"UpdateTime"`
	Type                    string            `json:"Type"`
	VpcID                   string            `json:"VpcId"`
	SubnetID                string            `json:"SubnetId"`
	EipID                   string            `json:"EipID"`
	EipAddress              string            `json:"EipAddress"`
	EniID                   string            `json:"EniID"`
	EniAddress              string            `json:"EniAddress"`
	BusinessStatus          string            `json:"BusinessStatus"`
	LockReason              string            `json:"LockReason"`
	OverdueTime             string            `json:"OverdueTime"`
	DeletedTime             string            `json:"DeletedTime"`
	LoadBalancerBillingType int               `json:"LoadBalancerBillingType"`
	AccessLog               *AccessLog        `json:"AccessLog"`
	Eip                     *Eip              `json:"Eip"`
	Listeners               []*SimpleListener `json:"Listeners"`
	DNSName                 string            `json:"DNSName"`
	ZoneMappings            []*ZoneMapping    `json:"ZoneMappings"`
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
type SimpleListener struct {
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
	RequestID  string      `json:"RequestId"`
	PageNumber int         `json:"PageNumber"`
	PageSize   int         `json:"PageSize"`
	TotalCount int         `json:"TotalCount"`
	Listeners  []*Listener `json:"Listeners"`
}

type Listener struct {
	CreateTime       time.Time          `json:"CreateTime"`
	UpdateTime       time.Time          `json:"UpdateTime"`
	ListenerID       string             `json:"ListenerId"`
	ListenerName     string             `json:"ListenerName"`
	Enabled          string             `json:"Enabled"`
	Protocol         string             `json:"Protocol"`
	Port             int                `json:"Port"`
	Status           string             `json:"Status"`
	ServerGroupID    string             `json:"ServerGroupId"`
	ServerGroups     []*ServerGroup     `json:"ServerGroups"`
	Description      string             `json:"Description,omitempty"`
	ACLStatus        string             `json:"AclStatus,omitempty"`
	ACLType          string             `json:"AclType,omitempty"`
	ACLIds           []string           `json:"AclIds,omitempty"`
	CertificateID    string             `json:"CertificateId,omitempty"`
	CACertificateId  string             `json:"CACertificateId,omitempty"` // 监听器关联的CA证书 ID，创建 HTTPS 监听器时指定 CA 证书，则支持双向认证，否则为单向认证。
	DomainExtensions []*DomainExtension `json:"DomainExtensions,omitempty"`
}

type DomainExtension struct {
	DomainExtensionId string `json:"DomainExtensionId,omitempty"`
	CertificateId     string `json:"CertificateId,omitempty"`
	Domain            string `json:"Domain,omitempty"`
	ListenerId        string `json:"ListenerId,omitempty"`
}

type ServerGroup struct {
	ServerGroupID   string `json:"ServerGroupId"`
	ServerGroupName string `json:"ServerGroupName"`
}

type ModifyListenerAttributesRequest struct {
	ListenerId       string                `json:"ListenerId"`
	CertificateId    *string               `json:"CertificateId,omitempty"`
	CACertificateId  *string               `json:"CACertificateId,omitempty"`
	DomainExtensions []*DomainExtensionReq `json:"DomainExtensions,omitempty"` // 仅支持HTTPS协议
}

type DomainExtensionReq struct {
	DomainExtensionId string       `json:"DomainExtensionId,omitempty"`
	CertificateId     string       `json:"CertificateId,omitempty"`
	Domain            string       `json:"Domain,omitempty"`
	Action            DomainAction `json:"Action,omitempty"` // 请求时参数
}

type DomainAction = string

const (
	DomainActionCreate DomainAction = "create" // 新增
	DomainActionModify DomainAction = "modify" // 修改
	DomainActionDelete DomainAction = "delete" // 删除
)

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

type DescribeCertificatesRequest struct {
	PageNumber *int64 `json:",omitempty"` // default: 1
	PageSize   *int64 `json:",omitempty"` // default: 20
}

type DescribeCertificatesResponse struct {
	ResponseMetadata *ResponseMetadata          `json:"ResponseMetadata"`
	Result           DescribeCertificatesResult `json:"Result"`
}

type DescribeCertificatesResult struct {
	RequestId    string         `json:"RequestId"`
	PageNumber   int            `json:"PageNumber"`
	PageSize     int            `json:"PageSize"`
	TotalCount   int            `json:"TotalCount"`
	Certificates []*Certificate `json:"Certificates"`
}

type Certificate struct {
	CertificateId   string    `json:"CertificateId"`
	CertificateName string    `json:"CertificateName"`
	CertificateType string    `json:"CertificateType"`
	Description     string    `json:"Description"`
	CreateTime      time.Time `json:"CreateTime"`
	ExpiredAt       time.Time `json:"ExpiredAt"`
	DomainName      string    `json:"DomainName"`
	Listeners       []string  `json:"Listeners"` // listener id list
}