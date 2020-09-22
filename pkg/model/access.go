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

type Access struct {
	Id           int    `gorm:"not null;primary_key;AUTO_INCREMENT"json:"id" form:"id"` // ID
	Client       string `gorm:"not null"json:"client" form:"client"`                    // client
	Authorize    string `gorm:"not null"json:"authorize" form:"authorize"`              // authorize
	Previous     string `gorm:"not null"json:"previous" form:"previous"`                // previous
	AccessToken  string `gorm:"not null"json:"accessToken" form:"accessToken"`          // access_token
	RefreshToken string `gorm:"not null"json:"refreshToken" form:"refreshToken"`        // refresh_token
	ExpiresIn    int    `gorm:"not null"json:"expiresIn" form:"expiresIn"`              // expires_in
	Scope        string `gorm:"not null"json:"scope" form:"scope"`                      // scope
	RedirectUri  string `gorm:"not null"json:"redirectUri" form:"redirectUri"`          // redirect_uri
	Extra        string `gorm:"not null;type:longtext"json:"extra" form:"extra"`        // extra
	Ctime        int64  `gorm:"not null"json:"ctime" form:"ctime"`                      // 创建时间

}

type Accesss []*Access

func (t *Access) TableName() string {
	return "access"
}

// AddAccess insert a new Access into database and returns
// last inserted Id on success.
func AccessCreate(db *gorm.DB, data *Access) (err error) {
	data.Ctime = time.Now().Unix()

	if err = db.Create(data).Error; err != nil {
		mus.Logger.Error("create access error", zap.Error(err))
		return
	}
	return
}

func AccessUpdate(db *gorm.DB, paramId int, ups Ups) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}

	if err = db.Table("access").Where(sql, binds...).Updates(ups).Error; err != nil {
		mus.Logger.Error("access update error", zap.Error(err))
		return
	}
	return
}

// UpdateX Update的扩展方法，根据Cond更新一条或多条记录
func AccessUpdateX(db *gorm.DB, conds Conds, ups Ups) (err error) {
	sql, binds := BuildQuery(conds)
	if err = db.Table("access").Where(sql, binds...).Updates(ups).Error; err != nil {
		mus.Logger.Error("access update error", zap.Error(err))
		return
	}
	return
}

// Delete 根据主键删除一条记录。如果有delete_time则软删除，否则硬删除。
func AccessDelete(db *gorm.DB, paramId int) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}

	if err = db.Table("access").Where(sql, binds...).Delete(&Access{}).Error; err != nil {
		mus.Logger.Error("access delete error", zap.Error(err))
		return
	}

	return
}

// DeleteX Delete的扩展方法，根据Cond删除一条或多条记录。如果有delete_time则软删除，否则硬删除。
func AccessDeleteX(db *gorm.DB, conds Conds) (err error) {
	sql, binds := BuildQuery(conds)

	if err = db.Table("access").Where(sql, binds...).Delete(&Access{}).Error; err != nil {
		mus.Logger.Error("access delete error", zap.Error(err))
		return
	}

	return
}

// Info 根据PRI查询单条记录
func AccessInfo(db *gorm.DB, paramId int) (resp Access, err error) {

	var sql = "`id`= ?"

	var binds = []interface{}{paramId}

	if err = db.Table("access").Where(sql, binds...).First(&resp).Error; err != nil {
		mus.Logger.Error("access info error", zap.Error(err))
		return
	}
	return
}

// InfoX Info的扩展方法，根据Cond查询单条记录
func AccessInfoX(db *gorm.DB, conds Conds) (resp Access, err error) {

	sql, binds := BuildQuery(conds)

	if err = db.Table("access").Where(sql, binds...).First(&resp).Error; err != nil {
		mus.Logger.Error("access info error", zap.Error(err))
		return
	}
	return
}

// List 查询list，extra[0]为sorts
func AccessList(conds Conds, extra ...string) (resp []*Access, err error) {

	sql, binds := BuildQuery(conds)

	sorts := ""
	if len(extra) >= 1 {
		sorts = extra[0]
	}
	if err = mus.Db.Table("access").Where(sql, binds...).Order(sorts).Find(&resp).Error; err != nil {
		mus.Logger.Error("access info error", zap.Error(err))
		return
	}
	return
}

// ListMap 查询map，map遍历的时候是无序的，所以指定sorts参数没有意义
func AccessListMap(conds Conds) (resp map[int]*Access, err error) {

	sql, binds := BuildQuery(conds)

	mysqlSlice := make([]*Access, 0)
	resp = make(map[int]*Access, 0)
	if err = mus.Db.Table("access").Where(sql, binds...).Find(&mysqlSlice).Error; err != nil {
		mus.Logger.Error("access info error", zap.Error(err))
		return
	}
	for _, value := range mysqlSlice {
		resp[value.Id] = value
	}
	return
}

// ListPage 根据分页条件查询list
func AccessListPage(conds Conds, reqList *trans.ReqPage) (total int, respList Accesss) {
	respList = make(Accesss, 0)

	if reqList.PageSize == 0 {
		reqList.PageSize = 10
	}
	if reqList.Current == 0 {
		reqList.Current = 1
	}
	sql, binds := BuildQuery(conds)

	db := mus.Db.Table("access").Where(sql, binds...)
	respList = make([]*Access, 0)
	db.Count(&total)
	db.Order(reqList.Sort).Offset((reqList.Current - 1) * reqList.PageSize).Limit(reqList.PageSize).Find(&respList)
	return
}
