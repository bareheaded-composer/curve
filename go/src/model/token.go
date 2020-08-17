package model

import "github.com/dgrijalva/jwt-go"

type UserClaims struct {
	jwt.StandardClaims
	UID int `json:"uid"`
}

const (
	FlagOfInvalidTokenString       = ""
	FlagOfInvalidSecretTokenString = ""
	FlagOfInvalidUID               = -1
	KeyForUid = "uid"
)

const (
	KeyForTokenInCookies = "token"
)
