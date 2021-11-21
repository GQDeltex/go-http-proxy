package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

func ValidateToken(token string, url string, secret string, expiry int64) error {
	if time.Now().Unix() > expiry {
		return errors.New("Token expired")
	}
	hash := sha256.Sum256([]byte(url + secret + strconv.FormatInt(expiry, 10)))
	compareHash := hex.EncodeToString(hash[:])
	log.Debug("Got token: " + token)
	log.Debug("Gen token: " + compareHash)
	if compareHash != token {
		return errors.New("Tokens/Hashes do not match")
	}
	return nil
}
