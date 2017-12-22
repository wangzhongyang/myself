package main

//func main() {
//	//str := `YW3oBSa5NKKMSqLDvGQNoUSBDSF5kzPMUUBU7Dv590hVWHShodLogLaPioho95Iq1sKP3mLnfaYVjTqtRId7tpn6nDtUvehzv/IcjSKMKcQK4Ji+UAWjFL1bUI3FZwYF3qWbq8VevflSckPz2QuIbcO8VbITmv0yWwvMZy5kYzWq2KKS3fm88Be97lwhE7knZLN7xgFDNWVYRFTDnv+twbQRGfvw+BnfybiQYSKGhHxFxKYIEfJSVLwA0gJzDk4NDQI+JIsU/OqkXRWWlG2p1g6NWXoWDD78rl4erU2FRvxlWf0rnoOcoHiP9uuqwRZ279bZb+4Y6Ng5OU06BbX2Eg==`
//	//fmt.Println(len(str))
//	type Email struct {
//		Where string `xml:"where,attr"`
//		Addr  string
//	}
//	type Address struct {
//		City, State string
//	}
//	type Result struct {
//		XMLName xml.Name `xml:"Person"`
//		Name    string   `xml:"FullName"`
//		Company string   `xml:"Company"`
//		State   string
//		City    string
//		Phone   string
//		Email   []Email
//		Groups  []string `xml:"Group>Value"`
//		Code    string   `xml:"code,attr"`
//		Address
//	}
//	v := Result{Name: "none", Phone: "none"}
//	data := `
//        <Person>
//            <FullName>Grace R. Emlin</FullName>
//            <Company>Example Inc.</Company>
//            <Email where="home">
//                <Addr>gre@example.com</Addr>
//            </Email>
//            <Email where='work'>
//                <Addr>gre@work.com</Addr>
//            </Email>
//            <Group code="en">
//                <Value >Friends</Value>
//                <Value>Squash</Value>
//            </Group>
//            <City>Hanga Roa</City>
//            <State>Easter Island</State>
//        </Person>
//    `
//	err := xml.Unmarshal([]byte(data), &v)
//	if err != nil {
//		fmt.Printf("error: %v", err)
//		return
//	}
//	//fmt.Printf("XMLName: %#v\n", v.XMLName)
//	//fmt.Printf("Name: %q\n", v.Name)
//	//fmt.Printf("Phone: %q\n", v.Phone)
//	//fmt.Printf("Email: %v\n", v.Email)
//	//fmt.Printf("Groups: %v\n", v.Groups)
//	//fmt.Printf("Address: %v\n", v.Address)
//	fmt.Printf("%+v\n", v)
//
//	v2 := Result{
//		Name: "wzy",
//		Email: []Email{
//			{Where: "maodian", Addr: "1"},
//			{Where: "shenzhen", Addr: "2"},
//		},
//		Company: "tengxun",
//		State:   "state",
//		City:    "city",
//		Phone:   "phone",
//		Groups:  []string{"nihao", "wohao"},
//		Address: Address{
//			City:  "000",
//			State: "state",
//		},
//	}
//	bytestr, err := xml.MarshalIndent(v2, "  ", "	")
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(0)
//	}
//	fmt.Println(string(bytestr))
//}
