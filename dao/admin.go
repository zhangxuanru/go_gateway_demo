/*
@Time : 2020/9/29 16:53
@Author : zxr
@File : admin
@Software: GoLand
*/
package dao

import (
	"time"

	"github.com/pkg/errors"

	"github.com/zhangxuanru/go_gateway_demo/dto"

	"github.com/zhangxuanru/go_gateway_demo/public"

	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
)

type Admin struct {
	Id        int       `json:"id" gorm:"primary_key" description:"自增主键"`
	UserName  string    `json:"user_name" gorm:"column:user_name" description:"用户名"`
	Salt      string    `json:"salt" gorm:"column:salt" description:"盐"`
	Password  string    `json:"password" gorm:"column:password" description:"密码"`
	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at" description:"更新时间"`
	CreatedAt time.Time `json:"create_at" gorm:"column:create_at" description:"创建时间"`
	IsDelete  int       `json:"is_delete" gorm:"column:is_delete" description:"是否删除"`
}

func (a *Admin) TableName() string {
	return "gateway_admin"
}

func (a *Admin) LoginCheck(c *gin.Context, tx *gorm.DB, param *dto.AdminLoginInput) (*Admin, error) {
	adminInfo, err := a.Find(c, tx, &Admin{UserName: param.UserName, IsDelete: 0})
	if err != nil {
		return nil, errors.New("user is not found")
	}
	saltPassword := public.GenSaltPassword(adminInfo.Salt, param.Password)
	if saltPassword != adminInfo.Password {
		return nil, errors.New("password is error, try ")
	}
	return adminInfo, nil
}

func (a *Admin) Find(c *gin.Context, tx *gorm.DB, search *Admin) (*Admin, error) {
	out := &Admin{}
	err := tx.SetCtx(public.GetGinTraceContext(c)).Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}
