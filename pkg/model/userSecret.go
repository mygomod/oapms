// @BeeOverwrite NO
// @BeeGenerateTime 20200831_101837
package model

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"oapms/pkg/mus"
	"oapms/pkg/trans"

	"time"
)

type UserSecret struct {
	Id       int    `orm:"auto"json:"id" form:"id"`              // id
	Uid      int    `json:"uid" form:"uid"`                      // uid
	Secret   string `orm:"size(255)"json:"secret" form:"secret"` // 秘钥
	IsBind   int    `json:"isBind" form:"isBind"`                // 是否绑定
	Ctime    int64  `json:"ctime" form:"ctime"`                  // 创建时间
	Utime    int64  `json:"utime" form:"utime"`                  // 更新时间
	Nickname string `gorm:"-",json:"nickname"`
}

func (t *UserSecret) TableName() string {
	return "user_secret"
}

// AddUserSecret insert a new UserSecret into database and returns
// last inserted Id on success.
func UserSecretCreate(db *gorm.DB, data *UserSecret) (err error) {
	data.Ctime = time.Now().Unix()
	data.Utime = time.Now().Unix()

	if err = db.Create(data).Error; err != nil {
		mus.Logger.Error("create user_secret error", zap.Error(err))
		return
	}
	return
}

func UserSecretUpdate(db *gorm.DB, paramId int, ups Ups) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}

	if err = db.Table("user_secret").Where(sql, binds...).Updates(ups).Error; err != nil {
		mus.Logger.Error("user_secret update error", zap.Error(err))
		return
	}
	return
}

// UpdateX Update的扩展方法，根据Cond更新一条或多条记录
func UserSecretUpdateX(db *gorm.DB, conds Conds, ups Ups) (err error) {
	sql, binds := BuildQuery(conds)
	if err = db.Table("user_secret").Where(sql, binds...).Updates(ups).Error; err != nil {
		mus.Logger.Error("user_secret update error", zap.Error(err))
		return
	}
	return
}

// Delete 根据主键删除一条记录。如果有delete_time则软删除，否则硬删除。
func UserSecretDelete(db *gorm.DB, paramId int) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}

	if err = db.Table("user_secret").Where(sql, binds...).Delete(&UserSecret{}).Error; err != nil {
		mus.Logger.Error("user_secret delete error", zap.Error(err))
		return
	}

	return
}

// DeleteX Delete的扩展方法，根据Cond删除一条或多条记录。如果有delete_time则软删除，否则硬删除。
func UserSecretDeleteX(db *gorm.DB, conds Conds) (err error) {
	sql, binds := BuildQuery(conds)

	if err = db.Table("user_secret").Where(sql, binds...).Delete(&UserSecret{}).Error; err != nil {
		mus.Logger.Error("user_secret delete error", zap.Error(err))
		return
	}

	return
}

// Info 根据PRI查询单条记录
func UserSecretInfo(db *gorm.DB, paramId int) (resp UserSecret, err error) {

	var sql = "`id`= ?"

	var binds = []interface{}{paramId}

	if err = db.Table("user_secret").Where(sql, binds...).First(&resp).Error; err != nil {
		mus.Logger.Error("user_secret info error", zap.Error(err))
		return
	}
	return
}

// InfoX Info的扩展方法，根据Cond查询单条记录
func UserSecretInfoX(db *gorm.DB, conds Conds) (resp UserSecret, err error) {

	sql, binds := BuildQuery(conds)

	if err = db.Table("user_secret").Where(sql, binds...).First(&resp).Error; err != nil {
		mus.Logger.Error("user_secret info error", zap.Error(err))
		return
	}
	return
}

// List 查询list，extra[0]为sorts
func UserSecretList(conds Conds, extra ...string) (resp []*UserSecret, err error) {

	sql, binds := BuildQuery(conds)

	sorts := ""
	if len(extra) >= 1 {
		sorts = extra[0]
	}
	if err = mus.Db.Table("user_secret").Where(sql, binds...).Order(sorts).Find(&resp).Error; err != nil {
		mus.Logger.Error("user_secret info error", zap.Error(err))
		return
	}
	return
}

// ListMap 查询map，map遍历的时候是无序的，所以指定sorts参数没有意义
func UserSecretListMap(conds Conds) (resp map[int]*UserSecret, err error) {

	sql, binds := BuildQuery(conds)

	mysqlSlice := make([]*UserSecret, 0)
	resp = make(map[int]*UserSecret, 0)
	if err = mus.Db.Table("user_secret").Where(sql, binds...).Find(&mysqlSlice).Error; err != nil {
		mus.Logger.Error("user_secret info error", zap.Error(err))
		return
	}
	for _, value := range mysqlSlice {
		resp[value.Id] = value
	}
	return
}

// ListPage 根据分页条件查询list
func UserSecretListPage(conds Conds, reqList *trans.ReqPage) (total int, respList []*UserSecret) {

	if reqList.PageSize == 0 {
		reqList.PageSize = 10
	}
	if reqList.Current == 0 {
		reqList.Current = 1
	}
	sql, binds := BuildQuery(conds)

	db := mus.Db.Table("user_secret").Where(sql, binds...)
	respList = make([]*UserSecret, 0)
	db.Count(&total)
	db.Order(reqList.Sort).Offset((reqList.Current - 1) * reqList.PageSize).Limit(reqList.PageSize).Find(&respList)
	return
}
