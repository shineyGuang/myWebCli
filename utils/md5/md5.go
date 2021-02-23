package md5

import (
	"crypto/md5"
	"encoding/hex"
)

var secret = "BenxiGG"

func EncryptPassword(pwd string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(pwd)))
}
