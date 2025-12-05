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

	ok := true
	pem := `-----BEGIN CERTIFICATE-----
MIIDqjCCApKgAwIBAgIIJpn9ADtNNQMwDQYJKoZIhvcNAQELBQAwejELMAkGA1UE
BhMCQ04xFzAVBgNVBAoTDktleU1hbmFnZXIub3JnMTEwLwYDVQQLEyhLZXlNYW5h
Z2VyIFRlc3QgUm9vdCAtIEZvciBUZXN0IFVzZSBPbmx5MR8wHQYDVQQDExZLZXlN
YW5hZ2VyIFRlc3QgUlNBIENBMB4XDTI1MTIwNTAyMzY1MloXDTI2MTIwNTAyMzY1
MlowIzELMAkGA1UEBhMCQ04xFDASBgNVBAMTC2V4YW1wbGUuY29tMIIBIjANBgkq
hkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwdheoq47lLv1uCT6N7vznm4DtfWCXICl
LqB6RoIDEhRU1KgkTwHDPM3DtWl1FOetY3m3yEaKJXlXxU1qFHUde3PRzM/mDtsR
SHdHIRdbwG2OG55YMK1KSqitBYC1LLLykLs27RZTD/LyEndwlBbj0Hk0dFXAn5uz
T8tMRpYLtSE/tg2dlWXufc+oCrXJij4+5VxP8y7tAg3yoYxUg1PXBDVOamzzAzT7
+6twAeclKSeJh0QduwO30ir2nk0QQqh5CCMyLYxQyHC6LcH7DSbHecBzdz9GYt0G
5BZR78qRUqpRE/W+XSaQH9faBgUTQJRbqL+qNmq4l/SVJ1KQQsJAEwIDAQABo4GK
MIGHMA4GA1UdDwEB/wQEAwIFoDAdBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUH
AwIwHQYDVR0OBBYEFGIPuz78JLdXVXuZrplEyj+lxU5LMB8GA1UdIwQYMBaAFOY/
ZWRjMmm7mgoHwof9NkwmsLEWMBYGA1UdEQQPMA2CC2V4YW1wbGUuY29tMA0GCSqG
SIb3DQEBCwUAA4IBAQA4+HKD6uxmxc/ZRg2X81yT7IA1J59R7z4/d75BErETooYZ
xpixDDALfEoFoOdcEiFVT+iPgDh5uMoQBuDGz3yAkY8aBNJrHkQVZX8ffCcwcfkF
AfoZgvdtWYKMoKTk5Vr3OUWyDYAkxRb8aLJ5E5S/thU8l7JDAEA5qXVoKU32uZIM
OGjYAtQbVhbB25P2BBh/HrYaeuBkiOEuj6yo7+kCj502+UIAC+2vAICC+IrcoI97
d2iRBmlZIOdAmayJwsBp0uccEIIKARjNJ5Bu6sLVtWsI0oATCfMb0KN6H1TvDbJF
JOTGsi0gldiLkdjHT32MpSyIUM/uK3nJPWXlL2VA
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIID2jCCAsKgAwIBAgIILGel7/K6JaYwDQYJKoZIhvcNAQELBQAwezELMAkGA1UE
BhMCQ04xFzAVBgNVBAoTDktleU1hbmFnZXIub3JnMTEwLwYDVQQLEyhLZXlNYW5h
Z2VyIFRlc3QgUm9vdCAtIEZvciBUZXN0IFVzZSBPbmx5MSAwHgYDVQQDExdLZXlN
YW5hZ2VyIFRlc3QgUm9vdCBDQTAeFw0yNTA1MTMwODMyMTlaFw0zNTA1MTMwODMy
MTlaMHoxCzAJBgNVBAYTAkNOMRcwFQYDVQQKEw5LZXlNYW5hZ2VyLm9yZzExMC8G
A1UECxMoS2V5TWFuYWdlciBUZXN0IFJvb3QgLSBGb3IgVGVzdCBVc2UgT25seTEf
MB0GA1UEAxMWS2V5TWFuYWdlciBUZXN0IFJTQSBDQTCCASIwDQYJKoZIhvcNAQEB
BQADggEPADCCAQoCggEBAO368YnvNt6cLWZpkmeGkEi7ChlM/SO4/KIVaKG2JBwu
skM1yaoNDMNZXfTjdTQKLDa3FkqTV8XFacahCSTT5bqefkdu2KFnvLOndDcS8mM7
Paro6m1QXriFc8PlgLsbZytDCM3OWxEHkD0IqYPNB0PIl/QxOFLhoOxFHDxVF43l
0Q+odjBSAAvFrLHDtpm4C7AZ/7mPePS/Wuy9qlWImltyHudNRDjaH9c9Mzwbcn2Q
TKnSvh3AuaFZ5c7rym+dPTW2DhIaV7zT78+DozSYjFqjDCvMEyEyN26K6roG9/v7
SvPq30rmDVWhWJe0wESbJ5u9Mitzja56UcUIUnUGPlMCAwEAAaNjMGEwDgYDVR0P
AQH/BAQDAgGGMA8GA1UdEwEB/wQFMAMBAf8wHQYDVR0OBBYEFOY/ZWRjMmm7mgoH
wof9NkwmsLEWMB8GA1UdIwQYMBaAFAWzNuVs4IzyRdQBvqasULa1kEAYMA0GCSqG
SIb3DQEBCwUAA4IBAQBKeJCMGcpWIedR/OfXgjt7R+Ic7ifb5N1Ul6IFMZafSMJn
0Tjg2l/aUkCJ6qANr9RfHG5Ro8jUnCrXnTQ3Yzr3hI0dnF0ZYQBj1P6LAEJ8Dv9l
JROK0EZGgxRlV5X1/94Gf3MHNWAMp8paFE4PZuHdWKia2xwl2dqHsgOQ/v04I7vi
XzD+DcqIGSPLsd7ARD7a3n9njsFdMa4ZiLGH3mxm4NbudT0NffB3xqeRqJ5bmoOG
s3ivtUY9cdTEG7sltLvt4x/bpt7mWpQ/SEFL92iwKkK03l48W1JQKycJjdINyeUW
yHsXEBG7TiyPV++B1WGU8K7q5qEsE7N2aor6JGE2
-----END CERTIFICATE-----`
	key := `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAwdheoq47lLv1uCT6N7vznm4DtfWCXIClLqB6RoIDEhRU1Kgk
TwHDPM3DtWl1FOetY3m3yEaKJXlXxU1qFHUde3PRzM/mDtsRSHdHIRdbwG2OG55Y
MK1KSqitBYC1LLLykLs27RZTD/LyEndwlBbj0Hk0dFXAn5uzT8tMRpYLtSE/tg2d
lWXufc+oCrXJij4+5VxP8y7tAg3yoYxUg1PXBDVOamzzAzT7+6twAeclKSeJh0Qd
uwO30ir2nk0QQqh5CCMyLYxQyHC6LcH7DSbHecBzdz9GYt0G5BZR78qRUqpRE/W+
XSaQH9faBgUTQJRbqL+qNmq4l/SVJ1KQQsJAEwIDAQABAoIBABW0TlRUweMTcJ3f
tgyLhOmWU+AJz0DgPp6L5WUDpahekmkFMpuODPg/6cvkNPHURvu+Yy+PhmRPw94Z
nukPe++YtkMCfDXpnXPJtzMjVGilW6YIAsPgg1L8s/LLyH1qw/IkpC/DsiU1BJnA
ueA69+Mdzxu/hkPYwSkMEExO/JyZHGcwDdYxIi+Xh7vkHnFVbXtOrxLaxjstKxtY
EdsyH1AT2o072b4LhSZI9Va4yY3V5Yxopt47iC+0meakBw/cdhcwB0ok6VONByQB
TicgaPHMUWKs9D32GDN6dchPlfGcsS1Qwgt2MtcF8DdcmAE3vVQO0R973aDPNYGN
reOgj4ECgYEAzZMT4x53sHAe6mgwJjq1QxvR5XU38VVbjDAG5L/4QZhKjhTG8DYG
4BvH2OFy7hiJxEZwpOcirYGytoM8NLppBUoMq2aqd2LVPMnGoJrEJpRP0HQ+Dy+Q
2vQ5Rlg5LqfW2Hov1pa7b0Uk1re632Xt9D2iH19HqjHy2TOYds0hZ/MCgYEA8WTB
mDJy3oIDYF4UPm/YSbLtnvpgcXRXMio9vK3wHUOd6I5cS7f+nvx+FSY9pHLOojdU
cGyHtxnPM8g/Mp0geWgNgk4qpEht+SIZj1H9l/U8DkoFyecXxA8HflfRFHqXFtjr
rODar6N7JpDlQxJExXsEzR1vNd3rL64fAwQ672ECgYEAsbq6st0gdXY6BEbXefW/
yKlAqrTRDODTKTxMeXBO3ccgjf4AkKTlBNR/z03640ZrtTmBIdCC9qbvL0P+xUd0
jJIcqgM01cqouMlp5GhptUD8UChvjwl8nxAOhHa52VBaOwZIXaormmOZGdn2wjWJ
qGOwypbBiq9sWcR1yCBJBWkCgYB7RQOjsCd+vw/wVzqpQ75ErG2x4JP9e/bHNFEA
mCZlG1F5H7G8LQKGt9O/Ax85ajNhHTYujw2jEDgCboDvXNPni55Sa45VBHO8ZJB+
4tTIk8AOX/iiyTMWuMIP39JyTcP4M9/uEYk7ZhWOdTHOVXbvDco+4IgZDvF6EMXf
aXqbIQKBgQCsusgcZtnD2Kz0cba+keSZYr7831g5Mq2x40U1uy3iriGysPy8idbV
Rc4N6cLdJ/E36JuK6p6Piq1eeHhLj2RZfobWujTfk0mlO6BeldTktFK5GzxnSad1
tlsiB3Rj9A2gHtvc/6PzMKbMZgFBxpgfO+F8yy8ADyZE/dkGaXjI/A==
-----END RSA PRIVATE KEY-----`
	gotResponseBody, err := s.ImportCertificate(&ImportCertificateRequest{
		Repeatable: &ok,
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
