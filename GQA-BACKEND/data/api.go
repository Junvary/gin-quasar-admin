package data

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
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
			fmt.Println("[Gin-Quasar-Admin] --> sys_api 表的初始数据已存在，跳过初始化数据！数据量：", count)
			global.GqaLogger.Warn("[Gin-Quasar-Admin] --> sys_api 表的初始数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysApiData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[Gin-Quasar-Admin] --> sys_api 表初始数据成功！")
		global.GqaLogger.Info("[Gin-Quasar-Admin] --> sys_api 表初始数据成功！")
		return nil
	})
}

var sysApiData = []model.SysApi{
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10001, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取用户的菜单（必选）",
	}}, ApiGroup: "user", ApiMethod: "POST", ApiPath: "/user/get-user-menu"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10002, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取用户列表",
	}}, ApiGroup: "user", ApiMethod: "POST", ApiPath: "/user/get-user-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10003, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "编辑用户信息",
	}}, ApiGroup: "user", ApiMethod: "POST", ApiPath: "/user/edit-user"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10004, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "新增用户",
	}}, ApiGroup: "user", ApiMethod: "POST", ApiPath: "/user/add-user"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10005, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "删除用户",
	}}, ApiGroup: "user", ApiMethod: "POST", ApiPath: "/user/delete-user-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10006, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "根据ID查找用户",
	}}, ApiGroup: "user", ApiMethod: "POST", ApiPath: "/user/query-user-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10007, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "重置用户密码",
	}}, ApiGroup: "user", ApiMethod: "POST", ApiPath: "/user/reset-password"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10008, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "用户修改密码",
	}}, ApiGroup: "user", ApiMethod: "POST", ApiPath: "/user/change-password"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10009, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "用户修改昵称",
	}}, ApiGroup: "user", ApiMethod: "POST", ApiPath: "/user/change-nickname"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10010, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取角色列表",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/get-role-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10011, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "编辑角色信息",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/edit-role"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10012, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "新增角色",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/add-role"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10013, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "删除角色",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/delete-role-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10014, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "根据ID查找角色",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/query-role-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10015, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取角色菜单",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/get-role-menu-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10016, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "编辑角色菜单",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/edit-role-menu"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10017, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取角色API",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/get-role-api-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10018, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "编辑角色Api",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/edit-role-api"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10019, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "根据角色查找用户",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/query-user-by-role"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10020, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "从角色中移除某个用户",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/remove-role-user"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10021, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "添加用户到某个角色",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/add-role-user"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10022, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "修改角色部门数据权限",
	}}, ApiGroup: "role", ApiMethod: "POST", ApiPath: "/role/edit-role-dept-data-permission"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10023, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取菜单列表",
	}}, ApiGroup: "menu", ApiMethod: "POST", ApiPath: "/menu/get-menu-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10024, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "编辑菜单信息",
	}}, ApiGroup: "menu", ApiMethod: "POST", ApiPath: "/menu/edit-menu"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10025, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "新增菜单",
	}}, ApiGroup: "menu", ApiMethod: "POST", ApiPath: "/menu/add-menu"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10026, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "删除菜单",
	}}, ApiGroup: "menu", ApiMethod: "POST", ApiPath: "/menu/delete-menu-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10027, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "根据ID查找菜单",
	}}, ApiGroup: "menu", ApiMethod: "POST", ApiPath: "/menu/query-menu-by-id"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10028, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取部门列表",
	}}, ApiGroup: "dept", ApiMethod: "POST", ApiPath: "/dept/get-dept-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10029, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "编辑部门信息",
	}}, ApiGroup: "dept", ApiMethod: "POST", ApiPath: "/dept/edit-dept"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10030, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "新增部门",
	}}, ApiGroup: "dept", ApiMethod: "POST", ApiPath: "/dept/add-dept"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10031, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "删除部门",
	}}, ApiGroup: "dept", ApiMethod: "POST", ApiPath: "/dept/delete-dept-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10032, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "根据ID查找部门",
	}}, ApiGroup: "dept", ApiMethod: "POST", ApiPath: "/dept/query-dept-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10033, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "根据部门查找用户",
	}}, ApiGroup: "dept", ApiMethod: "POST", ApiPath: "/dept/query-user-by-dept"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10034, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "从部门中移除某个用户",
	}}, ApiGroup: "dept", ApiMethod: "POST", ApiPath: "/dept/remove-dept-user"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10035, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "添加用户到部门",
	}}, ApiGroup: "dept", ApiMethod: "POST", ApiPath: "/dept/add-dept-user"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10036, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取根字典列表",
	}}, ApiGroup: "dict", ApiMethod: "POST", ApiPath: "/dict/get-dict-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10037, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "编辑字典信息",
	}}, ApiGroup: "dict", ApiMethod: "POST", ApiPath: "/dict/edit-dict"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10038, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "新增字典",
	}}, ApiGroup: "dict", ApiMethod: "POST", ApiPath: "/dict/add-dict"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10039, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "删除字典",
	}}, ApiGroup: "dict", ApiMethod: "POST", ApiPath: "/dict/delete-dict-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10040, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "根据ID查找字典",
	}}, ApiGroup: "dict", ApiMethod: "POST", ApiPath: "/dict/query-dict-by-id"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10041, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取api列表",
	}}, ApiGroup: "api", ApiMethod: "POST", ApiPath: "/api/get-api-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10042, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "编辑api信息",
	}}, ApiGroup: "api", ApiMethod: "POST", ApiPath: "/api/edit-api"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10043, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "新增api",
	}}, ApiGroup: "api", ApiMethod: "POST", ApiPath: "/api/add-api"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10044, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "根据id删除api",
	}}, ApiGroup: "api", ApiMethod: "POST", ApiPath: "/api/delete-api-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10045, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "根据id查找api",
	}}, ApiGroup: "api", ApiMethod: "POST", ApiPath: "/api/query-api-by-id"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10046, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "上传头像",
	}}, ApiGroup: "upload", ApiMethod: "POST", ApiPath: "/upload/upload-avatar"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10047, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "上传文件",
	}}, ApiGroup: "upload", ApiMethod: "POST", ApiPath: "/upload/upload-file"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10048, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "上传首页大图",
	}}, ApiGroup: "upload", ApiMethod: "POST", ApiPath: "/upload/upload-banner-image"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10049, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "上传网站Logo",
	}}, ApiGroup: "upload", ApiMethod: "POST", ApiPath: "/upload/upload-logo"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10050, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "上传标签页Logo",
	}}, ApiGroup: "upload", ApiMethod: "POST", ApiPath: "/upload/upload-favicon"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10051, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取后台配置列表",
	}}, ApiGroup: "config-backend", ApiMethod: "POST", ApiPath: "/config-backend/get-config-backend-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10052, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "编辑后台配置",
	}}, ApiGroup: "config-backend", ApiMethod: "POST", ApiPath: "/config-backend/edit-config-backend"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10053, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "新增后台配置",
	}}, ApiGroup: "config-backend", ApiMethod: "POST", ApiPath: "/config-backend/add-config-backend"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10054, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "删除后台配置",
	}}, ApiGroup: "config-backend", ApiMethod: "POST", ApiPath: "/config-backend/delete-config-backend-by-id"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10055, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取前台配置列表",
	}}, ApiGroup: "config-frontend", ApiMethod: "POST", ApiPath: "/config-frontend/get-config-frontend-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10056, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "编辑前台配置",
	}}, ApiGroup: "config-frontend", ApiMethod: "POST", ApiPath: "/config-frontend/edit-config-frontend"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10057, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "新增前台配置",
	}}, ApiGroup: "config-frontend", ApiMethod: "POST", ApiPath: "/config-frontend/add-config-frontend"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10058, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "删除前台配置",
	}}, ApiGroup: "config-frontend", ApiMethod: "POST", ApiPath: "/config-frontend/delete-config-frontend-by-id"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10059, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取登录日志列表",
	}}, ApiGroup: "log", ApiMethod: "POST", ApiPath: "/log/get-log-login-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10060, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "删除登录日志",
	}}, ApiGroup: "log", ApiMethod: "POST", ApiPath: "/log/delete-log-login-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10061, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取操作日志列表",
	}}, ApiGroup: "log", ApiMethod: "POST", ApiPath: "/log/get-log-operation-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10062, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "删除登录日志",
	}}, ApiGroup: "log", ApiMethod: "POST", ApiPath: "/log/delete-log-operation-by-id"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10063, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取消息列表",
	}}, ApiGroup: "notice", ApiMethod: "POST", ApiPath: "/notice/get-notice-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10064, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "编辑消息信息",
	}}, ApiGroup: "notice", ApiMethod: "POST", ApiPath: "/notice/edit-notice"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10065, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "新增消息"}},
		ApiGroup: "notice", ApiMethod: "POST", ApiPath: "/notice/add-notice"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10066, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "删除消息",
	}}, ApiGroup: "notice", ApiMethod: "POST", ApiPath: "/notice/delete-notice-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10067, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "根据ID查找消息",
	}}, ApiGroup: "notice", ApiMethod: "POST", ApiPath: "/notice/query-notice-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10068, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "根据ID查找消息并阅读",
	}}, ApiGroup: "notice", ApiMethod: "POST", ApiPath: "/notice/query-notice-read-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10069, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "发送消息",
	}}, ApiGroup: "notice", ApiMethod: "POST", ApiPath: "/notice/send-notice"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10070, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取待办便签列表",
	}}, ApiGroup: "todoNote", ApiMethod: "POST", ApiPath: "/note-todo/get-note-todo-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10071, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "编辑待办便签信息",
	}}, ApiGroup: "todoNote", ApiMethod: "POST", ApiPath: "/note-todo/edit-note-todo"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10072, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "新增待办便签",
	}}, ApiGroup: "todoNote", ApiMethod: "POST", ApiPath: "/note-todo/add-note-todo"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10073, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "删除待办便签",
	}}, ApiGroup: "todoNote", ApiMethod: "POST", ApiPath: "/note-todo/delete-note-todo-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10074, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "根据ID查找待办便签",
	}}, ApiGroup: "todoNote", ApiMethod: "POST", ApiPath: "/note-todo/query-note-todo-by-id"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10075, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件代码生成",
	}}, ApiGroup: "genPlugin", ApiMethod: "POST", ApiPath: "/gen-plugin/gen-plugin"},

	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10076, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "获取在线用户列表",
	}}, ApiGroup: "user-online", ApiMethod: "POST", ApiPath: "/user-online/get-user-online-list"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10077, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "踢出在线用户",
	}}, ApiGroup: "user-online", ApiMethod: "POST", ApiPath: "/user-online/kick-online-user"},
}
