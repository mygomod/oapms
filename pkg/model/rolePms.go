// @BeeOverwrite YES
// @BeeGenerateTime 20200831_174645
package model

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"oapms/pkg/mus"
	"oapms/pkg/trans"
)

type RolePms struct {
	Id     int `gorm:"auto"json:"id" form:"id"` // id
	RoleId int `json:"roleId" form:"roleId"`    // 角色id
	PmsId  int `json:"pmsId" form:"pmsId"`      // 数据权限id

}

type RolePmss []*RolePms

func (t *RolePms) TableName() string {
	return "role_pms"
}

// AddRolePms insert a new RolePms into database and returns
// last inserted Id on success.
func RolePmsCreate(db *gorm.DB, data *RolePms) (err error) {

	if err = db.Create(data).Error; err != nil {
		mus.Logger.Error("create role_pms error", zap.Error(err))
		return
	}
	return
}

func RolePmsUpdate(db *gorm.DB, paramId int, ups Ups) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}

	if err = db.Table("role_pms").Where(sql, binds...).Updates(ups).Error; err != nil {
		mus.Logger.Error("role_pms update error", zap.Error(err))
		return
	}
	return
}

// UpdateX Update的扩展方法，根据Cond更新一条或多条记录
func RolePmsUpdateX(db *gorm.DB, conds Conds, ups Ups) (err error) {
	sql, binds := BuildQuery(conds)
	if err = db.Table("role_pms").Where(sql, binds...).Updates(ups).Error; err != nil {
		mus.Logger.Error("role_pms update error", zap.Error(err))
		return
	}
	return
}

// Delete 根据主键删除一条记录。如果有delete_time则软删除，否则硬删除。
func RolePmsDelete(db *gorm.DB, paramId int) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}

	if err = db.Table("role_pms").Where(sql, binds...).Delete(&RolePms{}).Error; err != nil {
		mus.Logger.Error("role_pms delete error", zap.Error(err))
		return
	}

	return
}

// DeleteX Delete的扩展方法，根据Cond删除一条或多条记录。如果有delete_time则软删除，否则硬删除。
func RolePmsDeleteX(db *gorm.DB, conds Conds) (err error) {
	sql, binds := BuildQuery(conds)

	if err = db.Table("role_pms").Where(sql, binds...).Delete(&RolePms{}).Error; err != nil {
		mus.Logger.Error("role_pms delete error", zap.Error(err))
		return
	}

	return
}

// Info 根据PRI查询单条记录
func RolePmsInfo(db *gorm.DB, paramId int) (resp RolePms, err error) {

	var sql = "`id`= ?"

	var binds = []interface{}{paramId}

	if err = db.Table("role_pms").Where(sql, binds...).First(&resp).Error; err != nil {
		mus.Logger.Error("role_pms info error", zap.Error(err))
		return
	}
	return
}

// InfoX Info的扩展方法，根据Cond查询单条记录
func RolePmsInfoX(db *gorm.DB, conds Conds) (resp RolePms, err error) {

	sql, binds := BuildQuery(conds)

	if err = db.Table("role_pms").Where(sql, binds...).First(&resp).Error; err != nil {
		mus.Logger.Error("role_pms info error", zap.Error(err))
		return
	}
	return
}

// List 查询list，extra[0]为sorts
func RolePmsList(conds Conds, extra ...string) (resp []*RolePms, err error) {

	sql, binds := BuildQuery(conds)

	sorts := ""
	if len(extra) >= 1 {
		sorts = extra[0]
	}
	if err = mus.Db.Table("role_pms").Where(sql, binds...).Order(sorts).Find(&resp).Error; err != nil {
		mus.Logger.Error("role_pms info error", zap.Error(err))
		return
	}
	return
}

// ListMap 查询map，map遍历的时候是无序的，所以指定sorts参数没有意义
func RolePmsListMap(conds Conds) (resp map[int]*RolePms, err error) {

	sql, binds := BuildQuery(conds)

	mysqlSlice := make([]*RolePms, 0)
	resp = make(map[int]*RolePms, 0)
	if err = mus.Db.Table("role_pms").Where(sql, binds...).Find(&mysqlSlice).Error; err != nil {
		mus.Logger.Error("role_pms info error", zap.Error(err))
		return
	}
	for _, value := range mysqlSlice {
		resp[value.Id] = value
	}
	return
}

// ListPage 根据分页条件查询list
func RolePmsListPage(conds Conds, reqList *trans.ReqPage) (total int, respList RolePmss) {
	respList = make(RolePmss, 0)

	if reqList.PageSize == 0 {
		reqList.PageSize = 10
	}
	if reqList.Current == 0 {
		reqList.Current = 1
	}
	sql, binds := BuildQuery(conds)

	db := mus.Db.Table("role_pms").Where(sql, binds...)
	respList = make([]*RolePms, 0)
	db.Count(&total)
	db.Order(reqList.Sort).Offset((reqList.Current - 1) * reqList.PageSize).Limit(reqList.PageSize).Find(&respList)
	return
}
