package main

//type Recurlyservers struct {
//	XMLName xml.Name `xml:"servers"`
//	//xml:"serverName"
//	//称为 strcut tag
//	Version     string   `xml:"version,attr"`
//	Svs         []server `xml:"server"`
//	Description string   `xml:",innerxml"`
//}
//
//type server struct {
//	XMLName    xml.Name `xml:"server"`
//	ServerName string   `xml:"serverName"`
//	ServerIP   string   `xml:"serverIP"`
//}
//
////type TopupRequestCollection struct {
////	TopupRequest xml.Name `xml:"topupRequest"`
////	Xmlns        xml
////}
//
//type encryptedContent struct {
//	Xmlns             string       `xml:"xmlns,attr"`
//	Iv                string       `xml:"iv"`
//	Ciphertext        string       `xml:"ciphertext"`
//	AuthenticationTag string       `xml:"authenticationTag"`
//	EncryptedKey      EncryptedKey `xml:"encryptedKey"`
//}
//type EncryptedKey struct {
//	KeyID        string `xml:"keyId,attr"`
//	EncryptedKey string `xml:",innerxml"`
//}

//func main() {
//	str := `<encryptedContent xmlns="http://namespace.oos.online.octopus.com.hk/mls/"><iv>V2HVMeUNt8FiQfcs</iv><ciphertext>o1s5i8N</ciphertext><authenticationTag>HtI2oNUOc6Ikab8kNV+UA==</authenticationTag><encryptedKey keyId="1">==</encryptedKey></encryptedContent>`
//	var v1 encryptedContent
//
//	err := xml.Unmarshal([]byte(str), v1)
//	fmt.Println(err)
//	fmt.Printf("%+v\n", v1)
//
//	v := encryptedContent{
//		Xmlns:             "http://namespace.oos.online.octopus.com.hk/mls/",
//		Iv:                "V2HVMeUNt8FiQfcs",
//		Ciphertext:        "o1s5i8N",
//		AuthenticationTag: "HtI2oNUOc6Ikab8kNV+UA==",
//		EncryptedKey: EncryptedKey{
//			KeyID:        "1",
//			EncryptedKey: "==",
//		},
//	}
//	// str1, _ := xml.MarshalIndent(v, "  ", "  ")
//	str1, _ := xml.Marshal(v)
//	fmt.Println(string(str1))
//
//	str2 := CreateString()
//	fmt.Println(len("V2HVMeUNt8FiQfcs"), len(str2), str2)
//}
