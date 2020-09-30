/*
@Time : 2020/9/25 19:08
@Author : zxr
@File : admin_login
@Software: GoLand
*/
package controller

import (
	"encoding/json"
	"time"

	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	dao2 "github.com/zhangxuanru/go_gateway_demo/dao"
	"github.com/zhangxuanru/go_gateway_demo/dto"
	"github.com/zhangxuanru/go_gateway_demo/middleware"
	"github.com/zhangxuanru/go_gateway_demo/public"
)

type AdminLoginController struct {
}

func RegisterAdminLogin(router *gin.RouterGroup) {
	adminLogin := &AdminLoginController{}
	router.POST("/login", adminLogin.AdminLogin)
	router.GET("/login_out", adminLogin.AdminLoginOut)
}

// AdminLogin godoc
// @Summary 管理员登录
// @Description 管理员登录
// @Tags 管理员接口
// @ID /admin_login/login
// @Accept  json
// @Produce  json
// @Param body body dto.AdminLoginInput true "body"
// @Success 200 {object} middleware.Response{data=dto.AdminLoginOutput} "success"
// @Router /admin_login/login [post]
func (a *AdminLoginController) AdminLogin(c *gin.Context) {
	params := &dto.AdminLoginInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 1001, err)
		return
	}
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	dao := &dao2.Admin{}
	admin, err := dao.LoginCheck(c, tx, params)
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	session := sessions.Default(c)
	sessionAdmin := &dto.AdminSessionInfo{
		ID:        admin.Id,
		UserName:  admin.UserName,
		LoginTime: time.Now(),
	}
	sessBytes, err := json.Marshal(sessionAdmin)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	session.Set(public.SESSION_ADMIN_INFO_KEY, string(sessBytes))
	session.Save()
	out := &dto.AdminLoginOutput{Token: admin.UserName}
	middleware.ResponseSuccess(c, out)
}

// AdminLogin godoc
// @Summary 管理员退出
// @Description 管理员退出
// @Tags 管理员接口
// @ID /admin_login/login_out
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /admin_login/login_out [get]
func (a *AdminLoginController) AdminLoginOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete(public.SESSION_ADMIN_INFO_KEY)
	session.Save()
	middleware.ResponseSuccess(c, "")
}
