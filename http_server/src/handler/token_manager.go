package handler

import (
	"curve/src/model"
	"curve/src/utils"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type TokenManager struct {
	coder         *utils.Coder
	tokenDuration time.Duration
	secretKey     string
	keyForUID     string
}

func NewTokenManager(coder *utils.Coder, tokenDuration time.Duration, secretKey string, keyForUID string) *TokenManager {
	return &TokenManager{
		coder:         coder,
		tokenDuration: tokenDuration,
		secretKey:     secretKey,
		keyForUID:     keyForUID,
	}
}

func (t *TokenManager) GetSecretTokenString(uid int) (string, error) {
	tokenString, err := t.getTokenString(uid)
	if err != nil {
		logs.Error(err)
		return model.FlagOfInvalidSecretTokenString, err
	}
	secretTokenString, err := t.coder.Encrypt(tokenString)
	if err != nil {
		logs.Error(err)
		return model.FlagOfInvalidSecretTokenString, err
	}
	return secretTokenString, nil
}

func (t *TokenManager) GetUidFromSecretTokenString(secretTokenString string) (int, error) {
	tokenString, err := t.coder.Decrypt(secretTokenString)
	if err != nil {
		logs.Error(err)
		return model.FlagOfInvalidUID, err
	}
	uid, err := t.getUid(tokenString)
	if err != nil {
		logs.Error(err)
		return model.FlagOfInvalidUID, err
	}
	return uid, nil
}

func (t *TokenManager) getTokenString(uid int) (string, error) {
	userClaims := &model.UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(t.tokenDuration).Unix(),
		},
		UID: uid,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	tokenString, err := token.SignedString([]byte(t.secretKey))
	if tokenString, err = token.SignedString([]byte(t.secretKey)); err != nil {
		logs.Error(err)
		return model.FlagOfInvalidTokenString, err
	}
	return tokenString, nil
}

func (t *TokenManager) getUid(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(*jwt.Token) (interface{}, error) {
		return []byte(t.secretKey), nil
	})
	if err != nil {
		logs.Error(err)
		return 0, err
	}
	mapClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err := fmt.Errorf("asserting for jwt.MapClaims failed")
		logs.Error(err)
		return model.FlagOfInvalidUID, err
	}
	if err := mapClaims.Valid(); err != nil {
		logs.Error(err)
		return model.FlagOfInvalidUID, err
	}
	if _, ok = mapClaims[t.keyForUID].(float64); !ok {
		err := fmt.Errorf("asserting for float64 failed")
		logs.Error(err)
		return model.FlagOfInvalidUID, err
	}
	return int(mapClaims[t.keyForUID].(float64)), nil
}
