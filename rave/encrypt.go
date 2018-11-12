// Implements Rave Encryption Algorithm

package rave

import (
	"bytes"
	"crypto/des"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"strings"
)

// Get a key for encryption
func (r Rave) getKey(seckey string) string {
	keymd5 := md5.Sum([]byte(seckey))
	keymd5Last12 := keymd5[len(keymd5)-6:] // -6 because it's a hex byte array not a string
	seckeyAdjusted := strings.Replace(seckey, "FLWSECK-", "", 1)
	seckeyAdjustedFirst12 := seckeyAdjusted[:12]

	return seckeyAdjustedFirst12 + hex.EncodeToString(keymd5Last12[:])
}

func (r Rave) encrypt(payload string) string {
	seckey := r.getSecretKey()
	encryptedSecKey := r.getKey(seckey)

	return r.encrypt3Des(encryptedSecKey, payload)
	
}

// Encrypt3Des : Encrypts the data using 3Des encryption
func (r Rave) encrypt3Des(key string, payload string) string {
	block, err := des.NewTripleDESCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	bs := block.BlockSize() // block size is 8 by default
	payloadBytes := pkcs5Padding([]byte(payload), bs)

	if len(payloadBytes)%bs != 0 {
		panic("Need a multiple of the blocksize")
	}
	encrypted := make([]byte, len(payloadBytes))
	dst := encrypted

	for len(payloadBytes) > 0 {
		block.Encrypt(dst, payloadBytes[:bs])
		payloadBytes = payloadBytes[bs:]
		dst = dst[bs:]
	}

	return base64.StdEncoding.EncodeToString(encrypted)
}

// pkcs5Padding : Implements PKCS5 padding
func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
