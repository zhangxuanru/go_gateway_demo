/*
@Time : 2020/9/28 19:19 
@Author : zxr
@File : err_test
@Software: GoLand
*/
package test

import (
	"errors"
	"fmt"
	"testing"
)

func TestErr(t *testing.T)  {
	err:=errors.New("hello world")
	stack := fmt.Sprintf("%+v",err)
	fmt.Printf(stack)
}
