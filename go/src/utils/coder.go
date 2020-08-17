package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

type Coder struct {
	secretKey string
	paddingSize int
}

func NewCoder(secretKey string) *Coder{
	return &Coder{
		secretKey:secretKey,
		paddingSize:16,
	}
}

func (c *Coder)Encrypt(originString string) (string, error) {
	paddingBytes := c.getPaddingBytes([]byte(originString))
	resultBytes := make([]byte, len(paddingBytes))
	block, err := aes.NewCipher([]byte(c.secretKey))
	if err != nil {
		return "", err
	}
	blockMode := cipher.NewCBCEncrypter(block, []byte(c.secretKey))
	blockMode.CryptBlocks(resultBytes, paddingBytes)
	return base64.StdEncoding.EncodeToString(resultBytes), nil // 对加密后的结果进行Base64编码
}

func (c *Coder)Decrypt(encryptedString string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(encryptedString)
	if err != nil {
		return "", err
	}
	paddingBytes := make([]byte, len(decodedBytes))
	block, err := aes.NewCipher([]byte(c.secretKey))
	if err != nil {
		return "", err
	}
	blockMode := cipher.NewCBCDecrypter(block, []byte(c.secretKey))
	blockMode.CryptBlocks(paddingBytes, decodedBytes)
	return string(c.reducePaddingBytes(paddingBytes)), nil
}

func (c *Coder)getPaddingBytes(originBytes []byte) []byte {
	lengthenLength := c.paddingSize - len(originBytes)%c.paddingSize
	additionBytes := bytes.Repeat([]byte{byte(lengthenLength)}, lengthenLength)
	return append(originBytes, additionBytes...)
}

func (c *Coder)reducePaddingBytes(paddingBytes []byte) []byte {
	if len(paddingBytes) == 0 {
		return paddingBytes
	}
	additionBytesLength := int(paddingBytes[len(paddingBytes)-1])
	return paddingBytes[:len(paddingBytes)-additionBytesLength]
}

