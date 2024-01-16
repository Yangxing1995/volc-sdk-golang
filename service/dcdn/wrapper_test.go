package dcdn

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

var (
	_testAk string
	_testSk string
)

func init() {

	_testAk = os.Getenv("TEST_AK")
	_testSk = os.Getenv("TEST_SK")

	logrus.SetLevel(logrus.DebugLevel)
}

func TestDCDN_DescribeUserDomains(t *testing.T) {
	s := NewInstance()
	s.Client.SetAccessKey(_testAk)
	s.Client.SetSecretKey(_testSk)

	page := int64(1)
	pageSize := int64(2)

	gotResponseBody, err := s.DescribeUserDomains(&DescribeUserDomainsRequest{
		PageNum:  &page,
		PageSize: &pageSize,
	})
	if err != nil {
		t.Errorf("DCDN.DescribeUserDomains() error = %v", err)
		return
	}

	jsonBts, _ := json.MarshalIndent(gotResponseBody, "", "  ")

	t.Logf("%s\n", string(jsonBts))
}

func TestDCDN_UploadSelfCert(t *testing.T) {
	s := NewInstance()
	s.Client.SetAccessKey(_testAk)
	s.Client.SetSecretKey(_testSk)

	pem := ``
	key := ``
	gotResponseBody, err := s.UploadSelfCert(&UploadSelfCertRequest{
		CertName: GetStrPtr("yx759.fun"),
		CertPEM:  GetStrPtr(pem),
		KeyPEM:   GetStrPtr(key),
	})

	if err != nil {
		t.Fatal(err)
	}

	jsonBts, _ := json.MarshalIndent(gotResponseBody, "", "  ")

	t.Logf("%s\n", string(jsonBts))
}

func TestDCDN_DescribeDomainConfig(t *testing.T) {
	s := NewInstance()
	s.Client.SetAccessKey(_testAk)
	s.Client.SetSecretKey(_testSk)

	gotResponseBody, err := s.DescribeDomainConfig(&DescribeDomainConfigRequest{
		Domains: []string{"fast2.ldlb.site"},
	})
	if err != nil {
		t.Errorf("DCDN.DescribeDomainConfig() error = %v", err)
		return
	}

	jsonBts, _ := json.MarshalIndent(gotResponseBody, "", "  ")

	t.Logf("%s\n", string(jsonBts))

}
