package clb

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

	_testAk = os.Getenv("BS_AK")
	_testSk = os.Getenv("BS_SK")

	logrus.SetLevel(logrus.DebugLevel)
}

func TestCLB_DescribeLoadBalancers(t *testing.T) {
	s := NewInstance()
	s.Client.SetAccessKey(_testAk)
	s.Client.SetSecretKey(_testSk)
	// sh 41
	// bj 1
	// gz 0
	s.Client.ServiceInfo.Credentials.Region = "cn-guangzhou"

	page := int64(1)
	pageSize := int64(2)

	gotResponseBody, err := s.DescribeLoadBalancers(&DescribeLoadBalancersRequest{
		PageNum:  &page,
		PageSize: &pageSize,
	})
	if err != nil {
		t.Errorf("CLB.DescribeUserDomains() error = %v", err)
		return
	}

	jsonBts, _ := json.MarshalIndent(gotResponseBody, "", "  ")

	t.Logf("%s\n", string(jsonBts))
}

func TestCLB_UploadCertificate(t *testing.T) {
	s := NewInstance()
	s.Client.SetAccessKey(_testAk)
	s.Client.SetSecretKey(_testSk)

	pem := ""
	key := ""
	gotResponseBody, err := s.UploadCertificate(&UploadCertificateRequest{
		CertificateName: GetStrPtr("yx759.fun"),
		PublicKey:       GetStrPtr(pem),
		PrivateKey:      GetStrPtr(key),
	}, UseGet())

	if err != nil {
		t.Fatal(err)
	}

	jsonBts, _ := json.MarshalIndent(gotResponseBody, "", "  ")

	t.Logf("%s\n", string(jsonBts))
}

func TestCLB_DescribeListeners(t *testing.T) {
	s := NewInstance()
	s.Client.SetAccessKey(_testAk)
	s.Client.SetSecretKey(_testSk)
	s.Client.ServiceInfo.Credentials.Region = "cn-shanghai"

	page := int64(1)
	pageSize := int64(2)

	gotResponseBody, err := s.DescribeListeners(&DescribeListenersRequest{
		PageNum:  &page,
		PageSize: &pageSize,
	})
	if err != nil {
		t.Errorf("DescribeCertificates() error = %v", err)
		return
	}

	jsonBts, _ := json.MarshalIndent(gotResponseBody, "", "  ")

	t.Logf("%s\n", string(jsonBts))
}

func TestCLB_DescribeCertificates(t *testing.T) {
	s := NewInstance()
	s.Client.SetAccessKey(_testAk)
	s.Client.SetSecretKey(_testSk)
	s.Client.ServiceInfo.Credentials.Region = "cn-beijing"

	page := int64(1)
	pageSize := int64(100)

	gotResponseBody, err := s.DescribeCertificates(&DescribeCertificatesRequest{
		PageNumber: &page,
		PageSize:   &pageSize,
	})
	if err != nil {
		t.Errorf("DescribeCertificates() error = %v", err)
		return
	}

	jsonBts, _ := json.MarshalIndent(gotResponseBody, "", "  ")

	t.Logf("%s\n", string(jsonBts))
}
