package auth

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	secret        = "BMCCbMdA1SCbbRYbr7/6jYzOHb3o39RCLTKeuEPCUSoe91FcDaCHfw=="
	lifetime      = 30 * time.Minute
	tokenHeadName = "Bearer"
	HeaderKey     = "Authorization"
)

type JwtCustomClaim struct {
	ID uuid.UUID `json:"id"`
	jwt.StandardClaims
}

var jwtSecret = []byte(secret)

// token 生成
func Generate(id uuid.UUID, now time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtCustomClaim{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(lifetime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	return token.SignedString([]byte(secret))
}

// token確認
func Validate(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &JwtCustomClaim{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there's a problem with the signing method")
		}

		return jwtSecret, nil
	})
}

// パースしてclaimを取得
func Parse(token string) (*JwtCustomClaim, error) {
	tokenRes, tokenErr := Validate(token)

	if tokenErr != nil {
		if ve, ok := tokenErr.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, fmt.Errorf("%s is expired %s", token, tokenErr)
			} else {
				return nil, fmt.Errorf("%s is invalid %s", token, tokenErr)
			}
		} else {
			return nil, fmt.Errorf("%s is invalid %s", token, tokenErr)
		}
	}

	if tokenRes == nil {
		return nil, fmt.Errorf("not found token in %s", token)
	}

	claims, ok := tokenRes.Claims.(*JwtCustomClaim)
	if !ok {
		return nil, fmt.Errorf("not found claims in %s", token)
	}

	return claims, nil
}

func GetHeaderToken(c *gin.Context) (string, error) {
	authHeader := c.Request.Header.Get(HeaderKey)

	fmt.Println("authHeader", authHeader)

	if authHeader == "" {
		return "", fmt.Errorf("auth header is empty")
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == tokenHeadName) {
		return "", fmt.Errorf("auth header is invalid")
	}

	return parts[1], nil
}
