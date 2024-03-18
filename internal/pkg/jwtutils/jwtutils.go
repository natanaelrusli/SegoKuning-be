package jwtutils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/natanaelrusli/segokuning-be/internal/apperror"
	"github.com/natanaelrusli/segokuning-be/internal/config"
)

type JWTUtil interface {
	Sign(userId int64) (string, error)
	Parse(tokenStr string) (*MyAuthClaims, error)
}

type jwtUtil struct {
	config config.JwtConfig
}

func NewJWTUtil(jwtConfig config.JwtConfig) *jwtUtil {
	return &jwtUtil{
		config: jwtConfig,
	}
}

type MyAuthClaims struct {
	UserId int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func (j *jwtUtil) Sign(userId int64) (string, error) {
	currentTime := time.Now()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyAuthClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(currentTime),
			ExpiresAt: jwt.NewNumericDate(currentTime.Add(time.Duration(j.config.TokenDuration) * time.Minute)),
			Issuer:    j.config.Issuer,
		},
	})

	s, err := token.SignedString([]byte(j.config.SecretKey))
	if err != nil {
		return "", err
	}

	return s, nil
}

func (j *jwtUtil) Parse(tokenStr string) (*MyAuthClaims, error) {
	res := MyAuthClaims{}

	parser := jwt.NewParser(
		jwt.WithValidMethods(j.config.AllowedAlgs),
		jwt.WithIssuer(j.config.Issuer),
		jwt.WithIssuedAt(),
	)

	token, err := parser.ParseWithClaims(tokenStr, &MyAuthClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.config.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyAuthClaims); ok && token.Valid {
		res = *claims
	} else {
		return nil, apperror.ErrInvalidToken
	}

	return &res, nil
}
