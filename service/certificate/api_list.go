package certificate

import (
	"net/http"
	"net/url"

	"github.com/volcengine/volc-sdk-golang/base"
)

var ApiInfoList = map[string]*base.ApiInfo{
	"ImportCertificate": {
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"ImportCertificate"},
			"Version": []string{ServiceVersion},
		},
	},
	"CertificateGetInstance": {
		Method: http.MethodGet,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{"CertificateGetInstance"},
			"Version": []string{ServiceVersion},
		},
	},
}
