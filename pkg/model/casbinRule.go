// @BeeOverwrite YES
// @BeeGenerateTime 20200831_174645
package model

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"oapms/pkg/mus"
	"oapms/pkg/trans"
)

type CasbinRule struct {
	Id int    `gorm:"not null;primary_key;AUTO_INCREMENT"json:"id" form:"id"` // ID
	P  string `gorm:"not null"json:"p" form:"p"`                              // 策略、用户组
	V0 string `gorm:"not null"json:"v0" form:"v0"`                            // v0
	V1 string `gorm:"not null"json:"v1" form:"v1"`                            // v1
	V2 string `gorm:"not null"json:"v2" form:"v2"`                            // v2
	V3 string `gorm:"not null"json:"v3" form:"v3"`                            // v3
	V4 string `gorm:"not null"json:"v4" form:"v4"`                            // v4
	V5 string `gorm:"not null"json:"v5" form:"v5"`                            // v5

}

type CasbinRules []*CasbinRule

func (t *CasbinRule) TableName() string {
	return "casbin_rule"
}

// AddCasbinRule insert a new CasbinRule into database and returns
// last inserted Id on success.
func CasbinRuleCreate(db *gorm.DB, data *CasbinRule) (err error) {

	if err = db.Create(data).Error; err != nil {
		mus.Logger.Error("create casbin_rule error", zap.Error(err))
		return
	}
	return
}

func CasbinRuleUpdate(db *gorm.DB, paramId int, ups Ups) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}

	if err = db.Table("casbin_rule").Where(sql, binds...).Updates(ups).Error; err != nil {
		mus.Logger.Error("casbin_rule update error", zap.Error(err))
		return
	}
	return
}

// UpdateX Update的扩展方法，根据Cond更新一条或多条记录
func CasbinRuleUpdateX(db *gorm.DB, conds Conds, ups Ups) (err error) {
	sql, binds := BuildQuery(conds)
	if err = db.Table("casbin_rule").Where(sql, binds...).Updates(ups).Error; err != nil {
		mus.Logger.Error("casbin_rule update error", zap.Error(err))
		return
	}
	return
}

// Delete 根据主键删除一条记录。如果有delete_time则软删除，否则硬删除。
func CasbinRuleDelete(db *gorm.DB, paramId int) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}

	if err = db.Table("casbin_rule").Where(sql, binds...).Delete(&CasbinRule{}).Error; err != nil {
		mus.Logger.Error("casbin_rule delete error", zap.Error(err))
		return
	}

	return
}

// DeleteX Delete的扩展方法，根据Cond删除一条或多条记录。如果有delete_time则软删除，否则硬删除。
func CasbinRuleDeleteX(db *gorm.DB, conds Conds) (err error) {
	sql, binds := BuildQuery(conds)

	if err = db.Table("casbin_rule").Where(sql, binds...).Delete(&CasbinRule{}).Error; err != nil {
		mus.Logger.Error("casbin_rule delete error", zap.Error(err))
		return
	}

	return
}

// Info 根据PRI查询单条记录
func CasbinRuleInfo(db *gorm.DB, paramId int) (resp CasbinRule, err error) {

	var sql = "`id`= ?"

	var binds = []interface{}{paramId}

	if err = db.Table("casbin_rule").Where(sql, binds...).First(&resp).Error; err != nil {
		mus.Logger.Error("casbin_rule info error", zap.Error(err))
		return
	}
	return
}

// InfoX Info的扩展方法，根据Cond查询单条记录
func CasbinRuleInfoX(db *gorm.DB, conds Conds) (resp CasbinRule, err error) {

	sql, binds := BuildQuery(conds)

	if err = db.Table("casbin_rule").Where(sql, binds...).First(&resp).Error; err != nil {
		mus.Logger.Error("casbin_rule info error", zap.Error(err))
		return
	}
	return
}

// List 查询list，extra[0]为sorts
func CasbinRuleList(conds Conds, extra ...string) (resp []*CasbinRule, err error) {

	sql, binds := BuildQuery(conds)

	sorts := ""
	if len(extra) >= 1 {
		sorts = extra[0]
	}
	if err = mus.Db.Table("casbin_rule").Where(sql, binds...).Order(sorts).Find(&resp).Error; err != nil {
		mus.Logger.Error("casbin_rule info error", zap.Error(err))
		return
	}
	return
}

// ListMap 查询map，map遍历的时候是无序的，所以指定sorts参数没有意义
func CasbinRuleListMap(conds Conds) (resp map[int]*CasbinRule, err error) {

	sql, binds := BuildQuery(conds)

	mysqlSlice := make([]*CasbinRule, 0)
	resp = make(map[int]*CasbinRule, 0)
	if err = mus.Db.Table("casbin_rule").Where(sql, binds...).Find(&mysqlSlice).Error; err != nil {
		mus.Logger.Error("casbin_rule info error", zap.Error(err))
		return
	}
	for _, value := range mysqlSlice {
		resp[value.Id] = value
	}
	return
}

// ListPage 根据分页条件查询list
func CasbinRuleListPage(conds Conds, reqList *trans.ReqPage) (total int, respList CasbinRules) {
	respList = make(CasbinRules, 0)

	if reqList.PageSize == 0 {
		reqList.PageSize = 10
	}
	if reqList.Current == 0 {
		reqList.Current = 1
	}
	sql, binds := BuildQuery(conds)

	db := mus.Db.Table("casbin_rule").Where(sql, binds...)
	respList = make([]*CasbinRule, 0)
	db.Count(&total)
	db.Order(reqList.Sort).Offset((reqList.Current - 1) * reqList.PageSize).Limit(reqList.PageSize).Find(&respList)
	return
}
