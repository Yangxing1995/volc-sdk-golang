package certificate

type ErrorObj struct {
	CodeN   int64
	Code    string
	Message string
}

type ReqInterface interface {
	// 获取参数位置
	GetPosition() string
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

type ImportCertificateRequest struct {
	Tag                 string           `json:"tag,omitempty"`
	Project             string           `json:"project,omitempty"`
	NoVerifyAndFixChain *bool            `json:"no_verify_and_fix_chain,omitempty"` // 跳过检查证书的合法性
	Repeatable          *bool            `json:"repeatable,omitempty"`
	CertificateInfo     *CertificateInfo `json:"certificate_info,omitempty"`

	GmCertificateInfo *GmCertificateInfo `json:"gm_certificate_info,omitempty"`
}

type CertificateInfo struct {
	Certificate string `json:"certificate"`
	PrivateKey  string `json:"private_key"`
}

type GmCertificateInfo struct {
	EncryptCertificate string `json:"encrypt_certificate"`
	EncryptPrivateKey  string `json:"encrypt_private_key"`
	SignCertificate    string `json:"sign_certificate"`
	SignPrivateKey     string `json:"sign_private_key"`
}

type ImportCertificateResponse struct {
	ResponseMetadata *ResponseMetadata        `json:",omitempty"`
	Result           *ImportCertificateResult `json:",omitempty"`
}

type ImportCertificateResult struct {
	ID       string `json:"id"`
	RepeatID string `json:"repeat_id"` // 重复证书的ID
}

type CertificateGetInstanceListRequest struct {
	InstanceIds             []string     `json:"InstanceIds,omitempty"`             // 证书实例ID列表
	Status                  []string     `json:"Status,omitempty"`                  // 证书状态列表: NotSubmitted待提交, Pending验证中, Issued已签发, Cancelling取消中, Canceled已取消, Revoking吊销中, Revoked已吊销, Failed申请失败, Unknown未知
	Tag                     *string      `json:"Tag,omitempty"`                     // 证书备注,支持模糊匹配
	CommonName              *string      `json:"CommonName,omitempty"`              // 证书公用名称(CN),支持模糊匹配
	Domain                  *string      `json:"Domain,omitempty"`                  // 证书主题备用名称(SAN),支持模糊匹配
	InstanceType            *string      `json:"InstanceType,omitempty"`            // 证书实例类型: Free免费, Test测试(付费), Paid正式(付费), Imported上传的证书
	IsRevoked               *bool        `json:"IsRevoked,omitempty"`               // 是否只返回已吊销的证书,默认false
	IsValid                 *bool        `json:"IsValid,omitempty"`                 // 是否只返回有效的证书,默认false
	CertificateExpireBefore *string      `json:"CertificateExpireBefore,omitempty"` // 证书过期时间的结束时间,格式: yyyy-mm-dd hh:mm:ss
	CertificateExpireAfter  *string      `json:"CertificateExpireAfter,omitempty"`  // 证书过期时间的开始时间,格式: yyyy-mm-dd hh:mm:ss
	PageNumber              *int64       `json:"PageNumber,omitempty"`              // 页码,默认1
	PageSize                *int64       `json:"PageSize,omitempty"`                // 单页最大数量,默认10,最大100
	ProjectName             *string      `json:"ProjectName,omitempty"`             // 证书实例所属项目名称
	TagFilters              []*TagFilter `json:"TagFilters,omitempty"`              // 资源标签列表
}

type TagFilter struct {
	Key    string   `json:"Key"`              // 标签键
	Values []string `json:"Values,omitempty"` // 标签值列表
}

type CertificateGetInstanceListResponse struct {
	ResponseMetadata *ResponseMetadata           `json:",omitempty"`
	Result           *CertificateGetInstanceList `json:",omitempty"`
}

type CertificateGetInstanceList struct {
	PageNumber int `json:"PageNumber"`
	PageSize   int `json:"PageSize"`
	TotalCount int `json:"TotalCount"`

	Instances []*CertificateResult `json:"Instances"`
}

type CertificateGetInstanceRequest struct {
	InstanceId string
}

type CertificateGetInstanceResponse struct {
	ResponseMetadata *ResponseMetadata  `json:",omitempty"`
	Result           *CertificateDetail `json:",omitempty"`
}

type CertificateResult struct {
	InstanceId              string   `json:"InstanceId"`
	AccountId               string   `json:"AccountId"`
	SourceId                string   `json:"SourceId"`
	Tag                     string   `json:"Tag"`
	Status                  string   `json:"Status"`
	InstanceType            string   `json:"InstanceType"`
	InstanceLevel           string   `json:"InstanceLevel"`
	OrderBrand              string   `json:"OrderBrand"`
	OrderPeriod             int      `json:"OrderPeriod"`
	IsCertificateSm         bool     `json:"IsCertificateSm"`
	IsCertificateRevoked    bool     `json:"IsCertificateRevoked"`
	CommonName              string   `json:"CommonName"`
	San                     []string `json:"San"`
	Issuer                  string   `json:"Issuer"`
	CertificateDomainType   string   `json:"CertificateDomainType"`
	CertificateKeyAlgorithm string   `json:"CertificateKeyAlgorithm"`
	NotBefore               string   `json:"NotBefore"`
	NotAfter                string   `json:"NotAfter"`
	Tags                    []struct {
		Key   string `json:"Key"`
		Value string `json:"Value"`
	} `json:"Tags"`
	ProjectName string `json:"ProjectName"`
	CreatedTime string `json:"CreatedTime"`
}

type CertificateDetail struct {
	InstanceId            string   `json:"InstanceId"`
	AccountId             string   `json:"AccountId"`
	SourceId              string   `json:"SourceId"`
	Tag                   string   `json:"Tag"`
	Status                string   `json:"Status"`
	InstanceType          string   `json:"InstanceType"`
	InstanceLevel         string   `json:"InstanceLevel"`
	OrderBrand            string   `json:"OrderBrand"`
	OrderOrganizationId   string   `json:"OrderOrganizationId"`
	OrderPeriod           int      `json:"OrderPeriod"`
	OrderPlan             string   `json:"OrderPlan"`
	CertificateDomainType string   `json:"CertificateDomainType"`
	IsCertificateSm       bool     `json:"IsCertificateSm"`
	IsCertificateRevoked  bool     `json:"IsCertificateRevoked"`
	CommonName            string   `json:"CommonName"`
	San                   []string `json:"San"`
	NotBefore             string   `json:"NotBefore"`
	NotAfter              string   `json:"NotAfter"`
	CertificateDetail     struct {
		Subject struct {
			CommonName       string `json:"CommonName"`
			Organization     string `json:"Organization"`
			OrganizationUnit string `json:"OrganizationUnit"`
			Country          string `json:"Country"`
			Province         string `json:"Province"`
			Locality         string `json:"Locality"`
			Address          string `json:"Address"`
			PostalCode       string `json:"PostalCode"`
		} `json:"Subject"`
		Issuer struct {
			CommonName       string `json:"CommonName"`
			Organization     string `json:"Organization"`
			OrganizationUnit string `json:"OrganizationUnit"`
			Country          string `json:"Country"`
			Province         string `json:"Province"`
			Locality         string `json:"Locality"`
			Address          string `json:"Address"`
			PostalCode       string `json:"PostalCode"`
		} `json:"Issuer"`
		Chain              []string `json:"Chain"`
		PrivateKey         string   `json:"PrivateKey"`
		SerialNumber       string   `json:"SerialNumber"`
		KeyAlgorithm       string   `json:"KeyAlgorithm"`
		SignatureAlgorithm string   `json:"SignatureAlgorithm"`
		FingerPrintSha256  string   `json:"FingerPrintSha256"`
		FingerPrintSha1    string   `json:"FingerPrintSha1"`
	} `json:"CertificateDetail"`
	EncryptionCertificateDetail struct {
		Subject struct {
			CommonName       string `json:"CommonName"`
			Organization     string `json:"Organization"`
			OrganizationUnit string `json:"OrganizationUnit"`
			Country          string `json:"Country"`
			Province         string `json:"Province"`
			Locality         string `json:"Locality"`
			Address          string `json:"Address"`
			PostalCode       string `json:"PostalCode"`
		} `json:"Subject"`
		Issuer struct {
			CommonName       string `json:"CommonName"`
			Organization     string `json:"Organization"`
			OrganizationUnit string `json:"OrganizationUnit"`
			Country          string `json:"Country"`
			Province         string `json:"Province"`
			Locality         string `json:"Locality"`
			Address          string `json:"Address"`
			PostalCode       string `json:"PostalCode"`
		} `json:"Issuer"`
		Chain              []string `json:"Chain"`
		PrivateKey         string   `json:"PrivateKey"`
		SerialNumber       string   `json:"SerialNumber"`
		KeyAlgorithm       string   `json:"KeyAlgorithm"`
		SignatureAlgorithm string   `json:"SignatureAlgorithm"`
		FingerPrintSha256  string   `json:"FingerPrintSha256"`
		FingerPrintSha1    string   `json:"FingerPrintSha1"`
	} `json:"EncryptionCertificateDetail"`
	ProjectName string `json:"ProjectName"`
	Tags        []struct {
		Key   string `json:"Key"`
		Value string `json:"Value"`
	} `json:"Tags"`
	CreatedTime string `json:"CreatedTime"`
}

type DeployService struct {
	BindingDomains []string `json:"binding_domains"`
	Service        string   `json:"service"` //DCDN、CDN
}
