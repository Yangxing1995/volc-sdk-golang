package clb

import (
	"net/http"
	"strings"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	DefaultRegion  = "cn-beijing"
	ServiceVersion = "2020-04-01"
	ServiceName    = "CLB"
)

var (
	AllRegions = []string{
		DefaultRegion,
		"cn-shanghai",
		"cn-guangzhou",
	}
)

var (
	ServiceInfo = map[string]*base.ServiceInfo{
		DefaultRegion: {
			Host:    "open.volcengineapi.com",
			Timeout: time.Minute * 5,
			Header: http.Header{
				"Accept":       []string{"application/json"},
				"Content-Type": []string{"application/json"},
			},
		},
	}
)

// DefaultInstance Package level default instance
var DefaultInstance = NewInstance()

type CLB struct {
	Client *base.Client
}

func NewInstance() *CLB {
	instance := new(CLB)
	instance.Client = base.NewClient(ServiceInfo[DefaultRegion], ApiInfoList)
	instance.Client.ServiceInfo.Credentials.Service = ServiceName
	instance.Client.ServiceInfo.Credentials.Region = DefaultRegion
	return instance
}

func (s *CLB) GetServiceInfo(region string) *base.ServiceInfo {
	if serviceInfo, ok := ServiceInfo[region]; ok {
		return serviceInfo
	}
	return nil
}

func (s *CLB) GetAPIInfo(api string) *base.ApiInfo {
	if apiInfo, ok := ApiInfoList[api]; ok {
		return apiInfo
	}
	return nil
}

func (s *CLB) GetMethod(api string) string {
	if apiInfo, ok := ApiInfoList[api]; ok {
		return apiInfo.Method
	}
	return ""
}

func (s *CLB) SetRegion(region string) {
	s.Client.ServiceInfo.Credentials.Region = region
}

func (s *CLB) SetHost(host string) {
	s.Client.ServiceInfo.Host = host
}

// SetSchema .
func (s *CLB) SetSchema(schema string) {
	s.Client.ServiceInfo.Scheme = schema
}

func (s *CLB) SetMethod(api, method string) bool {
	if _, ok := ApiInfoList[api]; ok {
		ApiInfoList[api].Method = strings.ToUpper(method)
		return true
	}
	return false
}

func (s *CLB) SetAPIRetrySetting(apis []string, retryTime int, interval time.Duration) bool {
	if len(apis) <= 0 {
		return false
	}
	_retryTime, _interval := uint64(retryTime), interval
	for _, api := range apis {
		if _, ok := s.Client.ApiInfoList[api]; ok {
			s.Client.ApiInfoList[api].Retry = base.RetrySettings{
				AutoRetry:     true,
				RetryTimes:    &_retryTime,
				RetryInterval: &_interval,
			}
		} else {
			return false
		}
	}
	s.Client.SetRetrySettings(&base.RetrySettings{AutoRetry: true})
	return true
}

func (s *CLB) SetNewAPI(info base.ApiInfo) {
	if len(info.Query.Get("Action")) <= 0 {
		return
	}
	action := info.Query.Get("Action")
	if _, ok := s.Client.ApiInfoList[action]; ok {
		return
	}
	s.Client.ApiInfoList[action] = &info
}
