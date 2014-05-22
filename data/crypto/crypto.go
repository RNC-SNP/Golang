package main

import (
	"fmt"
	"io"
	"bytes"
	"encoding/base64"
	"hash"
	"crypto/sha256"
	"crypto/sha1"
	"crypto/md5"
	"crypto/aes"
	"crypto/cipher"
)

func hashDecode(data string, algorithms string) string {
	var h hash.Hash
	if algorithms == "sha1" {
		h = sha1.New()
	} else if algorithms == "sha256" {
		h = sha256.New()
	} else {
		h = md5.New()
	}	
	io.WriteString(h, data)
	return string(h.Sum(nil))
}

func base64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func base64Decode(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func aesEncrypt(origData, key []byte) ([]byte, error) {
     block, err := aes.NewCipher(key)
     if err != nil {
          return nil, err
     }
     blockSize := block.BlockSize()
     origData = PKCS5Padding(origData, blockSize)
     blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
     crypted := make([]byte, len(origData))
     blockMode.CryptBlocks(crypted, origData)
     return crypted, nil
}

func aesDecrypt(crypted, key []byte) ([]byte, error) {
     block, err := aes.NewCipher(key)
     if err != nil {
          return nil, err
     }
     blockSize := block.BlockSize()
     blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
     origData := make([]byte, len(crypted))
     blockMode.CryptBlocks(origData, crypted)
     origData = PKCS5UnPadding(origData)
     return origData, nil
}

func main() {
	data0 := "Rinc is learning golang."
	fmt.Printf("Source: %s\n", data0)
	
	fmt.Printf("SHA1 decode: %s\n", hashDecode(data0, "sha1"))
	fmt.Printf("SHA256 decode: %s\n", hashDecode(data0, "sha256"))
	fmt.Printf("MD5 decode: %s\n", hashDecode(data0, "md5"))
	
	data1 := base64Encode(data0)
	fmt.Printf("base64 Encode: %s\n", data1)
	
	data2, err := base64Decode(data1)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("base64 Decode: %s\n", data2)
	
	key := "i@RincLiu.com"
	data3,_ := aesEncrypt([]byte(data0), []byte(key))
	fmt.Printf("EAS encrypt: %s\n", string(data3))
	data4,_ := aesDecrypt(data3, []byte(key))
	fmt.Printf("EAS decrypt: %s\n", string(data4))
}
