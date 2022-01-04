package system

import (
	"errors"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
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
	//创建循环节点
	Loop:
	for _, v := range tList {
		switch v {
		case "all":
			//如果权限列表中包含 all，那么直接放弃其他判断，跳出循环
			permissionDb = db
			break Loop
		case "user":
			//用户自己的数据权限，直接查找创建人为自己的数据
			permissionDb = tempDb.Where("created_by = ?", username)
		case "dept":
			//部门数据权限
			var deptList []system.SysDept
			if err = global.GqaDb.Model(&user).Association("Dept").Find(&deptList); err != nil {
				return err, nil
			}
			var deptUserList []string
			//通过用户所在部门查找这些部门的用户集
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
			//包含子部门的部门数据权限
			var deptList []system.SysDept
			if err = global.GqaDb.Model(&user).Association("Dept").Find(&deptList); err != nil {
				return err, nil
			}
			//用户所在部门，且包含所有子部门的用户集
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
			//自定义部门数据权限
			var deptUserList []string
			for _, dept := range cList {
				var deptUser []system.SysDeptUser
				global.GqaDb.Where("sys_dept_dept_code = ?", dept).Find(&deptUser)
				for _, u := range deptUser {
					deptUserList = append(deptUserList, u.Username)
				}
			}
			allUser := utils.RemoveDuplicateElementFromSlice(deptUserList)
			permissionDb = tempDb.Or(tempDb.Where("created_by in ?", allUser))
		default:
			permissionDb = tempDb
			return errors.New("没有数据权限配置！"), nil
		}
	}
	return nil, permissionDb
}

func GetChildrenFromDept(deptCode string) (dl []string) {
	//递归函数，查出所有子部门和部门
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

