package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/hex"
    "fmt"
    "io"
)

func Run() {
    key := []byte("12345678901234567890123456789012") // 32 bytes for AES-256
    plaintext := []byte("some 33 long plaintext")
    fmt.Printf("%s\n", plaintext)
    ciphertext := Encrypt(key, plaintext)
    fmt.Printf("%s\n", ciphertext)
    ciphertext = "00ac5bd31b24c83e3be51328cd1526c587972eaac177a8dfb2c6f496373dfc3adaf29186e7e409f1"
    result := Decrypt(key, []byte(ciphertext))
    fmt.Printf("%s\n", result)
}

// See recommended IV creation from ciphertext below
//var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}


func Encrypt(key, text []byte) string {
    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }
    ciphertext := make([]byte, aes.BlockSize+len(text))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        panic(err)
    }
    cfb := cipher.NewCFBEncrypter(block, iv)
    cfb.XORKeyStream(ciphertext[aes.BlockSize:], text)
    return hex.EncodeToString(ciphertext)
}

func Decrypt(key, text []byte) string {
    ciphertext, _ := hex.DecodeString(string(text))
    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }
    if len(ciphertext) < aes.BlockSize {
        panic("ciphertext too short")
    }
    iv := ciphertext[:aes.BlockSize]
    ciphertext = ciphertext[aes.BlockSize:]
    cfb := cipher.NewCFBDecrypter(block, iv)
    cfb.XORKeyStream(ciphertext, ciphertext)
    return string(ciphertext)
}
