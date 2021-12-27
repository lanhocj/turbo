package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"time"
)

func Hash(email, in string) string {
	h := hmac.New(sha256.New, []byte(email))
	h.Write([]byte(in))
	return hex.EncodeToString(h.Sum(nil))
}

func Vaild(email, in, enc string) bool {
	h := hmac.New(sha256.New, []byte(email))
	h.Write([]byte(in))
	expected := h.Sum(nil)
	hash, _ := hex.DecodeString(enc)
	return hmac.Equal(expected, hash)
}

func SecretHash(email string) string {
	h := hmac.New(sha256.New, nil)
	key := email + time.Now().String()
	p := []byte(key)
	h.Write(p)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
