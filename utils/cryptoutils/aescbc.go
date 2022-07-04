package cryptoutils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"

	"github.com/daqiancode/gocommons/utils/hashutils"
)

type AESBase struct{}

func (AESBase) PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func (AESBase) PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

type AESCBC struct {
	AESBase
	key []byte
}

//Symmetric key length must be a multiple of 16
func NewAESCBC(key []byte) AESCBC {
	return AESCBC{key: key}
}

//Symmetric key length must be a multiple of 16
func NewAESCBCMd5Key(key string) AESCBC {
	return AESCBC{key: []byte(hashutils.Md5Str(key))}
}

func (s AESCBC) Encrypt(originalData []byte) ([]byte, error) {
	block, err := aes.NewCipher(s.key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	paddedData := s.PKCS7Padding(originalData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, s.key[:blockSize])
	crypted := make([]byte, len(paddedData))
	blockMode.CryptBlocks(crypted, paddedData)
	return crypted, nil
}
func (s AESCBC) Decrypt(cryptedData []byte) ([]byte, error) {
	block, err := aes.NewCipher(s.key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, s.key[:blockSize])
	origData := make([]byte, len(cryptedData))
	blockMode.CryptBlocks(origData, cryptedData)
	origData = s.PKCS7UnPadding(origData)
	return origData, nil
}

func (s AESCBC) EncryptStr(originalData string) (string, error) {
	r, err := s.Encrypt([]byte(originalData))
	return hex.EncodeToString(r), err
}

func (s AESCBC) DecryptStr(cryptedData string) (string, error) {
	bs, err := hex.DecodeString(cryptedData)
	if err != nil {
		return "", err
	}
	r, err := s.Decrypt(bs)
	return string(r), err
}
