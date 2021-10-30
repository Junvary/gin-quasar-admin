package data

import (
	"fmt"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var SysApi = new(sysApi)

type sysApi struct {
}

var sysApiData = []system.SysApi{
	{GqaModel: global.GqaModel{ Status: "on", Sort: 1, Desc: "获取用户列表", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "user", Path: "/user/user-list", Method: "POST",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 2, Desc: "编辑用户信息", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "user", Path: "/user/user-edit", Method: "PUT",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 3, Desc: "新增用户", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "user", Path: "/user/user-add", Method: "POST",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 4, Desc: "删除用户", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "user", Path: "/user/user-delete", Method: "DELETE",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 5, Desc: "根据ID查找用户", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "user", Path: "/user/user-id", Method: "POST",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 6, Desc: "获取用户的菜单", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "user", Path: "/user/user-menu", Method: "GET",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 7, Desc: "获取用户的角色列表", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "user", Path: "/user/user-role", Method: "GET",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 8, Desc: "获取角色列表", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "role", Path: "/role/role-list", Method: "POST",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 9, Desc: "编辑角色信息", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "role", Path: "/role/role-edit", Method: "PUT",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 10, Desc: "新增角色", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "role", Path: "/role/role-add", Method: "POST",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 11, Desc: "删除角色", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "role", Path: "/role/role-delete", Method: "DELETE",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 12, Desc: "根据ID查找角色", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "role", Path: "/role/role-id", Method: "POST",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 13, Desc: "获取角色菜单", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "role", Path: "/role/role-menu", Method: "POST",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 14, Desc: "编辑角色菜单", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "role", Path: "/role/role-menu-edit", Method: "PUT",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 15, Desc: "获取角色API", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "role", Path: "/role/role-api", Method: "POST",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 16, Desc: "编辑角色Api", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "role", Path: "/role/role-api-edit", Method: "PUT",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 16, Desc: "根据角色查找用户", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "role", Path: "role-user", Method: "POST",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 16, Desc: "从角色中移除某个用户", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "role", Path: "role-user-remove", Method: "POST",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 17, Desc: "获取菜单列表", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "menu", Path: "/menu/menu-list", Method: "POST",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 18, Desc: "编辑菜单信息", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "menu", Path: "/menu/menu-edit", Method: "PUT",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 19, Desc: "新增菜单", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "menu", Path: "/menu/menu-add", Method: "POST",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 20, Desc: "删除菜单", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "menu", Path: "/menu/menu-delete", Method: "DELETE",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 21, Desc: "根据ID查找菜单", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "menu", Path: "/menu/menu-id", Method: "POST",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 22, Desc: "获取部门列表", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "dept", Path: "/dept/dept-list", Method: "POST",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 23, Desc: "编辑部门信息", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "dept", Path: "/dept/dept-edit", Method: "PUT",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 24, Desc: "新增部门", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "dept", Path: "/dept/dept-add", Method: "POST",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 25, Desc: "删除部门", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "dept", Path: "/dept/dept-delete", Method: "DELETE",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 26, Desc: "根据ID查找部门", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "dept", Path: "/dept/dept-id", Method: "POST",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 27, Desc: "获取根字典列表", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "dict", Path: "/dict/dict-list", Method: "POST",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 28, Desc: "编辑字典信息", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "dict", Path: "/dict/dict-edit", Method: "PUT",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 29, Desc: "新增字典", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "dict", Path: "/dict/dict-add", Method: "POST",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 30, Desc: "删除字典", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "dict", Path: "/dict/dict-delete", Method: "DELETE",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 31, Desc: "根据ID查找字典", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "dict", Path: "/dict/dict-id", Method: "POST",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 32, Desc: "根据ParentId获取字典详情", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "dict", Path: "/dict/dict-parent-id", Method: "POST",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 33, Desc: "获取api列表", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "api", Path: "/api/api-list", Method: "POST",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 34, Desc: "编辑api信息", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "api", Path: "/api/api-edit", Method: "PUT",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 35, Desc: "新增api", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "api", Path: "/api/api-add", Method: "POST",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 36, Desc: "删除api", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "api", Path: "/api/api-delete", Method: "DELETE",
	},
	{GqaModel: global.GqaModel{ Status: "on", Sort: 37, Desc: "根据ID查找api", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		Group: "api", Path: "/api/api-id", Method: "POST",
	},
}

func (s *sysApi) Init() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&system.SysApi{}).Count(&count)
		if count != 0 {
			fmt.Println("[Gin-Quasar-Admin] --> sys_api 表的初始数据已存在，跳过初始化数据！数据量：", count)
			global.GqaLog.Error("sys_api 表的初始数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysApiData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[Gin-Quasar-Admin] --> sys_api 表初始数据成功！")
		global.GqaLog.Error("sys_api 表初始数据成功！")
		return nil
	})
}
