package main

import (
	"PaymentLib/UtilsLib/algorithmUtils"
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"errors"
	"fmt"
)

type BaiWangUploadPublic struct {
	Method    string `json:"method"`
	AppKey    string `json:"appKey"`
	Sign      string `json:"sign"`
	Token     string `json:"token"`
	Timestamp string `json:"timestamp"`
	Format    string `json:"format"`
	Version   string `json:"version"`
	Type      string `json:"type"`
}

type BaiQangUploadBusiness struct {
	SellerTaxNo          string               `json:"sellerTaxNo"`
	InvoiceUploadType    string               `json:"invoiceUploadType"`
	InvoiceInvalidDate   string               `json:"invoiceInvalidDate"`
	DeviceType           string               `json:"deviceType"`
	OrganizationCode     string               `json:"organizationCode"`
	OrganizationName     string               `json:"organizationName"`
	AutoOpen             string               `json:"autoOpen"`
	ReturnType           string               `json:"returnType"`
	SerialNo             string               `json:"serialNo"`
	InvoiceSpecialMark   string               `json:"invoiceSpecialMark"`
	InvoiceTypeCode      string               `json:"invoiceTypeCode"`
	InvoiceTerminalCode  string               `json:"invoiceTerminalCode"`
	InvoiceCode          string               `json:"invoiceCode"`
	InvoiceNo            string               `json:"invoiceNo"`
	InvoiceDate          string               `json:"invoiceDate"`
	InvoiceClosingDate   string               `json:"invoiceClosingDate"`
	TaxControlCode       string               `json:"taxControlCode"`
	InvoiceCheckCode     string               `json:"invoiceCheckCode"`
	InvoiceQrCode        string               `json:"invoiceQrCode"`
	BuyerTaxNo           string               `json:"buyerTaxNo"`
	BuyerName            string               `json:"buyerName"`
	BuyerEmail           string               `json:"buyerEmail"`
	BuyerPhone           string               `json:"buyerPhone"`
	BuyerAddressPhone    string               `json:"buyerAddressPhone"`
	BuyerBankAccount     string               `json:"buyerBankAccount"`
	SellerName           string               `json:"sellerName"`
	SellerAddressPhone   string               `json:"sellerAddressPhone"`
	SellerBankAccount    string               `json:"sellerBankAccount"`
	Drawer               string               `json:"drawer"`
	Checker              string               `json:"checker"`
	Payee                string               `json:"payee"`
	InvoiceType          string               `json:"invoiceType"`
	InvoiceListMark      string               `json:"invoiceListMark"`
	RedInfoNo            string               `json:"redInfoNo"`
	OriginalInvoiceCode  string               `json:"originalInvoiceCode"`
	OriginalInvoiceNo    string               `json:"originalInvoiceNo"`
	TaxationMode         string               `json:"taxationMode"`
	DeductibleAmount     float64              `json:"deductibleAmount"`
	InvoiceTotalPrice    float64              `json:"invoiceTotalPrice"`
	InvoiceTotalTax      float64              `json:"invoiceTotalTax"`
	InvoiceTotalPriceTax float64              `json:"invoiceTotalPriceTax"`
	MachineNo            string               `json:"machineNo"`
	SignatureParameter   string               `json:"signatureParameter"`
	TaxDiskNo            string               `json:"taxDiskNo"`
	GoodsCodeVersion     string               `json:"goodsCodeVersion"`
	ConsolidatedTaxRate  float64              `json:"consolidatedTaxRate"`
	NotificationNo       string               `json:"notificationNo"`
	Remarks              string               `json:"remarks"`
	ReqSerialNumber      string               `json:"reqSerialNumber"`
	InvoiceDetailsList   []InvoiceDetailsList `json:"invoiceDetailsList"`
}

