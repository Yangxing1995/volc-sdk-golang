package certificate

import (
	"net/http"
	"strings"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	DefaultRegion  = "cn-north-1"
	ServiceVersion = "2021-06-01"
	ServiceName    = "certificate_service"
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

type Certificate struct {
	Client *base.Client
}

func NewInstance() *Certificate {
	instance := new(Certificate)
	instance.Client = base.NewClient(ServiceInfo[DefaultRegion], ApiInfoList)
	instance.Client.ServiceInfo.Credentials.Service = ServiceName
	instance.Client.ServiceInfo.Credentials.Region = DefaultRegion
	return instance
}

func (s *Certificate) GetServiceInfo(region string) *base.ServiceInfo {
	if serviceInfo, ok := ServiceInfo[region]; ok {
		return serviceInfo
	}
	return nil
}

func (s *Certificate) GetAPIInfo(api string) *base.ApiInfo {
	if apiInfo, ok := ApiInfoList[api]; ok {
		return apiInfo
	}
	return nil
}

func (s *Certificate) GetMethod(api string) string {
	if apiInfo, ok := ApiInfoList[api]; ok {
		return apiInfo.Method
	}
	return ""
}

func (s *Certificate) SetRegion(region string) {
	s.Client.ServiceInfo.Credentials.Region = region
}

func (s *Certificate) SetHost(host string) {
	s.Client.ServiceInfo.Host = host
}

// SetSchema .
func (s *Certificate) SetSchema(schema string) {
	s.Client.ServiceInfo.Scheme = schema
}

func (s *Certificate) SetMethod(api, method string) bool {
	if _, ok := ApiInfoList[api]; ok {
		ApiInfoList[api].Method = strings.ToUpper(method)
		return true
	}
	return false
}

func (s *Certificate) SetAPIRetrySetting(apis []string, retryTime int, interval time.Duration) bool {
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

func (s *Certificate) SetNewAPI(info base.ApiInfo) {
	if len(info.Query.Get("Action")) <= 0 {
		return
	}
	action := info.Query.Get("Action")
	if _, ok := s.Client.ApiInfoList[action]; ok {
		return
	}
	s.Client.ApiInfoList[action] = &info
}