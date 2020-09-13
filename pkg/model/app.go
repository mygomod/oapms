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

type App struct {
	Aid         int    `gorm:"auto"json:"aid" form:"aid"`                      // 应用id
	ClientId    string `gorm:"size(255)"json:"clientId" form:"clientId"`       // 客户端
	Name        string `gorm:"size(255)"json:"name" form:"name"`               // 名称
	Secret      string `gorm:"size(255)"json:"secret" form:"secret"`           // 秘钥
	RedirectUri string `gorm:"size(255)"json:"redirectUri" form:"redirectUri"` // 跳转地址
	Url         string `gorm:"size(255)"json:"url" form:"url"`                 // 访问地址
	Extra       string `gorm:"type(longtext)"json:"extra" form:"extra"`        // 额外信息
	CallNo      int    `json:"callNo" form:"callNo"`                           // 号码
	State       int    `json:"state" form:"state"`                             // 状态
	Ctime       int64  `json:"ctime" form:"ctime"`                             // 创建时间
	Utime       int64  `json:"utime" form:"utime"`                             // 更新时间
	Dtime       int64  `json:"dtime" form:"dtime"`                             // 删除时间

}

type Apps []*App

func (t *App) TableName() string {
	return "app"
}

// AddApp insert a new App into database and returns
// last inserted Id on success.
func AppCreate(db *gorm.DB, data *App) (err error) {
	data.Ctime = time.Now().Unix()
	data.Utime = time.Now().Unix()

	if err = db.Create(data).Error; err != nil {
		mus.Logger.Error("create app error", zap.Error(err))
		return
	}
	return
}

func AppUpdate(db *gorm.DB, paramId int, ups Ups) (err error) {
	var sql = "`aid`=?"
	var binds = []interface{}{paramId}

	if err = db.Table("app").Where(sql, binds...).Updates(ups).Error; err != nil {
		mus.Logger.Error("app update error", zap.Error(err))
		return
	}
	return
}

// UpdateX Update的扩展方法，根据Cond更新一条或多条记录
func AppUpdateX(db *gorm.DB, conds Conds, ups Ups) (err error) {
	sql, binds := BuildQuery(conds)
	if err = db.Table("app").Where(sql, binds...).Updates(ups).Error; err != nil {
		mus.Logger.Error("app update error", zap.Error(err))
		return
	}
	return
}

// Delete 根据主键删除一条记录。如果有delete_time则软删除，否则硬删除。
func AppDelete(db *gorm.DB, paramId int) (err error) {
	var sql = "`aid`=?"
	var binds = []interface{}{paramId}

	if err = db.Table("app").Where(sql, binds...).Updates(map[string]interface{}{
		"dtime": time.Now().Unix(),
	}).Error; err != nil {
		mus.Logger.Error("app delete error", zap.Error(err))
		return
	}

	return
}

// DeleteX Delete的扩展方法，根据Cond删除一条或多条记录。如果有delete_time则软删除，否则硬删除。
func AppDeleteX(db *gorm.DB, conds Conds) (err error) {
	sql, binds := BuildQuery(conds)

	if err = db.Table("app").Where(sql, binds...).Updates(map[string]interface{}{
		"dtime": time.Now().Unix(),
	}).Error; err != nil {
		mus.Logger.Error("app delete error", zap.Error(err))
		return
	}

	return
}

// Info 根据PRI查询单条记录
func AppInfo(db *gorm.DB, paramId int) (resp App, err error) {

	var sql = "`aid`= ? and dtime = 0"

	var binds = []interface{}{paramId}

	if err = db.Table("app").Where(sql, binds...).First(&resp).Error; err != nil {
		mus.Logger.Error("app info error", zap.Error(err))
		return
	}
	return
}

// InfoX Info的扩展方法，根据Cond查询单条记录
func AppInfoX(db *gorm.DB, conds Conds) (resp App, err error) {

	conds["dtime"] = 0

	sql, binds := BuildQuery(conds)

	if err = db.Table("app").Where(sql, binds...).First(&resp).Error; err != nil {
		mus.Logger.Error("app info error", zap.Error(err))
		return
	}
	return
}

// List 查询list，extra[0]为sorts
func AppList(conds Conds, extra ...string) (resp []*App, err error) {

	conds["dtime"] = 0

	sql, binds := BuildQuery(conds)

	sorts := ""
	if len(extra) >= 1 {
		sorts = extra[0]
	}
	if err = mus.Db.Table("app").Where(sql, binds...).Order(sorts).Find(&resp).Error; err != nil {
		mus.Logger.Error("app info error", zap.Error(err))
		return
	}
	return
}

// ListMap 查询map，map遍历的时候是无序的，所以指定sorts参数没有意义
func AppListMap(conds Conds) (resp map[int]*App, err error) {

	conds["dtime"] = 0

	sql, binds := BuildQuery(conds)

	mysqlSlice := make([]*App, 0)
	resp = make(map[int]*App, 0)
	if err = mus.Db.Table("app").Where(sql, binds...).Find(&mysqlSlice).Error; err != nil {
		mus.Logger.Error("app info error", zap.Error(err))
		return
	}
	for _, value := range mysqlSlice {
		resp[value.Aid] = value
	}
	return
}

// ListPage 根据分页条件查询list
func AppListPage(conds Conds, reqList *trans.ReqPage) (total int, respList Apps) {
	respList = make(Apps, 0)

	conds["dtime"] = 0

	if reqList.PageSize == 0 {
		reqList.PageSize = 10
	}
	if reqList.Current == 0 {
		reqList.Current = 1
	}
	sql, binds := BuildQuery(conds)

	db := mus.Db.Table("app").Where(sql, binds...)
	respList = make([]*App, 0)
	db.Count(&total)
	db.Order(reqList.Sort).Offset((reqList.Current - 1) * reqList.PageSize).Limit(reqList.PageSize).Find(&respList)
	return
}
