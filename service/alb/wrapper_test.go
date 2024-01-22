package alb

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

func TestCLB_DescribeLoadBalancers(t *testing.T) {
	s := NewInstance()
	s.Client.SetAccessKey(_testAk)
	s.Client.SetSecretKey(_testSk)

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

func TestALB_DescribeCertificates(t *testing.T) {
	s := NewInstance()
	s.Client.SetAccessKey(_testAk)
	s.Client.SetSecretKey(_testSk)
	s.Client.ServiceInfo.Credentials.Region = "cn-shanghai"

	page := int64(1)
	pageSize := int64(2)

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
