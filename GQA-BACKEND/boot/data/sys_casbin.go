package data

import (
	"fmt"
	"gin-quasar-admin/global"
	adapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var SysCasbin = new(sysCasbin)

type sysCasbin struct{}

var sysCasbinData = []adapter.CasbinRule{
	// user组
	{Ptype: "p", V0: "super-admin", V1: "/user/user-list", V2: "POST", V3: "user", V4: "获取用户列表"},
	{Ptype: "p", V0: "super-admin", V1: "/user/user-edit", V2: "PUT", V3: "user", V4: "编辑用户信息"},
	{Ptype: "p", V0: "super-admin", V1: "/user/user-add", V2: "POST", V3: "user", V4: "新增用户"},
	{Ptype: "p", V0: "super-admin", V1: "/user/user-delete", V2: "DELETE", V3: "user", V4: "删除用户"},
	{Ptype: "p", V0: "super-admin", V1: "/user/user-id", V2: "POST", V3: "user", V4: "根据ID查找用户"},
	{Ptype: "p", V0: "super-admin", V1: "/user/user-menu", V2: "GET", V3: "user", V4: "获取用户的菜单"},
	{Ptype: "p", V0: "super-admin", V1: "/user/user-role", V2: "GET", V3: "user", V4: "获取用户的角色列表"},
	// role组
	{Ptype: "p", V0: "super-admin", V1: "/role/role-list", V2: "POST", V3: "role", V4: "获取角色列表"},
	{Ptype: "p", V0: "super-admin", V1: "/role/role-edit", V2: "PUT", V3: "role", V4: "编辑角色信息"},
	{Ptype: "p", V0: "super-admin", V1: "/role/role-add", V2: "POST", V3: "role", V4: "新增角色"},
	{Ptype: "p", V0: "super-admin", V1: "/role/role-delete", V2: "DELETE", V3: "role", V4: "删除角色"},
	{Ptype: "p", V0: "super-admin", V1: "/role/role-id", V2: "POST", V3: "role", V4: "根据ID查找角色"},
	{Ptype: "p", V0: "super-admin", V1: "/role/role-menu", V2: "POST", V3: "role", V4: "获取角色菜单"},
	{Ptype: "p", V0: "super-admin", V1: "/role/role-menu-edit", V2: "PUT", V3: "role", V4: "编辑角色菜单"},
	{Ptype: "p", V0: "super-admin", V1: "/role/role-api", V2: "POST", V3: "role", V4: "获取角色API"},
	{Ptype: "p", V0: "super-admin", V1: "/role/role-api-edit", V2: "PUT", V3: "role", V4: "编辑角色Api"},
	// menu组
	{Ptype: "p", V0: "super-admin", V1: "/menu/menu-list", V2: "POST", V3: "menu", V4: "获取菜单列表"},
	{Ptype: "p", V0: "super-admin", V1: "/menu/menu-edit", V2: "PUT", V3: "menu", V4: "编辑菜单信息"},
	{Ptype: "p", V0: "super-admin", V1: "/menu/menu-add", V2: "POST", V3: "menu", V4: "新增菜单"},
	{Ptype: "p", V0: "super-admin", V1: "/menu/menu-delete", V2: "DELETE", V3: "menu", V4: "删除菜单"},
	{Ptype: "p", V0: "super-admin", V1: "/menu/menu-id", V2: "POST", V3: "menu", V4: "根据ID查找菜单"},
	// dept组
	{Ptype: "p", V0: "super-admin", V1: "/dept/dept-list", V2: "POST", V3: "dept", V4: "获取部门列表"},
	{Ptype: "p", V0: "super-admin", V1: "/dept/dept-edit", V2: "PUT", V3: "dept", V4: "编辑部门信息"},
	{Ptype: "p", V0: "super-admin", V1: "/dept/dept-add", V2: "POST", V3: "dept", V4: "新增部门"},
	{Ptype: "p", V0: "super-admin", V1: "/dept/dept-delete", V2: "DELETE", V3: "dept", V4: "删除部门"},
	{Ptype: "p", V0: "super-admin", V1: "/dept/dept-id", V2: "POST", V3: "dept", V4: "根据ID查找部门"},
	// dict组
	{Ptype: "p", V0: "super-admin", V1: "/dict/dict-list", V2: "POST", V3: "dict", V4: "获取根字典列表"},
	{Ptype: "p", V0: "super-admin", V1: "/dict/dict-edit", V2: "PUT", V3: "dict", V4: "编辑字典信息"},
	{Ptype: "p", V0: "super-admin", V1: "/dict/dict-add", V2: "POST", V3: "dict", V4: "新增字典"},
	{Ptype: "p", V0: "super-admin", V1: "/dict/dict-delete", V2: "DELETE", V3: "dict", V4: "删除字典"},
	{Ptype: "p", V0: "super-admin", V1: "/dict/dict-id", V2: "POST", V3: "dict", V4: "根据ID查找字典"},
	{Ptype: "p", V0: "super-admin", V1: "/dict/dict-parent-id", V2: "POST", V3: "dict", V4: "根据ParentId获取字典详情"},

	// api组
	{Ptype: "p", V0: "super-admin", V1: "/api/api-list", V2: "POST", V3: "api", V4: "获取api列表"},
	{Ptype: "p", V0: "super-admin", V1: "/api/api-edit", V2: "PUT", V3: "api", V4: "编辑api信息"},
	{Ptype: "p", V0: "super-admin", V1: "/api/api-add", V2: "POST", V3: "api", V4: "新增api"},
	{Ptype: "p", V0: "super-admin", V1: "/api/api-delete", V2: "DELETE", V3: "api", V4: "删除api"},
	{Ptype: "p", V0: "super-admin", V1: "/api/api-id", V2: "POST", V3: "api", V4: "根据ID查找api"},
}

func (s *sysCasbin) Init() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&[]adapter.CasbinRule{}).Count(&count)
		if count != 0 {
			fmt.Println("[Gin-Quasar-Admin] --> casbin_rule 表的初始数据已存在！数据量：", count)
			global.GqaLog.Error("casbin_rule 表的初始数据已存在！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysCasbinData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[Gin-Quasar-Admin] --> casbin_rule 表初始数据成功！")
		global.GqaLog.Error("casbin_rule 表初始数据成功！")
		return nil
	})
}
