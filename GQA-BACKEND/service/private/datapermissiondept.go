package private

import (
	"errors"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
)

func DeptDataPermission(c *gin.Context, username string, db *gorm.DB) (err error, permissionDb *gorm.DB) {
	user := model.SysUser{
		Username: username,
	}
	var dataPermissionList []model.SysRole
	if err = global.GqaDb.Model(&user).Association("Role").Find(&dataPermissionList); err != nil {
		return err, nil
	}
	var permissionTypeList []string
	var permissionCustomList []string
	for _, v := range dataPermissionList {
		permissionTypeList = append(permissionTypeList, v.DeptDataPermissionType)
		if v.DeptDataPermissionCustom != "" {
			ss := strings.Split(v.DeptDataPermissionCustom, ",")
			permissionCustomList = append(permissionCustomList, ss...)
		}
	}
	tList := utils.SliceSortCompact(permissionTypeList)
	cList := utils.SliceSortCompact(permissionCustomList)
	tempDb := db

Loop:
	for _, v := range tList {
		switch v {
		case "deptDataPermissionType_all":
			// If "all" is included in the permission list, you can directly give up other judgments and jump out of the loop
			permissionDb = db
			break Loop
		case "deptDataPermissionType_user":
			// The user's own data permission can directly find the data created by himself
			permissionDb = tempDb.Where("created_by = ?", username)
		case "deptDataPermissionType_dept":
			// Department Data Permission
			var deptList []model.SysDept
			if err = global.GqaDb.Model(&user).Association("Dept").Find(&deptList); err != nil {
				return err, nil
			}
			var deptUserList []string
			// Find the user set of these departments by their departments
			for _, dept := range deptList {
				var deptUser []model.SysDeptUser
				global.GqaDb.Where("sys_dept_dept_code = ?", dept.DeptCode).Find(&deptUser)
				for _, u := range deptUser {
					deptUserList = append(deptUserList, u.SysUserUsername)
				}
			}
			allUser := utils.SliceSortCompact(deptUserList)
			permissionDb = tempDb.Or(tempDb.Where("created_by in ?", allUser))
		case "deptDataPermissionType_deptAndChildren":
			// Department data permission including sub departments
			var deptList []model.SysDept
			if err = global.GqaDb.Model(&user).Association("Dept").Find(&deptList); err != nil {
				return err, nil
			}
			// The user set of the user's department, including all sub departments
			var deptListTotal []string
			for _, dept := range deptList {
				deptListTotal = append(deptListTotal, dept.DeptCode)
				deptListTotal = append(deptListTotal, GetChildrenFromDept(dept.DeptCode)...)
			}
			deptListTotal = utils.SliceSortCompact(deptListTotal)
			var deptUserList []string
			for _, dept := range deptListTotal {
				var deptUser []model.SysDeptUser
				global.GqaDb.Where("sys_dept_dept_code = ?", dept).Find(&deptUser)
				for _, u := range deptUser {
					deptUserList = append(deptUserList, u.SysUserUsername)
				}
			}
			allUser := utils.SliceSortCompact(deptUserList)
			permissionDb = tempDb.Or(tempDb.Where("created_by in ?", allUser))
		case "deptDataPermissionType_custom":
			// User defined department data permission
			var deptUserList []string
			for _, dept := range cList {
				var deptUser []model.SysDeptUser
				global.GqaDb.Where("sys_dept_dept_code = ?", dept).Find(&deptUser)
				for _, u := range deptUser {
					deptUserList = append(deptUserList, u.SysUserUsername)
				}
			}
			allUser := utils.SliceSortCompact(deptUserList)
			permissionDb = tempDb.Or(tempDb.Where("created_by in ?", allUser))
		default:
			permissionDb = tempDb
			return errors.New(utils.GqaI18n(c, "NoConfig")), nil
		}
	}
	return nil, permissionDb
}

func GetChildrenFromDept(deptCode string) (dl []string) {
	// find out all sub departments and departments
	var deptList []model.SysDept
	global.GqaDb.Where("parent_code = ?", deptCode).Find(&deptList)
	var deptListString []string
	for _, d := range deptList {
		deptListString = append(deptListString, d.DeptCode)
	}
	if len(deptList) != 0 {
		for _, d := range deptList {
			deptListString = append(deptListString, GetChildrenFromDept(d.DeptCode)...)
		}
	}
	return deptListString
}
