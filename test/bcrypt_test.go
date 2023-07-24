/**
 @author : ikkk
 @date   : 2023/7/23
 @text   : nil
**/

package test

import (
	"car_pooling_admini/tools/tools"
	"testing"
)

func TestBcrypt(T *testing.T) {
	encryption := tools.PasswordEncryption("123456")
	print(": ", encryption, " end")
}
