package alb

import (
	"encoding/json"
	"testing"
)

func Test_structToMap(t *testing.T) {

	t.Run("struct and map", func(t *testing.T) {
		type param struct {
			Name   string `json:"name"`
			Person struct {
				UserName string `json:"user_name"`
				Age      int    `json:"age"`
			}

			Cats []struct {
				Name string `json:"name"`
				Age  int    `json:"age"`
			}
		}

		p := param{
			Name: "test",
			Person: struct {
				UserName string `json:"user_name"`
				Age      int    `json:"age"`
			}{
				UserName: "test",
				Age:      10,
			},
			Cats: []struct {
				Name string `json:"name"`
				Age  int    `json:"age"`
			}{
				{
					Name: "cat1",
					Age:  1,
				},
				{
					Name: "cat2",
					Age:  2,
				},
			},
		}

		m := structToMap(p)

		bts, _ := json.MarshalIndent(m, "", "  ")
		t.Log(string(bts))
	})

	t.Run("modifyListener", func(t *testing.T) {

		certID := "cert1"
		caCertID := ""
		d := DomainExtensionReq{
			DomainExtensionId: "d1",
			CertificateId:     "sniCert1",
			Action:            "create",
		}

		p := &ModifyListenerAttributesRequest{
			ListenerId:       "l1",
			CertificateId:    &certID,
			CACertificateId:  &caCertID,
			DomainExtensions: []*DomainExtensionReq{&d},
		}

		m := structToMap(p)

		bts, _ := json.MarshalIndent(m, "", "  ")
		t.Log(string(bts))
	})

}
