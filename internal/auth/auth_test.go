package auth

import (
	"testing"
	"time"

	"github.com/KsaweryZietara/garage/internal"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestCreateToken(t *testing.T) {
	auth := New("testKey")
	token, err := auth.CreateToken("john@example.com", internal.Owner)
	assert.NoError(t, err)
	assert.NotEmpty(t, token.JWT)
}

func TestVerifyToken(t *testing.T) {
	auth := New("testKey")

	t.Run("should return email and role for valid token", func(t *testing.T) {
		token, _ := auth.CreateToken("john@example.com", internal.Owner)
		email, role, err := auth.VerifyToken(token.JWT)
		assert.NoError(t, err)
		assert.Equal(t, "john@example.com", email)
		assert.Equal(t, internal.Owner, role)
	})

	t.Run("should return error for invalid token", func(t *testing.T) {
		_, _, err := auth.VerifyToken("invalidToken")
		assert.EqualError(t, err, "invalid token")
	})

	t.Run("should return error for token without email claim", func(t *testing.T) {
		noEmailToken := jwt.NewWithClaims(jwt.SigningMethodHS256,
			jwt.MapClaims{
				"role": internal.Owner,
				"exp":  time.Now().Add(time.Hour * 24).Unix(),
			})
		noEmailTokenString, _ := noEmailToken.SignedString(auth.key)
		_, _, err := auth.VerifyToken(noEmailTokenString)
		assert.EqualError(t, err, "unable to extract email")
	})

	t.Run("should return error for token without role claim", func(t *testing.T) {
		noRoleToken := jwt.NewWithClaims(jwt.SigningMethodHS256,
			jwt.MapClaims{
				"email": "john@example.com",
				"exp":   time.Now().Add(time.Hour * 24).Unix(),
			})
		noRoleTokenString, _ := noRoleToken.SignedString(auth.key)
		_, _, err := auth.VerifyToken(noRoleTokenString)
		assert.EqualError(t, err, "unable to extract role")
	})
}

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
