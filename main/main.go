//go:generate stringer -type=Pill
package main

import "fmt"

func main() {
	customerLoginAs := ""
	emailStr := "phone-bindo-43881b92ebe203738170016e2d5709f2-86-13410448384@hkt.com"
	if emailStr != "" {
		switch emailStr[:11] {
		case "phone-bindo": // guest(phone number)
			//sendTypeCase.IsGuest = true
			customerLoginAs = "1"
		case "email-bindo": // guest(email)
			//sendTypeCase.IsGuest = true
			customerLoginAs = "2"
		default: // member
			//sendTypeCase.IsGuest = false
			customerLoginAs = "3"
		}
	} else {
		//sendTypeCase.IsGuest = false
		customerLoginAs = "3"
	}
	fmt.Println(customerLoginAs)
}
