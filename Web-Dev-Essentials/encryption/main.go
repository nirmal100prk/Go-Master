package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
)

func encrypt(key []byte, value string, blockSegments bool) (string, error) {
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	var ciphertext []byte
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if blockSegments {
		// Padding the value to be a multiple of 16 bytes
		remainder := len(value) % aes.BlockSize
		paddedValue := value
		if remainder != 0 {
			paddedValue += string(bytes.Repeat([]byte{0}, aes.BlockSize-remainder))
		}

		stream := cipher.NewCFBEncrypter(block, iv)
		ciphertext = make([]byte, len(paddedValue))
		stream.XORKeyStream(ciphertext, []byte(paddedValue))
		ciphertext = ciphertext[:len(value)] // trim to original length
	} else {
		stream := cipher.NewCFBEncrypter(block, iv)
		ciphertext = make([]byte, len(value))
		stream.XORKeyStream(ciphertext, []byte(value))
	}

	// Return the base64 encoded string without padding
	encoded := base64.RawURLEncoding.EncodeToString(append(iv, ciphertext...))
	return encoded, nil
}

func decrypt(key []byte, encrypted string, blockSegments bool) (string, error) {
	// Decode the base64 encoded string
	decoded, err := base64.RawURLEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	// Extract the IV (Initialization Vector)
	if len(decoded) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := decoded[:aes.BlockSize]
	ciphertext := decoded[aes.BlockSize:]

	var plaintext []byte

	// Create a new AES cipher with the provided key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if blockSegments {
		// Decrypt with padding consideration
		stream := cipher.NewCFBDecrypter(block, iv)
		plaintext = make([]byte, len(ciphertext))
		stream.XORKeyStream(plaintext, ciphertext)

		// Trim padding added during encryption
		plaintext = trimPadding(plaintext)
	} else {
		// Decrypt directly
		stream := cipher.NewCFBDecrypter(block, iv)
		plaintext = make([]byte, len(ciphertext))
		stream.XORKeyStream(plaintext, ciphertext)
	}

	return string(plaintext), nil
}

// trimPadding removes the padding added during the encryption when blockSegments is true.
func trimPadding(plaintext []byte) []byte {
	// Find the first null byte and trim the plaintext up to that point
	for i := len(plaintext) - 1; i >= 0; i-- {
		if plaintext[i] != 0 {
			return plaintext[:i+1]
		}
	}
	return plaintext
}

func main() {
	key := "thisis32bitlongpassphraseimusing" // AES key must be 16, 24, or 32 bytes long
	value := "your message here"

	encrypted, err := encrypt([]byte(key), value, false)
	if err != nil {
		fmt.Println("Error encrypting:", err)
		return
	}
	decrypted, err := decrypt([]byte(key), encrypted, false)

	fmt.Println("Encrypted value:", encrypted)
	fmt.Println("Decrypted value:", decrypted)
}
