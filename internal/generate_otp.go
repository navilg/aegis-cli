package internal

import (
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func GenerateTOTP(secret string, algo string, digits otp.Digits, period uint) string {

	var otpalgo otp.Algorithm = otp.AlgorithmSHA1

	if algo == "SHA1" {
		otpalgo = otp.AlgorithmSHA1
	} else if algo == "SHA256" {
		otpalgo = otp.AlgorithmSHA256
	} else if algo == "SHA512" {
		otpalgo = otp.AlgorithmSHA512
	}

	token, err := totp.GenerateCodeCustom(secret, time.Now(), totp.ValidateOpts{
		Period:    period,
		Digits:    digits,
		Algorithm: otpalgo,
	})

	if err != nil {
		return ""
	}

	return token
}
