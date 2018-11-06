package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/gomodule/redigo/redis"
)

type PrefixQuery struct {
	AppKey      string
	AccessToken string
	Compress    string
	SignMethod  string
	DataType    string
	AppRate     string
	UserTax     string
	ContentType string
	AppSecret   string
}

type ClientReqParam struct {
	Private PrefixQueryBodyPrivate `json:"private"`
	Public  PrefixQueryBodyPublic  `json:"public"`
}

type PrefixQueryBodyPrivate struct {
	Servicedata []PrefixQueryBodyPrivateServicedata `json:"servicedata"`
}

type PrefixQueryBodyPrivateServicedata struct {
	Q    string `json:"q"`
	Code string `json:"code"`
}

type PrefixQueryBodyPublic struct {
	Method    string `json:"method"`
	Version   string `json:"version"`
	Timestamp string `json:"timestamp"`
}

type PrefixQueryResult struct {
	Code     string                    `json:"code"`
	Describe string                    `json:"describe"`
	Result   []PrefixQueryResultResult `json:"result"`
}

type PrefixQueryResultResult struct {
	Code       string `json:"code"`
	Name       string `json:"name"`
	NameLength int    `json:"name_length"`
}

type KPResult struct {
	Code     string `json:"code"`
	Describe string `json:"describe"`
	Result   KPResultResult
}

type KPResultResult struct {
	KpCode  string `json:"kpCode"`
	Code    string `json:"code"`
	KpName  string `json:"kpName"`
	Systype string `json:"systype"`
}

var NuonuoRefixQueryLocak sync.Mutex

const (
	NuoNuoAccessTokenUrl = "https://open.jss.com.cn/accessToken"
	NuoNuoApiUrl         = "https://sdk.jss.com.cn/openPlatform/services"
)

// PrefixQuery 查询请求开票信息接口
func main() {

	prefixQuery := PrefixQuery{
		AppKey:      "4yFWntOF",
		AccessToken: "335973ffebfa1853d2c0df0apnozr8gs",
		Compress:    "",
		SignMethod:  "AES/AES",
		DataType:    "JSON",
		AppRate:     "10",
		ContentType: "application/x-www-form-urlencoded",
		UserTax:     "440301999999018",
		AppSecret:   "DB9B4561BF944DD4",
	}

	//获取公司完整名称
	clientReqParam := ClientReqParam{
		Private: PrefixQueryBodyPrivate{
			Servicedata: []PrefixQueryBodyPrivateServicedata{{Q: "浙江爱信诺"}},
		},
		Public: PrefixQueryBodyPublic{
			Method:    "nuonuo.speedBilling.prefixQuery",
			Version:   "V1.0.0",
			Timestamp: "1541401843085",
		},
	}
	result, err := PrefixQueryName(clientReqParam, prefixQuery)
	if err != nil {
		fmt.Printf("\nresult:        %+v\n", result)
		if err.Error() == "accessToken不匹配/或appKey不匹配" && result.Code == "070301" {
			fmt.Println("不匹配：		", err.Error())
		}
		fmt.Println("err:		", err.Error())
	} else {
		fmt.Printf("\nPrefixQueryName:		+%v\n", result)
	}

	clientReqParam2 := ClientReqParam{
		Private: PrefixQueryBodyPrivate{
			Servicedata: []PrefixQueryBodyPrivateServicedata{{Code: result.Result[1].Code}},
		},
		Public: PrefixQueryBodyPublic{
			Method:    "nuonuo.speedBilling.queryNameAndTaxByCode",
			Version:   "V1.0.0",
			Timestamp: "1541401843085",
		},
	}

	result2, err := PrefixQueryTax(clientReqParam2, prefixQuery)
	if err != nil {
		fmt.Println("result2 error:	", err.Error())
	} else {
		fmt.Printf("\nPrefixQueryTax:		+%v\n", result2)
	}
}

