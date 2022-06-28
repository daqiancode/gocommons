package hashutils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

func Md5Str(text string) string {
	return Md5Bytes([]byte(text))
}

func Md5Bytes(bs []byte) string {
	hash := md5.Sum(bs)
	return hex.EncodeToString(hash[:])
}

func Sha1Bytes(bs []byte) string {
	hash := sha1.Sum(bs)
	return hex.EncodeToString(hash[:])
}

func Sha1Str(text string) string {
	return Sha1Bytes([]byte(text))
}
