package aescrypt

import (
	"encoding/base64"
	"testing"
)

var (
	key    = []byte("1234567890nnbbdd")
	data   = []byte("hello mike")
	enDoce = "Gh7pZizLefiQtQ1GcH2dLw=="
)

func TestAesEncrypt(t *testing.T) {
	AesEncrypt(data, key)
}

func BenchmarkAesEncrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AesEncrypt(data, key)
	}
}

func TestAesDecrypt(t *testing.T) {
	enDoc, _ := base64.StdEncoding.DecodeString(enDoce)
	AesDecrypt([]byte(enDoc), key)
}

func BenchmarkAesDecrypt(b *testing.B) {
	enDoc, _ := base64.StdEncoding.DecodeString(enDoce)
	for i := 0; i < b.N; i++ {
		AesDecrypt([]byte(enDoc), key)
	}
}
