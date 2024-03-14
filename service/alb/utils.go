package alb

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

type CDNError struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
	Status  int64  `json:"Status"`
}

func (e CDNError) Error() string {
	return fmt.Sprintf("status: %d, code: %s, message: %s", e.Status, e.Code, e.Message)
}

func wrapperError(errObj *ErrorObj) (err error) {
	if errObj == nil {
		return
	}
	err = CDNError{
		Message: errObj.Message,
		Status:  errObj.CodeN,
		Code:    errObj.Code,
	}
	return
}

func validateResponse(meta *ResponseMetadata) (err error) {
	if meta == nil {
		return errors.New("response meta is not found")
	}
	err = wrapperError(meta.Error)
	return
}

type OptionArg struct {
	useGet bool
}

func UseGet() OptionArg {
	return OptionArg{
		useGet: true,
	}
}

func (s *ALB) post(apiName string, requestBody interface{}, responseBody interface{}, options ...OptionArg) (err error) {
	var body []byte
	if requestBody == nil {
		requestBody = struct{}{}
	}
	body, err = json.Marshal(requestBody)
	if err != nil {
		err = fmt.Errorf("marshal request body failed, %w", err)
		return
	}
	query := url.Values{}
	option := OptionArg{useGet: false}
	option.useGet = s.GetMethod(apiName) == http.MethodGet

	if len(options) > 0 {
		option = options[0]
	}

	if option.useGet {
		if s.SetMethod(apiName, http.MethodGet) {
		} else {
			err = fmt.Errorf("set method falied")
			return
		}
		query, err = MergeQueryArgs(requestBody, query)
		if err != nil {
			return
		}
	}

	var data []byte
	if option.useGet {
		data, _, err = s.Client.Query(apiName, query)
	} else {
		data, _, err = s.Client.Json(apiName, query, string(body))
	}

	if err != nil {
		err = fmt.Errorf("request %s api failed: %w", apiName, err)
		return
	}
	if err = json.Unmarshal(data, responseBody); err != nil {
		err = fmt.Errorf("unmarshal response body failed: %w", err)
		return
	}
	return
}

func GetStrPtr(str string) *string {
	return &str
}

func MergeQueryArgs(body interface{}, query url.Values) (url.Values, error) {
	data := structToMap(body)

	for k, v := range data {
		query.Set(k, fmt.Sprintf("%s", v))
	}

	return query, nil
}

func structToMap(obj interface{}) map[string]string {
	result := make(map[string]string)
	valueOf := reflect.ValueOf(obj)

	if valueOf.Kind() == reflect.Ptr {
		valueOf = valueOf.Elem()
	}

	typeOf := valueOf.Type()

	for i := 0; i < valueOf.NumField(); i++ {
		field := valueOf.Field(i)
		fieldType := typeOf.Field(i)

		tag := fieldType.Tag.Get("json")

		// 字段申明了omitempty,且为零值,则跳过
		if tagParts := strings.Split(tag, ","); len(tagParts) > 1 && tagParts[1] == "omitempty" && isEmptyValue(field) {
			continue
		}

		switch field.Kind() {
		case reflect.Struct:
			subMap := structToMap(field.Interface())
			for key, value := range subMap {
				result[fieldType.Name+"."+key] = value
			}
		case reflect.Slice:
			// 如果是切片类型，递归处理每个元素
			for j := 0; j < field.Len(); j++ {
				subMap := structToMap(field.Index(j).Interface())
				for key, value := range subMap {
					result[fieldType.Name+"."+strconv.Itoa(j+1)+"."+key] = value
				}
			}
		case reflect.Ptr:
			// 如果是指针类型，判断具体类型处理
			if field.IsNil() {
				continue
			}
			elem := field.Elem()
			switch elem.Kind() {
			case reflect.Struct:
				// 如果是结构体指针类型，递归处理
				subMap := structToMap(elem.Interface())
				for key, value := range subMap {
					result[fieldType.Name+"."+key] = value
				}
			case reflect.Slice:
				// 如果是切片指针类型，递归处理每个元素
				for j := 0; j < field.Elem().Len(); j++ {
					subMap := structToMap(field.Elem().Index(j).Interface())
					for key, value := range subMap {
						result[fieldType.Name+"."+strconv.Itoa(j+1)+"."+key] = value
					}
				}

			default:
				// 其他类型直接取值
				result[fieldType.Name] = fmt.Sprintf("%v", elem.Interface())
			}
		default:
			// 其他类型直接添加到结果中
			result[fieldType.Name] = fmt.Sprintf("%v", field.Interface())
		}
	}

	return result
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	default:
		return false
	}
}
