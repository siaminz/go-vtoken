/**
 * Author: Amin Zamani
 * File: decrypt.go
 */

package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

// Decrypt method is to extract back the encrypted text
func Decrypt(text, mySecret string) (string, error) {
	block, err := aes.NewCipher([]byte(mySecret))
	if err != nil {
		return "", err
	}
	cipherText := decode(text)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}
