package certificate

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

type CertificateGetInstanceRequest struct {
	Limit            *int64 `json:",omitempty"` // max100
	Offset           *int64 `json:",omitempty"` // default 0
	Page             *int64 `json:",omitempty"` // 设置要返回的证书所在的页码。默认值：1。 该参数必须与limit同时使用。 page与offset二选一。如果您同时设置了page和offset，则只有page会生效。
	CertificateExist *bool  `json:",omitempty"` // 是否只返回已签发的证书。该参数有以下取值： true：是  false：否

}

type CertificateGetInstanceResponse struct {
	ResponseMetadata *ResponseMetadata       `json:",omitempty"`
	Result           *CertificateGetInstance `json:",omitempty"`
}

type CertificateGetInstance struct {
	Count   int                  `json:"count"`
	Content []*CertificateResult `json:"content"`
}

type CertificateResult struct {
	Id                   string      `json:"id"`
	ParentId             string      `json:"parent_id"`
	ChainId              string      `json:"chain_id"`
	Upstream             string      `json:"upstream"`
	UpstreamFilter       string      `json:"upstream_filter"`
	Deleted              int         `json:"deleted"`
	SourceParentId       string      `json:"source_parent_id"`
	SourceId             string      `json:"source_id"`
	Number               int         `json:"number"`
	CertType             int         `json:"cert_type"`
	InstanceType         int         `json:"instance_type"`
	Disabled             int         `json:"disabled"`
	Tag                  string      `json:"tag"`
	Purpose              string      `json:"purpose"`
	Type                 string      `json:"type"`
	CommonName           string      `json:"common_name"`
	ApplicableDomains    string      `json:"applicable_domains"`
	OrderExist           int         `json:"order_exist"`
	OrderBrand           string      `json:"order_brand"`
	OrderSanNumber       int         `json:"order_san_number"`
	OrderStatus          int         `json:"order_status"`
	OrderRequireProgress int         `json:"order_require_progress"`
	CertificateExist     int         `json:"certificate_exist"`
	CertificateRevoked   int         `json:"certificate_revoked"`
	DeployInfo           interface{} `json:"deploy_info"`
	ProjectName          string      `json:"project_name"`
	Issuer               string      `json:"issuer"`
	IsSm                 bool        `json:"is_sm"`
	OrderPeriod          int         `json:"order_period"`
	OrderOrganization    struct {
		Department           string `json:"department"`
		Name                 string `json:"name"`
		PostalCode           string `json:"postal_code"`
		Address              string `json:"address"`
		City                 string `json:"city"`
		Province             string `json:"province"`
		Country              string `json:"country"`
		Email                string `json:"email"`
		Phone                string `json:"phone"`
		BankAccountLicenseNo string `json:"bank_account_license_no"`
		BusinessLicenseNo    string `json:"business_license_no"`
		Contact              struct {
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Email     string `json:"email"`
			Phone     string `json:"phone"`
			Title     string `json:"title"`
			IdCardNo  string `json:"id_card_no"`
		} `json:"contact"`
		UpstreamExtension interface{} `json:"upstream_extension"`
	} `json:"order_organization"`
	Ssl struct {
		OrderSanNumber       int    `json:"order_san_number"`
		OrderValidationType  string `json:"order_validation_type"`
		OrderValidationReady bool   `json:"order_validation_ready"`
		OrderValidations     struct {
			WwwExampleCom struct {
				Key       []string `json:"key"`
				Value     string   `json:"value"`
				Validated bool     `json:"validated"`
			} `json:"www.example.com"`
		} `json:"order_validations"`
		San         []string `json:"san"`
		Certificate struct {
			Csr          string      `json:"csr"`
			PrivateKey   string      `json:"private_key"`
			KeyType      string      `json:"key_type"`
			Chain        []string    `json:"chain"`
			EncryptKey   string      `json:"encrypt_key"`
			EncryptChain interface{} `json:"encrypt_chain"`
		} `json:"certificate"`
	} `json:"ssl"`
	Log []struct {
		Time    string `json:"time"`
		Title   string `json:"title"`
		Content string `json:"content"`
	} `json:"log"`
	UpstreamExtension struct {
		Plan            string `json:"plan"`
		SubOrderId      string `json:"sub_order_id"`
		IncomeConfirmed bool   `json:"income_confirmed"`
		Dns             []struct {
			RecordID string `json:"RecordID"`
			ZoneName string `json:"ZoneName"`
			Line     string `json:"Line"`
			Host     string `json:"Host"`
			Type     string `json:"Type"`
			Value    string `json:"Value"`
		} `json:"dns"`
		IsCompleted bool `json:"is_completed"`
	} `json:"upstream_extension"`
	OrderProgressTime           string            `json:"order_progress_time"`
	CertificateNotBefore        string            `json:"certificate_not_before"`
	CertificateNotAfter         string            `json:"certificate_not_after"`
	CertificateNotAfterMs       int64             `json:"certificate_not_after_ms"`
	CertificateNotBeforeMs      int64             `json:"certificate_not_before_ms"`
	CertificateDetail           CertificateDetail `json:"certificate_detail"`
	EncryptionCertificateDetail CertificateDetail `json:"encryption_certificate_detail"` // SM2 国密
}

type CertificateDetail struct {
	Subject struct {
		CommonName       string   `json:"common_name"`
		San              []string `json:"san"`
		Organization     string   `json:"organization"`
		OrganizationUnit string   `json:"organization_unit"`
		Country          string   `json:"country"`
		Province         string   `json:"province"`
		City             string   `json:"city"`
		Address          string   `json:"address"`
		PostalCode       string   `json:"postal_code"`
	} `json:"subject"`
	Issuer struct {
		CommonName       string      `json:"common_name"`
		San              interface{} `json:"san"`
		Organization     string      `json:"organization"`
		OrganizationUnit string      `json:"organization_unit"`
		Country          string      `json:"country"`
		Province         string      `json:"province"`
		City             string      `json:"city"`
		Address          string      `json:"address"`
		PostalCode       string      `json:"postal_code"`
	} `json:"issuer"`
	SerialNumber        string `json:"serial_number"`
	NotBefore           string `json:"not_before"`
	NotAfter            string `json:"not_after"`
	EncryptionAlgorithm string `json:"encryption_algorithm"`
	SignatureAlgorithm  string `json:"signature_algorithm"`
	FingerprintSha1     string `json:"fingerprint_sha_1"`
	FingerprintSha256   string `json:"fingerprint_sha_256"`
}
