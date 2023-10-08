package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	message := []byte("Test Encrypt Devnus")

	fmt.Println("Original Message:", string(message))

	for i := 1; i <= 3; i++ {
		aesKey := make([]byte, 32)
		if _, err := io.ReadFull(rand.Reader, aesKey); err != nil {
			fmt.Println("Error generating AES key:", err)
			return
		}

		ciphertext, err := encryptAES(message, aesKey)
		if err != nil {
			fmt.Println("Error encrypting message with AES:", err)
			return
		}
		fmt.Printf("Encrypted Message %d: %x\n", i, ciphertext)

		decryptedMessage, err := decryptAES(ciphertext, aesKey)
		if err != nil {
			fmt.Println("Error decrypting message with AES:", err)
			return
		}
		fmt.Printf("Decrypted Message %d: %s\n", i, string(decryptedMessage))
	}
}

func encryptAES(plaintext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := aesGCM.Seal(nil, nonce, plaintext, nil)
	return append(nonce, ciphertext...), nil
}

func decryptAES(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("Ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
