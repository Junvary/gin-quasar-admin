package system

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gin-quasar-admin/utils"
	"gorm.io/gorm"
	"strings"
)

func DeptDataPermission(username string, db *gorm.DB) (err error, permissionDb *gorm.DB) {
	user := system.SysUser{
		Username: username,
	}
	var roleList []system.SysRole
	if err = global.GqaDb.Model(&user).Association("Role").Find(&roleList); err != nil {
		return err, nil
	}
	var permissionTypeList []string
	var permissionCustomList []string
	for _, v := range roleList {
		permissionTypeList = append(permissionTypeList, v.DeptDataPermissionType)
		if v.DeptDataPermissionCustom != "" {
			ss := strings.Split(v.DeptDataPermissionCustom, ",")
			permissionCustomList = append(permissionCustomList, ss...)
		}
	}
	tList := utils.RemoveDuplicateElementFromSlice(permissionTypeList)
	cList := utils.RemoveDuplicateElementFromSlice(permissionCustomList)
	tempDb := db
	Loop:
	for _, v := range tList {
		switch v {
		case "all":
			permissionDb = db
			break Loop
		case "user":
			permissionDb = tempDb.Where("created_by = ?", username)
		case "dept":
			var deptList []system.SysDept
			if err = global.GqaDb.Model(&user).Association("Dept").Find(&deptList); err != nil {
				return err, nil
			}
			var deptUserList []string
			for _, dept := range deptList {
				var deptUser []system.SysDeptUser
				global.GqaDb.Where("sys_dept_dept_code = ?", dept.DeptCode).Find(&deptUser)
				for _, u := range deptUser {
					deptUserList = append(deptUserList, u.Username)
				}
			}
			allUser := utils.RemoveDuplicateElementFromSlice(deptUserList)
			permissionDb = tempDb.Or(tempDb.Where("created_by in ?", allUser))
		case "deptWithChildren":
			var deptList []system.SysDept
			if err = global.GqaDb.Model(&user).Association("Dept").Find(&deptList); err != nil {
				return err, nil
			}
			var deptListTotal []string
			for _, dept := range deptList {
				deptListTotal = append(deptListTotal, dept.DeptCode)
				deptListTotal = append(deptListTotal, GetChildrenFromDept(dept.DeptCode)...)
			}
			deptListTotal = utils.RemoveDuplicateElementFromSlice(deptListTotal)
			var deptUserList []string
			for _, dept := range deptListTotal {
				var deptUser []system.SysDeptUser
				global.GqaDb.Where("sys_dept_dept_code = ?", dept).Find(&deptUser)
				for _, u := range deptUser {
					deptUserList = append(deptUserList, u.Username)
				}
			}
			allUser := utils.RemoveDuplicateElementFromSlice(deptUserList)
			permissionDb = tempDb.Or(tempDb.Where("created_by in ?", allUser))
		case "custom":
			var deptListTotal []string
			for _, cl:= range cList{
				deptListTotal = append(deptListTotal, cl)
				deptListTotal= append(deptListTotal, GetChildrenFromDept(cl)...)
			}
			deptListTotal = utils.RemoveDuplicateElementFromSlice(deptListTotal)
			var deptUserList []string
			for _, dept := range deptListTotal {
				var deptUser []system.SysDeptUser
				global.GqaDb.Where("sys_dept_dept_code = ?", dept).Find(&deptUser)
				for _, u := range deptUser {
					deptUserList = append(deptUserList, u.Username)
				}
			}
			allUser := utils.RemoveDuplicateElementFromSlice(deptUserList)
			permissionDb = tempDb.Or(tempDb.Where("created_by in ?", allUser))
		}
	}
	return nil, permissionDb
}

func GetChildrenFromDept(deptCode string) (dl []string) {
	var deptList []system.SysDept
	global.GqaDb.Where("parent_code = ?", deptCode).Find(&deptList)
	var deptListString []string
	for _, d :=range deptList{
		deptListString = append(deptListString, d.DeptCode)
	}
	if len(deptList) != 0{
		for _, d := range deptList{
			deptListString = append(deptListString, GetChildrenFromDept(d.DeptCode)...)
		}
	}
	return deptListString
}

