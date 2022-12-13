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
		Sort: 1, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取用户的菜单（必选）",
	}}, ApiGroup: "user", ApiMethod: "POST", ApiPath: "/user/get-user-menu"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 2, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取用户列表",
	}}, ApiGroup: "user", ApiMethod: "POST", ApiPath: "/user/get-user-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 3, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "编辑用户信息",
	}}, ApiGroup: "user", ApiMethod: "POST", ApiPath: "/user/edit-user"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 4, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "新增用户",
	}}, ApiGroup: "user", ApiMethod: "POST", ApiPath: "/user/add-user"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 5, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "删除用户",
	}}, ApiGroup: "user", ApiMethod: "POST", ApiPath: "/user/delete-user-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 6, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "根据ID查找用户",
	}}, ApiGroup: "user", ApiMethod: "POST", ApiPath: "/user/query-user-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 7, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "重置用户密码",
	}}, ApiGroup: "user", ApiMethod: "POST", ApiPath: "/user/reset-password"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 8, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "用户修改密码",
	}}, ApiGroup: "user", ApiMethod: "POST", ApiPath: "/user/change-password"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 9, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "用户修改昵称",
	}}, ApiGroup: "user", ApiMethod: "POST", ApiPath: "/user/change-nickname"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取角色列表",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/get-role-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 11, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "编辑角色信息",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/edit-role"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 12, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "新增角色",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/add-role"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 13, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "删除角色",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/delete-role-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 14, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "根据ID查找角色",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/query-role-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 15, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取角色菜单",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/get-role-menu-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 16, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "编辑角色菜单",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/edit-role-menu"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 17, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取角色API",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/get-role-api-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 18, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "编辑角色Api",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/edit-role-api"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 19, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "根据角色查找用户",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/query-user-by-role"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 20, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "从角色中移除某个用户",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/remove-role-user"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 21, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "添加用户到某个角色",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/add-role-user"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 22, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "修改角色部门数据权限",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/edit-role-dept-data-permission"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 23, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取角色按钮列表",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/get-role-button-list"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 24, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取菜单列表",
	}}, ApiGroup: "menu", ApiMethod: "POST", ApiPath: "/menu/get-menu-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 25, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "编辑菜单信息",
	}}, ApiGroup: "menu", ApiMethod: "POST", ApiPath: "/menu/edit-menu"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 26, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "新增菜单",
	}}, ApiGroup: "menu", ApiMethod: "POST", ApiPath: "/menu/add-menu"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 27, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "删除菜单",
	}}, ApiGroup: "menu", ApiMethod: "POST", ApiPath: "/menu/delete-menu-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 28, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "根据ID查找菜单",
	}}, ApiGroup: "menu", ApiMethod: "POST", ApiPath: "/menu/query-menu-by-id"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 29, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取部门列表",
	}}, ApiGroup: "dept", ApiMethod: "POST", ApiPath: "/dept/get-dept-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 30, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "编辑部门信息",
	}}, ApiGroup: "dept", ApiMethod: "POST", ApiPath: "/dept/edit-dept"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 31, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "新增部门",
	}}, ApiGroup: "dept", ApiMethod: "POST", ApiPath: "/dept/add-dept"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 32, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "删除部门",
	}}, ApiGroup: "dept", ApiMethod: "POST", ApiPath: "/dept/delete-dept-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 33, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "根据ID查找部门",
	}}, ApiGroup: "dept", ApiMethod: "POST", ApiPath: "/dept/query-dept-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 34, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "根据部门查找用户",
	}}, ApiGroup: "dept", ApiMethod: "POST", ApiPath: "/dept/query-user-by-dept"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 35, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "从部门中移除某个用户",
	}}, ApiGroup: "dept", ApiMethod: "POST", ApiPath: "/dept/remove-dept-user"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 36, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "添加用户到部门",
	}}, ApiGroup: "dept", ApiMethod: "POST", ApiPath: "/dept/add-dept-user"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 37, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取根字典列表",
	}}, ApiGroup: "dict", ApiMethod: "POST", ApiPath: "/dict/get-dict-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 38, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "编辑字典信息",
	}}, ApiGroup: "dict", ApiMethod: "POST", ApiPath: "/dict/edit-dict"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 39, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "新增字典",
	}}, ApiGroup: "dict", ApiMethod: "POST", ApiPath: "/dict/add-dict"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 40, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "删除字典",
	}}, ApiGroup: "dict", ApiMethod: "POST", ApiPath: "/dict/delete-dict-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 41, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "根据ID查找字典",
	}}, ApiGroup: "dict", ApiMethod: "POST", ApiPath: "/dict/query-dict-by-id"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 42, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取api列表",
	}}, ApiGroup: "api", ApiMethod: "POST", ApiPath: "/api/get-api-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 43, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "编辑api信息",
	}}, ApiGroup: "api", ApiMethod: "POST", ApiPath: "/api/edit-api"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 44, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "新增api",
	}}, ApiGroup: "api", ApiMethod: "POST", ApiPath: "/api/add-api"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 45, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "根据id删除api",
	}}, ApiGroup: "api", ApiMethod: "POST", ApiPath: "/api/delete-api-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 46, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "根据id查找api",
	}}, ApiGroup: "api", ApiMethod: "POST", ApiPath: "/api/query-api-by-id"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 47, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "上传头像",
	}}, ApiGroup: "upload", ApiMethod: "POST", ApiPath: "/upload/upload-avatar"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 48, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "上传文件",
	}}, ApiGroup: "upload", ApiMethod: "POST", ApiPath: "/upload/upload-file"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 49, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "上传首页大图",
	}}, ApiGroup: "upload", ApiMethod: "POST", ApiPath: "/upload/upload-banner-image"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 50, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "上传网站Logo",
	}}, ApiGroup: "upload", ApiMethod: "POST", ApiPath: "/upload/upload-logo"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 51, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "上传标签页Logo",
	}}, ApiGroup: "upload", ApiMethod: "POST", ApiPath: "/upload/upload-favicon"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 52, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取后台配置列表",
	}}, ApiGroup: "config-backend", ApiMethod: "POST", ApiPath: "/config-backend/get-config-backend-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 53, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "编辑后台配置",
	}}, ApiGroup: "config-backend", ApiMethod: "POST", ApiPath: "/config-backend/edit-config-backend"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 54, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "新增后台配置",
	}}, ApiGroup: "config-backend", ApiMethod: "POST", ApiPath: "/config-backend/add-config-backend"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 55, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "删除后台配置",
	}}, ApiGroup: "config-backend", ApiMethod: "POST", ApiPath: "/config-backend/delete-config-backend-by-id"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 56, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取前台配置列表",
	}}, ApiGroup: "config-frontend", ApiMethod: "POST", ApiPath: "/config-frontend/get-config-frontend-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 57, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "编辑前台配置",
	}}, ApiGroup: "config-frontend", ApiMethod: "POST", ApiPath: "/config-frontend/edit-config-frontend"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 58, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "新增前台配置",
	}}, ApiGroup: "config-frontend", ApiMethod: "POST", ApiPath: "/config-frontend/add-config-frontend"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 59, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "删除前台配置",
	}}, ApiGroup: "config-frontend", ApiMethod: "POST", ApiPath: "/config-frontend/delete-config-frontend-by-id"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 60, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取登录日志列表",
	}}, ApiGroup: "log", ApiMethod: "POST", ApiPath: "/log/get-log-login-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 61, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "删除登录日志",
	}}, ApiGroup: "log", ApiMethod: "POST", ApiPath: "/log/delete-log-login-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 62, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取操作日志列表",
	}}, ApiGroup: "log", ApiMethod: "POST", ApiPath: "/log/get-log-operation-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 63, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "删除操作日志",
	}}, ApiGroup: "log", ApiMethod: "POST", ApiPath: "/log/delete-log-operation-by-id"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 64, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取消息列表",
	}}, ApiGroup: "notice", ApiMethod: "POST", ApiPath: "/notice/get-notice-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 65, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "新增消息"}},
		ApiGroup: "notice", ApiMethod: "POST", ApiPath: "/notice/add-notice"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 66, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "删除消息",
	}}, ApiGroup: "notice", ApiMethod: "POST", ApiPath: "/notice/delete-notice-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 67, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "根据ID查找消息",
	}}, ApiGroup: "notice", ApiMethod: "POST", ApiPath: "/notice/query-notice-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 68, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "根据ID查找消息并阅读",
	}}, ApiGroup: "notice", ApiMethod: "POST", ApiPath: "/notice/query-notice-read-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 69, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "发送消息",
	}}, ApiGroup: "notice", ApiMethod: "POST", ApiPath: "/notice/send-notice"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 70, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取待办便签列表",
	}}, ApiGroup: "todoNote", ApiMethod: "POST", ApiPath: "/note-todo/get-note-todo-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 71, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "编辑待办便签信息",
	}}, ApiGroup: "todoNote", ApiMethod: "POST", ApiPath: "/note-todo/edit-note-todo"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 72, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "新增待办便签",
	}}, ApiGroup: "todoNote", ApiMethod: "POST", ApiPath: "/note-todo/add-note-todo"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 73, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "删除待办便签",
	}}, ApiGroup: "todoNote", ApiMethod: "POST", ApiPath: "/note-todo/delete-note-todo-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 74, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "根据ID查找待办便签",
	}}, ApiGroup: "todoNote", ApiMethod: "POST", ApiPath: "/note-todo/query-note-todo-by-id"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 75, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件代码生成",
	}}, ApiGroup: "genPlugin", ApiMethod: "POST", ApiPath: "/gen-plugin/gen-plugin"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 76, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取在线用户列表",
	}}, ApiGroup: "user-online", ApiMethod: "POST", ApiPath: "/user-online/get-user-online-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 77, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "踢出在线用户",
	}}, ApiGroup: "user-online", ApiMethod: "POST", ApiPath: "/user-online/kick-online-user"},
}
