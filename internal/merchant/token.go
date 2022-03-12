package merchant

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/eldadj/dgpg/internal/errors"
	"os"
	"time"
)

func Generate(claims jwt.MapClaims) (string, error) {
	key := os.Getenv("JWT_SECRET_KEY")
	var mySigningKey = []byte(key)
	// expires in 1day
	//claims.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", errors.LogError(fmt.Errorf("generate jwt error: %s", err.Error()), errors.ErrMerchantAuthenticationFailed)
	}
	return tokenString, nil
}

func Validate(ctx *context.Context, tokenString string) error {
	key := os.Getenv("JWT_SECRET_KEY")
	var mySigningKey = []byte(key)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing")
		}
		return mySigningKey, nil
	})
	if err != nil {
		return errors.LogError(err, errors.ErrMerchantInvalidToken)
	}

	if !token.Valid {
		return errors.ErrMerchantTokenExpired
	}

	if claims, ok := token.Claims.(jwt.MapClaims); !ok {
		return errors.ErrMerchantInvalidToken
	} else {
		//store merchantId in context
		intf := claims["mid"]
		var mid int64
		if fmid, ok := intf.(float64); ok {
			mid = int64(fmid)
		} else if fmid, ok := intf.(int64); ok {
			mid = fmid
		}
		*ctx = NewContext(*ctx, mid)
	}

	return nil
}
