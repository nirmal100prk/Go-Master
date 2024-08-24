package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	//sample JSON - API request
	message := `{"name":"John Doe","email":"john.doe@example.com"}`
	secret := "your-secret-key"

	receivedSignature := generateHMAC(message, secret)
	fmt.Println(verifyHMAC(message, receivedSignature, secret))

}

// HMAC using SHA256
func generateHMAC(message, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

func verifyHMAC(message, receivedSignature, secret string) bool {
	expectedSignature := generateHMAC(message, secret)
	return hmac.Equal([]byte(receivedSignature), []byte(expectedSignature))
}
