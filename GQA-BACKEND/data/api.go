package data

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var SysApi = new(sysApi)

type sysApi struct{}

func (s *sysApi) LoadData() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&model.SysApi{}).Count(&count)
		if count != 0 {
			fmt.Println(utils.GqaI18nWithData("SkipInsertWithData", "sys_api"), count)
			global.GqaLogger.Warn(utils.GqaI18nWithData("SkipInsertWithData", "sys_api"), zap.Any("count", count))
			return nil
		}
		if err := tx.Create(&sysApiData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println(utils.GqaI18nWithData("TableInitSuccess", "sys_api"))
		global.GqaLogger.Info(utils.GqaI18nWithData("TableInitSuccess", "sys_api"))
		return nil
	})
}

var sysApiData = []model.SysApi{
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 1, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取用户的菜单（必选）",
	}}, ApiGroup: "user", ApiMethod: "POST", ApiPath: "/user/get-user-menu"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 2, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取用户列表",
	}}, ApiGroup: "user", ApiMethod: "POST", ApiPath: "/user/get-user-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 3, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "编辑用户信息",
	}}, ApiGroup: "user", ApiMethod: "POST", ApiPath: "/user/edit-user"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 4, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "新增用户",
	}}, ApiGroup: "user", ApiMethod: "POST", ApiPath: "/user/add-user"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 5, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "删除用户",
	}}, ApiGroup: "user", ApiMethod: "POST", ApiPath: "/user/delete-user-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 6, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "根据ID查找用户",
	}}, ApiGroup: "user", ApiMethod: "POST", ApiPath: "/user/query-user-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 7, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "重置用户密码",
	}}, ApiGroup: "user", ApiMethod: "POST", ApiPath: "/user/reset-password"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 8, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "用户修改密码",
	}}, ApiGroup: "user", ApiMethod: "POST", ApiPath: "/user/change-password"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 9, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "用户修改昵称",
	}}, ApiGroup: "user", ApiMethod: "POST", ApiPath: "/user/change-nickname"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 10, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取角色列表",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/get-role-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 11, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "编辑角色信息",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/edit-role"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 12, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "新增角色",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/add-role"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 13, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "删除角色",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/delete-role-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 14, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "根据ID查找角色",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/query-role-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 15, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取角色菜单",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/get-role-menu-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 16, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "编辑角色菜单",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/edit-role-menu"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 17, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取角色API",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/get-role-api-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 18, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "编辑角色Api",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/edit-role-api"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 19, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "根据角色查找用户",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/query-user-by-role"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 20, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "从角色中移除某个用户",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/remove-role-user"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 21, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "添加用户到某个角色",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/add-role-user"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 22, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "修改角色部门数据权限",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/edit-role-dept-data-permission"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 23, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取角色按钮列表",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/get-role-button-list"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 24, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取菜单列表",
	}}, ApiGroup: "menu", ApiMethod: "POST", ApiPath: "/menu/get-menu-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 25, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "编辑菜单信息",
	}}, ApiGroup: "menu", ApiMethod: "POST", ApiPath: "/menu/edit-menu"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 26, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "新增菜单",
	}}, ApiGroup: "menu", ApiMethod: "POST", ApiPath: "/menu/add-menu"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 27, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "删除菜单",
	}}, ApiGroup: "menu", ApiMethod: "POST", ApiPath: "/menu/delete-menu-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 28, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "根据ID查找菜单",
	}}, ApiGroup: "menu", ApiMethod: "POST", ApiPath: "/menu/query-menu-by-id"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 29, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取部门列表",
	}}, ApiGroup: "dept", ApiMethod: "POST", ApiPath: "/dept/get-dept-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 30, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "编辑部门信息",
	}}, ApiGroup: "dept", ApiMethod: "POST", ApiPath: "/dept/edit-dept"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 31, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "新增部门",
	}}, ApiGroup: "dept", ApiMethod: "POST", ApiPath: "/dept/add-dept"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 32, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "删除部门",
	}}, ApiGroup: "dept", ApiMethod: "POST", ApiPath: "/dept/delete-dept-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 33, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "根据ID查找部门",
	}}, ApiGroup: "dept", ApiMethod: "POST", ApiPath: "/dept/query-dept-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 34, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "根据部门查找用户",
	}}, ApiGroup: "dept", ApiMethod: "POST", ApiPath: "/dept/query-user-by-dept"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 35, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "从部门中移除某个用户",
	}}, ApiGroup: "dept", ApiMethod: "POST", ApiPath: "/dept/remove-dept-user"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 36, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "添加用户到部门",
	}}, ApiGroup: "dept", ApiMethod: "POST", ApiPath: "/dept/add-dept-user"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 37, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取根字典列表",
	}}, ApiGroup: "dict", ApiMethod: "POST", ApiPath: "/dict/get-dict-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 38, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "编辑字典信息",
	}}, ApiGroup: "dict", ApiMethod: "POST", ApiPath: "/dict/edit-dict"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 39, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "新增字典",
	}}, ApiGroup: "dict", ApiMethod: "POST", ApiPath: "/dict/add-dict"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 40, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "删除字典",
	}}, ApiGroup: "dict", ApiMethod: "POST", ApiPath: "/dict/delete-dict-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 41, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "根据ID查找字典",
	}}, ApiGroup: "dict", ApiMethod: "POST", ApiPath: "/dict/query-dict-by-id"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 42, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取api列表",
	}}, ApiGroup: "api", ApiMethod: "POST", ApiPath: "/api/get-api-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 43, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "编辑api信息",
	}}, ApiGroup: "api", ApiMethod: "POST", ApiPath: "/api/edit-api"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 44, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "新增api",
	}}, ApiGroup: "api", ApiMethod: "POST", ApiPath: "/api/add-api"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 45, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "根据id删除api",
	}}, ApiGroup: "api", ApiMethod: "POST", ApiPath: "/api/delete-api-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 46, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "根据id查找api",
	}}, ApiGroup: "api", ApiMethod: "POST", ApiPath: "/api/query-api-by-id"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 47, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "上传头像",
	}}, ApiGroup: "upload", ApiMethod: "POST", ApiPath: "/upload/upload-avatar"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 48, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "上传文件",
	}}, ApiGroup: "upload", ApiMethod: "POST", ApiPath: "/upload/upload-file"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 49, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "上传首页大图",
	}}, ApiGroup: "upload", ApiMethod: "POST", ApiPath: "/upload/upload-banner-image"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 50, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "上传网站Logo",
	}}, ApiGroup: "upload", ApiMethod: "POST", ApiPath: "/upload/upload-logo"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 51, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "上传标签页Logo",
	}}, ApiGroup: "upload", ApiMethod: "POST", ApiPath: "/upload/upload-favicon"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 52, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取后台配置列表",
	}}, ApiGroup: "config-backend", ApiMethod: "POST", ApiPath: "/config-backend/get-config-backend-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 53, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "编辑后台配置",
	}}, ApiGroup: "config-backend", ApiMethod: "POST", ApiPath: "/config-backend/edit-config-backend"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 54, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "新增后台配置",
	}}, ApiGroup: "config-backend", ApiMethod: "POST", ApiPath: "/config-backend/add-config-backend"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 55, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "删除后台配置",
	}}, ApiGroup: "config-backend", ApiMethod: "POST", ApiPath: "/config-backend/delete-config-backend-by-id"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 56, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取前台配置列表",
	}}, ApiGroup: "config-frontend", ApiMethod: "POST", ApiPath: "/config-frontend/get-config-frontend-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 57, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "编辑前台配置",
	}}, ApiGroup: "config-frontend", ApiMethod: "POST", ApiPath: "/config-frontend/edit-config-frontend"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 58, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "新增前台配置",
	}}, ApiGroup: "config-frontend", ApiMethod: "POST", ApiPath: "/config-frontend/add-config-frontend"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 59, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "删除前台配置",
	}}, ApiGroup: "config-frontend", ApiMethod: "POST", ApiPath: "/config-frontend/delete-config-frontend-by-id"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 60, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取登录日志列表",
	}}, ApiGroup: "log", ApiMethod: "POST", ApiPath: "/log/get-log-login-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 61, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "删除登录日志",
	}}, ApiGroup: "log", ApiMethod: "POST", ApiPath: "/log/delete-log-login-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 62, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取操作日志列表",
	}}, ApiGroup: "log", ApiMethod: "POST", ApiPath: "/log/get-log-operation-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 63, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "删除操作日志",
	}}, ApiGroup: "log", ApiMethod: "POST", ApiPath: "/log/delete-log-operation-by-id"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 64, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取消息列表",
	}}, ApiGroup: "notice", ApiMethod: "POST", ApiPath: "/notice/get-notice-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 65, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "新增消息"}},
		ApiGroup: "notice", ApiMethod: "POST", ApiPath: "/notice/add-notice"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 66, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "删除消息",
	}}, ApiGroup: "notice", ApiMethod: "POST", ApiPath: "/notice/delete-notice-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 67, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "根据ID查找消息",
	}}, ApiGroup: "notice", ApiMethod: "POST", ApiPath: "/notice/query-notice-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 68, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "根据ID查找消息并阅读",
	}}, ApiGroup: "notice", ApiMethod: "POST", ApiPath: "/notice/query-notice-read-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 69, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "发送消息",
	}}, ApiGroup: "notice", ApiMethod: "POST", ApiPath: "/notice/send-notice"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 70, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取待办列表",
	}}, ApiGroup: "todo", ApiMethod: "POST", ApiPath: "/todo/get-todo-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 71, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "编辑待办信息",
	}}, ApiGroup: "todo", ApiMethod: "POST", ApiPath: "/todo/edit-todo"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 72, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "新增待办",
	}}, ApiGroup: "todo", ApiMethod: "POST", ApiPath: "/todo/add-todo"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 73, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "删除待办",
	}}, ApiGroup: "todo", ApiMethod: "POST", ApiPath: "/todo/delete-todo-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 74, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "根据ID查找待办",
	}}, ApiGroup: "todo", ApiMethod: "POST", ApiPath: "/todo/query-todo-by-id"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 75, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件代码生成",
	}}, ApiGroup: "genPlugin", ApiMethod: "POST", ApiPath: "/gen-plugin/gen-plugin"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 76, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取在线用户列表",
	}}, ApiGroup: "user-online", ApiMethod: "POST", ApiPath: "/user-online/get-user-online-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 77, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "踢出在线用户",
	}}, ApiGroup: "user-online", ApiMethod: "POST", ApiPath: "/user-online/kick-online-user"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 78, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取定时任务列表",
	}}, ApiGroup: "cron", ApiMethod: "POST", ApiPath: "/cron/get-cron-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 79, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "启动定时任务",
	}}, ApiGroup: "cron", ApiMethod: "POST", ApiPath: "/cron/start-cron"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 80, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "停止定时任务",
	}}, ApiGroup: "cron", ApiMethod: "POST", ApiPath: "/cron/stop-cron"},
}
