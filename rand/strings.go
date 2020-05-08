package rand

import (
	"crypto/rand"
	"encoding/base64"
)

const RememberTokenBytes = 64

//Bytes will help us generate n random bytes, or will
//return an error if there was one. This uses crypto/rand
//package so it is safe to use with things like remember token
func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func String(nBytes int) (string, error) {
	b, err := Bytes(nBytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

//Remember Token is a helper function that generate predetermined byte size
func RememberToken() (string, error) {
	return String(RememberTokenBytes)
}
