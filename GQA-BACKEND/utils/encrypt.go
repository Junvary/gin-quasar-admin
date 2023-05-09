package utils

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func EncodeMD5(str string, b ...byte) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(b))
}

func EncodeRsa(pk string, label string) (data []byte, err error) {
	ds, err := base64.StdEncoding.DecodeString(pk)
	if err != nil {
		return nil, err
	}
	db, err := x509.ParsePKIXPublicKey(ds)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	rsaPublicKey := db.(*rsa.PublicKey)
	oaep, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPublicKey, []byte(label))
	if err != nil {
		return nil, err
	}
	return oaep, nil
}
