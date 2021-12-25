package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
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
