package utils

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

func EncodeMD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func EncodeBcrypt(str string) (hs string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CompareBcrypt(hash, str string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
	return err == nil
}

func EncodeRsa(pk string, label string) (data string, err error) {
	ds, err := base64.StdEncoding.DecodeString(pk)
	if err != nil {
		return "", err
	}
	pub, err := x509.ParsePKIXPublicKey(ds)
	if err != nil {
		return "", err
	}
	var pubKey = pub.(*rsa.PublicKey)
	result, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, []byte(label))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(result), nil
}

func ContainsUpperLowerSpecCharDigit(s string) bool {
	// 字符串同时包含大写字母、小写字母、特殊符号和数字，且大于8位
	if len(s) <= 8 {
		return false
	}
	hasUpper := regexp.MustCompile("[A-Z]").MatchString(s)
	hasLower := regexp.MustCompile("[a-z]").MatchString(s)
	hasDigit := regexp.MustCompile("[0-9]").MatchString(s)
	hasSpecial := regexp.MustCompile("[~!@#$%^&*(),.?\":{}|<>]").MatchString(s)
	return hasUpper && hasLower && hasDigit && hasSpecial
}
