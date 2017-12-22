package main

import (
	"fmt"
	"net/http"
)

var path = "/Users/bindo/Desktop/doc/myselfCa/"

type myhandler struct {
}

func (h *myhandler) ServeHTTP(w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprintf(w,
		"Hi, This is an example of http service in golang!\n")
}

//func main() {
//	pool := x509.NewCertPool()
//	caCertPath := path + "ca.crt"
//
//	caCrt, err := ioutil.ReadFile(caCertPath)
//	if err != nil {
//		fmt.Println("ReadFile err:", err)
//		return
//	}
//	pool.AppendCertsFromPEM(caCrt)
//
//	s := &http.Server{
//		Addr:    ":8081",
//		Handler: &myhandler{},
//		TLSConfig: &tls.Config{
//			ClientCAs:  pool,
//			ClientAuth: tls.RequireAndVerifyClientCert,
//		},
//	}
//
//	err = s.ListenAndServeTLS(path+"server.crt", path+"server.key")
//	if err != nil {
//		fmt.Println("ListenAndServeTLS err:", err)
//	}
//}
