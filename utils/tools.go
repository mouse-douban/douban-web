package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"douban-webend/config"
	"encoding/base64"
	"encoding/hex"
	"math/rand"
	"time"
)

func GenerateRandomPassword() string {
	rand.Seed(time.Now().Unix())
	n := rand.Intn(10) + 8
	var b = make([]byte, 0)
	for i := 0; i < n; i++ {
		b = append(b, byte(rand.Intn(128)))
	}
	hash := hmac.New(sha256.New, []byte(config.Config.PasswordKey))
	hash.Write(b)
	return base64.StdEncoding.EncodeToString([]byte(hex.EncodeToString(hash.Sum(nil))))[0:8]
}

func GenerateRandomUUID() string {
	rand.Seed(time.Now().Unix())
	var randomBytes = make([]byte, 16)
	for i := 0; i < 16; i++ {
		randomBytes[i] = byte(rand.Intn(128))
	}
	randomBytes[6] &= 0x0f /* clear version        */
	randomBytes[6] |= 0x40 /* set to version 4     */
	randomBytes[8] &= 0x3f /* clear variant        */
	randomBytes[8] |= 0x80 /* set to IETF variant  */
	return hex.EncodeToString(randomBytes)
}

func GenerateRandomUserName() string {
	subIndex := rand.Intn(8)
	return "豆豆" + GenerateRandomUUID()[subIndex:subIndex+8]
}
