package tool

import (
	"testing"
)

var key = "12345678901234567890123456789012"
var plaintext = "some random long plaintext"
var ciphertext = "00ac5bd31b24c83e3be51328cd1526c587972eaac177a8dfb2c6f496373dfc3adaf29186e7e409f1"
var decipheredtext = "hahahahaha hahaha hahaha"

func TestEncrypt(t *testing.T) {
	encrypted := Encrypt(key, plaintext)
	decrypted := Decrypt(key, encrypted)
	if plaintext != decrypted {
		t.Errorf("For '%s' expected '%s', got '%s'.", plaintext, plaintext, decrypted)
	}
}

func TestDecrypt(t *testing.T) {
	decrypted := Decrypt(key, ciphertext)
	if decipheredtext != decrypted {
		t.Errorf("For '%s' expected '%s', got '%s'.", plaintext, plaintext, decrypted)
	}
}
