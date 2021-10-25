package system

import (
	"errors"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gorm.io/gorm"
)

type ServiceDept struct {
}

func (s *ServiceDept)GetDeptList(pageInfo system.RequestPage) (err error, role interface{}, total int64) {
	pageSize := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	db := global.GqaDb.Model(&system.SysDept{})
	var deptList []system.SysDept
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Find(&deptList).Error
	return err, deptList, total
}

func (s *ServiceDept) EditDept(dept system.SysDept) (err error) {
	err = global.GqaDb.Updates(&dept).Error
	return err
}

func (s *ServiceDept) AddDept(d system.SysDept) (err error) {
	var dept system.SysDept
	if !errors.Is(global.GqaDb.Where("dept_code = ?", d.DeptCode).First(&dept).Error, gorm.ErrRecordNotFound) {
		return errors.New("此部门已存在：" + d.DeptCode)
	}
	err = global.GqaDb.Create(&d).Error
	return err
}

func (s *ServiceDept) DeleteDept(id uint) (err error) {
	var dept system.SysDept
	err = global.GqaDb.Where("id = ?", id).Delete(&dept).Error
	return err
}

func (s *ServiceDept) QueryDeptById(id uint) (err error, deptInfo system.SysDept) {
	var dept system.SysDept
	err = global.GqaDb.Preload("Owner").First(&dept, "id = ?", id).Error
	return err, dept
}