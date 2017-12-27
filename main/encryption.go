package main

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io"
)

type Encryption struct {
	rawPublicKey    []byte
	rawPrivateKey   []byte
	rawOosPublicKey []byte
	privateKey      *rsa.PrivateKey
	publicKey       *rsa.PublicKey
	oosPublicKey    *rsa.PublicKey
	key             []byte
	nonce           []byte
	authTag         []byte
}

func NewEncryption(publicKey, privateKey, oosPublicKey string) (*Encryption, string) {
	b, _ := pem.Decode([]byte(privateKey))
	if b == nil {
		panic("error decode pem")
	}
	pb, _ := pem.Decode([]byte(publicKey))
	if pb == nil {
		panic("error decode pem")
	}
	pc, err := x509.ParseCertificate(pb.Bytes)
	if err != nil {
		panic(err)
	}
	priv, err := x509.ParsePKCS1PrivateKey(b.Bytes)
	if err != nil {
		panic(err)
	}

	oospb, _ := pem.Decode([]byte(oosPublicKey))
	if oospb == nil {
		panic("oos error decode pem")
	}

	oospbc, err := x509.ParseCertificate(oospb.Bytes)
	if err != nil {
		panic(err)
	}
	iv, _ := base64.StdEncoding.DecodeString("V/wi4VI3VmKHFBb1")
	tag, _ := base64.StdEncoding.DecodeString("Nfw1yuXChPU6DPr15tG46g==")
	key, _ := base64.StdEncoding.DecodeString("/MzU5PCGbFHB7oBxvZAs6ADLuR8zgwx3nVQN+yOTkWM=")
	enc := &Encryption{
		rawPrivateKey:   []byte(privateKey),
		rawPublicKey:    []byte(publicKey),
		rawOosPublicKey: []byte(oosPublicKey),
		publicKey:       pc.PublicKey.(*rsa.PublicKey),
		privateKey:      priv,
		oosPublicKey:    oospbc.PublicKey.(*rsa.PublicKey),
		//nonce:           randSlice(12),
		//aes256Key:       randSlice(32),
		//authTag:         randSlice(16),
		nonce:   iv,
		authTag: tag,
		key:     key,
	}
	//t := time.Now()
	//now := t.Format(time.RFC3339Nano)
	//var documentTime = now[:23] + now[26:]
	//mm, _ := time.ParseDuration("1m")
	//t1 := t.Add(mm * 10)
	//now = t1.Format(time.RFC3339Nano)
	//fmt.Println(now[:23] + now[26:])
	//var expiryTime = now[:23] + now[26:]
	//var businessDate = t.Format("2006-01-02")
	//var amount = "1000"
	//var gatewayId = "196030"
	//var merchantId = "196031"

	//var pgRef = "PG" + (new SimpleDateFormat("yyyyMMddHHmmssS").format(new Date())); // unique PG reference number
	//now = t.Format("20060102150405.000")
	//var pgRef = "PG" + now[:14] + now[15:]

	//var requestDocument = "<paymentRequestCollection xmlns=\"http://namespace.oos.online.octopus.com.hk/transaction/\" documentTime=\"" + documentTime + "\"><paymentRequest><gatewayId>" + gatewayId + "</gatewayId><gatewayRef>" + pgRef + "</gatewayRef><merchantId>" + merchantId + "</merchantId><expiryTime>" + expiryTime + "</expiryTime><businessDate>" + businessDate + "</businessDate><amount>" + amount + "</amount><currency code=\"USD\">" + "<exchangeRate>7.8</exchangeRate><localAmount>128.20512</localAmount></currency><mpos>1</mpos><returnUrl>https://payment.example.octopus.com.hk/result</returnUrl></paymentRequest></paymentRequestCollection>"
	requestDocument = `<paymentRequestCollection xmlns="http://namespace.oos.online.octopus.com.hk/transaction/" documentTime="2017-12-03T19:21:20.371+08:00"><paymentRequest><gatewayId>196030</gatewayId><gatewayRef>PG20171203192120322</gatewayRef><merchantId>196031</merchantId><expiryTime>2017-12-03T19:31:20.371+08:00</expiryTime><businessDate>2017-12-03</businessDate><amount>1000</amount><currency code="USD"><exchangeRate>7.8</exchangeRate><localAmount>128.20512</localAmount></currency><mpos>1</mpos><returnUrl>https://payment.example.octopus.com.hk/result</returnUrl></paymentRequest></paymentRequestCollection>`
	return enc, requestDocument
}

