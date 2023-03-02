package auth

import (
	"auction/internal/model"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

type TokenManager interface {
	NewJWT(body model.TokenBody, ttl time.Duration) (string, error)
	ParseJWT(accessToken string) (model.TokenBody, error)
}

type Manager struct {
	signingKey string
}

type JWTClaim struct {
	Role string `json:"role"`
	jwt.StandardClaims
}

func NewManager(signingKey string) (*Manager, error) {
	if signingKey == "" {
		return nil, errors.New("empty signing key")
	}

	return &Manager{signingKey: signingKey}, nil
}

func (m *Manager) NewJWT(body model.TokenBody, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &JWTClaim{
		body.Role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
			Subject:   body.Uuid,
		},
	})

	return token.SignedString([]byte(m.signingKey))
}

func (m *Manager) ParseJWT(accessToken string) (model.TokenBody, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(m.signingKey), nil
	})

	if err != nil {
		return model.TokenBody{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return model.TokenBody{}, fmt.Errorf("error get user claims from token")
	}

	body := model.TokenBody{
		Role: claims["role"].(string),
		Uuid: claims["sub"].(string),
	}

	return body, nil
}
