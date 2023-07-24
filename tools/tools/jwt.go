/**
 @author : ikkk
 @date   : 2023/7/22
 @text   : 完成使用jwt进行token加密和解密部分
**/

package tools

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 定义加密和解密所需的密钥
var secretKey = []byte("soft_ware_417_2023")

// GenerateToken 生成JWT加密的Token
func GenerateToken(id string) (string, error) {
	// 设置过期时间为当前时间加上2天
	expirationTime := time.Now().Add(2 * 24 * time.Hour)
	// 创建一个新的Token对象
	token := jwt.New(jwt.SigningMethodHS256)

	// 设置Token的Claims，这里我们将id作为一个自定义的Claim
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = expirationTime

	// 使用密钥对Token进行签名
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ParseToken 解析并验证JWT Token，并返回解析出的id以及过期时间
func ParseToken(tokenString string) (string, time.Time, error) {
	// 解析Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法是否为HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return "", time.Time{}, err
	}

	// 验证Token是否有效
	if !token.Valid {
		return "", time.Time{}, fmt.Errorf("invalid token")
	}

	// 获取Claims中的id值
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", time.Time{}, fmt.Errorf("invalid claims")
	}

	id, ok := claims["id"].(string)
	if !ok {
		return "", time.Time{}, fmt.Errorf("invalid id")
	}
	exp, ok := claims["exp"].(float64)
	if !ok {
		return "", time.Time{}, fmt.Errorf("invalid exp")
	}

	// 获取过期时间
	expirationTime := time.Unix(int64(exp), 0)

	return id, expirationTime, nil
}
