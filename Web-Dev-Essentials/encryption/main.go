package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
)

func Encrypt(key string, value string, blockSegments bool) (string, error) {
	fmt.Println("1:", key)
	decodedKey, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}
	fmt.Println("2:", decodedKey)
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	var ciphertext []byte
	block, err := aes.NewCipher(decodedKey)
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

func Decrypt(key []byte, encrypted string, blockSegments bool) (string, error) {
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
	//key := "thisis32bitlongpassphraseimusing" // AES key must be 16, 24, or 32 bytes long
	value := "your message here"

	// Generate a 32-byte key for AES-256,24-byte for AES-192, 16-byte for AES-128
	key1, err := generateAESKey(16)
	if err != nil {
		log.Fatal("Error generating AES key:", err)
	}
	encodedkey1 := hex.EncodeToString(key1)
	fmt.Println("hexencodedkey", encodedkey1)
	encrypted, err := Encrypt(encodedkey1, value, false)
	if err != nil {
		fmt.Println("Error encrypting:", err)
		return
	}
	decrypted, err := Decrypt([]byte(key1), encrypted, false)

	fmt.Println("Encrypted value:", encrypted)
	fmt.Println("Decrypted value:", decrypted)
	// Convert key to a hex string for easier display/storage
	fmt.Println("raw", key1)
	keyHex := hex.EncodeToString(key1)
	keyU, _ := hex.DecodeString(keyHex)
	fmt.Println("decodedkey:", keyU)
	fmt.Println("Generated AES Key:", keyHex)
}

func generateAESKey(keySize int) ([]byte, error) {
	key := make([]byte, keySize)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}
	return key, nil
}
