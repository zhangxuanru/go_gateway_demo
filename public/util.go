/*
@Time : 2020/9/29 18:17
@Author : zxr
@File : util
@Software: GoLand
*/
package public

import (
	"crypto/sha256"
	"fmt"
)

func GenSaltPassword(slat, password string) string {
	s1 := sha256.New()
	s1.Write([]byte(password))
	str1 := fmt.Sprintf("%x", s1.Sum(nil))
	s2 := sha256.New()
	s2.Write([]byte(str1 + slat))
	return fmt.Sprintf("%x", s2.Sum(nil))
}
