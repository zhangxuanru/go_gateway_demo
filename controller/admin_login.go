/*
@Time : 2020/9/25 19:08
@Author : zxr
@File : admin_login
@Software: GoLand
*/
package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangxuanru/go_gateway_demo/dto"
	"github.com/zhangxuanru/go_gateway_demo/middleware"
)

type AdminLoginController struct {
}

func RegisterAdminLogin(router *gin.RouterGroup) {
	adminLogin := &AdminLoginController{}
	router.GET("/login", adminLogin.AdminLogin)
}

func (adminLogin *AdminLoginController) AdminLogin(c *gin.Context) {
	params := &dto.AdminLoginInput{}
	if err := params.BindValidParam(c); err != nil {
	    middleware.ResponseError(c, 1001, err)
		return
	}
	middleware.ResponseSuccess(c, "")
}
