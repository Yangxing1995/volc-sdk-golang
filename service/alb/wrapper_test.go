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

	_testAk = os.Getenv("BS_AK")
	_testSk = os.Getenv("BS_SK")

	logrus.SetLevel(logrus.DebugLevel)
}

func TestALB_DescribeLoadBalancers(t *testing.T) {
	s := NewInstance()
	s.Client.SetAccessKey(_testAk)
	s.Client.SetSecretKey(_testSk)

	for _, r := range []string{"cn-shanghai", "cn-beijing", "cn-guangzhou", "ap-southeast-1", "cn-beijing2"} {
		page := int64(1)
		pageSize := int64(2)

		s.Client.ServiceInfo.Credentials.Region = r

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
