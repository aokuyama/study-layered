package path

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

type Encrypted struct {
	Data []byte
	Iv   []byte
}

func generateIV() ([]byte, error) {
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}
	return iv, nil
}

func pkcs7Pad(data []byte) []byte {
	length := aes.BlockSize - (len(data) % aes.BlockSize)
	trailing := bytes.Repeat([]byte{byte(length)}, length)
	return append(data, trailing...)
}

func encrypt(data string, key string) (*Encrypted, error) {
	k := []byte(key)
	//paddedKey := pkcs7Pad(k)
	d := []byte(data)

	iv, err := generateIV()
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(k)
	if err != nil {
		return nil, err
	}
	padded := pkcs7Pad(d)
	encrypted := make([]byte, len(padded))
	cbcEncrypter := cipher.NewCBCEncrypter(block, iv)
	cbcEncrypter.CryptBlocks(encrypted, padded)
	e := Encrypted{encrypted, iv}
	return &e, nil
}

func pkcs7Unpad(data []byte) []byte {
	dataLength := len(data)
	padLength := int(data[dataLength-1])
	return data[:dataLength-padLength]
}

func decrypt(e *Encrypted, key string) (*string, error) {
	k := []byte(key)
	//paddedKey := pkcs7Pad(k)
	block, err := aes.NewCipher(k)
	if err != nil {
		return nil, err
	}
	decrypted := make([]byte, len(e.Data))
	cbcDecrypter := cipher.NewCBCDecrypter(block, e.Iv)
	cbcDecrypter.CryptBlocks(decrypted, e.Data)
	d := pkcs7Unpad(decrypted)
	data := string(d)
	return &data, nil
}
