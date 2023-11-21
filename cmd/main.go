package main

import (
	"crypto/tls"
	"github.com/gin-gonic/gin"
	"github.com/h3adex/fp"
	"log"
)

func main() {
	handle := gin.Default()
	//fp url in : https://127.0.0.1:8999/
	handle.GET("/", fp.GinHandlerFunc)
	err := fp.Server(nil, handle.Handler(), fp.Option{Addr: ":8999", GetCertificate: func(clientHello *tls.ClientHelloInfo) (*tls.Certificate, error) {
		certPem := []byte(`-----BEGIN CERTIFICATE-----
MIIDazCCAlOgAwIBAgIUWCyEkV65Dh6HTm8KfZL5d4vI8z0wDQYJKoZIhvcNAQEL
BQAwRTELMAkGA1UEBhMCQVUxEzARBgNVBAgMClNvbWUtU3RhdGUxITAfBgNVBAoM
GEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0yMzExMjAyMDQ0MzVaFw0yNDEx
MTkyMDQ0MzVaMEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEw
HwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQwggEiMA0GCSqGSIb3DQEB
AQUAA4IBDwAwggEKAoIBAQC9D0204AJsHdi5sVel2u23Ic721ub+W8leU7/zdPC8
GNhYS8yBnRZlJGnnEFE+NLEteuuBZnyAGZErZR1eSmuhdsjuTkwhy7LSor82/tbW
CkG9Q8rx52jIRfmXo1v7baWEc8XrI7+nzAaBuTgERWF9B2ZLDacRzNziejrBpWuA
OTPIakStBIvBs2GJVfQYtckaCNK84qSMO4wfLqcuJrAXpuligNHp5VsWFfo/zOyN
lyhMZPmUv1UVxuJHWet6/uI+xXzxvQzbxRIdYFGIoDUWI64lV3vBtOXj6u/0mUN1
r63KwYHf4N8hTjNsNzinqutJvQ8OjWib57PSRAWjyCrlAgMBAAGjUzBRMB0GA1Ud
DgQWBBTJWhoqxNYbNlunuQt0qeY+qglFejAfBgNVHSMEGDAWgBTJWhoqxNYbNlun
uQt0qeY+qglFejAPBgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQBA
RDkVwN6BdFe15HCFgLApgW8M7PWaYiBT66Iq/f4jJel8fPgcdva7yHgDSf3jkw8K
Q+EAbbe2GZLyWn4vJMT4hPafXDlvWOzQSe6xBdMDywo0yzudAlm6jHCYlvRaSiT9
wF4y50qcfRoMygxN3QWvKqufzwGfLjWKM80BZZ66DEi3fFo5h1/yAemx5X3cToUU
GhkzjOAklf3PUgYRc804Oc9s7O+yun8OAMfsxhSL4IRzHcISf/0TWZarQEdfirJf
e1s9KeFLKrbt7ibUFYFTNuzIG2yrM/SoTvl0EjIVTebdvHPVDdVv/XyOHTWMRK0e
JMxl6uB2i0jOVr85Kay7
-----END CERTIFICATE-----
`)
		keyPem := []byte(`-----BEGIN PRIVATE KEY-----
MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQC9D0204AJsHdi5
sVel2u23Ic721ub+W8leU7/zdPC8GNhYS8yBnRZlJGnnEFE+NLEteuuBZnyAGZEr
ZR1eSmuhdsjuTkwhy7LSor82/tbWCkG9Q8rx52jIRfmXo1v7baWEc8XrI7+nzAaB
uTgERWF9B2ZLDacRzNziejrBpWuAOTPIakStBIvBs2GJVfQYtckaCNK84qSMO4wf
LqcuJrAXpuligNHp5VsWFfo/zOyNlyhMZPmUv1UVxuJHWet6/uI+xXzxvQzbxRId
YFGIoDUWI64lV3vBtOXj6u/0mUN1r63KwYHf4N8hTjNsNzinqutJvQ8OjWib57PS
RAWjyCrlAgMBAAECggEAGQZVTA+Jqotx1DSB7EF0DTVDqp8sALegCiUOR5ivQ8qL
GnbgCTkEjZs6DFx30ILTf/hhA2YQLTmVIlgWQNSbdgod0xNYlvGaSDDEHDCzua2u
YXG/g3EUyMugW77DDl8HVWaoqDT5aanDI7kjTcdsPcs0slMKjfvesfipXdf6SISv
caAT7fnInbaEsfRkwSd4/co2SOWdM8gozkA3MFTvbo923o4IO+fUdaJ0RjsnRimg
qJ42k8RGA0IJqeJiRRCi4CgSuBdaUtt6C8uAWaRGL+HeaZaqTvO0cNwT60N9aif6
m9gezdIBgyXwVeFhX88aFlD6hMYA7C5tXWqrz9Nv6QKBgQDrt7am2KnKd9i2oITb
SqQHKUFGYx/qijLfP6vSwKmdqLTn5Cq4vJkVAQdY453UMwJBog60dpLHDGbHiFEf
NpCPjdiP+zhSY5stcq1kAMB6kG4JavXS1Wr3KD4/+cyPrSKs5g89sf4oHbax5p+p
eih7bdoxseCjOg/miCQUTqpYPQKBgQDNU9S7yfXavAGBuGiqkMDx6QMZX2/mtnP9
PcxBOcoDlQZe1gYKmbJ1mI2U6IwoYFbUfoCVOdSHvDAk5nlt4q3CLqLfIdpolcik
jbuIjCgfvX2umcjL1p64e5OFaE9dM42fbmrJVZc+4LnhGS1fFYhXmnVgaaMMe7xi
h+1PWNefyQKBgGIpU2sP90VVu23yUuFvp/dDeudxCC2H9794qHlPulLpmsym/BOK
lsVkdEbUIznnNB1Y+36zklRKGdMmNYImGvVtQK5VFBNbX5gBlat7lKx10R0i6dQv
BCiBHctOn52FoFcYR2iN4yWZmidjv0G4mXstOBxR0xama1C5iSzbxZyRAoGAX/Uh
onGMOKFMgvdP0wa7ZVLkY+M6RLRYGK3c207KnknzJDcZs7KIuHSHFmRnvCbp7X0B
UrwoGxdT2KrtbPFXCz0IWQdVLzNxBZZHYlU+GzdPmV29faXbn5QBKYPKM2B4e137
AgTCHlygAhbmIV3KzkYdPOWhkZ/yNooX2dFHVCECgYAEzbs+MkI5Xi5g9tO19byT
Xq92XPrdtD+a+SPvnNRvZNUG6yYSF4gSLHbsWzoEAJr+rAYoXzcIo2QfFDMIKaFV
JhsR8m1waRdam/2O/Vof6JNsqAh3sw6wtk5BkASWGHBuuXq+zaaqAH2hZOUFf9g+
sYkLozqJpi8fn6+IMDY3jA==
-----END PRIVATE KEY-----
`)
		cert, err := tls.X509KeyPair(certPem, keyPem)
		if err != nil {
			log.Fatalf("Error creating X509KeyPair: %s", err)
		}

		return &cert, nil
	}})
	log.Print(err)
}
