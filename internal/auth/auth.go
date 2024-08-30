package auth

import (
	"time"

	"github.com/KsaweryZietara/garage/internal"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	key []byte
}

func New(key string) *Auth {
	return &Auth{
		key: []byte(key),
	}
}

func (a *Auth) CreateToken(email string, role internal.Role) (internal.Token, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"role":  role,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(a.key)
	if err != nil {
		return internal.Token{}, err
	}

	return internal.Token{JWT: tokenString}, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
