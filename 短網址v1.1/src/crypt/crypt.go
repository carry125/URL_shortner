package crypt

import (
	"crypto/md5"
	"fmt"
)

func MD5(long_url string) (output string) {
	bytes := md5.Sum([]byte(long_url))
	code := fmt.Sprintf("%x", bytes)[:6]
	return code
}
