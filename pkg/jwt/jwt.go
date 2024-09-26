package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	jwtKey  []byte
	issuer  string
	expTime time.Duration
}

func New(jwtKey []byte, issuer string, expirationTime time.Duration) *JWTService {
	return &JWTService{
		jwtKey,
		issuer,
		expirationTime,
	}
}

func (s *JWTService) CreateJWTtoken(subject string) (string, error) {
	now := time.Now().UTC()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.RegisteredClaims{
			Issuer:    s.issuer,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.expTime)),
			Subject:   subject,
		})
	return t.SignedString(s.jwtKey)
}

func (s *JWTService) GetSubjectFromJWT(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&jwt.RegisteredClaims{},
		func(*jwt.Token) (any, error) { return s.jwtKey, nil },
	)
	if err != nil {
		return "", err
	}

	return token.Claims.GetSubject()
}