func getAccessToken(storeID int, prefixQuery PrefixQuery) (string, error) {
	NuonuoRefixQueryLocak.Lock()
	defer NuonuoRefixQueryLocak.Unlock()
	client, err := redis.Dial("tcp", ":34560")
	if err != nil {
		panic("redis conn error:	" + err.Error())
	}
	redisKey := fmt.Sprintf("prefix:%d:access-token", storeID)
	accessToken, err := redis.String(client.Do("get", redisKey))
	if err != nil || accessToken == "" {
		requestBody := fmt.Sprintf("client_id=%s&client_secret=%s&grant_type=client_credentials", prefixQuery.AppKey, prefixQuery.AppSecret)
		req, _ := http.NewRequest("POST", NuoNuoAccessTokenUrl, strings.NewReader(requestBody))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "", err
		}
		type TokenResult struct {
			ExpiresIn   int    `json:"expires_in"`
			AccessToken string `json:"access_token"`
		}
		var result TokenResult
		if err := json.Unmarshal(body, &result); err != nil {
			return "", errors.New("json unmarshal error:	" + err.Error())
		}
		if _, err := client.Do("set", redisKey, result.AccessToken); err != nil {
			return "", errors.New("set redis value error:	" + err.Error())
		}
		return result.AccessToken, nil
	}
	return accessToken, nil
}

func PrefixQueryTax(clientReqParam ClientReqParam, prefixQuery PrefixQuery) (*KPResult, error) {
	bodyDecrypt, err := request(clientReqParam, prefixQuery)
	if err != nil {
		return nil, errors.New("request result error:" + err.Error())
	}
	var result KPResult
	if err := json.Unmarshal(bodyDecrypt, &result); err != nil {
		return nil, errors.New("cant json unmarshal:" + err.Error() + "       " + string(bodyDecrypt))
	}
	if result.Code != "S0000" {
		return &result, errors.New(result.Describe)
	}
	return &result, nil
}

// 关键字查询企业税号接口1
func PrefixQueryName(clientReqParam ClientReqParam, prefixQuery PrefixQuery) (*PrefixQueryResult, error) {
	bodyDecrypt, err := request(clientReqParam, prefixQuery)
	if err != nil {
		return nil, errors.New("request result error:" + err.Error())
	}
	var result PrefixQueryResult
	if err := json.Unmarshal(bodyDecrypt, &result); err != nil {
		return nil, errors.New("cant json unmarshal:" + err.Error() + "       " + string(bodyDecrypt))
	}
	if result.Code != "S0000" {
		return &result, errors.New(result.Describe)
	}
	return &result, nil
}

func request(clientReqParam ClientReqParam, prefixQuery PrefixQuery) ([]byte, error) {
	b, err := json.Marshal(clientReqParam)
	if err != nil {
		return nil, errors.New("json marshal error:		" + err.Error())
	}
	// 加密
	crypted, err := AesEncrypt(string(b), prefixQuery.AppSecret)
	if err != nil {
		return nil, errors.New("aes encrypt error:		" + err.Error())
	}
	encodedString := base64.StdEncoding.EncodeToString(crypted)
	encodedString = url.QueryEscape(url.QueryEscape(encodedString))
	// 组装数据，发送请求
	requestBody := "param=" + encodedString
	//fmt.Println(requestBody)
	req, _ := http.NewRequest(http.MethodPost, NuoNuoApiUrl, bytes.NewReader([]byte(requestBody)))
	{
		req.Header.Add("appKey", prefixQuery.AppKey)
		req.Header.Add("accessToken", prefixQuery.AccessToken)
		req.Header.Add("compress", "")
		req.Header.Add("signMethod", prefixQuery.SignMethod)
		req.Header.Add("dataType", prefixQuery.DataType)
		req.Header.Add("appRate", prefixQuery.AppRate)
		req.Header.Add("userTax", prefixQuery.UserTax)
		req.Header.Add("Content-Type", prefixQuery.ContentType)
	}

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	bodyStr, err := url.QueryUnescape(string(body))
	bb, err := base64.StdEncoding.DecodeString(bodyStr)
	if err != nil {
		return nil, errors.New("body decode String " + err.Error())
	}
	bodyDecrypt, err := AesDecrypt(bb, []byte(prefixQuery.AppSecret))
	if err != nil {
		return nil, errors.New("aes decrypt error:	" + err.Error())
	}
	return bodyDecrypt, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := NewECBDecrypter(block)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func AesEncrypt(src, key string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, errors.New("key error " + err.Error())
	}
	if src == "" {
		return nil, errors.New("plain content empty")
	}
	ecb := NewECBEncrypter(block)
	content := []byte(src)
	content = PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)
	return crypted, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncrypter ecb

// NewECBEncrypter returns a BlockMode which encrypts in electronic code book
// mode, using the given Block.
func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}
func (x *ecbEncrypter) BlockSize() int { return x.blockSize }
func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecrypter ecb

// NewECBDecrypter returns a BlockMode which decrypts in electronic code book
// mode, using the given Block.
func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}
func (x *ecbDecrypter) BlockSize() int { return x.blockSize }
func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
