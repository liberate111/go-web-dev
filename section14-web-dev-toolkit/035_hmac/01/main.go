package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	message, message2 := "test@test.com", "test@test1.com"

	mHMAC := getHmac(message)
	fmt.Println(mHMAC)
	mHMAC2 := getHmac(message2)
	fmt.Println(mHMAC2)

	check2 := checkHMAC(message, mHMAC, []byte("secretkey"))
	fmt.Println(check2)

	check3 := getHmac(message2) == mHMAC2
	fmt.Println(check3)

	// check := ValidMAC([]byte(message), []byte(mHMAC), []byte("secretkey"))
	// fmt.Println(check)
}

func getHmac(s string) string {
	h := hmac.New(sha256.New, []byte("secretkey"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func checkHMAC(message, messageMAC string, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	io.WriteString(mac, message)
	expectedMAC := fmt.Sprintf("%x", mac.Sum(nil))
	return expectedMAC == messageMAC
}

// func ValidMAC(message, messageMAC, key []byte) bool {
// 	mac := hmac.New(sha256.New, key)
// 	mac.Write(message)
// 	expectedMAC := mac.Sum(nil)
// 	return hmac.Equal(messageMAC, expectedMAC)
// }
