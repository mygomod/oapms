// @BeeOverwrite YES
// @BeeGenerateTime 20200902_214035
package model

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"oapms/pkg/mus"
	"oapms/pkg/trans"

	"time"
)

type MenuPms struct {
	Id      int    `gorm:"not null;primary_key;AUTO_INCREMENT"json:"id" form:"id"` // id
	PmsCode string `gorm:"not null"json:"pmsCode" form:"pmsCode"`                  // 标识
	Key     string `gorm:"not null"json:"key" form:"key"`                          // api或者button
	AppId   int    `gorm:"not null"json:"appId" form:"appId"`                      // 应用id
	Ctime   int64  `gorm:"not null"json:"ctime" form:"ctime"`                      // 创建时间
	Utime   int64  `gorm:"not null"json:"utime" form:"utime"`                      // 更新时间

}

type MenuPmss []*MenuPms

func (t *MenuPms) TableName() string {
	return "menu_pms"
}

// AddMenuPms insert a new MenuPms into database and returns
// last inserted Id on success.
func MenuPmsCreate(db *gorm.DB, data *MenuPms) (err error) {
	data.Ctime = time.Now().Unix()
	data.Utime = time.Now().Unix()

	if err = db.Create(data).Error; err != nil {
		mus.Logger.Error("create menu_pms error", zap.Error(err))
		return
	}
	return
}

func MenuPmsUpdate(db *gorm.DB, paramId int, ups Ups) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}

	if err = db.Table("menu_pms").Where(sql, binds...).Updates(ups).Error; err != nil {
		mus.Logger.Error("menu_pms update error", zap.Error(err))
		return
	}
	return
}

// UpdateX Update的扩展方法，根据Cond更新一条或多条记录
func MenuPmsUpdateX(db *gorm.DB, conds Conds, ups Ups) (err error) {
	sql, binds := BuildQuery(conds)
	if err = db.Table("menu_pms").Where(sql, binds...).Updates(ups).Error; err != nil {
		mus.Logger.Error("menu_pms update error", zap.Error(err))
		return
	}
	return
}

// Delete 根据主键删除一条记录。如果有delete_time则软删除，否则硬删除。
func MenuPmsDelete(db *gorm.DB, paramId int) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}

	if err = db.Table("menu_pms").Where(sql, binds...).Delete(&MenuPms{}).Error; err != nil {
		mus.Logger.Error("menu_pms delete error", zap.Error(err))
		return
	}

	return
}

// DeleteX Delete的扩展方法，根据Cond删除一条或多条记录。如果有delete_time则软删除，否则硬删除。
func MenuPmsDeleteX(db *gorm.DB, conds Conds) (err error) {
	sql, binds := BuildQuery(conds)

	if err = db.Table("menu_pms").Where(sql, binds...).Delete(&MenuPms{}).Error; err != nil {
		mus.Logger.Error("menu_pms delete error", zap.Error(err))
		return
	}

	return
}

// Info 根据PRI查询单条记录
func MenuPmsInfo(db *gorm.DB, paramId int) (resp MenuPms, err error) {

	var sql = "`id`= ?"

	var binds = []interface{}{paramId}

	if err = db.Table("menu_pms").Where(sql, binds...).First(&resp).Error; err != nil {
		mus.Logger.Error("menu_pms info error", zap.Error(err))
		return
	}
	return
}

// InfoX Info的扩展方法，根据Cond查询单条记录
func MenuPmsInfoX(db *gorm.DB, conds Conds) (resp MenuPms, err error) {

	sql, binds := BuildQuery(conds)

	if err = db.Table("menu_pms").Where(sql, binds...).First(&resp).Error; err != nil {
		mus.Logger.Error("menu_pms info error", zap.Error(err))
		return
	}
	return
}

// List 查询list，extra[0]为sorts
func MenuPmsList(conds Conds, extra ...string) (resp []*MenuPms, err error) {

	sql, binds := BuildQuery(conds)

	sorts := ""
	if len(extra) >= 1 {
		sorts = extra[0]
	}
	if err = mus.Db.Table("menu_pms").Where(sql, binds...).Order(sorts).Find(&resp).Error; err != nil {
		mus.Logger.Error("menu_pms info error", zap.Error(err))
		return
	}
	return
}

// ListMap 查询map，map遍历的时候是无序的，所以指定sorts参数没有意义
func MenuPmsListMap(conds Conds) (resp map[int]*MenuPms, err error) {

	sql, binds := BuildQuery(conds)

	mysqlSlice := make([]*MenuPms, 0)
	resp = make(map[int]*MenuPms, 0)
	if err = mus.Db.Table("menu_pms").Where(sql, binds...).Find(&mysqlSlice).Error; err != nil {
		mus.Logger.Error("menu_pms info error", zap.Error(err))
		return
	}
	for _, value := range mysqlSlice {
		resp[value.Id] = value
	}
	return
}

// ListPage 根据分页条件查询list
func MenuPmsListPage(conds Conds, reqList *trans.ReqPage) (total int, respList MenuPmss) {
	respList = make(MenuPmss, 0)

	if reqList.PageSize == 0 {
		reqList.PageSize = 10
	}
	if reqList.Current == 0 {
		reqList.Current = 1
	}
	sql, binds := BuildQuery(conds)

	db := mus.Db.Table("menu_pms").Where(sql, binds...)
	respList = make([]*MenuPms, 0)
	db.Count(&total)
	db.Order(reqList.Sort).Offset((reqList.Current - 1) * reqList.PageSize).Limit(reqList.PageSize).Find(&respList)
	return
}
