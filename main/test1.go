package main

import (
	"bufio"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fatih/structs"
)

const (
	path1     = "/Users/bindo/Downloads/new_keys/"
	olt_crt   = path1 + "bindo-tls.chained.crt"
	tls_Pem   = path1 + "tls.key"
	urlPath   = "https://integration.online.octopus.com.hk:7443/webapi-restricted/payment/"
	urlJava   = "http://localhost:8080/octopus/"
	gatewayId = "196030"
)

type Int optionalInt

type optionalInt []int

type Wang struct {
	Name string
	Age  Int
}

//

// func main() {
// 	str := "BINDO_COMMON_TEST_CONFIG_FILE"
// 	fmt.Println(os.Getenv(str))
// 	fmt.Println(os.Getenv("GOPATH"))
// 	fmt.Println(os.LookupEnv(str))
// }

func FieldVoid(w Wang, field string) bool {
	s := structs.New(w)
	f := s.Field(field)
	fmt.Println("kind:		", f.Kind().String())
	switch f.Kind().String() {
	case "int":
		if f.Value().(int) == 0 {
			return false
		}
	case "string":
		if f.Value().(string) == "" {
			return false
		}
	}
	return true
}

//<paymentEnquiryResponseCollection xmlns="http://namespace.oos.online.octopus.com.hk/transaction/" documentTime="2017-12-07T16:05:56.855+08:00">
//<paymentEnquiryResponse>
//<gatewayId>196030</gatewayId>
//<gatewayRef>PG20171207160556039</gatewayRef>
//<status code="NOT_FOUND"/>
//</paymentEnquiryResponse>
//</paymentEnquiryResponseCollection>
type paymentEnquiryResponseCollection struct {
	Xmlns                  string                 `xml:"xmlns,attr"`
	DocumentTime           string                 `xml:"documentTime,attr"`
	PaymentEnquiryResponse PaymentEnquiryResponse `xml:"paymentEnquiryResponse"`
}
type PaymentEnquiryResponse struct {
	GatewayId  string `xml:"gatewayId"`
	GatewayRef string `xml:"gatewayRef"`
	Status     Status `xml:"status"`
}
type Status struct {
	Code string `xml:"code,attr"`
}

//<paymentCancellationResponseCollection xmlns="http://namespace.oos.online.octopus.com.hk/transaction/" documentTime="2017-12-07T16:40:11.188+08:00">
//<paymentCancellationResponse>
//<gatewayId>196030</gatewayId>
//<gatewayRef>PG20171207163952538</gatewayRef>
//<status code="NOT_FOUND"/>
//</paymentCancellationResponse>
//</paymentCancellationResponseCollection>
type paymentCancellationResponseCollection struct {
	Xmlns        string                      `xml:"xmlns,attr"`
	DocumentTime string                      `xml:"documentTime,attr"`
	Cancellation PaymentCancellationResponse `xml:"paymentCancellationResponse"`
}
type PaymentCancellationResponse struct {
	GatewayId  string `xml:"gatewayId"`
	GatewayRef string `xml:"gatewayRef"`
	Status     Status `xml:"status"`
}
type paymentResult struct {
	Xmlns           string `xml:"xmlns,attr"`
	Amount          string `xml:"amount"`
	GatewayId       string `xml:"gatewayId"`
	GatewayRef      string `xml:"gatewayRef"`
	ReceiptId       string `xml:"receiptId"`
	Status          Status `xml:"status"`
	TransactionTime string `xml:"transactionTime"`
}

// 构建Enquiry请求字符串
func SetEnquiryRequestDoc() (string, error) {
	t := time.Now()
	now := t.Format(time.RFC3339Nano)
	var documentTime = now[:23] + now[26:]
	var gatewayId = gatewayId
	now = t.Format("20060102150405.000")
	var gatewayRef = "PG" + now[:14] + now[15:]
	var requestDocument = "<paymentEnquiryCollection xmlns=\"http://namespace.oos.online.octopus.com.hk/transaction/\" documentTime=\"" + documentTime + "\"><paymentEnquiry><gatewayId>" + gatewayId + "</gatewayId><gatewayRef>" + gatewayRef + "</gatewayRef></paymentEnquiry></paymentEnquiryCollection>"
	return requestDocument, nil
}

// 构建Cancellation请求字符串
func SetCancellationRequestDoc() (string, error) {
	t := time.Now()
	now := t.Format(time.RFC3339Nano)
	var documentTime = now[:23] + now[26:]
	var gatewayId = gatewayId
	now = t.Format("20060102150405.000")
	var gatewayRef = "PG" + now[:14] + now[15:]
	var requestDocument = "<paymentCancellationCollection xmlns=\"http://namespace.oos.online.octopus.com.hk/transaction/\" documentTime=\"" + documentTime + "\"><paymentCancellation><gatewayId>" + gatewayId + "</gatewayId><gatewayRef>" + gatewayRef + "</gatewayRef></paymentCancellation></paymentCancellationCollection>"
	return requestDocument, nil
}

// 用于加解密得到返回XML字符串
func GetResponseXml(requestDoc string, urlType string) (string, error) {
	// step 1 加密请求文本
	requestString, err := HttpPost(requestDoc, gatewayId, "encrypt")
	if err != nil {
		return "", err
	}
	setFile(requestString, "req.txt")
	// step 2 获取response
	encReq, err := HttpsRequest(requestString, urlType)
	if err != nil {
		return "", err
	}
	// step 3 解密
	responseXml, err := HttpPost(encReq, "", "decode")
	if err != nil {
		return "", err
	}
	return responseXml, nil
}

// HTTPS 请求，用于八达通
func HttpsRequest(requestDoc string, urlType string) (string, error) {

	cliCrt, err := tls.LoadX509KeyPair(olt_crt, tls_Pem)
	if err != nil {
		return "", errors.New("Loadx509keypair err")
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates: []tls.Certificate{cliCrt},
		},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Post(urlPath+urlType+"/", "application/xml", strings.NewReader(requestDoc))
	fmt.Println(resp)
	if err != nil || resp.StatusCode != 200 {
		return "", errors.New("Post err,Response code:" + resp.Status)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("Body error,Response code:" + resp.Status)
	}
	return string(body), nil
}

// http post 请求，用于对Java代码
func HttpPost(request, gatewayId, genre string) (string, error) {
	urlValue := map[string][]string{}
	if genre == "encrypt" {
		urlValue = map[string][]string{
			"reqDoc":    {request},
			"gatewayId": {gatewayId},
		}
	} else {
		urlValue = map[string][]string{
			"encReq": {request},
		}
	}
	resp, err := http.PostForm(urlJava+genre,
		urlValue)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), err
}

// 输出内容到文件
func setFile(outputString, fileName string) {
	outputFile, outputError := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)

	outputWriter.WriteString(outputString)

	outputWriter.Flush()
}
