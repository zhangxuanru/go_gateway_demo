/*
@Time : 2020/9/29 19:20
@Author : zxr
@File : admin
@Software: GoLand
*/
package dto

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zhangxuanru/go_gateway_demo/public"
)

//管理员信息
type AdminInfoOutput struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	LoginTime    time.Time `json:"login_time"`
	Avatar       string    `json:"avatar"`
	Introduction string    `json:"introduction"`
	Roles        []string  `json:"roles"`
}

//修改密码
type AdminChangePwdInput struct {
	OldPassWd string `json:"old_pass_wd" form:"old_pass_wd"  comment:"旧密码" example:"123456" validate:"required"`
	NewPassWd string `json:"new_pass_wd" form:"new_pass_wd"  comment:"新密码" example:"123456" validate:"required"`
}

func (param *AdminChangePwdInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}
