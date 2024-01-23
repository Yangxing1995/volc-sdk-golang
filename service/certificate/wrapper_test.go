package certificate

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

var (
	_testAk, _testSk string
)

func init() {

	_testAk = os.Getenv("TEST_AK")
	_testSk = os.Getenv("TEST_SK")

	logrus.SetLevel(logrus.DebugLevel)
}

func TestCertificate_ImportCertificate(t *testing.T) {
	s := NewInstance()
	s.Client.SetAccessKey(_testAk)
	s.Client.SetSecretKey(_testSk)

	pem := ""
	key := ""
	gotResponseBody, err := s.ImportCertificate(&ImportCertificateRequest{
		CertificateInfo: &CertificateInfo{
			Certificate: pem,
			PrivateKey:  key,
		},
		Tag: "测试证书",
	})

	if err != nil {
		t.Fatal(err)
	}

	jsonBts, _ := json.MarshalIndent(gotResponseBody, "", "  ")

	t.Logf("%s\n", string(jsonBts))
}

func TestCertificate_CertificateGetInstance(t *testing.T) {
	s := NewInstance()
	s.Client.SetAccessKey(_testAk)
	s.Client.SetSecretKey(_testSk)

	limit := int64(4)
	page := int64(1)
	gotResponseBody, err := s.CertificateGetInstance(&CertificateGetInstanceRequest{
		Limit: &limit,
		Page:  &page,
	})

	if err != nil {
		t.Fatal(err)
	}

	jsonBts, _ := json.MarshalIndent(gotResponseBody, "", "  ")

	t.Logf("%s\n", string(jsonBts))
}
