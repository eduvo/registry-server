package main

import "testing"

var key = "12345678901234567890123456789012" // 32 bytes for AES-256
var plaintext = "some random long plaintext"
var ciphertext = "98b6686cdb661dab66b461b07cb2b70e228ff3907dabf8ceaffd7ba2d53cec384ef0eebd0f1aa542c2d7"

func TestEncrypt(t *testing.T) {
  encrypted := Encrypt([]byte(key), []byte(plaintext))
  decrypted := Decrypt([]byte(key), []byte(encrypted))
  if plaintext != decrypted {
    t.Errorf( "For '%s' expected '%s', got '%s'.", plaintext, plaintext, decrypted )
  }
}
