// @BeeOverwrite YES
// @BeeGenerateTime 20200831_174645
package model

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"oapms/pkg/mus"
	"oapms/pkg/trans"
)

type Department struct {
	Id          int    `gorm:"not null;primary_key;AUTO_INCREMENT"json:"id" form:"id"`      // id
	Name        string `gorm:"not null"json:"name" form:"name"`                             // 名称
	Pid         int    `gorm:"not null"json:"pid" form:"pid"`                               // 上级部门id
	OrderNum    int    `gorm:"not null"json:"orderNum" form:"orderNum"`                     // 排序
	ExtendField string `gorm:"not null;type:longtext"json:"extendField" form:"extendField"` // 扩展字段
	Intro       string `gorm:"not null"json:"intro" form:"intro"`                           // 介绍
	CreatedAt   int64  `gorm:"not null"json:"createdAt" form:"createdAt"`                   // 创建时间
	UpdatedAt   int64  `gorm:"not null"json:"updatedAt" form:"updatedAt"`                   // 更新时间

}

type Departments []*Department

func (t *Department) TableName() string {
	return "department"
}

// AddDepartment insert a new Department into database and returns
// last inserted Id on success.
func DepartmentCreate(db *gorm.DB, data *Department) (err error) {

	if err = db.Create(data).Error; err != nil {
		mus.Logger.Error("create department error", zap.Error(err))
		return
	}
	return
}

func DepartmentUpdate(db *gorm.DB, paramId int, ups Ups) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}

	if err = db.Table("department").Where(sql, binds...).Updates(ups).Error; err != nil {
		mus.Logger.Error("department update error", zap.Error(err))
		return
	}
	return
}

// UpdateX Update的扩展方法，根据Cond更新一条或多条记录
func DepartmentUpdateX(db *gorm.DB, conds Conds, ups Ups) (err error) {
	sql, binds := BuildQuery(conds)
	if err = db.Table("department").Where(sql, binds...).Updates(ups).Error; err != nil {
		mus.Logger.Error("department update error", zap.Error(err))
		return
	}
	return
}

// Delete 根据主键删除一条记录。如果有delete_time则软删除，否则硬删除。
func DepartmentDelete(db *gorm.DB, paramId int) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}

	if err = db.Table("department").Where(sql, binds...).Delete(&Department{}).Error; err != nil {
		mus.Logger.Error("department delete error", zap.Error(err))
		return
	}

	return
}

// DeleteX Delete的扩展方法，根据Cond删除一条或多条记录。如果有delete_time则软删除，否则硬删除。
func DepartmentDeleteX(db *gorm.DB, conds Conds) (err error) {
	sql, binds := BuildQuery(conds)

	if err = db.Table("department").Where(sql, binds...).Delete(&Department{}).Error; err != nil {
		mus.Logger.Error("department delete error", zap.Error(err))
		return
	}

	return
}

// Info 根据PRI查询单条记录
func DepartmentInfo(db *gorm.DB, paramId int) (resp Department, err error) {

	var sql = "`id`= ?"

	var binds = []interface{}{paramId}

	if err = db.Table("department").Where(sql, binds...).First(&resp).Error; err != nil {
		mus.Logger.Error("department info error", zap.Error(err))
		return
	}
	return
}

// InfoX Info的扩展方法，根据Cond查询单条记录
func DepartmentInfoX(db *gorm.DB, conds Conds) (resp Department, err error) {

	sql, binds := BuildQuery(conds)

	if err = db.Table("department").Where(sql, binds...).First(&resp).Error; err != nil {
		mus.Logger.Error("department info error", zap.Error(err))
		return
	}
	return
}

// List 查询list，extra[0]为sorts
func DepartmentList(conds Conds, extra ...string) (resp []*Department, err error) {

	sql, binds := BuildQuery(conds)

	sorts := ""
	if len(extra) >= 1 {
		sorts = extra[0]
	}
	if err = mus.Db.Table("department").Where(sql, binds...).Order(sorts).Find(&resp).Error; err != nil {
		mus.Logger.Error("department info error", zap.Error(err))
		return
	}
	return
}

// ListMap 查询map，map遍历的时候是无序的，所以指定sorts参数没有意义
func DepartmentListMap(conds Conds) (resp map[int]*Department, err error) {

	sql, binds := BuildQuery(conds)

	mysqlSlice := make([]*Department, 0)
	resp = make(map[int]*Department, 0)
	if err = mus.Db.Table("department").Where(sql, binds...).Find(&mysqlSlice).Error; err != nil {
		mus.Logger.Error("department info error", zap.Error(err))
		return
	}
	for _, value := range mysqlSlice {
		resp[value.Id] = value
	}
	return
}

// ListPage 根据分页条件查询list
func DepartmentListPage(conds Conds, reqList *trans.ReqPage) (total int, respList Departments) {
	respList = make(Departments, 0)

	if reqList.PageSize == 0 {
		reqList.PageSize = 10
	}
	if reqList.Current == 0 {
		reqList.Current = 1
	}
	sql, binds := BuildQuery(conds)

	db := mus.Db.Table("department").Where(sql, binds...)
	respList = make([]*Department, 0)
	db.Count(&total)
	db.Order(reqList.Sort).Offset((reqList.Current - 1) * reqList.PageSize).Limit(reqList.PageSize).Find(&respList)
	return
}
