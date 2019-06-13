package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

// 诺诺发票
const (
	keyBase64 = "LmMGStGtOpF4xNyvYt54EQ=="
	identity  = "93363DCC6064869708F1F3C72A0CE72A713A9D425CD50CDE"
)

func main() {

	orderNo, _ := randomHex(10)
	// 开票请求
	paymentRespones, err := PostPayment(orderNo)
	if err != nil {
		fmt.Println("PostPayment:		", err.Error())
	}
	fmt.Printf("\npaymentRespones:        %+v\n", *paymentRespones)

	// 查询开票结果
	for true {
		status, err := GetPaymentStatus(paymentRespones.Fpqqlsh)
		if err != nil {
			fmt.Println("GetPaymentStatus:		", err.Error())
		}
		fmt.Printf("\nGetPaymentStatus:        %+v\n", status)
		if status.List[0].CStatus == "2" {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}

	// 根据order number 查询发票状态
	status, err := GetPaymentStatusByOrderNo(orderNo)
	if err != nil {
		fmt.Println("GetPaymentStatusByOrderNo error:	", err.Error())
	}
	fmt.Printf("\nGetPaymentStatusByOrderNo:        %+v\n", status)

	//
}

func GetPaymentStatusByOrderNo(orderNo string) (*PaymentQueryResponse, error) {
	url := "https://nnfpdev.jss.com.cn/shop/buyer/allow/ecOd/queryElectricKp.action"
	req := PaymentQueryByOrderNo{
		Identity: identity,
		Orderno:  []string{orderNo},
	}
	reqByte, _ := json.Marshal(req)
	orderString, err := Encode(string(reqByte))
	if err != nil {
		return nil, errors.New("encode error:		" + err.Error())
	}
	body, err := Request(orderString, url)
	if err != nil {
		return nil, err
	}
	fmt.Println("GetPaymentStatusByOrderNo string:		", string(body))
	var res PaymentQueryResponse
	_ = json.Unmarshal(body, &res)
	return &res, nil
}

// 发票状态查询
func GetPaymentStatus(paymentNo string) (*PaymentQueryResponse, error) {
	url := "https://nnfpdev.jss.com.cn/shop/buyer/allow/ecOd/queryElectricKp.action"
	req := PaymentQuery{
		Identity: identity,
		Fpqqlsh:  []string{paymentNo},
	}
	reqByte, _ := json.Marshal(req)
	orderString, err := Encode(string(reqByte))
	if err != nil {
		return nil, errors.New("encode error:		" + err.Error())
	}
	body, err := Request(orderString, url)
	if err != nil {
		return nil, err
	}
	fmt.Println("PaymentQueryResponse string:		", string(body))
	var res PaymentQueryResponse
	_ = json.Unmarshal(body, &res)
	return &res, nil
}

func PostPayment(orderNo string) (*PaymentResponse, error) {
	requestPayMent := `{"identity":"93363DCC6064869708F1F3C72A0CE72A713A9D425CD50CDE","order":{"orderno":"` + orderNo + `","saletaxnum":"339901999999142","saleaddress":"&*杭州市中河中路 222 号平海国际商务大厦 5 楼 ","salephone":"0571-87022168","saleaccount":"东亚银行杭州分行 131001088303400","clerk":"袁 牧庆","payee":"络克","checker":"朱燕","invoicedate":"2016-06-15 01:51:41","kptype":"1","address":"","phone":"13185029520","taxnum":"110101TRDX8RQU1","buyername":" 个 人 ","account":"","fpdm":"","fphm":"","message":"开票机号为02 请前往诺诺网 (www.jss.com.cn)查询发票详情","qdbz":"1","qdxmmc":"1111","detail":[{"goodsname":"1","spec":"1","unit":"1","hsbz":"1","num":"1","price":"19.99","spbm":"1090511030000000000","fphxz":"0","yhzcbs":"0","zzstsgl":"222222","l slbs":"","taxrate":"0.16"},{"goodsname":"2","spec":"2","unit":"2","hsbz":"1","num":"1","price":"20","spbm":"1090511030000000000","fphxz":"0","yhzcbs":"0","zzstsgl":"222222","l slbs":"","taxrate":"0.16"}]}}`
	url := "https://nnfpdev.jss.com.cn/shop/buyer/allow/cxfKp/cxfServerKpOrderSync.action"

	orderString, err := Encode(requestPayMent)
	if err != nil {
		return nil, errors.New("encode error:		" + err.Error())
	}
	body, err := Request(orderString, url)
	if err != nil {
		return nil, err
	}
	var res PaymentResponse
	_ = json.Unmarshal(body, &res)
	return &res, nil
}

func Request(param, urlPath string) ([]byte, error) {
	form := url.Values{}
	form.Set("order", param)

	client := &http.Client{}
	req, err := http.NewRequest("POST", urlPath, bytes.NewReader([]byte(form.Encode())))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func Encode(content string) (string, error) {
	// 取需要加密内容的utf-8编码。
	encrypt := []byte(content)
	// 取MD5Hash码，并组合加密数组
	md5Hasn := getMD5Hash(encrypt)
	// 组合消息体
	totalByte := AddMD5(md5Hasn, encrypt)
	// 取密钥和偏转向量
	key, iv := getKeyIv()

	// 使用DES算法使用加密消息体
	temp, err := AesEncrypt(totalByte, key, iv)
	if err != nil {
		return "", err
	}
	return temp, nil
}

func Decode(decode string) (string, error) {
	//解决URL里加号变空格
	decode = replace(decode)

	// base64解码
	encBuf, err := base64.StdEncoding.DecodeString(decode)
	if err != nil {
		return "", errors.New("base64 decode err:	" + err.Error())
	}

	// 取密钥和偏转向量
	key, iv := getKeyIv()

	// 使用DES算法解密
	temp, err := AesDecrypt(encBuf, key, iv)
	if err != nil {
		return "", errors.New("Aes decrypt err:	" + err.Error())
	}

	// 进行解密后的md5Hash校验
	md5Hash := getMD5Hash(temp[16:])

	// 进行解密校检
	for i := 0; i < len(md5Hash); i++ {
		if md5Hash[i] != temp[i] {
			return "", errors.New("MD5校验错误。")
		}
	}

	// 返回解密后的数组，其中前16位MD5Hash码要除去
	return string(temp[16:]), nil
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

//解决URL里加号变空格
func replace(str string) string {
	str = regexp.MustCompile(" ").ReplaceAllString(str, "+")
	return str
}

func AddMD5(md5Byte, bodyByte []byte) []byte {
	md5Byte = append(md5Byte, bodyByte...)
	return md5Byte
}

func getMD5Hash(encrypt []byte) []byte {
	h2 := md5.New()
	h2.Write(encrypt)
	return h2.Sum(nil)
}

func getKeyIv() (key, iv []byte) {
	keyByte, _ := base64.StdEncoding.DecodeString(keyBase64)
	key = keyByte[:8]
	iv = keyByte[8:]
	return
}

func AesEncrypt(encodeStr, key, iv []byte) (string, error) {
	encodeBytes := []byte(encodeStr)
	//根据key 生成密文
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	encodeBytes = PKCS5Padding(encodeBytes, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(encodeBytes))
	blockMode.CryptBlocks(crypted, encodeBytes)

	return base64.StdEncoding.EncodeToString(crypted), nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//填充
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(ciphertext, padtext...)
}

func AesDecrypt(decodeBytes, key, iv []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	origData := make([]byte, len(decodeBytes))

	blockMode.CryptBlocks(origData, decodeBytes)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// 发票状态查询 Response
type PaymentQueryResponse struct {
	Result   string                     `json:"result"`
	ErrorMsg string                     `json:"error_msg"`
	List     []PaymentQueryResponseList `json:"list"`
}

type PaymentQueryResponseList struct {
	CStatus    string  `json:"c_status"`
	CFpdm      string  `json:"c_fpdm"`
	CKprq      float64 `json:"c_kprq"`
	CBhsje     float64 `json:"c_bhsje"`
	COrderno   string  `json:"c_orderno"`
	CInvoiceid string  `json:"c_invoiceid"`
	CMsg       string  `json:"c_msg"`
	CFpqqlsh   string  `json:"c_fpqqlsh"`
	CFphm      string  `json:"c_fphm"`
	CResultmsg string  `json:"c_resultmsg"`
	CUrl       string  `json:"c_url"`
	CJym       string  `json:"c_jym"`
	CJpgUrl    string  `json:"c_jpg_url"`
	CHjse      float64 `json:"c_hjse"`
	CBuyername string  `json:"c_buyername"`
	CTaxnum    string  `json:"c_taxnum"`
}

// 发票状态查询 ByOrderNo
type PaymentQueryByOrderNo struct {
	Identity string   `json:"identity"`
	Orderno  []string `json:"orderno"`
}

// 发票状态查询
type PaymentQuery struct {
	Identity string   `json:"identity"`
	Fpqqlsh  []string `json:"fpqqlsh"`
}

// 发票请求 Response
type PaymentResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Fpqqlsh string `json:"fpqqlsh"`
}

// 发票请求
type Payment struct {
	Identity string `json:"identity"`
	Order    Order  `json:"order"`
}

type Order struct {
	Orderno     string `json:"orderno"`
	saletaxnum  string `json:"saletaxnum"`
	Saleaddress string `json:"saleaddress"`
	Salephone   string `json:"salephone"`
	Saleaccount string `json:"saleaccount"`
	Clerk       string `json:"clerk"`
	Payee       string `json:"payee"`
	Checker     string `json:"checker"`
	Invoicedate string `json:"invoicedate"`
	Kptype      string `json:"kptype"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	Taxnum      string `json:"taxnum"`
	Buyername   string `json:"buyername"`
	Account     string `json:"account"`
	Fpdm        string `json:"fpdm"`
	Fphm        string `json:"fphm"`
	Message     string `json:"message"`
	Qdbz        string `json:"qdbz"`
	Qdxmmc      string `json:"qdxmmc"`
	Detail      Detail `json:"detail"`
}

type Detail struct {
	Goodsname string `json:"goodsname"`
	Spec      string `json:"spec"`
	Unit      string `json:"unit"`
	Hsbz      string `json:"hsbz"`
	Num       string `json:"num"`
	Price     string `json:"price"`
	Spbm      string `json:"spbm"`
	Fphxz     string `json:"fphxz"`
	Yhzcbs    string `json:"yhzcbs"`
	Zzstsgl   string `json:"zzstsgl"`
	Lslbs     string `json:"lslbs"`
	Taxrate   string `json:"taxrate"`
}
