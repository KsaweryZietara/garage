package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	t.Run("should successfully hash a password", func(t *testing.T) {
		password := "MySecurePassword1!"
		hash, err := HashPassword(password)
		assert.NoError(t, err)
		assert.NotEmpty(t, hash)
	})

	t.Run("should return different hashes for the same password", func(t *testing.T) {
		password := "MySecurePassword1!"
		hash1, _ := HashPassword(password)
		hash2, _ := HashPassword(password)
		assert.NotEqual(t, hash1, hash2, "hashing the same password should result in different hashes due to salting")
	})
}

func TestVerifyPassword(t *testing.T) {
	t.Run("should return true for correct password and hash", func(t *testing.T) {
		password := "MySecurePassword1!"
		hash, _ := HashPassword(password)
		isValid := VerifyPassword(password, hash)
		assert.True(t, isValid)
	})

	t.Run("should return false for incorrect password", func(t *testing.T) {
		password := "MySecurePassword1!"
		wrongPassword := "WrongPassword"
		hash, _ := HashPassword(password)
		isValid := VerifyPassword(wrongPassword, hash)
		assert.False(t, isValid)
	})

	t.Run("should return false for incorrect hash", func(t *testing.T) {
		password := "MySecurePassword1!"
		wrongHash := "$2a$14$wronghash"
		isValid := VerifyPassword(password, wrongHash)
		assert.False(t, isValid)
	})
}
