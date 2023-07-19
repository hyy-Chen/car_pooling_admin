/**
 @author : ikkk
 @date   : 2023/7/19
 @text   : nil
**/

package tools

import "golang.org/x/crypto/bcrypt"

// PasswordEncryption 使用bcrypt加密密码数据，保证数据安全性
func PasswordEncryption(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

// PasswordValidation 验证密码正确性，password 用户输入密码, encryptPassword 哈希加密密码
func PasswordValidation(password, encryptPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(encryptPassword), []byte(password))
	return err
}
