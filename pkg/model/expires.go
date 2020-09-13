// @BeeOverwrite YES
// @BeeGenerateTime 20200831_174645
package model

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"oapms/pkg/mus"
	"oapms/pkg/trans"
)

type Expires struct {
	Id        int    `gorm:"auto"json:"id" form:"id"`            // 客户端
	Token     string `gorm:"size(255)"json:"token" form:"token"` // token
	ExpiresAt int64  `json:"expiresAt" form:"expiresAt"`         // 过期时间

}

type Expiress []*Expires

func (t *Expires) TableName() string {
	return "expires"
}

// AddExpires insert a new Expires into database and returns
// last inserted Id on success.
func ExpiresCreate(db *gorm.DB, data *Expires) (err error) {

	if err = db.Create(data).Error; err != nil {
		mus.Logger.Error("create expires error", zap.Error(err))
		return
	}
	return
}

func ExpiresUpdate(db *gorm.DB, paramId int, ups Ups) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}

	if err = db.Table("expires").Where(sql, binds...).Updates(ups).Error; err != nil {
		mus.Logger.Error("expires update error", zap.Error(err))
		return
	}
	return
}

// UpdateX Update的扩展方法，根据Cond更新一条或多条记录
func ExpiresUpdateX(db *gorm.DB, conds Conds, ups Ups) (err error) {
	sql, binds := BuildQuery(conds)
	if err = db.Table("expires").Where(sql, binds...).Updates(ups).Error; err != nil {
		mus.Logger.Error("expires update error", zap.Error(err))
		return
	}
	return
}

// Delete 根据主键删除一条记录。如果有delete_time则软删除，否则硬删除。
func ExpiresDelete(db *gorm.DB, paramId int) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}

	if err = db.Table("expires").Where(sql, binds...).Delete(&Expires{}).Error; err != nil {
		mus.Logger.Error("expires delete error", zap.Error(err))
		return
	}

	return
}

// DeleteX Delete的扩展方法，根据Cond删除一条或多条记录。如果有delete_time则软删除，否则硬删除。
func ExpiresDeleteX(db *gorm.DB, conds Conds) (err error) {
	sql, binds := BuildQuery(conds)

	if err = db.Table("expires").Where(sql, binds...).Delete(&Expires{}).Error; err != nil {
		mus.Logger.Error("expires delete error", zap.Error(err))
		return
	}

	return
}

// Info 根据PRI查询单条记录
func ExpiresInfo(db *gorm.DB, paramId int) (resp Expires, err error) {

	var sql = "`id`= ?"

	var binds = []interface{}{paramId}

	if err = db.Table("expires").Where(sql, binds...).First(&resp).Error; err != nil {
		mus.Logger.Error("expires info error", zap.Error(err))
		return
	}
	return
}

// InfoX Info的扩展方法，根据Cond查询单条记录
func ExpiresInfoX(db *gorm.DB, conds Conds) (resp Expires, err error) {

	sql, binds := BuildQuery(conds)

	if err = db.Table("expires").Where(sql, binds...).First(&resp).Error; err != nil {
		mus.Logger.Error("expires info error", zap.Error(err))
		return
	}
	return
}

// List 查询list，extra[0]为sorts
func ExpiresList(conds Conds, extra ...string) (resp []*Expires, err error) {

	sql, binds := BuildQuery(conds)

	sorts := ""
	if len(extra) >= 1 {
		sorts = extra[0]
	}
	if err = mus.Db.Table("expires").Where(sql, binds...).Order(sorts).Find(&resp).Error; err != nil {
		mus.Logger.Error("expires info error", zap.Error(err))
		return
	}
	return
}

// ListMap 查询map，map遍历的时候是无序的，所以指定sorts参数没有意义
func ExpiresListMap(conds Conds) (resp map[int]*Expires, err error) {

	sql, binds := BuildQuery(conds)

	mysqlSlice := make([]*Expires, 0)
	resp = make(map[int]*Expires, 0)
	if err = mus.Db.Table("expires").Where(sql, binds...).Find(&mysqlSlice).Error; err != nil {
		mus.Logger.Error("expires info error", zap.Error(err))
		return
	}
	for _, value := range mysqlSlice {
		resp[value.Id] = value
	}
	return
}

// ListPage 根据分页条件查询list
func ExpiresListPage(conds Conds, reqList *trans.ReqPage) (total int, respList Expiress) {
	respList = make(Expiress, 0)

	if reqList.PageSize == 0 {
		reqList.PageSize = 10
	}
	if reqList.Current == 0 {
		reqList.Current = 1
	}
	sql, binds := BuildQuery(conds)

	db := mus.Db.Table("expires").Where(sql, binds...)
	respList = make([]*Expires, 0)
	db.Count(&total)
	db.Order(reqList.Sort).Offset((reqList.Current - 1) * reqList.PageSize).Limit(reqList.PageSize).Find(&respList)
	return
}
