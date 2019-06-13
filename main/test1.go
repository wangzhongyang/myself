package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Invoices struct {
	//InvoiceNumber      string  `json:"invoice_number"`
	InvoiceDate        string   `json:"invoice_date"`
	InvoiceTime        string   `json:"invoice_time"`
	Buyer              Buyer    `json:"buyer"`
	TaxType            string   `json:"tax_type"`
	TaxAmount          int      `json:"tax_amount"`
	SalesAmount        int      `json:"sales_amount"`
	TaxRate            float64  `json:"tax_rate"`
	FreeTaxSalesAmount int      `json:"free_tax_sales_amount"`
	ZeroTaxSalesAmount int      `json:"zero_tax_sales_amount"`
	TotalAmount        int      `json:"total_amount"`
	PrintMark          string   `json:"print_mark"` //Y or N
	RandomNumber       string   `json:"random_number"`
	Details            []Detail `json:"details"`
	OrderId            string   `json:"order_id"`
}
type Buyer struct {
	Identifier      string `json:"identifier"`
	Name            string `json:"name"`
	TelephoneNumber string `json:"telephone_number"`
	EmailAddress    string `json:"email_address"`
}
type Detail struct {
	Description string  `json:"description"`
	Quantity    float64 `json:"quantity"`
	//Unit string //单位
	UnitPrice      float64 `json:"unit_price"` //单价
	Amount         float64 `json:"amount"`
	SequenceNumber string  `json:"sequence_number"`
	//remark string
}

type TwReponese struct {
	ProcessId                    string        `json:"process_id"`
	AutoAssignInvoiceTrackResult []AutoAssigns `json:"auto_assign_invoice_track_result"`
	PrintData                    []Printdatas  `json:"print_data"`
}
type AutoAssigns struct {
	OrderId       string `json:"order_id"`
	InvoiceNumber string `json:"invoice_number"`
	InvoiceYear   string `json:"invoice_year"`
	InvoicePeriod string `json:"invoice_period"`
}
type Printdatas struct {
	InvoiceNumber string `json:"invoice_number"`
	Barcode       string `json:"barcode"`
	Qr1           string `json:"qr1"` //左侧二维码
	Qr2           string `json:"qr2"` //右侧二维码
}

type Reqinvo struct {
	Invoices []Invoices `json:"invoices"`
}

const apikey = "YTM4Njc4OTItYmUwNy00MDNjLTkxZTQtMzFiMDViY2QwNWU2MjAxOTAxMDgxNjQ5"
const apisecret = "ZjY0ZDJkYWEtYzFmMy00YTU0LWE1MzYtOTA4ZTNhMjRjMGIyMjAxOTAxMDgxNjQ5"

func main() {
	//init invoices
	by1 := Buyer{
		Identifier:      "12345678",
		Name:            "Donald",
		TelephoneNumber: "0921234567",
		EmailAddress:    "trump@usa.gov",
	}
	devs := make([]Detail, 0)
	dev := Detail{
		Description:    "好棒棒+禮物券",
		Quantity:       3,
		SequenceNumber: "001",
		UnitPrice:      50,
		Amount:         150,
	}
	devs = append(devs, dev)
	//invs := make([]Invoices, 0)
	inv := Reqinvo{
		Invoices: []Invoices{
			{
				InvoiceDate:        "20190118",
				InvoiceTime:        "174536",
				Buyer:              by1,
				TaxAmount:          7,
				TaxRate:            0.05,
				SalesAmount:        143,
				FreeTaxSalesAmount: 0,
				ZeroTaxSalesAmount: 0,
				TotalAmount:        150,
				PrintMark:          "Y",
				RandomNumber:       "1523",
				Details:            devs,
				OrderId:            "test",
				TaxType:            "1",
			},
		},
	}
	reqByte, _ := json.Marshal(inv)

	form := url.Values{}

	form.Set("api_key", apikey)
	form.Set("auto_assign_invoice_track", "true")
	form.Set("invoice", string(reqByte))
	timesmp := strconv.FormatInt(time.Now().Unix(), 10)
	form.Set("timestamp", timesmp)

	encodeSign := getHmacHash(form)
	fmt.Println(encodeSign)
	content := getContent(form)
	Twrequest(encodeSign, content)
}

func getContent(form url.Values) string {
	str := ""
	{
		var postitem []string
		for k := range form {
			postitem = append(postitem, k)
		}
		sort.Strings(postitem)
		//var str string
		for _, kk := range postitem {
			str = str + "&" + kk + "=" + form.Get(kk)
		}
		str = str[1 : len(str)-0]
		textQuoted := strconv.QuoteToASCII(str)
		textUnquoted := textQuoted[1 : len(textQuoted)-1]
		sOld := "\\\""
		sNew := "\""
		str = strings.Replace(textUnquoted, sOld, sNew, -1)
		str = strings.Replace(strings.Replace(url.QueryEscape(str), "%3D", "=", -1), "%26", "&", -1)
		fmt.Println(" getContent str: 		", str)
	}
	return str
}

func getHmacHash(form url.Values) string {
	//apiSecret := "ZjY0ZDJkYWEtYzFmMy00YTU0LWE1MzYtOTA4ZTNhMjRjMGIyMjAxOTAxMDgxNjQ5"
	str := ""
	{
		var postitem []string
		for k, _ := range form {
			postitem = append(postitem, k)
		}
		sort.Strings(postitem)
		for _, kk := range postitem {
			str = str + "&" + kk + "=" + form.Get(kk)
		}
		str = str[1 : len(str)-0]
		textQuoted := strconv.QuoteToASCII(str)
		textUnquoted := textQuoted[1 : len(textQuoted)-1]
		sOld := "\\\""
		sNew := "\""
		str = strings.Replace(textUnquoted, sOld, sNew, -1)

	}
	fmt.Println("getHmacHash: 		", str)
	key := []byte(apisecret)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(str))
	a := mac.Sum(nil)
	temp := base64.StdEncoding.EncodeToString(a)
	fmt.Println("getHmacHash string:		", temp)
	return temp
}

func Twrequest(signaturte string, content string) string {
	url := "https://boxtest.ecloudlife.com/customer/api/C0401"

	println("content:		", content)
	payload := strings.NewReader(content)
	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("signature", signaturte)
	req.Header.Add("cache-control", "no-cache")
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	return string(body)
}
