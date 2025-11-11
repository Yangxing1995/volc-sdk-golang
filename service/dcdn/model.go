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
	PageNumber int      `json:"PageNumber,omitempty"` // 分页参数，页码，大于0小于100000，不输入时默认为1。
	PageSize   int      `json:"PageSize,omitempty"`   // 分页参数，页大小，大于0小于10000，不输入或输入为0时默认为查询所有域名。
	Domains    []string `json:"Domains"`
}

type DescribeDomainConfigResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeDomainConfigResult
}

type DescribeDomainConfigResult struct {
	DomainList []*DomainConfig `json:"DomainList"`
	PageNumber int             `json:"PageNumber"`
	PageSize   int             `json:"PageSize"`
	Total      int             `json:"Total"`
}

type DomainConfig struct {
	Domain     string `json:"Domain"`
	Cname      string `json:"Cname"`
	Status     string `json:"Status"`
	UserName   string `json:"UserName"`
	CreateTime string `json:"CreateTime"`
	UpdateTime string `json:"UpdateTime"`
	Origin     struct {
		Origins []struct {
			Name     string `json:"Name"`
			Weight   int    `json:"Weight"`
			Type     string `json:"Type"`
			Strategy string `json:"Strategy"`
		} `json:"Origins"`
		OriginType         string `json:"OriginType"`
		OriginProtocolType string `json:"OriginProtocolType"`
		BackupOriginType   string `json:"BackupOriginType"`
		ResponseTimeout    int    `json:"ResponseTimeout"`
		ResponseHeader     struct {
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
		RequestHeader struct {
			Enable     bool          `json:"Enable"`
			HeaderInfo []interface{} `json:"HeaderInfo"`
		} `json:"RequestHeader"`
		FollowRedirect struct {
			Enable   bool `json:"Enable"`
			MaxTries int  `json:"MaxTries"`
		} `json:"FollowRedirect"`
		GAOriginProbe struct {
			Enable bool   `json:"Enable"`
			Host   string `json:"Host"`
			Url    string `json:"Url"`
		} `json:"GAOriginProbe"`
		OriginIPv6 struct {
			Enable   bool   `json:"Enable"`
			Strategy string `json:"Strategy"`
		} `json:"OriginIPv6"`
		VEFaaSConf struct {
			Enable          bool          `json:"Enable"`
			Region          string        `json:"Region"`
			FunctionWeights []interface{} `json:"FunctionWeights"`
		} `json:"VEFaaSConf"`
	} `json:"Origin"`
	Https struct {
		EnableHttps bool `json:"EnableHttps"`
		Http2       bool `json:"Http2"`
		Hsts        struct {
			Enable           bool `json:"Enable"`
			MaxAge           int  `json:"MaxAge"`
			IncludeSubDomain bool `json:"IncludeSubDomain"`
		} `json:"Hsts"`
		TlsVersions struct {
			Enable     bool     `json:"Enable"`
			TlsVersion []string `json:"TlsVersion"`
		} `json:"TlsVersions"`
		CertBind struct {
			CertSource     string    `json:"CertSource"`
			CertId         string    `json:"CertId"`
			CertName       string    `json:"CertName"`
			DomainName     string    `json:"DomainName"`
			DomainId       string    `json:"DomainId"`
			DeployStatus   string    `json:"DeployStatus"`
			Expire         time.Time `json:"Expire"`
			AccountId      string    `json:"AccountId"`
			CertResourceId string    `json:"CertResourceId"`
			CertStatus     string    `json:"CertStatus"`
		} `json:"CertBind"`
		ForceRedirect struct {
			Enable       bool   `json:"Enable"`
			RedirectType string `json:"RedirectType"`
			RedirectCode int    `json:"RedirectCode"`
		} `json:"ForceRedirect"`
		QUICSwitch   bool `json:"QUICSwitch"`
		EarlyData    bool `json:"EarlyData"`
		OCSPStapling bool `json:"OCSPStapling"`
	} `json:"Https"`
	StrategyType   string `json:"StrategyType"`
	EnableFailOver bool   `json:"EnableFailOver"`
	IpAccess       struct {
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
	UrlAccess struct {
		Enable  bool   `json:"Enable"`
		GenType string `json:"GenType"`
		GenKey  string `json:"GenKey"`
		GenTTL  int    `json:"GenTTL"`
	} `json:"UrlAccess"`
	IPv6Switch bool `json:"IPv6Switch"`
	Cache      struct {
		Enable     bool `json:"Enable"`
		CacheRules []struct {
			Type                  string `json:"Type"`
			Contents              string `json:"Contents"`
			CacheTime             int    `json:"CacheTime"`
			ParamsFilterType      string `json:"ParamsFilterType"`
			IgnoreCase            bool   `json:"IgnoreCase"`
			CacheTimeUnit         string `json:"CacheTimeUnit"`
			CacheTimeFollowOrigin bool   `json:"CacheTimeFollowOrigin"`
			ParamsReserveList     string `json:"ParamsReserveList"`
			Policy                string `json:"Policy"`
			IgnoreSetCookie       bool   `json:"IgnoreSetCookie"`
		} `json:"CacheRules,omitempty"`
		StatusCodeCacheRule struct {
			Enable bool `json:"Enable"`
			Rules  []struct {
				Status        string `json:"Status"`
				CacheTime     int    `json:"CacheTime"`
				CacheTimeUnit string `json:"CacheTimeUnit"`
				RespUrl       string `json:"RespUrl"`
				RespCode      string `json:"RespCode"`
			} `json:"Rules"`
		} `json:"StatusCodeCacheRule"`
		AdaptCache   bool `json:"AdaptCache"`
		CacheKeyRule struct {
			IgnoreCase        bool   `json:"IgnoreCase"`
			ParamsFilterType  string `json:"ParamsFilterType"`
			ParamsReserveList string `json:"ParamsReserveList"`
		} `json:"CacheKeyRule"`
		CacheKeyRules []struct {
			Type              string `json:"Type"`
			Contents          string `json:"Contents"`
			IgnoreCase        bool   `json:"IgnoreCase"`
			ParamsFilterType  string `json:"ParamsFilterType"`
			ParamsReserveList string `json:"ParamsReserveList"`
		} `json:"CacheKeyRules"`
	} `json:"Cache"`
	RecordFiling string `json:"RecordFiling"`
	WebSocket    struct {
		Enable  bool `json:"Enable"`
		Timeout int  `json:"Timeout"`
	} `json:"WebSocket"`
	UrlRedirect struct {
		Enable bool          `json:"Enable"`
		Rules  []interface{} `json:"Rules"`
	} `json:"UrlRedirect"`
	IsCNAMEResolved bool   `json:"IsCNAMEResolved"`
	Scope           string `json:"Scope"`
	VerifyClient    bool   `json:"VerifyClient"`
	CaCertId        string `json:"CaCertId"`
	GzipCompress    struct {
		Enable bool `json:"Enable"`
	} `json:"GzipCompress"`
	BrCompress struct {
		Enable bool `json:"Enable"`
	} `json:"BrCompress"`
	ProbeSetting struct {
		Host   string `json:"Host"`
		Url    string `json:"Url"`
		Switch string `json:"Switch"`
	} `json:"ProbeSetting"`
	ProjectName        string `json:"ProjectName"`
	StaticOptimization struct {
		H2Priority bool `json:"H2Priority"`
	} `json:"StaticOptimization"`
	ScheduleType struct {
		Type string `json:"Type"`
	} `json:"ScheduleType"`
	ServiceType string `json:"ServiceType"`
	RTTOptimize struct {
		Enable bool `json:"Enable"`
	} `json:"RTTOptimize"`
	PreConnect struct {
		Enable bool `json:"Enable"`
	} `json:"PreConnect"`
	UploadOptimize struct {
		Enable bool `json:"Enable"`
	} `json:"UploadOptimize"`
	ManualLockInfo struct {
		Reason string `json:"Reason"`
		Enable bool   `json:"Enable"`
	} `json:"ManualLockInfo"`
	IsBP bool `json:"IsBP"`
}

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

type ListCertBindRequest struct {
	SearchKey   string   `json:"SearchKey"`
	PageSize    int      `json:"PageSize"`
	PageNumber  int      `json:"PageNumber"`
	ProjectName []string `json:"ProjectName"`
}

type ListCertBindResult struct {
	BindList   []*BindInfo `json:"BindList"`
	PageNumber int         `json:"PageNumber"`
	PageSize   int         `json:"PageSize"`
	Total      int         `json:"Total"`
}

type BindInfo struct {
	CertSource   string    `json:"CertSource"`
	CertId       string    `json:"CertId"`
	CertName     string    `json:"CertName"` //证书的名称 不是证书的通用名称
	DomainName   string    `json:"DomainName"`
	DomainId     string    `json:"DomainId"`
	DeployStatus string    `json:"DeployStatus"`
	Expire       time.Time `json:"Expire"`
	CertStatus   string    `json:"CertStatus"`
}

type ListCertBindResponse struct {
	ResponseMetadata *ResponseMetadata   `json:"ResponseMetadata"`
	Result           *ListCertBindResult `json:"Result"`
}
