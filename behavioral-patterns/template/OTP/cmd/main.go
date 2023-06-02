package main

import (
	"fmt"

	"design-patterns/behavioral-patterns/template/OTP/concrete"
	"design-patterns/behavioral-patterns/template/OTP/template"
)

func main() {
	smsOTP := &concrete.Sms{}
	o := template.Otp{
		IOtp: smsOTP,
	}
	o.GenAndSendOTP(4)

	fmt.Println("")
	emailOTP := &concrete.Email{}
	o = template.Otp{
		IOtp: emailOTP,
	}
	o.GenAndSendOTP(4)

}
