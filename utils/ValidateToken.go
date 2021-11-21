package utils

import (
	"crypto/sha256"
	"errors"
	"strconv"
	"time"
)

func ValidateToken(token string, url string, secret string, expiry int64) error {
	if expiry > time.Now().Unix() {
		return errors.New("Token expired")
	}
	hash := sha256.Sum256([]byte(url + secret + strconv.FormatInt(expiry, 10)))
	compareHash := string(hash[:])
	if compareHash != token {
		return errors.New("Tokens/Hashes do not match")
	}
	return nil
}
