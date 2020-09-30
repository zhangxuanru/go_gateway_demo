/*
@Time : 2020/9/25 19:08
@Author : zxr
@File : admin_login
@Software: GoLand
*/
package controller

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/zhangxuanru/go_gateway_demo/dao"

	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zhangxuanru/go_gateway_demo/dto"
	"github.com/zhangxuanru/go_gateway_demo/middleware"
	"github.com/zhangxuanru/go_gateway_demo/public"
)

type AdminController struct {
}

func RegisterAdmin(router *gin.RouterGroup) {
	admin := &AdminController{}
	router.GET("/admin_info", admin.AdminInfo)
	router.POST("/change_pwd", admin.ChangePwd)
}

// AdminLogin godoc
// @Summary 管理员信息获取
// @Description 管理员信息获取
// @Tags 管理员接口
// @ID /admin/admin_info
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=dto.AdminInfoOutput} "success"
// @Router /admin/admin_info [get]
func (a *AdminController) AdminInfo(c *gin.Context) {
	sess := sessions.Default(c)
	sessInfo := sess.Get(public.SESSION_ADMIN_INFO_KEY)
	adminSess := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(fmt.Sprint(sessInfo)), adminSess); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	out := &dto.AdminInfoOutput{
		ID:           adminSess.ID,
		Name:         adminSess.UserName,
		LoginTime:    adminSess.LoginTime,
		Avatar:       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		Introduction: "I am a super administrator",
		Roles:        []string{"admin"},
	}
	middleware.ResponseSuccess(c, out)
}

// AdminLogin godoc
// @Summary 管理员修改密码
// @Description 管理员修改密码
// @Tags 管理员接口
// @ID /admin/change_pwd
// @Accept  json
// @Produce  json
// @Param body body dto.AdminChangePwdInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /admin/change_pwd [post]
func (a *AdminController) ChangePwd(c *gin.Context) {
	sessData := sessions.Default(c).Get(public.SESSION_ADMIN_INFO_KEY)
	sessionInfo := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(fmt.Sprint(sessData)), sessionInfo); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	param := &dto.AdminChangePwdInput{}
	if err := param.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	admin := &dao.Admin{
		Id: sessionInfo.ID,
	}
	admin, err = admin.Find(c, tx, admin)
	if err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	password := public.GenSaltPassword(admin.Salt, param.OldPassWd)
	if password != admin.Password {
		middleware.ResponseError(c, 2003, errors.New("原密码不对"))
		return
	}
	saltPassword := public.GenSaltPassword(admin.Salt, param.NewPassWd)
	admin.Password = saltPassword
	if err = admin.Edit(c, tx); err != nil {
		middleware.ResponseError(c, 2004, err)
		return
	}
	middleware.ResponseSuccess(c, "")
	return
}
