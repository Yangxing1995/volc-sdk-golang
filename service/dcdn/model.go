package dcdn

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

type DescribeUserDomainsRequest struct {
	PageNum  *int64    `json:",omitempty"`
	PageSize *int64    `json:",omitempty"`
	Project  *[]string `json:",omitempty"`
}

type DescribeUserDomainsResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeUserDomainsResult
}

type DescribeUserDomainsResult struct {
	AllDomainNum    int `json:"AllDomainNum"`
	OnlineDomainNum int `json:"OnlineDomainNum"`
	Domains         []struct {
		Domain string `json:"Domain"`
		Status string `json:"Status"`
		Scope  string `json:"Scope"`
	} `json:"Domains"`
	PageNum  int `json:"PageNum"`
	PageSize int `json:"PageSize"`
}

type DescribeDomainConfigRequest struct {
	Domains []string `json:"Domains"`
}

type DescribeDomainConfigResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeDomainConfigResult
}

type DescribeDomainConfigResult = []*DescribeDomainDetailResult

type DescribeDomainDetailRequest struct {
	Domain *string `json:"Domain,omitempty"`
}

type DescribeDomainDetailResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeDomainDetailResult
}

type DescribeDomainDetailResult struct {
	BrCompress struct {
		Enable bool `json:"Enable"`
	} `json:"BrCompress"`
	CaCertId string `json:"CaCertId"`
	Cache    struct {
		AdaptCache   bool `json:"AdaptCache"`
		CacheKeyRule struct {
			IgnoreCase        bool   `json:"IgnoreCase"`
			ParamsFilterType  string `json:"ParamsFilterType"`
			ParamsReserveList string `json:"ParamsReserveList"`
		} `json:"CacheKeyRule"`
		CacheKeyRules []struct {
			Contents          string `json:"Contents"`
			IgnoreCase        bool   `json:"IgnoreCase"`
			ParamsFilterType  string `json:"ParamsFilterType"`
			ParamsReserveList string `json:"ParamsReserveList"`
			Type              string `json:"Type"`
		} `json:"CacheKeyRules"`
		Enable              bool `json:"Enable"`
		StatusCodeCacheRule struct {
			Enable bool `json:"Enable"`
			Rules  []struct {
				CacheTime     int    `json:"CacheTime"`
				CacheTimeUnit string `json:"CacheTimeUnit"`
				RespCode      string `json:"RespCode"`
				RespUrl       string `json:"RespUrl"`
				Status        string `json:"Status"`
			} `json:"Rules"`
		} `json:"StatusCodeCacheRule"`
	} `json:"Cache"`
	Domain       string `json:"Domain"`
	Cname        string `json:"Cname"`
	Status       string `json:"Status"`
	UserName     string `json:"UserName"`
	CreateTime   string `json:"CreateTime"`
	UpdateTime   string `json:"UpdateTime"`
	GzipCompress struct {
		Enable bool `json:"Enable"`
	} `json:"GzipCompress"`
	Origin struct {
		Origins []struct {
			Name   string `json:"Name"`
			Weight int    `json:"Weight"`
		} `json:"Origins"`
		OriginType         string `json:"OriginType"`
		OriginProtocolType string `json:"OriginProtocolType"`
		BackupOrigins      []struct {
			Name   string `json:"Name"`
			Weight int    `json:"Weight"`
		} `json:"BackupOrigins"`
		BackupOriginType string `json:"BackupOriginType"`
		ResponseTimeout  int    `json:"ResponseTimeout"`
		ResponseHeader   struct {
			Enable     bool          `json:"Enable"`
			HeaderInfo []interface{} `json:"HeaderInfo"`
		} `json:"ResponseHeader"`
		OriginHost struct {
			Enable   bool   `json:"Enable"`
			HostInfo string `json:"HostInfo"`
		} `json:"OriginHost"`
		OriginSni struct {
			Enable  bool   `json:"Enable"`
			SniInfo string `json:"SniInfo"`
		} `json:"OriginSni"`
		OriginRange struct {
			Enable bool   `json:"Enable"`
			Size   string `json:"Size"`
		} `json:"OriginRange"`
	} `json:"Origin"`
	HTTPS struct {
		HTTP2 bool `json:"Http2"`
		Hsts  struct {
			Enable           bool `json:"Enable"`
			MaxAge           int  `json:"MaxAge"`
			IncludeSubDomain bool `json:"IncludeSubDomain"`
		} `json:"Hsts"`
		TLSVersions struct {
			Enable     bool     `json:"Enable"`
			TLSVersion []string `json:"TlsVersion"`
		} `json:"TlsVersions"`
		CertBind struct {
			CertSource   string    `json:"CertSource"`
			CertID       string    `json:"CertId"`
			CertName     string    `json:"CertName"`
			DomainName   string    `json:"DomainName"`
			DomainID     string    `json:"DomainId"`
			DeployStatus string    `json:"DeployStatus"`
			Expire       time.Time `json:"Expire"`
		} `json:"CertBind"`
		ForceRedirect struct {
			Enable       bool   `json:"Enable"`
			RedirectType string `json:"RedirectType"`
		} `json:"ForceRedirect"`
		QUICSwitch bool `json:"QUICSwitch"`
	} `json:"Https"`
	StrategyType   string `json:"StrategyType"`
	EnableFailOver bool   `json:"EnableFailOver"`
	IPAccess       struct {
		Enable     bool          `json:"Enable"`
		FilterType string        `json:"FilterType"`
		FilterList []interface{} `json:"FilterList"`
	} `json:"IpAccess"`
	UserAgentAccess struct {
		Enable     bool          `json:"Enable"`
		FilterType string        `json:"FilterType"`
		FilterList []interface{} `json:"FilterList"`
	} `json:"UserAgentAccess"`
	RefererAccess struct {
		Enable     bool          `json:"Enable"`
		AllowNone  bool          `json:"AllowNone"`
		FilterType string        `json:"FilterType"`
		FilterList []interface{} `json:"FilterList"`
	} `json:"RefererAccess"`
	URLAccess struct {
		Enable  bool   `json:"Enable"`
		GenType string `json:"GenType"`
		GenKey  string `json:"GenKey"`
		GenTTL  int    `json:"GenTTL"`
	} `json:"UrlAccess"`
	IPv6Switch   bool   `json:"IPv6Switch"`
	RecordFiling string `json:"RecordFiling"`
	WebSocket    struct {
		Enable  bool `json:"Enable"`
		Timeout int  `json:"Timeout"`
	} `json:"WebSocket"`
	URLRedirect struct {
		Enable bool          `json:"Enable"`
		Rules  []interface{} `json:"Rules"`
	} `json:"UrlRedirect"`
}

type UploadSelfCertRequest struct {
	CertName   *string `json:",omitempty"`
	CertPEM    *string `json:",omitempty"`
	CertSource *string `json:",omitempty"` // volc
	KeyPEM     *string `json:",omitempty"`
}

type UploadSelfCertResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           UploadSelfCertResult
}

type UploadSelfCertResult struct {
	CertId string `json:"CertId"`
}

type CreateCertBindRequest struct {
	CertID string `json:"CertId"`

	DomainIds   []string `json:"DomainIds"`
	DomainNames []string `json:"DomainNames"`
}

type CreateCertBindResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
}
