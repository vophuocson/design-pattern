package template

import "design-pattern-practice/behavioral-pattern/template/otp"

func main() {
	t := otp.Otp{}
	insOtp := otp.CreateOtp("sms")
	t.SetImOtp(insOtp)
	t.GenAndSendOTP(4)
	eOtp := otp.CreateOtp("email")
	t.SetImOtp(eOtp)
	t.GenAndSendOTP(4)

}
