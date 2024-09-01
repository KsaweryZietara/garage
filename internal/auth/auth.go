package auth

import (
	"errors"
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

func (a *Auth) VerifyToken(tokenString string) (string, internal.Role, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return a.key, nil
	})
	if err != nil || !token.Valid {
		return "", "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", errors.New("unable to extract claims")
	}
	if exp, ok := claims["exp"].(float64); ok {
		if time.Now().Unix() > int64(exp) {
			return "", "", errors.New("token has expired")
		}
	}
	email, ok := claims["email"].(string)
	if !ok {
		return "", "", errors.New("unable to extract email")
	}
	role, ok := claims["role"].(string)
	if !ok {
		return "", "", errors.New("unable to extract role")
	}

	return email, internal.Role(role), nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
