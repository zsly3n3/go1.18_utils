package gauth

import (
	"github.com/928799934/googleAuthenticator"
	"strings"
)

var GoogleAuthenticator = googleAuthenticator.NewGAuth()

func CreateSecret(ga *googleAuthenticator.GAuth) string {
	secret, err := ga.CreateSecret(16)
	if err != nil {
		return ""
	}
	secret = strings.ReplaceAll(secret, `%`, `A`)
	secret = strings.ReplaceAll(secret, `=`, `A`)
	return secret
}

func GetCode(ga *googleAuthenticator.GAuth, secret string) string {
	code, err := ga.GetCode(secret)
	if err != nil {
		return "*"
	}
	return code
}

//验证令牌
func VerifyCode(ga *googleAuthenticator.GAuth, secret, code string) bool {
	// 1:30sec
	ret, err := ga.VerifyCode(secret, code, 1)
	if err != nil {
		return false
	}
	return ret
}

//生成身份认证二维码
func CreateProvisionURIWithIssuer(user string, secret string) string {
	auth := "totp/"
	secret = `secret=` + secret
	return "otpauth://" + auth + user + "?" + secret
}
