// @BeeOverwrite YES
// @BeeGenerateTime 20200831_174645
package model

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"oapms/pkg/mus"
	"oapms/pkg/trans"
)

type Refresh struct {
	Id     int    `gorm:"auto"json:"id" form:"id"`              // ID
	Token  string `gorm:"size(255)"json:"token" form:"token"`   // token
	Access string `gorm:"size(255)"json:"access" form:"access"` // access

}

type Refreshs []*Refresh

func (t *Refresh) TableName() string {
	return "refresh"
}

// AddRefresh insert a new Refresh into database and returns
// last inserted Id on success.
func RefreshCreate(db *gorm.DB, data *Refresh) (err error) {

	if err = db.Create(data).Error; err != nil {
		mus.Logger.Error("create refresh error", zap.Error(err))
		return
	}
	return
}

func RefreshUpdate(db *gorm.DB, paramId int, ups Ups) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}

	if err = db.Table("refresh").Where(sql, binds...).Updates(ups).Error; err != nil {
		mus.Logger.Error("refresh update error", zap.Error(err))
		return
	}
	return
}

// UpdateX Update的扩展方法，根据Cond更新一条或多条记录
func RefreshUpdateX(db *gorm.DB, conds Conds, ups Ups) (err error) {
	sql, binds := BuildQuery(conds)
	if err = db.Table("refresh").Where(sql, binds...).Updates(ups).Error; err != nil {
		mus.Logger.Error("refresh update error", zap.Error(err))
		return
	}
	return
}

// Delete 根据主键删除一条记录。如果有delete_time则软删除，否则硬删除。
func RefreshDelete(db *gorm.DB, paramId int) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}

	if err = db.Table("refresh").Where(sql, binds...).Delete(&Refresh{}).Error; err != nil {
		mus.Logger.Error("refresh delete error", zap.Error(err))
		return
	}

	return
}

// DeleteX Delete的扩展方法，根据Cond删除一条或多条记录。如果有delete_time则软删除，否则硬删除。
func RefreshDeleteX(db *gorm.DB, conds Conds) (err error) {
	sql, binds := BuildQuery(conds)

	if err = db.Table("refresh").Where(sql, binds...).Delete(&Refresh{}).Error; err != nil {
		mus.Logger.Error("refresh delete error", zap.Error(err))
		return
	}

	return
}

// Info 根据PRI查询单条记录
func RefreshInfo(db *gorm.DB, paramId int) (resp Refresh, err error) {

	var sql = "`id`= ?"

	var binds = []interface{}{paramId}

	if err = db.Table("refresh").Where(sql, binds...).First(&resp).Error; err != nil {
		mus.Logger.Error("refresh info error", zap.Error(err))
		return
	}
	return
}

// InfoX Info的扩展方法，根据Cond查询单条记录
func RefreshInfoX(db *gorm.DB, conds Conds) (resp Refresh, err error) {

	sql, binds := BuildQuery(conds)

	if err = db.Table("refresh").Where(sql, binds...).First(&resp).Error; err != nil {
		mus.Logger.Error("refresh info error", zap.Error(err))
		return
	}
	return
}

// List 查询list，extra[0]为sorts
func RefreshList(conds Conds, extra ...string) (resp []*Refresh, err error) {

	sql, binds := BuildQuery(conds)

	sorts := ""
	if len(extra) >= 1 {
		sorts = extra[0]
	}
	if err = mus.Db.Table("refresh").Where(sql, binds...).Order(sorts).Find(&resp).Error; err != nil {
		mus.Logger.Error("refresh info error", zap.Error(err))
		return
	}
	return
}

// ListMap 查询map，map遍历的时候是无序的，所以指定sorts参数没有意义
func RefreshListMap(conds Conds) (resp map[int]*Refresh, err error) {

	sql, binds := BuildQuery(conds)

	mysqlSlice := make([]*Refresh, 0)
	resp = make(map[int]*Refresh, 0)
	if err = mus.Db.Table("refresh").Where(sql, binds...).Find(&mysqlSlice).Error; err != nil {
		mus.Logger.Error("refresh info error", zap.Error(err))
		return
	}
	for _, value := range mysqlSlice {
		resp[value.Id] = value
	}
	return
}

// ListPage 根据分页条件查询list
func RefreshListPage(conds Conds, reqList *trans.ReqPage) (total int, respList Refreshs) {
	respList = make(Refreshs, 0)

	if reqList.PageSize == 0 {
		reqList.PageSize = 10
	}
	if reqList.Current == 0 {
		reqList.Current = 1
	}
	sql, binds := BuildQuery(conds)

	db := mus.Db.Table("refresh").Where(sql, binds...)
	respList = make([]*Refresh, 0)
	db.Count(&total)
	db.Order(reqList.Sort).Offset((reqList.Current - 1) * reqList.PageSize).Limit(reqList.PageSize).Find(&respList)
	return
}
