package dcdn

import (
	"net/http"
	"net/url"

	"github.com/volcengine/volc-sdk-golang/base"
)

var ApiInfoList = map[string]*base.ApiInfo{
	"DescribeUserDomains": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"DescribeUserDomains"},
			"Version": []string{"2023-01-01"},
		},
	},
	"DescribeDomainDetail": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"DescribeDomainDetail"},
			"Version": []string{ServiceVersion},
		},
	},
	"ListDomainConfig": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"ListDomainConfig"},
			"Version": []string{ServiceVersion},
		},
	},
	"UploadSelfCert": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"UploadSelfCert"},
			"Version": []string{ServiceVersion},
		},
	},
	"CreateCertBind": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"CreateCertBind"},
			"Version": []string{ServiceVersion},
		},
	},
	"ListCertBind": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"ListCertBind"},
			"Version": []string{ServiceVersion},
		},
	},
}
