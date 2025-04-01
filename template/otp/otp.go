package otp

import (
	"fmt"
)

type IOtp interface {
	genRandomOTP(int) string
	saveOTPCache(string)
	getMessage(string) string
	sendNotification(string) error
}

type Otp struct {
	iOtp IOtp
}

func (o *Otp) GenAndSendOTP(otpLength int) error {
	otp := o.iOtp.genRandomOTP(otpLength)
	o.iOtp.saveOTPCache(otp)
	message := o.iOtp.getMessage(otp)
	err := o.iOtp.sendNotification(message)
	if err != nil {
		return err
	}
	return nil
}

func (o *Otp) SetImOtp(i IOtp) {
	o.iOtp = i
}

type SMSOtp struct{}

func (o *SMSOtp) genRandomOTP(int) string {
	return "SMSOtp"
}

func (o *SMSOtp) saveOTPCache(otp string) {
	fmt.Printf("SMS: saving otp: %s to cache\n", otp)
}

func (o *SMSOtp) getMessage(otp string) string {
	return "SMS OTP for login is " + otp
}

func (o *SMSOtp) sendNotification(message string) error {
	fmt.Printf("SMS: sending sms: %s\n", message)
	return nil
}

type EmailOtp struct{}

func (o *EmailOtp) genRandomOTP(int) string {
	return "EmailOtp"
}

func (o *EmailOtp) saveOTPCache(otp string) {
	fmt.Printf("Email: saving otp: %s to cache\n", otp)
}

func (o *EmailOtp) getMessage(otp string) string {
	return "Email OTP for login is " + otp
}

func (o *EmailOtp) sendNotification(message string) error {
	fmt.Printf("Email: sending sms: %s\n", message)
	return nil
}

func CreateOtp(s string) IOtp {
	if s == "sms" {
		return &SMSOtp{}
	}
	return &EmailOtp{}
}