type InvoiceDetailsList struct {
	GoodsLineNo          string  `json:"goodsLineNo"`
	GoodsLineNature      string  `json:"goodsLineNature"`
	GoodsCode            string  `json:"goodsCode"`
	GoodsExtendCode      string  `json:"goodsExtendCode"`
	GoodsName            string  `json:"goodsName"`
	GoodsTaxItem         string  `json:"goodsTaxItem"`
	GoodsSpecification   string  `json:"goodsSpecification"`
	GoodsUnit            string  `json:"goodsUnit"`
	GoodsQuantity        string  `json:"goodsQuantity"`
	GoodsPrice           float64 `json:"goodsPrice"`
	GoodsTotalPrice      float64 `json:"goodsTotalPrice"`
	GoodsTotalTax        float64 `json:"goodsTotalTax"`
	GoodsTaxRate         float64 `json:"goodsTaxRate"`
	GoodsDiscountLineNo  string  `json:"goodsDiscountLineNo"`
	PriceTaxMark         string  `json:"priceTaxMark"`
	VatSpecialManagement string  `json:"vatSpecialManagement"`
	FreeTaxMark          string  `json:"freeTaxMark"`
	PreferentialMark     string  `json:"preferentialMark"`
}

type BaiWangUploadOut struct {
	Method        string                        `json:"method"`
	RequestID     string                        `json:"requestId"`
	Response      BaiWangUploadOutResponse      `json:"response"`
	ErrorResponse BaiWangUploadOutErrorResponse `json:"error_response"`
}

type BaiWangUploadOutResponse struct {
	Message string `json:"message"`
}

type BaiWangUploadOutErrorResponse struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	SubCode    int    `json:"subCode"`
	SubMessage string `json:"subMessage"`
}

func main() {
	//str := "Base64编/解码"
	//a := base64.RawStdEncoding.EncodeToString([]byte(str))
	//fmt.Println(string(a))
	//test3Des()
	qr := QrElement{
		ServiceUrl:  "http://mtest.baiwang-inner.com",
		SellerTaxNo: "512345678900000040",
		SerialNo:    "6650763601",
		UploadTime:  "24",
		Source:      "LS01",
		UserSalt:    "3703d042709f4770a33a859a9c2d2c36",
	}
	//http://mtest.baiwang-inner.com/mix/soi?t=y2CsJ85cDQ9IF4GCdUmf5aqUVAdQY1ky&s=NSUR3RQym9O67-3iJuMNeQ&h=24&d=LS01
	//http://mtest.baiwang-inner.com/mix/soi?t=y2CsJ85cDQ9IF4GCdUmf5aqUVAdQY1ky&s=NSUR3RQym9O67+3iJuMNeQ&h=24&d=LS01
	fmt.Println(des3EncryptRtnUrl(qr))
}

type QrElement struct {
	ServiceUrl  string `json:"service_url"`
	SellerTaxNo string `json:"seller_tax_no"`
	SerialNo    string `json:"serial_no"`
	UploadTime  string `json:"upload_time"`
	Source      string `json:"source"`
	UserSalt    string `json:"user_salt"`
}

func des3EncryptRtnUrl(qr QrElement) (string, error) {
	t, err := encryptMode(qr.SellerTaxNo, qr.UserSalt)
	if err != nil {
		return "", errors.New(fmt.Sprintf("get t error:%s", err.Error()))
	}
	s, _ := encryptMode(qr.SerialNo, qr.UserSalt)
	if err != nil {
		return "", errors.New(fmt.Sprintf("get s error:%s", err.Error()))
	}
	url := fmt.Sprintf("%s/mix/soi?t=%s&s=%s&h=%s&d=%s", qr.ServiceUrl, t, s, qr.UploadTime, qr.Source)
	return url, nil
}

func encryptMode(content, keyStr string) (string, error) {
	key := []byte(keyStr)
	key = GetKey(key)
	temp := GetContent(content)
	res, err := algorithmUtils.EncryptDesECB(temp, key[:24])
	if err != nil {
		return "", err
	}
	return base64.RawStdEncoding.EncodeToString(res), nil
}

func GetContent(content string) []byte {
	contentByte := []byte(content)
	length := len(contentByte) // 原长度
	remainder := 8 - length%8
	if remainder == 0 {
		remainder = 8
	}
	for i := 0; i < remainder; i++ {
		contentByte = append(contentByte, byte(remainder))
	}
	return contentByte
}

func GetKey(key []byte) []byte {
	if len(key) >= 24 {
		return key[:24]
	} else {
		t := 24 - len(key)
		tArray := make([]byte, t)
		key = append(key, tArray...)
		return key
	}
}

// 3DES加密
func TripleDesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:8])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// 3DES解密
func TripleDesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key[:8])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
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
