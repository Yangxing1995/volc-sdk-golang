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

func newSmokeALB(t *testing.T) *ALB {
	t.Helper()

	if _testAk == "" || _testSk == "" {
		t.Skip("set TEST_AK and TEST_SK to run ALB smoke tests")
	}

	region := os.Getenv("ALB_SMOKE_REGION")
	if region == "" {
		region = "cn-shanghai"
	}

	s := NewInstance()
	s.Client.SetAccessKey(_testAk)
	s.Client.SetSecretKey(_testSk)
	s.Client.ServiceInfo.Credentials.Region = region

	return s
}

func TestALB_DescribeLoadBalancers(t *testing.T) {
	s := NewInstance()
	s.Client.SetAccessKey(_testAk)
	s.Client.SetSecretKey(_testSk)

	for _, r := range []string{"cn-shanghai", "cn-beijing", "cn-guangzhou", "ap-southeast-1"} {
		page := int64(1)
		pageSize := int64(2)

		s.Client.ServiceInfo.Credentials.Region = r

		gotResponseBody, err := s.DescribeLoadBalancers(&DescribeLoadBalancersRequest{
			PageNumber: &page,
			PageSize:   &pageSize,
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

func TestALB_SmokeDescribeListeners(t *testing.T) {
	defaultLBID := "alb-3vu0pck2qbjsw6096psmpow31"
	//                 "LoadBalancerId": "alb-1vyoe1907xk3k3766qaiowrj6",
	//                 "LoadBalancerId": "alb-3vu0pck2qbjsw6096psmpow31",

	loadBalancerID := os.Getenv("ALB_SMOKE_LOAD_BALANCER_ID")
	if loadBalancerID == "" {
		loadBalancerID = defaultLBID
	}

	s := newSmokeALB(t)
	page := int64(1)
	pageSize := int64(100)

	gotResponseBody, err := s.DescribeListeners(&DescribeListenersRequest{
		PageNumber:     &page,
		PageSize:       &pageSize,
		LoadBalancerId: &loadBalancerID,
	})
	if err != nil {
		t.Fatalf("DescribeListeners() error = %v", err)
	}
	if gotResponseBody.Result.TotalCount == 0 || len(gotResponseBody.Result.Listeners) == 0 {
		t.Fatalf("DescribeListeners() returned no listeners for load balancer %s", loadBalancerID)
	}

	listenerID := os.Getenv("ALB_SMOKE_LISTENER_ID")
	if listenerID == "" {
		listenerID = "lsn-3vu0pcryurrpc6096ptxfbwzt"
	}

	for _, listener := range gotResponseBody.Result.Listeners {
		bts, _ := json.MarshalIndent(listener, "", " ")
		t.Log(string(bts))

		if listener.ListenerID != listenerID {
			continue
		}
		if listener.Protocol == "" {
			t.Fatalf("DescribeListeners() listener %s has empty Protocol", listenerID)
		}
		if listener.Port == 0 {
			t.Fatalf("DescribeListeners() listener %s has empty Port", listenerID)
		}
		if listener.ServerGroupID == "" {
			t.Fatalf("DescribeListeners() listener %s has empty ServerGroupID", listenerID)
		}
		return
	}

	t.Fatalf("DescribeListeners() did not return listener %s under load balancer %s", listenerID, loadBalancerID)
}

func TestALB_SmokeDescribeListenerAttributes(t *testing.T) {
	// 标准款的监听器
	//listenerID := "lsn-3vu0pcryurrpc6096ptxfbwzt"
	// 基础款的监听器， DomainExtensions.Domain 应该不会空
	listenerID := "lsn-2daxu5yqq9fy83rvq31aarz95"

	s := newSmokeALB(t)

	gotResponseBody, err := s.DescribeListenerAttributes(&DescribeListenerAttributesRequest{
		ListenerId: listenerID,
	})
	if err != nil {
		t.Fatalf("DescribeListenerAttributes() error = %v", err)
	}

	bts, _ := json.MarshalIndent(gotResponseBody, "", " ")
	t.Log(string(bts))

	if gotResponseBody.Result.ListenerID != listenerID {
		t.Fatalf("DescribeListenerAttributes() ListenerID = %s, want %s", gotResponseBody.Result.ListenerID, listenerID)
	}
	if gotResponseBody.Result.Protocol == "" {
		t.Fatalf("DescribeListenerAttributes() listener %s has empty Protocol", listenerID)
	}
	if gotResponseBody.Result.Port == 0 {
		t.Fatalf("DescribeListenerAttributes() listener %s has empty Port", listenerID)
	}
	if gotResponseBody.Result.LoadBalancerID == "" {
		t.Fatalf("DescribeListenerAttributes() listener %s has empty LoadBalancerID", listenerID)
	}
	if gotResponseBody.Result.ServerGroupID == "" {
		t.Fatalf("DescribeListenerAttributes() listener %s has empty ServerGroupID", listenerID)
	}
}
