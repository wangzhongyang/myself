package main

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"time"
)

// func main() {
//var str *string
//err := Signature1(str)
//if err != nil {
//	fmt.Println(err)
//	os.Exit(0)
//}
//fmt.Println(*str)
// }

func Signature1(str *string) error {
	// 构建XML
	topupRequest := ChargeExtraOctopus{
		Amount:     1,
		ExpiryTime: "2015-08-13T04:00:00+08:00",
		Lang:       "cn",
		Currency:   CurrencyJson{},
	}
	topup := topupRequestCollection{
		Xmlns:        "https://integration.oos.online.octopus.com.hk:7443/msl/",
		DocumentTime: "DocumentTime",
		TopupRequest: TopupRequest{
			Requester: Requester{
				Id: "id",
				RequesterName: RequesterName{
					Default: "default",
					Lang:    topupRequest.Lang,
					Name:    "name",
				},
			},
			RequestId:        "requestId",
			RequestTime:      "2015-08-13T03:00:00+08:00",
			ExpiryTime:       topupRequest.ExpiryTime,
			BusinessDate:     "2015-08-13",
			Amount:           strconv.Itoa(topupRequest.Amount),
			Currency:         Currency{},
			ReturnUrl:        "returnUrl",
			StatusEnquiryUrl: "statusEnquiryUrl",
			CodeID:           123456,
		},
	}
	topupByte, err := xml.MarshalIndent(topup, "    ", "    ")
	if err != nil {
		return err
	}
	// 读取私钥
	inputFile := "/Users/bindo/Desktop/doc/new_keys/ws_pg.pem"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return err
	}
	block, _ := pem.Decode(buf)
	if block == nil {
		return errors.New("privateKey is empty")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return err
	}
	// sign todo is ok
	signature, err := PssSignXml(string(topupByte), priv)
	if err != nil {
		return err
	}
	signed := SignedContent{
		Xmlns:   "https://integration.oos.online.octopus.com.hk:7443/mls/",
		Content: base64.StdEncoding.EncodeToString(topupByte),
		Signature: Signature{
			SignerID: "192001",
			KeyID:    "1",
			Content:  signature,
		},
	}

	// 加密
	// step 1: 生成随机密钥 256bit
	randPrivKey, err := rsa.GenerateKey(rand.Reader, 256)
	derStream := x509.MarshalPKCS1PrivateKey(randPrivKey)

	log.Println(derStream, len(derStream))
	//randBlock := &pem.Block{
	//	Type:  "RSA PRIVATE KEY",
	//	Bytes: derStream,
	//}
	//randPrivKeyByte := pem.EncodeToMemory(randBlock)

	// step 2: 生成96bit的iv
	ivStr := CreateString()

	// step 3: 加密signedContent
	requestDocument, err := xml.MarshalIndent(signed, "    ", "    ")
	ciphertextStr, err := AesGcmEncrypt(requestDocument, derStream, ivStr)
	//if err != nil {
	//	return err
	//}

	// step 4: Construct obj
	iv := base64.StdEncoding.EncodeToString([]byte(ivStr))
	encryptedKey, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, randPrivKey, derStream, nil)
	if err != nil {
		fmt.Println("----")
		return err
	}
	fmt.Println("===")

	authenticationTag := "" // todo what's this?
	encryptedContent := encryptedContent{
		Xmlns:             "https://integration.oos.online.octopus.com.hk:7443/mls/",
		Iv:                iv,
		Ciphertext:        base64.StdEncoding.EncodeToString([]byte(ciphertextStr)),
		AuthenticationTag: authenticationTag,
		EncryptedKey: EncryptedKey{
			KeyID:        "1",
			EncryptedKey: base64.StdEncoding.EncodeToString([]byte(encryptedKey)),
		},
	}
	encryptedRequeststr, err := xml.MarshalIndent(encryptedContent, "  ", "  ")
	if err != nil {
		return err
	}
	str2 := string(encryptedRequeststr)
	str = &str2
	return nil
}

func PssSignXml(content string, priv *rsa.PrivateKey) (string, error) {
	message := []byte(content)
	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto
	pssMessage := message
	//newHash := crypto.SHA256
	pssh := crypto.SHA256.New()
	pssh.Write(pssMessage)
	hashed := pssh.Sum(nil)
	signature, err := rsa.SignPSS(rand.Reader, priv, crypto.SHA256, hashed, &opts)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), err
}

