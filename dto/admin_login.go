/*
@Time : 2020/9/25 19:10
@Author : zxr
@File : admin_login
@Software: GoLand
*/
package dto

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zhangxuanru/go_gateway_demo/public"
)

type AdminLoginInput struct {
	UserName string `json:"user_name" form:"user_name"  comment:"用户名" example:"admin" validate:"required,is-validuser"`
	Password string `json:"password"  form:"password"   comment:"密码" example:"123456" validate:"required"`
}

//看到第5-6， 06:34秒
func (param *AdminLoginInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

type AdminLoginOutput struct {
	Token string `json:"token" form:"token"  comment:"Token" example:"123456" validate:""`
}

type AdminSessionInfo struct {
	ID        int       `json:"id"`
	UserName  string    `json:"user_name"`
	LoginTime time.Time `json:"login_time"`
}
