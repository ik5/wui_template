package auth

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/crypto/scrypt"
)

// type of crypt
const (
	SCrypt = 0x1
)

// The default values for generation
const (
	DefaultSaltSize = 16
)

// GenSalt creates a random crypto based salt, based on a given size.
// if not given, then DefaultSaltSize is in use
func GenSalt(size int) []byte {
	s := size
	if s == 0 {
		s = DefaultSaltSize
	}

	salt := make([]byte, s)
	rand.Read(salt)
	return salt
}

// GenPassword generate a new password string
func GenPassword(cryptoType int, str string, salt []byte) string {
	padding := make([]byte, 8)
	rand.Read(padding)
	dk, err := scrypt.Key([]byte(str), salt, 1<<15, 8, 1, 32)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%04x$%x$%x$%x", cryptoType, salt, dk, padding)
}

// ValidPassword Validates a given password with the result of the string
func ValidPassword(pass, str string) (bool, error) {
	elements := strings.Split(str, "$")
	if len(elements) != 4 {
		return false, errors.New("Invalid encrypted password")
	}
	cryptoType, err := strconv.Atoi(elements[0])
	if err != nil {
		return false, err
	}
	if cryptoType == 0 {
		return false, errors.New("Invalid encrypted password")
	}
	salt, err := hex.DecodeString(elements[1])
	if err != nil {
		return false, err
	}
	password, err := hex.DecodeString(elements[2])
	if err != nil {
		return false, err
	}

	newPass := GenPassword(cryptoType, pass, salt)
	elements = strings.Split(newPass, "$")
	newPassword, err := hex.DecodeString(elements[2])
	return bytes.Equal(password, newPassword), nil
}
