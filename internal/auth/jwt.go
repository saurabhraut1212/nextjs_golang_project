package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string `json:"uid"`
	jwt.RegisteredClaims
}

func GenerateTokens(secret, userID string, accessTTL, refreshTTL time.Duration) (access, refresh string, aExp, rExp time.Time, err error) {
	now := time.Now()
	aExp = now.Add(accessTTL)
	rExp = now.Add(refreshTTL)

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(aExp),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	})
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(rExp),
			IssuedAt:  jwt.NewNumericDate(now),
			Subject:   "refresh",
		},
	})
	access, err = accessToken.SignedString([]byte(secret))
	if err != nil {
		return
	}
	refresh, err = refreshToken.SignedString([]byte(secret))
	return
}

func ParseToken(secret, token string) (*Claims, error) {
	t, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if c, ok := t.Claims.(*Claims); ok && t.Valid {
		return c, nil
	}
	return nil, jwt.ErrTokenInvalidClaims
}
