/*
@Time : 2020/9/29 19:20
@Author : zxr
@File : admin
@Software: GoLand
*/
package dto

import "time"

type AdminInfoOutput struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	LoginTime    time.Time `json:"login_time"`
	Avatar       string    `json:"avatar"`
	Introduction string    `json:"introduction"`
	Roles        []string  `json:"roles"`
}