func randSlice(size int) []byte {
	slice := make([]byte, size)
	if _, err := io.ReadFull(rand.Reader, slice); err != nil {
		panic(err)
	}
	return slice
}

//RSASSA-PSS
//response is signedContent
func (e *Encryption) Signature(content []byte) (string, error) {
	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto
	contentHash := crypto.SHA256.New()
	contentHash.Write(content)

	sign, err := rsa.SignPSS(rand.Reader, e.privateKey, crypto.SHA256, contentHash.Sum(nil), &opts)
	if err != nil {
		return "", errors.New("Encryption.Signature:" + err.Error())
	}
	return base64.StdEncoding.EncodeToString(sign), nil
}

func (e *Encryption) Verify(content []byte) (string, error) {
	return "", nil
}

// RSA-OAEP 加密随机key
func (e *Encryption) KeyTransport() ([]byte, error) {
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, e.oosPublicKey, e.key, nil)
}

// RSA-OAEP 解密随机key
func (e *Encryption) keyDecode(key []byte) ([]byte, error) {
	return rsa.DecryptOAEP(sha256.New(), rand.Reader, e.privateKey, key, nil)
}

// RSA-GCM 加密密文
func (e *Encryption) Encrypt(content []byte) ([]byte, error) {
	//block, err := aes.NewCipher(e.aes256Key)
	//if err != nil {
	//	return nil, errors.New("Encryption.Encrypt:" + err.Error())
	//}
	//aead, err := cipher.NewGCMWithNonceSize(block, len(e.nonce))
	//if err != nil {
	//	return nil, errors.New("Encryption.Encrypt:" + err.Error())
	//}
	//return aead.Seal(nil, e.nonce, content, e.authTag), nil
	block, err := aes.NewCipher(e.key)
	if err != nil {
		return nil, errors.New("Encryption.Encrypt:" + err.Error())
	}
	gcm, err := cipher.NewGCMWithNonceSize(block, 12)
	if err != nil {
		return nil, errors.New("Encryption.Encrypt:" + err.Error())
	}
	return gcm.Seal(nil, e.nonce, content, []byte("nihao")), nil
}

// RSA-GCM 解密密文
//func (e *Encryption) Decode(content []byte) ([]byte, error) {
//	block, err := aes.NewCipher(e.nonce)
//	if err != nil {
//		return nil, errors.New("Encryption.Encrypt:" + err.Error())
//	}
//	aead, err := cipher.NewGCMWithNonceSize(block, len(e.authTag))
//	if err != nil {
//		return nil, errors.New("Encryption.Encrypt:" + err.Error())
//	}
//	return aead.Open(nil, e.authTag, content, e.aes256Key)
//}

// https 请求
func https(content string) (string, error) {

	// var path1 = "/Users/bindo/Downloads/new_keys/"
	// var olt_crt = path1 + "bindo-tls.chained.crt"
	// var tls_Pem = path1 + "tls.key"
	// var urlPath = "https://integration.online.octopus.com.hk:7443/webapi-restricted/payment/request/"
	// urlPath = "https://www.baidu.com"
	// cliCrt, err := tls.LoadX509KeyPair(olt_crt, tls_Pem)
	// if err != nil {
	// 	return "", err
	// }

	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{
	// 		Certificates: []tls.Certificate{cliCrt},
	// 	},
	// }
	// client := &http.Client{Transport: tr}
	// resp, err := client.Post(urlPath, "application/xml;charset=UTF-8", strings.NewReader("333"))
	// return "", err
	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return "", err
	// }
	// return string(body), nil
	return "", nil
}
