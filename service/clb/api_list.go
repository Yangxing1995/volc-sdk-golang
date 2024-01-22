package clb

import (
	"net/http"
	"net/url"

	"github.com/volcengine/volc-sdk-golang/base"
)

var ApiInfoList = map[string]*base.ApiInfo{
	"DescribeLoadBalancers": {
		Method: http.MethodGet,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"DescribeLoadBalancers"},
			"Version": []string{ServiceVersion},
		},
	},
	"DescribeLoadBalancerAttributes": {
		Method: http.MethodGet,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"DescribeLoadBalancerAttributes"},
			"Version": []string{ServiceVersion},
		},
	},
	"DescribeListeners": {
		Method: http.MethodGet,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"DescribeListeners"},
			"Version": []string{ServiceVersion},
		},
	},
	"ModifyListenerAttributes": {
		Method: http.MethodGet,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"ModifyListenerAttributes"},
			"Version": []string{ServiceVersion},
		},
	},
	"DescribeRules": {
		Method: http.MethodGet,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"DescribeRules"},
			"Version": []string{ServiceVersion},
		},
	},
	"UploadCertificate": {
		Method: http.MethodGet,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"UploadCertificate"},
			"Version": []string{ServiceVersion},
		},
	},
	"DescribeCertificates": {
		Method: http.MethodGet,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"DescribeCertificates"},
			"Version": []string{ServiceVersion},
		},
	},
}