func Nonce(nBytes int) ([]byte, error) {
	b := make([]byte, nBytes)

	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func AesGcmEncrypt(content, buf []byte, iv string) (string, error) {
	blockGcm, err := aes.NewCipher(buf)
	if err != nil {
		fmt.Println("2222")
		return "", err
	}
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(blockGcm)
	if err != nil {
		panic(err.Error())
	}
	// 加密signedContent
	ciphertextByte := aesgcm.Seal(nil, nonce, content, nil)
	fmt.Println("00000000000", string(ciphertextByte))
	return string(ciphertextByte), nil

}

// 生成16位的string
func CreateString() string {
	t := time.Now()
	h := md5.New()
	io.WriteString(h, "crazyof.me")
	io.WriteString(h, t.String())
	str := fmt.Sprintf("%x", h.Sum(nil))
	return str[:12]
}

/*--------------------------------------------结构体定义---------------------------------------------------------------*/
// topupRequestCollection
type topupRequestCollection struct {
	Xmlns        string       `xml:"xmlns,attr"`
	DocumentTime string       `xml:"documentTime,attr"`
	TopupRequest TopupRequest `xml:"topupRequest"`
}

type TopupRequest struct {
	Requester        Requester `xml:"requester"`
	RequestId        string    `xml:"requestId"`
	RequestTime      string    `xml:"requestTime"`
	ExpiryTime       string    `xml:"expiryTime"`
	BusinessDate     string    `xml:"businessDate"`
	Amount           string    `xml:"amount"`
	Currency         Currency  `xml:"currency"`
	ReturnUrl        string    `xml:"returnUrl"`
	StatusEnquiryUrl string    `xml:"statusEnquiryUrl"`
	CodeID           int       `xml:"codeId"`
}

type Requester struct {
	Id            string        `xml:"id,attr"`
	RequesterName RequesterName `xml:"name"`
}

type RequesterName struct {
	Default string `xml:"default,attr"`
	Lang    string `xml:"xml:lang,attr"`
	Name    string `xml:",innerxml"`
}

type Currency struct {
	Code         string  `xml:"code,attr"`
	ExchangeRate float64 `xml:"exchangeRate"`
	LocalAmount  float64 `xml:"localAmount"`
}

type SignedContent struct {
	Xmlns     string    `xml:"xmlns,attr"`
	Content   string    `xml:"content"`
	Signature Signature `xml:"sognature"`
}

type Signature struct {
	SignerID string `xml:"signerID,attr"`
	KeyID    string `xml:"keyID,attr"`
	Content  string `xml:",innerxml"`
}

// encryptedContent
type encryptedContent struct {
	Xmlns             string       `xml:"xmlns,attr"`
	Iv                string       `xml:"iv"`
	Ciphertext        string       `xml:"ciphertext"`
	AuthenticationTag string       `xml:"authenticationTag"`
	EncryptedKey      EncryptedKey `xml:"encryptedKey"`
}
type EncryptedKey struct {
	KeyID        string `xml:"keyId,attr"`
	EncryptedKey string `xml:",innerxml"`
}

type ChargeExtraOctopus struct {
	Amount             int          `json:"amount"`
	RequestDescription string       `json:"request_description"`
	ExpiryTime         string       `json:"expiry_time"`
	Lang               string       `json:"lang"`
	Currency           CurrencyJson `json:"currency"`
	LocationId         int          `json:"location_id"`
	Fee                int          `json:"fee"`
	Mpos               bool         `json:"mpos"`
}
type CurrencyJson struct {
	Code         string  `json:"code"`
	ExchangeRate float64 `json:"exchange_rate"`
	LocalAmount  float64 `json:"local_amount"`
}

type paymentRequestCollection struct {
	Xmlns          string         `xml:"xmlns,attr"`
	DocumrntTime   string         `xml:"documrntTime,attr"`
	PaymentRequest PaymentRequest `xml:"paymentRequest"`
}

type PaymentRequest struct {
	GatewayId    int      `xml:"gatewayId"`
	GatewayRef   string   `xml:"gatewayRef"`
	MerchantId   int      `xml:"merchantId"`
	ExpiryTime   string   `xml:"expiryTime"`
	BusinessDate string   `xml:"businessDate"`
	Amount       int      `xml:"amount"`
	Currency     Currency `xml:"currency"`
	Mpos         int      `xml:"mpos"`
	ReturnUrl    string   `xml:"returnUrl"`
}
