package data

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var SysRoleApi = new(sysRoleApi)

type sysRoleApi struct{}

func (s *sysRoleApi) LoadData() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&model.SysRoleApi{}).Count(&count)
		if count != 0 {
			fmt.Println(utils.GqaI18nWithData("SkipInsertWithData", "sys_role_api"), count)
			global.GqaLogger.Warn(utils.GqaI18nWithData("SkipInsertWithData", "sys_role_api"), zap.Any("count", count))
			return nil
		}
		if err := tx.Create(&sysRoleApiData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println(utils.GqaI18nWithData("TableInitSuccess", "sys_role_api"))
		global.GqaLogger.Info(utils.GqaI18nWithData("TableInitSuccess", "sys_role_api"))
		return nil
	})
}

var sysRoleApiData = []model.SysRoleApi{
	{RoleCode: "super-admin", ApiGroup: "user", ApiMethod: "POST", ApiPath: "/user/get-user-menu"},
	{RoleCode: "super-admin", ApiGroup: "user", ApiMethod: "POST", ApiPath: "/user/get-user-list"},
	{RoleCode: "super-admin", ApiGroup: "user", ApiMethod: "POST", ApiPath: "/user/edit-user"},
	{RoleCode: "super-admin", ApiGroup: "user", ApiMethod: "POST", ApiPath: "/user/add-user"},
	{RoleCode: "super-admin", ApiGroup: "user", ApiMethod: "POST", ApiPath: "/user/delete-user-by-id"},
	{RoleCode: "super-admin", ApiGroup: "user", ApiMethod: "POST", ApiPath: "/user/query-user-by-id"},
	{RoleCode: "super-admin", ApiGroup: "user", ApiMethod: "POST", ApiPath: "/user/reset-password"},
	{RoleCode: "super-admin", ApiGroup: "user", ApiMethod: "POST", ApiPath: "/user/change-password"},
	{RoleCode: "super-admin", ApiGroup: "user", ApiMethod: "POST", ApiPath: "/user/change-nickname"},

	{RoleCode: "super-admin", ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/get-role-list"},
	{RoleCode: "super-admin", ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/edit-role"},
	{RoleCode: "super-admin", ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/add-role"},
	{RoleCode: "super-admin", ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/delete-role-by-id"},
	{RoleCode: "super-admin", ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/query-role-by-id"},
	{RoleCode: "super-admin", ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/get-role-menu-list"},
	{RoleCode: "super-admin", ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/edit-role-menu"},
	{RoleCode: "super-admin", ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/get-role-api-list"},
	{RoleCode: "super-admin", ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/edit-role-api"},
	{RoleCode: "super-admin", ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/query-user-by-role"},
	{RoleCode: "super-admin", ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/remove-role-user"},
	{RoleCode: "super-admin", ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/add-role-user"},
	{RoleCode: "super-admin", ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/edit-role-dept-data-permission"},
	{RoleCode: "super-admin", ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/get-role-button-list"},

	{RoleCode: "super-admin", ApiGroup: "menu", ApiMethod: "POST", ApiPath: "/menu/get-menu-list"},
	{RoleCode: "super-admin", ApiGroup: "menu", ApiMethod: "POST", ApiPath: "/menu/edit-menu"},
	{RoleCode: "super-admin", ApiGroup: "menu", ApiMethod: "POST", ApiPath: "/menu/add-menu"},
	{RoleCode: "super-admin", ApiGroup: "menu", ApiMethod: "POST", ApiPath: "/menu/delete-menu-by-id"},
	{RoleCode: "super-admin", ApiGroup: "menu", ApiMethod: "POST", ApiPath: "/menu/query-menu-by-id"},

	{RoleCode: "super-admin", ApiGroup: "dept", ApiMethod: "POST", ApiPath: "/dept/get-dept-list"},
	{RoleCode: "super-admin", ApiGroup: "dept", ApiMethod: "POST", ApiPath: "/dept/edit-dept"},
	{RoleCode: "super-admin", ApiGroup: "dept", ApiMethod: "POST", ApiPath: "/dept/add-dept"},
	{RoleCode: "super-admin", ApiGroup: "dept", ApiMethod: "POST", ApiPath: "/dept/delete-dept-by-id"},
	{RoleCode: "super-admin", ApiGroup: "dept", ApiMethod: "POST", ApiPath: "/dept/query-dept-by-id"},
	{RoleCode: "super-admin", ApiGroup: "dept", ApiMethod: "POST", ApiPath: "/dept/query-user-by-dept"},
	{RoleCode: "super-admin", ApiGroup: "dept", ApiMethod: "POST", ApiPath: "/dept/remove-dept-user"},
	{RoleCode: "super-admin", ApiGroup: "dept", ApiMethod: "POST", ApiPath: "/dept/add-dept-user"},

	{RoleCode: "super-admin", ApiGroup: "dict", ApiMethod: "POST", ApiPath: "/dict/get-dict-list"},
	{RoleCode: "super-admin", ApiGroup: "dict", ApiMethod: "POST", ApiPath: "/dict/edit-dict"},
	{RoleCode: "super-admin", ApiGroup: "dict", ApiMethod: "POST", ApiPath: "/dict/add-dict"},
	{RoleCode: "super-admin", ApiGroup: "dict", ApiMethod: "POST", ApiPath: "/dict/delete-dict-by-id"},
	{RoleCode: "super-admin", ApiGroup: "dict", ApiMethod: "POST", ApiPath: "/dict/query-dict-by-id"},

	{RoleCode: "super-admin", ApiGroup: "api", ApiMethod: "POST", ApiPath: "/api/get-api-list"},
	{RoleCode: "super-admin", ApiGroup: "api", ApiMethod: "POST", ApiPath: "/api/edit-api"},
	{RoleCode: "super-admin", ApiGroup: "api", ApiMethod: "POST", ApiPath: "/api/add-api"},
	{RoleCode: "super-admin", ApiGroup: "api", ApiMethod: "POST", ApiPath: "/api/delete-api-by-id"},
	{RoleCode: "super-admin", ApiGroup: "api", ApiMethod: "POST", ApiPath: "/api/query-api-by-id"},

	{RoleCode: "super-admin", ApiGroup: "upload", ApiMethod: "POST", ApiPath: "/upload/upload-avatar"},
	{RoleCode: "super-admin", ApiGroup: "upload", ApiMethod: "POST", ApiPath: "/upload/upload-file"},
	{RoleCode: "super-admin", ApiGroup: "upload", ApiMethod: "POST", ApiPath: "/upload/upload-banner-image"},
	{RoleCode: "super-admin", ApiGroup: "upload", ApiMethod: "POST", ApiPath: "/upload/upload-logo"},
	{RoleCode: "super-admin", ApiGroup: "upload", ApiMethod: "POST", ApiPath: "/upload/upload-favicon"},

	{RoleCode: "super-admin", ApiGroup: "config-backend", ApiMethod: "POST", ApiPath: "/config-backend/get-config-backend-list"},
	{RoleCode: "super-admin", ApiGroup: "config-backend", ApiMethod: "POST", ApiPath: "/config-backend/edit-config-backend"},
	{RoleCode: "super-admin", ApiGroup: "config-backend", ApiMethod: "POST", ApiPath: "/config-backend/add-config-backend"},
	{RoleCode: "super-admin", ApiGroup: "config-backend", ApiMethod: "POST", ApiPath: "/config-backend/delete-config-backend-by-id"},

	{RoleCode: "super-admin", ApiGroup: "config-frontend", ApiMethod: "POST", ApiPath: "/config-frontend/get-config-frontend-list"},
	{RoleCode: "super-admin", ApiGroup: "config-frontend", ApiMethod: "POST", ApiPath: "/config-frontend/edit-config-frontend"},
	{RoleCode: "super-admin", ApiGroup: "config-frontend", ApiMethod: "POST", ApiPath: "/config-frontend/add-config-frontend"},
	{RoleCode: "super-admin", ApiGroup: "config-frontend", ApiMethod: "POST", ApiPath: "/config-frontend/delete-config-frontend-by-id"},

	{RoleCode: "super-admin", ApiGroup: "log", ApiMethod: "POST", ApiPath: "/log/get-log-login-list"},
	{RoleCode: "super-admin", ApiGroup: "log", ApiMethod: "POST", ApiPath: "/log/delete-log-login-by-id"},
	{RoleCode: "super-admin", ApiGroup: "log", ApiMethod: "POST", ApiPath: "/log/get-log-operation-list"},
	{RoleCode: "super-admin", ApiGroup: "log", ApiMethod: "POST", ApiPath: "/log/delete-log-operation-by-id"},

	{RoleCode: "super-admin", ApiGroup: "notice", ApiMethod: "POST", ApiPath: "/notice/get-notice-list"},
	{RoleCode: "super-admin", ApiGroup: "notice", ApiMethod: "POST", ApiPath: "/notice/add-notice"},
	{RoleCode: "super-admin", ApiGroup: "notice", ApiMethod: "POST", ApiPath: "/notice/delete-notice-by-id"},
	{RoleCode: "super-admin", ApiGroup: "notice", ApiMethod: "POST", ApiPath: "/notice/query-notice-by-id"},
	{RoleCode: "super-admin", ApiGroup: "notice", ApiMethod: "POST", ApiPath: "/notice/query-notice-read-by-id"},
	{RoleCode: "super-admin", ApiGroup: "notice", ApiMethod: "POST", ApiPath: "/notice/send-notice"},

	{RoleCode: "super-admin", ApiGroup: "todo", ApiMethod: "POST", ApiPath: "/todo/get-todo-list"},
	{RoleCode: "super-admin", ApiGroup: "todo", ApiMethod: "POST", ApiPath: "/todo/edit-todo"},
	{RoleCode: "super-admin", ApiGroup: "todo", ApiMethod: "POST", ApiPath: "/todo/add-todo"},
	{RoleCode: "super-admin", ApiGroup: "todo", ApiMethod: "POST", ApiPath: "/todo/delete-todo-by-id"},
	{RoleCode: "super-admin", ApiGroup: "todo", ApiMethod: "POST", ApiPath: "/todo/query-todo-by-id"},

	{RoleCode: "super-admin", ApiGroup: "genPlugin", ApiMethod: "POST", ApiPath: "/gen-plugin/gen-plugin"},

	{RoleCode: "super-admin", ApiGroup: "user-online", ApiMethod: "POST", ApiPath: "/user-online/get-user-online-list"},
	{RoleCode: "super-admin", ApiGroup: "user-online", ApiMethod: "POST", ApiPath: "/user-online/kick-online-user"},

	{RoleCode: "super-admin", ApiGroup: "cron", ApiMethod: "POST", ApiPath: "/cron/get-cron-list"},
	{RoleCode: "super-admin", ApiGroup: "cron", ApiMethod: "POST", ApiPath: "/cron/start-cron"},
	{RoleCode: "super-admin", ApiGroup: "cron", ApiMethod: "POST", ApiPath: "/cron/stop-cron"},
}
