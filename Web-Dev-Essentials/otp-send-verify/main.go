package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"math/big"
	"strconv"
)

func main() {

	otp, err := generateOTP(6)
	fmt.Println(otp, "err", err)
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

// numeric otp generator
func generateNumericOTP(maxDigit int) (int32, error) {
	b := make([]byte, maxDigit)
	n, err := io.ReadAtLeast(rand.Reader, b, maxDigit)
	if n != maxDigit {
		return 0, err
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	number, err := strconv.ParseInt(string(b), 10, 32)
	if err != nil {
		return 0, err
	}

	return int32(number), nil
}

const otpChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

// generates alphanumeric otp of specified length
func generateOTP(length int) (string, error) {
	otp := make([]byte, length)
	for i := range otp {
		randIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(otpChars))))
		fmt.Println(randIndex, big.NewInt(int64(len(otpChars))))
		if err != nil {
			return "", err
		}
		otp[i] = otpChars[randIndex.Int64()]
	}
	return string(otp), nil
}

func hashOTP(otp string) string {
	hash := sha256.New()
	hash.Write([]byte(otp))
	return hex.EncodeToString(hash.Sum(nil))
}

func verifyOTP(inputOTP, storedHash string) bool {
	return hashOTP(inputOTP) == storedHash
}
