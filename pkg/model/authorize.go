// @BeeOverwrite YES
// @BeeGenerateTime 20200831_174645
package model

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"oapms/pkg/mus"
	"oapms/pkg/trans"

	"time"
)

type Authorize struct {
	Id          int    `gorm:"not null;primary_key;AUTO_INCREMENT"json:"id" form:"id"` // ID
	Client      string `gorm:"not null"json:"client" form:"client"`                    // 客户端
	Code        string `gorm:"not null"json:"code" form:"code"`                        // 状态码
	ExpiresIn   int32  `gorm:"not null"json:"expiresIn" form:"expiresIn"`              // 过期时间
	Scope       string `gorm:"not null"json:"scope" form:"scope"`                      // 范围
	RedirectUri string `gorm:"not null"json:"redirectUri" form:"redirectUri"`          // 跳转地址
	State       string `gorm:"not null"json:"state" form:"state"`                      // 状态
	Extra       string `gorm:"not null;type:longtext"json:"extra" form:"extra"`        // 额外信息
	Ctime       int64  `gorm:"not null"json:"ctime" form:"ctime"`                      // 创建时间

}

type Authorizes []*Authorize

func (t *Authorize) TableName() string {
	return "authorize"
}

// AddAuthorize insert a new Authorize into database and returns
// last inserted Id on success.
func AuthorizeCreate(db *gorm.DB, data *Authorize) (err error) {
	data.Ctime = time.Now().Unix()

	if err = db.Create(data).Error; err != nil {
		mus.Logger.Error("create authorize error", zap.Error(err))
		return
	}
	return
}

func AuthorizeUpdate(db *gorm.DB, paramId int, ups Ups) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}

	if err = db.Table("authorize").Where(sql, binds...).Updates(ups).Error; err != nil {
		mus.Logger.Error("authorize update error", zap.Error(err))
		return
	}
	return
}

// UpdateX Update的扩展方法，根据Cond更新一条或多条记录
func AuthorizeUpdateX(db *gorm.DB, conds Conds, ups Ups) (err error) {
	sql, binds := BuildQuery(conds)
	if err = db.Table("authorize").Where(sql, binds...).Updates(ups).Error; err != nil {
		mus.Logger.Error("authorize update error", zap.Error(err))
		return
	}
	return
}

// Delete 根据主键删除一条记录。如果有delete_time则软删除，否则硬删除。
func AuthorizeDelete(db *gorm.DB, paramId int) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}

	if err = db.Table("authorize").Where(sql, binds...).Delete(&Authorize{}).Error; err != nil {
		mus.Logger.Error("authorize delete error", zap.Error(err))
		return
	}

	return
}

// DeleteX Delete的扩展方法，根据Cond删除一条或多条记录。如果有delete_time则软删除，否则硬删除。
func AuthorizeDeleteX(db *gorm.DB, conds Conds) (err error) {
	sql, binds := BuildQuery(conds)

	if err = db.Table("authorize").Where(sql, binds...).Delete(&Authorize{}).Error; err != nil {
		mus.Logger.Error("authorize delete error", zap.Error(err))
		return
	}

	return
}

// Info 根据PRI查询单条记录
func AuthorizeInfo(db *gorm.DB, paramId int) (resp Authorize, err error) {

	var sql = "`id`= ?"

	var binds = []interface{}{paramId}

	if err = db.Table("authorize").Where(sql, binds...).First(&resp).Error; err != nil {
		mus.Logger.Error("authorize info error", zap.Error(err))
		return
	}
	return
}

// InfoX Info的扩展方法，根据Cond查询单条记录
func AuthorizeInfoX(db *gorm.DB, conds Conds) (resp Authorize, err error) {

	sql, binds := BuildQuery(conds)

	if err = db.Table("authorize").Where(sql, binds...).First(&resp).Error; err != nil {
		mus.Logger.Error("authorize info error", zap.Error(err))
		return
	}
	return
}

// List 查询list，extra[0]为sorts
func AuthorizeList(conds Conds, extra ...string) (resp []*Authorize, err error) {

	sql, binds := BuildQuery(conds)

	sorts := ""
	if len(extra) >= 1 {
		sorts = extra[0]
	}
	if err = mus.Db.Table("authorize").Where(sql, binds...).Order(sorts).Find(&resp).Error; err != nil {
		mus.Logger.Error("authorize info error", zap.Error(err))
		return
	}
	return
}

// ListMap 查询map，map遍历的时候是无序的，所以指定sorts参数没有意义
func AuthorizeListMap(conds Conds) (resp map[int]*Authorize, err error) {

	sql, binds := BuildQuery(conds)

	mysqlSlice := make([]*Authorize, 0)
	resp = make(map[int]*Authorize, 0)
	if err = mus.Db.Table("authorize").Where(sql, binds...).Find(&mysqlSlice).Error; err != nil {
		mus.Logger.Error("authorize info error", zap.Error(err))
		return
	}
	for _, value := range mysqlSlice {
		resp[value.Id] = value
	}
	return
}

// ListPage 根据分页条件查询list
func AuthorizeListPage(conds Conds, reqList *trans.ReqPage) (total int, respList Authorizes) {
	respList = make(Authorizes, 0)

	if reqList.PageSize == 0 {
		reqList.PageSize = 10
	}
	if reqList.Current == 0 {
		reqList.Current = 1
	}
	sql, binds := BuildQuery(conds)

	db := mus.Db.Table("authorize").Where(sql, binds...)
	respList = make([]*Authorize, 0)
	db.Count(&total)
	db.Order(reqList.Sort).Offset((reqList.Current - 1) * reqList.PageSize).Limit(reqList.PageSize).Find(&respList)
	return
}
