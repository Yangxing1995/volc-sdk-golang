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

func TestCertificate_CertificateGetInstanceList(t *testing.T) {
	s := NewInstance()
	s.Client.SetAccessKey(_testAk)
	s.Client.SetSecretKey(_testSk)

	limit := int64(4)
	page := int64(1)
	gotResponseBody, err := s.CertificateGetInstanceList(&CertificateGetInstanceListRequest{
		PageSize:   &limit,
		PageNumber: &page,
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

	gotResponseBody, err := s.CertificateGetInstance(&CertificateGetInstanceRequest{
		InstanceId: "cert-63412976b4864fd49dd33acf5c80984b",
	})

	if err != nil {
		t.Fatal(err)
	}

	jsonBts, _ := json.MarshalIndent(gotResponseBody, "", "  ")

	t.Logf("%s\n", string(jsonBts))
}
