// @BeeOverwrite NO
// @BeeGenerateTime 20200821_113539
package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"oapms/pkg/mus"
	"oapms/pkg/trans"
)

type Role struct {
	Id         int         `gorm:"not null;primary_key;AUTO_INCREMENT"json:"id" form:"id"`    // id
	Name       string      `gorm:"not null;"json:"name" form:"name"`                          // 角色名称
	AppId      int         `gorm:"not null;"json:"appId" form:"appId"`                        // 应用名称
	Intro      string      `gorm:"not null;"json:"intro" form:"intro"`                        // 说明
	MenuIds    MenuIdsJson `gorm:"not null;type:json"json:"menuIds" form:"menuIds"`           // menu_ids
	MenuIdsEle string      `gorm:"not null;type:longtext"json:"menuIdsEle" form:"menuIdsEle"` // menu_ids_ele
	Pms        []Pms       `gorm:"-"json:"pms"`
}

type DataPms struct {
	RoleId    int
	DataPmsId int
}

func (t *Role) TableName() string {
	return "role"
}

type MenuIdsJson []int

func (c MenuIdsJson) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *MenuIdsJson) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

// GetRolesByNames
func RolesByRoleIds(ids []string) []*Role {
	roles := make([]*Role, 0)
	mus.Db.Where("id in (?)", ids).Find(&roles)

	for _, value := range roles {
		rolePmsList, err := RolePmsList(Conds{"role_id": value.Id})
		if err != nil {
			continue
		}
		ids := make([]int, 0)
		for _, rolePms := range rolePmsList {
			ids = append(ids, rolePms.PmsId)
		}
		pmsList := make([]Pms, 0)
		mus.Db.Where("id in (?)", ids).Find(&pmsList)
		value.Pms = pmsList
	}
	return roles
}

// AddRole insert a new Role into database and returns
// last inserted Id on success.
func RoleCreate(db *gorm.DB, data *Role) (err error) {

	if err = db.Create(data).Error; err != nil {
		mus.Logger.Error("create role error", zap.Error(err))
		return
	}
	return
}

func RoleUpdate(db *gorm.DB, paramId int, ups Ups) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}

	if err = db.Table("role").Where(sql, binds...).Updates(ups).Error; err != nil {
		mus.Logger.Error("role update error", zap.Error(err))
		return
	}
	return
}

// UpdateX Update的扩展方法，根据Cond更新一条或多条记录
func RoleUpdateX(db *gorm.DB, conds Conds, ups Ups) (err error) {
	sql, binds := BuildQuery(conds)
	if err = db.Table("role").Where(sql, binds...).Updates(ups).Error; err != nil {
		mus.Logger.Error("role update error", zap.Error(err))
		return
	}
	return
}

// Delete 根据主键删除一条记录。如果有delete_time则软删除，否则硬删除。
func RoleDelete(db *gorm.DB, paramId int) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}

	if err = db.Table("role").Where(sql, binds...).Delete(&Role{}).Error; err != nil {
		mus.Logger.Error("role delete error", zap.Error(err))
		return
	}

	return
}

// DeleteX Delete的扩展方法，根据Cond删除一条或多条记录。如果有delete_time则软删除，否则硬删除。
func RoleDeleteX(db *gorm.DB, conds Conds) (err error) {
	sql, binds := BuildQuery(conds)

	if err = db.Table("role").Where(sql, binds...).Delete(&Role{}).Error; err != nil {
		mus.Logger.Error("role delete error", zap.Error(err))
		return
	}

	return
}

// Info 根据PRI查询单条记录
func RoleInfo(db *gorm.DB, paramId int) (resp Role, err error) {

	var sql = "`id`= ?"

	var binds = []interface{}{paramId}

	if err = db.Table("role").Where(sql, binds...).First(&resp).Error; err != nil {
		mus.Logger.Error("role info error", zap.Error(err))
		return
	}
	return
}

// InfoX Info的扩展方法，根据Cond查询单条记录
func RoleInfoX(db *gorm.DB, conds Conds) (resp Role, err error) {

	sql, binds := BuildQuery(conds)

	if err = db.Table("role").Where(sql, binds...).First(&resp).Error; err != nil {
		mus.Logger.Error("role info error", zap.Error(err))
		return
	}
	return
}

// List 查询list，extra[0]为sorts
func RoleList(conds Conds, extra ...string) (resp []*Role, err error) {

	sql, binds := BuildQuery(conds)

	sorts := ""
	if len(extra) >= 1 {
		sorts = extra[0]
	}
	if err = mus.Db.Table("role").Where(sql, binds...).Order(sorts).Find(&resp).Error; err != nil {
		mus.Logger.Error("role info error", zap.Error(err))
		return
	}
	return
}

// ListMap 查询map，map遍历的时候是无序的，所以指定sorts参数没有意义
func RoleListMap(conds Conds) (resp map[int]*Role, err error) {

	sql, binds := BuildQuery(conds)

	mysqlSlice := make([]*Role, 0)
	resp = make(map[int]*Role, 0)
	if err = mus.Db.Table("role").Where(sql, binds...).Find(&mysqlSlice).Error; err != nil {
		mus.Logger.Error("role info error", zap.Error(err))
		return
	}
	for _, value := range mysqlSlice {
		resp[value.Id] = value
	}
	return
}

// ListPage 根据分页条件查询list
func RoleListPage(conds Conds, reqList *trans.ReqPage) (total int, respList []*Role) {

	if reqList.PageSize == 0 {
		reqList.PageSize = 10
	}
	if reqList.Current == 0 {
		reqList.Current = 1
	}
	sql, binds := BuildQuery(conds)

	db := mus.Db.Table("role").Where(sql, binds...)
	respList = make([]*Role, 0)
	db.Count(&total)
	db.Order(reqList.Sort).Offset((reqList.Current - 1) * reqList.PageSize).Limit(reqList.PageSize).Find(&respList)
	return
}

// batch delete
func RolePmsDeleteMulti(db *gorm.DB, roleId int, dataPermIds []int) error {
	return db.Where("role_id = ? and pms_id in(?)", roleId, dataPermIds).
		Delete(RolePms{}).
		Error
}

// assign data permission
func RolePmsInsertMulti(db *gorm.DB, datas []DataPms) error {
	sql := "insert into role_pms (role_id,pms_id) values "
	for key, v := range datas {
		if len(datas)-1 == key {
			sql += fmt.Sprintf("(%d,%d);", v.RoleId, v.DataPmsId)
		} else {
			sql += fmt.Sprintf("(%d,%d),", v.RoleId, v.DataPmsId)
		}
	}

	return db.Exec(sql).Error
}
