package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is invalid")
)

const symmetricKey = "abcd1234"

// Payload contains the payload data of the token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	USERID    int32     `json:"user_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func CreateToken(userId int32, duration time.Duration) (string, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	payload := &Payload{
		ID:        tokenId,
		USERID:    userId,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(symmetricKey))
}

// // JWT needs this method
func (payload *Payload) Valid() error {

	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}

	return nil
}

func VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(symmetricKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)

	if err != nil {
		verr, ok := err.(*jwt.ValidationError)

		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)

	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
