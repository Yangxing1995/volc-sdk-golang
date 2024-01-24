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

	//_testAk = os.Getenv("TEST_AK")
	//_testSk = os.Getenv("TEST_SK")

	_testAk = os.Getenv("BS_AK")
	_testSk = os.Getenv("BS_SK")

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

	t.Run("ok", func(t *testing.T) {
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
	})

	t.Run("多个", func(t *testing.T) {
		s := NewInstance()
		s.Client.SetAccessKey(_testAk)
		s.Client.SetSecretKey(_testSk)

		page := int64(1)
		pageSize := int64(100)

		gotResponseBody, err := s.DescribeUserDomains(&DescribeUserDomainsRequest{
			PageNum:  &page,
			PageSize: &pageSize,
		})

		if err != nil {
			t.Fatal(err)
		}

		t.Logf("%+v\n", gotResponseBody.Result.AllDomainNum)
		t.Logf("%+v\n", gotResponseBody.Result.OnlineDomainNum)

		domains := []string{}

		for _, v := range gotResponseBody.Result.Domains {
			t.Logf("%+v\n", v.Domain)
			domains = append(domains, v.Domain)
		}

		res, err := s.DescribeDomainConfig(&DescribeDomainConfigRequest{
			Domains: domains,
		})
		if err != nil {
			t.Error(err)
			return
		}

		t.Logf("%d\n", len(res.Result))

		for _, v := range res.Result {
			jsonBts, _ := json.MarshalIndent(v.HTTPS.CertBind, "", "  ")

			t.Logf("%s\n", string(jsonBts))
		}
	})
}
