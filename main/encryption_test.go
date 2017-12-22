package main

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"strconv"
	"testing"
)

func TestEncryption(t *testing.T) {
	privatekey := `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAucCf4bwg0NQEHPpQ2TJyTAkPz0y/BccnlVdVRVdG0qgYcWn6
y8BZeC2HxXHPA4rgVuraiJjDI8rMg9wVOwMTbIA4h32U2RbI6Hd+RNaWrMD/wivN
rgDRccnp9TyTGkRd5Ma8W2RbbZ9sZ+dWFEvl6seQZJHlvO6CoeQeYjEie+Qzyk2w
hRmqYTpd7BV9xsiCV8WRw0aU7g94D2OdQ22qAh3iTCdlYc83c06oL++sZPSuSCU2
t1QYdjCa4Ml2bN09M5qxwSAtL9SWsKLivjonlN7RLVy8Z2TUiuc94exniXzQFzMW
mZ03uAIJmWEPU9JnFL+BGaa+JSJbMKMyQEV84QIDAQABAoIBAEl+Z0PfNXSqjj4Q
5DArf4GKDFFO4j2dAJJcDYbz8zeh/pnQ/sPjBQNBsHh0gR27sutw3KozFvJwaN67
E0NYAjVpvfQNwfjqxO8FaFZAOTl82zSuNCDmfffxlbnMD7/S0PuVjizy1iHXdALg
SvSY2w07jGveNfG8xL7dDRB4tFYj1NeY8BmKtJdkKAq/W0q6HwwkRYAV32LRe9dQ
DOL262YveOMsJWXxJnE288LyOqqIP1mRiNTSbuWHwo9bsmjWgnAO89ykqT2eqJYv
1Pew6aK5aCIHHttg6bakLGcXbEfCUU7LKKqjFJnr0xWiKRLfzXMfL2uSECgOLtHO
++OCpKkCgYEA8OOLwJOWCnLWUj+HDaGUMtOeIfn30VQtaUvgKMxaSUrN5iCYtzFO
MSarpjFv8eVZfgTlzGZDPBjFhlNVIqThUGCxQlGCVEa3w07gtzIoadQWNMLOk8Vl
sysIATG5KbG5lTCYZ0X7DHL8hKkSTqa75chtv5Gv9HP6zGnZSvj2368CgYEAxWej
bvskWf3nQoSeBFCDDsm5W/VYpFthSnxkMiZbrp7rDivxKknzU8v2JXU/jz+VFp15
zHfGNLnrWJCNSa+b7gXSWoeQC97x1nk7USCiqOzYYREirOcbptu3eqglTrapzqAb
uhH7dy5ZfKGXA5phW2Cwvv8Ix1dUO6r9L3aNgG8CgYANW3C6YvSk6606rk8c0GLZ
VqakF6pILzS9a/moCXzQJ5e3NQOC1PcS/qPx+TfN3/vQYxEi/mCoCm+ZfTFxVFcy
D8qEdOSXK7yw1cTcI6neBGae0laGFyIGh1JQTqOHzlUOEr3ArD65d/7MlFtxhQlS
OoTrZHavRWcYwp3L7HYz/QKBgQCxESpO5SSROCdUyiH3GsTD4fvK2YK+Ql09c+Bn
/3IjatbKqm2zEgwZ8QyEQuxVMKIpW+2hkxoNt2q70UV6f/NtCHnLzGdPzpW8XJfx
SEW/Iltgjf89ejuaauDkO6jjNwOPnJviRjj6iW+pVERh7fs//LXtTFPygonCz7g0
97ErLwKBgBaM2DYTp4d+cf2n86rmXagI2TACaiVdWmkbkXKgc4EK+B/pPOZieWiR
VgT4lbNorPRjoDqGS1tPToro9r93ukOfqePQ/mLxRt8FdbTKE1FoCCkIGW2qwUHU
JYJSbfht7thNvjVcaL4Fqy6015wWNCfCRwEhrhKHtYYXUx13abfj
-----END RSA PRIVATE KEY-----`

	//	var oclPublicKey string = `-----BEGIN CERTIFICATE-----
	//MIID+DCCAuCgAwIBAgIKPrLj1ScaDIfIFjANBgkqhkiG9w0BAQsFADCBijELMAkG
	//A1UEBhMCSEsxEjAQBgNVBAgMCUhvbmcgS29uZzEeMBwGA1UECgwVT2N0b3B1cyBD
	//YXJkcyBMaW1pdGVkMR0wGwYDVQQLDBRUZWNobmljYWwgRGVwYXJ0bWVudDEoMCYG
	//A1UEAwwfT09TIEF1dGhlbnRpY2F0aW9uIENBIChUZXN0aW5nKTAiGA8yMDE3MTEx
	//NzAwMDAwMFoYDzIwMjAxMTE2MjM1OTU5WjBtMQswCQYDVQQGEwJISzESMBAGA1UE
	//CBMJSG9uZyBLb25nMRIwEAYDVQQHEwlIb25nIEtvbmcxDjAMBgNVBAoTBUJpbmRv
	//MRMwEQYDVQQLEwpCaW5kbyBMYWJzMREwDwYDVQQDEwhiaW5kby5pbzCCASIwDQYJ
	//KoZIhvcNAQEBBQADggEPADCCAQoCggEBAJ+qW6lNqaANtZe0N5J/oiSlHV8FmUwG
	//qjqMnDm8pT5d6gjyW/vzF3awQYWrBFU8NDa75IKa3YyZNiRx8YJbcNQj6RLXI8bv
	//BkKxy9W/ZmnYT2quGxQzlldKeMjjtyQ6+nMHIvMPh9EX1EDZvwuCHzwvYUiTl55t
	//vCghcoprjNlDivAZ6XD77OWIvEMMBD+//VzxLfTW82V299nDk2qEhAY7hrDxK/qb
	//UAb3ATgAiNSbuNwKN5D/aeRTpzarm3Ye8rTFw+Di/qC4d6cxiQLYafi7IuZDVLTE
	//vMebh+ritxdcHcRnPMYD5AcYMjhBhcZKkFX0dAbq5kAMCFktdNiILHkCAwEAAaN4
	//MHYwDAYDVR0TAQH/BAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYDVR0lAQH/BAwwCgYI
	//KwYBBQUHAwIwHQYDVR0OBBYEFAr91FT9CkxGuAP3pjXmOsE5WwqfMB8GA1UdIwQY
	//MBaAFCaARbo6LE0sIieaKvhO5ZosekfQMA0GCSqGSIb3DQEBCwUAA4IBAQCgclhb
	//Z2kfuAM//RDz9+T4M7QRIAH8i+DciNJoqXZJacZ5mWub1Ycb91aW0Ae3r/X/WRoN
	//oC8oYOJc+PNlidvRW5SWjDx6cvVQHJmnplGf6sjZASfKRVaBaJ1hAbaobnNpD6ZE
	//ja6w5srC0C5vOvwGKGnOeLSjs+NXXz7wl90HhJT7wL5ujlBx2K3iBsgr9x4QEsr2
	//7qsCxx30P3fMxkZB1IiM30VDr4tvwYNwDO/hLOmq1H1FytSNlb/tzhsw0MIZ2F+s
	//1+I9Wujemst/EmJzgC5Zkg6zElVyypxC51Omyw1hZNbEeY6dzEEfeZp6/xI4wx5p
	//fgdq2yVoXoERziLS
	//-----END CERTIFICATE-----`

	var oclPublicKey string = `-----BEGIN CERTIFICATE-----
MIIElTCCA32gAwIBAgIJAKYpBACKP/p+MA0GCSqGSIb3DQEBCwUAMIGNMQswCQYD
VQQGEwJISzESMBAGA1UECBMJSG9uZyBLb25nMRIwEAYDVQQHEwlIb25nIEtvbmcx
DjAMBgNVBAoTBUJpbmRvMRMwEQYDVQQLEwpCaW5kbyBMYWJzMREwDwYDVQQDEwhi
aW5kby5pbzEeMBwGCSqGSIb3DQEJARYPYWRtaW5AYmluZG8uY29tMB4XDTE3MTEx
NjA5MjE1OFoXDTI3MTExNDA5MjE1OFowgY0xCzAJBgNVBAYTAkhLMRIwEAYDVQQI
EwlIb25nIEtvbmcxEjAQBgNVBAcTCUhvbmcgS29uZzEOMAwGA1UEChMFQmluZG8x
EzARBgNVBAsTCkJpbmRvIExhYnMxETAPBgNVBAMTCGJpbmRvLmlvMR4wHAYJKoZI
hvcNAQkBFg9hZG1pbkBiaW5kby5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAw
ggEKAoIBAQC5wJ/hvCDQ1AQc+lDZMnJMCQ/PTL8FxyeVV1VFV0bSqBhxafrLwFl4
LYfFcc8DiuBW6tqImMMjysyD3BU7AxNsgDiHfZTZFsjod35E1paswP/CK82uANFx
yen1PJMaRF3kxrxbZFttn2xn51YUS+Xqx5BkkeW87oKh5B5iMSJ75DPKTbCFGaph
Ol3sFX3GyIJXxZHDRpTuD3gPY51DbaoCHeJMJ2VhzzdzTqgv76xk9K5IJTa3VBh2
MJrgyXZs3T0zmrHBIC0v1JawouK+OieU3tEtXLxnZNSK5z3h7GeJfNAXMxaZnTe4
AgmZYQ9T0mcUv4EZpr4lIlswozJARXzhAgMBAAGjgfUwgfIwHQYDVR0OBBYEFLau
Qfe4UAanl9uXig9MTvNzdjOzMIHCBgNVHSMEgbowgbeAFLauQfe4UAanl9uXig9M
TvNzdjOzoYGTpIGQMIGNMQswCQYDVQQGEwJISzESMBAGA1UECBMJSG9uZyBLb25n
MRIwEAYDVQQHEwlIb25nIEtvbmcxDjAMBgNVBAoTBUJpbmRvMRMwEQYDVQQLEwpC
aW5kbyBMYWJzMREwDwYDVQQDEwhiaW5kby5pbzEeMBwGCSqGSIb3DQEJARYPYWRt
aW5AYmluZG8uY29tggkApikEAIo/+n4wDAYDVR0TBAUwAwEB/zANBgkqhkiG9w0B
AQsFAAOCAQEAF+P669IeLlll2hiQUePJV5b5urC4LWCKdAWYrv/6dlHUCI8JSOBR
PiiHKMfFInrJIhHM/ox1zBXB+15P5O0Tn/JLBFvpOB9P5IkI683R5BB+Ou7LsIOf
mXxAd95NUatocWsNGkM9SS3wafPEWD3waDD/QSFFg6LkT0gjgR5gBS71WbT+yOnu
z0HnQibWC0bRte4R7mQFvOHelkHTqlwjcygQVYcGd1r47mOuHuEjD/4DS3O2KfLj
rp/J9GXR41546JRRraSjvJBZXz6j+Dg+w6WT/geDIB/RWkaUwfU99EMEtcantsja
MU8Voji937hNMO7jYLqD5FR9X7sB7VXJ8g==
-----END CERTIFICATE-----`

	var oosPublicKey = `-----BEGIN CERTIFICATE-----
MIIC1DCCAbygAwIBAgIEVZ+aBDANBgkqhkiG9w0BAQsFADAsMSowKAYDVQQDDCFP
T1MgVHJhbnNhY3Rpb24gV2ViIEFQSSAoVGVzdGluZykwHhcNMTUwNzEwMTAxMDU1
WhcNMTgwNzEwMTAxMDU1WjAsMSowKAYDVQQDDCFPT1MgVHJhbnNhY3Rpb24gV2Vi
IEFQSSAoVGVzdGluZykwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCj
OQI0MSFMe7pQ2L2jpXVrzQJfoguoN/BghOpw06AeimRqGDGxi+/wLLCb0l+QTLA1
eWtmroaoumFqtR0pO4TDDoBgD8B4VHN3yJZwCBfqZL1hr29Wy08zRpCS02dPkVk4
nQDDOcZ3JI8mDjbI7ZgM2boE4PNLVh0eyUhJBrx4dMXXU5sBG9osHX6x8NSjeheB
5pmSfh8MimQGOXcZ7hFxqsn15p+SkRF2kPFVpwe1O4OTsrsMAnNHbBPNbRxG+Lix
TQnl+KlwjLtSPAO0VwmB9FvoOzblMSPhdvTW6NIICakABnEUBbQSmrdCm422gM67
cHzeKMY0Oau6Xu4AVqXVAgMBAAEwDQYJKoZIhvcNAQELBQADggEBAGgfGbvD7htk
M9/A5ZzaGh6MzAjbbFafxATFSdnaRFENyMpKT7MLmlLOE8TLX0Pd25kB8dNNCEGv
wAHa60MhhmjDfhWWRHAIlIuHnSs9Tu9NLlDrtWIkzV0KYYgl1oBKvOiS4deFK01o
Vl3seMBx/NDiH4DI1U/Bfv94YH3N7PwAOCYDonCrs8zD27hSaCcefJnURY36nwVo
UwVggXCAupSWG7OVth9RAGafKurkzQG7uMvOtL/UUcWx5ydlDvL2gybrcwMVmuGS
tJ267v/vN7N+/MloD8DAGAJXxCDQ30ZQaNQz4Oxvf5xU544OqsbUZy/SLx33+d8j
bHwQpJExucU=
-----END CERTIFICATE-----
`

	//var requestDocument = `<paymentRequestCollection xmlns="http://namespace.oos.online.octopus.com.hk/transaction/" documentTime="2017-11-29T15:45:40.309+08:00"><paymentRequest><gatewayId>196030</gatewayId><gatewayRef>PG20171129154540278</gatewayRef><merchantId>196031</merchantId><expiryTime>2017-11-29T15:55:40.309+08:00</expiryTime><businessDate>2017-11-29</businessDate><amount>1000</amount><returnUrl>https://payment.example.octopus.com.hk/result</returnUrl></paymentRequest></paymentRequestCollection>`

	fmt.Println("---------------------------------------")

	//Step 1: Sign XML Content
	//Step 2: Encrypt Signed Content with random AES-256 key
	//Step 3: Encrypt AES-256 Key with public key

	xmlBytes := []byte(requestDocument)
	// 1 签名内容 todo it`s ok ! (都是以长度为判断依据)
	enc, requestDocument := NewEncryption(oclPublicKey, privatekey, oosPublicKey)

	var requestObj paymentRequestCollection
	err := xml.Unmarshal([]byte(requestDocument), &requestObj)
	//fmt.Printf("%v",)
	if err != nil {
		fmt.Println(err)
	}
	setFile(requestDocument, "requestDocument.txt")
	signed, err := enc.Signature(xmlBytes)
	if err != nil {
		fmt.Println(err)
		return
	}
	signObg := SignedContent{
		Xmlns:   "http://namespace.oos.online.octopus.com.hk/mls/",
		Content: base64.StdEncoding.EncodeToString(xmlBytes),
		Signature: Signature{
			SignerID: strconv.Itoa(requestObj.PaymentRequest.GatewayId),
			KeyID:    "1",
			Content:  signed,
		},
	}
	signedByte, _ := xml.Marshal(signObg)
	fmt.Println(string(signedByte))
	fmt.Println("------------2 begin------")

	// 2 加密内容，使用随机aes256密钥 ciphertext
	xmlBytes, err = xml.Marshal(signObg)
	enced, err := enc.Encrypt(xmlBytes)
	if err != nil {
		fmt.Println("密文加密错误：	", err)
	}
	encryptedKeyDefulet := "qTAri/df6EErUl0Og1F+A9Q4iknI1mY/Fad8rUAu0oN+WJKAV3TvpRWApDpJL0Puu5hvhGVxujROFwMBEkzrd65uiQVnrnASlQ0cR0r7goe/P0zoI/eos5Kd4sO9mUB4TzVeM9w3HJbleyyvzGlWI7PEKVHfaSUMV2UW2C1Nxq12E5aaiSIcMEA2+WfFtrsNpkawlaiwAfEBaOBT4CLV/aBJeG6rlujAClxp0tgcqhC9hYhJ6a0OrIh0SGt8IEKY6rIHf+PLRZdJNnIkdhaZ+8tc+MoO9Rse1TxG5E2TGJRlZBIOW4Gj/Ai9vXc4Cj+4YTbC1GyUPTYPpWtFoem6CDPRS1fLO22fdncve678ZjdJKGKzVCHXN52Nd8pQC0s1CEPr9gfisAHRawcsaHyCZDWEsm5Q0H86q+nXunyZsA7AG6AXb3DjZfvMW6GUy2UJMN02MC7yAJRZmhuu63Hx39RXCsVrIrbl7hfE8imHGEfyBjmB3Jpr8dLz/Lrr4BKAhWeiA6MqlY4spXuXHops3JcnGq+KRMdxj3MgzPmA2mdcZdv64BHt1Vi3UzV5T1SVeP2I9OXOW1h4FkUf0q1kCVrTYItaxqP0Ch076tukwuLUz9Ee1lECwdj0sFtDsg+nT0brl7U+biNs1DMEr/jjnrTBUDjvhEJdpBhyHTyYPykXBulqRcKXToYQTO6DI4LKEXrR8mxUF5ZvqvBfoFb7GlZ4EKzBulrgQS6f6/YN2woaymaMiDv1LUFsijc/+S2tn5yfRhTa+wvjYR4+K97W0WcqonDR61a5U0AjaYLEP3KkuBNhLztLjwQXLZeMLl4iWjGXVP1M5vl2lFbrMcXIpad9J7V8fLIS45eNkrNsH7RsH/ZejacpXTnsyj9hHPZDsziXOCNoQGR30U5m4ZXX+zBRb6R9GpoDJGrAvs4m53MuH+rPIGIiqr143U7hgW3lpRLlQU7uNhhZAXOJ9ZiEsaFxay4E1X9njFCl2OqZJAwe14VH84HBdzHBduuMYCRMdvqzR6iMXt8JhPG8SNlYEvwp27FohbpUrqBHVZDctbrbfFW9I2fA/HebMZT3J0OSAyWH3QQZ2IgK7rnPBlBK1U/2EviuqTmy8/TbOqY3e+QF4hZO/DxuXZVp9tYwLWFkWDZi4UT3syhGh0i7oRbf+4sTooG+OVXD1zDTGsQNdKVwuvkAgmh+evAZ0MO4z3sKIqpj2qn8OeWsrlH0cyqjz9hEvA8faH/SpHi7OqRGZrKfM/Bg6SMnr3qsERE9oIZ9o5laFXkPXJ/wHcSXSrcUpUmp/8mPNWHATub2EyerHf37+qs9pK0lDisOkry60XiFU22DDP11MaBPZnsjU/Ew26jvNVoD6lfEmTNDFxvQsgZr0+CC4Npt24YmpdRzEiimGtDcpFuTUtOfSywoaot86jBhJi7aiVw7qW7x3xz5ff7UCD/jaUq0JHe6dWGV68Xz1IR+59leM9G/huCBH8JygR8nbhPlUd4uvj2eC/smZ0WSHcEdbRhdcnX/qj/AvFbJHjHhqXeS6rKGiy7KVSHmA2BITN7p/XiOsnAhzuTo83quHSapTxQImLh4yfAkF93vlJZqty0AmHSirStV/raDhqDuZkvQ0JMLXPyO4uQdeYxabjGLk5dnnt86+tmk6VTEQU4sOtxoOYTaHtEd8jecCmxjk9AJY3TpxlsiNFKrmc7vm9jrvBuERgeCwWKfnZHp+g=="
	str, _ := base64.StdEncoding.DecodeString(encryptedKeyDefulet)
	fmt.Println("密文原文：	", hex.EncodeToString(enced))
	str1 := base64.StdEncoding.EncodeToString(enced)
	fmt.Println(len(str1), len(encryptedKeyDefulet))

	//return
	// 3 使用公钥加密aes-256密钥
	encryptedKey, err := enc.KeyTransport() // todo it`s ok !
	fmt.Println(len(str) == len(encryptedKey))

	// 4 set obj 构建结构体并生成XML
	iv := enc.nonce // todo it`s ok !
	encryptedContent := encryptedContent{
		Xmlns:             "http://namespace.oos.online.octopus.com.hk/mls/",
		Iv:                base64.StdEncoding.EncodeToString(iv),
		Ciphertext:        base64.StdEncoding.EncodeToString([]byte(enced)),
		AuthenticationTag: base64.StdEncoding.EncodeToString(enc.authTag),
		EncryptedKey: EncryptedKey{
			KeyID:        "3",
			EncryptedKey: base64.StdEncoding.EncodeToString([]byte(encryptedKey)),
		},
	}
	_, err = xml.Marshal(encryptedContent)
	if err != nil {
		fmt.Println(err)
		return
	}
	setFile(string(base64.StdEncoding.EncodeToString([]byte(enced))), "output.txt")

	//response, err := https(string(encryptedRequeststr))
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(response)

	//encryptedKey := "duIgVxiRKIkFPYRsf9QGNlhEgSlxXdqblITfGa2wgxpq2gq8bkTlntjw0XX0swS69NgQHfI60oECeGWa5MvYlWOYtg3mrFg74nwuCEwzVcyhkZGNAFviiiOsuoYLCzkeMnsT5Q88dsP2CmR0l/Oyo+u09/rNeIYtlPhfdjc9aqvG3j/QQU+ooIKzIIyJsHmDQdPdVe3Vdg3E96mGwLpzemyK0pcPdHifg6e1lJtiKOO4Iyi8I/WT3Brz/g1ApUL4ybPN0wv29tRbaBKZOFPOfmEpTHCnqk8aDcZ47FeSCd/AhwMPVfGQIv0wsWhqRbjaGSjrwmWEtzXaftHOO/qcfw=="

}
