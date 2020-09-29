/*
@Time : 2020/9/25 19:08
@Author : zxr
@File : admin_login
@Software: GoLand
*/
package controller

import (
	"encoding/json"
	"fmt"

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
