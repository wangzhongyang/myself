package main

//
//import (
//	"crypto/tls"
//	"fmt"
//	"io/ioutil"
//	"net/http"
//	"strings"
//)
//
//var path1 = "/Users/bindo/Downloads/new_keys/"
//var olt_crt = path1 + "bindo-tls.chained.crt"
//var tls_Pem = path1 + "tls.key"
//var ws_pg_crt = path1 + "ws_pg.crt"
//var urlPath = "https://integration.online.octopus.com.hk:7443/webapi-restricted/payment/request/"
//
//func main() {
//	https()
//}
//
//func https() {
//
//	cliCrt, err := tls.LoadX509KeyPair(olt_crt, tls_Pem)
//	if err != nil {
//		fmt.Println("Loadx509keypair err:", err)
//		return
//	}
//
//	tr := &http.Transport{
//		TLSClientConfig: &tls.Config{
//			Certificates: []tls.Certificate{cliCrt},
//		},
//	}
//	client := &http.Client{Transport: tr}
//	resp, err := client.Post(urlPath, "application/xml", strings.NewReader("333"))
//	if err != nil {
//		fmt.Println("Get error:", err)
//		return
//	}
//	defer resp.Body.Close()
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		fmt.Println("body error:	", err)
//	}
//	fmt.Printf("%v\n", resp)
//	fmt.Println("-----")
//	fmt.Println(string(body))
//}
