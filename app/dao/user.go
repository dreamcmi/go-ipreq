package dao

import "github.com/gogf/gf/os/gtime"

// User 存储结构
type User struct {
	UserIP   string      `orm:"userIP"   json:"userIP"`     // 用户IP
	UserPort string      `orm:"userPort" json:"userPort"`   // 用户端口
	CreateAt *gtime.Time `orm:"createAt"  json:"createAt"` // 创建时间
}
