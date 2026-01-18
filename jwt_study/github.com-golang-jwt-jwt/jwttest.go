package github_com_golang_jwt_jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var hmacKey = []byte("your-secret-key")

func TokenTest() {
	tokenStr, err := SignToken("user123")
	if err != nil {
		panic(err)
	}
	fmt.Println("token:", tokenStr)

	claims, err := ParseToken(tokenStr)
	if err != nil {
		fmt.Println("parse error:", err)
		return
	}
	fmt.Printf("ok: sub=%s, exp=%v\n", claims["sub"], claims["exp"])
}

// 生成 token
func SignToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"iss": "your-service",
		"aud": "your-audience",
		"exp": time.Now().Add(15 * time.Minute).Unix(), // 过期时间
		"iat": time.Now().Unix(),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(hmacKey)
}

// 解析 + 验证 token
func ParseToken(tokenStr string) (jwt.MapClaims, error) {
	parser := jwt.NewParser(jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	token, err := parser.Parse(tokenStr, func(t *jwt.Token) (any, error) {
		return hmacKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("claims type mismatch")
	}
	return claims, nil
}
