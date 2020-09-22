// @BeeOverwrite NO
// @BeeGenerateTime 20200821_113539
package model

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"oapms/pkg/mus"
	"oapms/pkg/trans"

	"time"
)

type Pms struct {
	Id       int    `gorm:"not null;primary_key;AUTO_INCREMENT"json:"id" form:"id"` // ID
	AppId    int    `gorm:"not null;"json:"appId" form:"appId"`                     // 应用id
	Pid      int    `gorm:"not null;"json:"pid" form:"pid"`                         // 菜单id
	Name     string `gorm:"not null;"json:"name" form:"name"`                       // 名称
	PmsCode  string `gorm:"not null;"json:"pmsCode" form:"pmsCode"`                 // 权限标识
	PmsRule  string `gorm:"not null;"json:"pmsRule" form:"pmsRule"`                 // 数据规则
	PmsType  int    `gorm:"not null;"json:"pmsType" form:"pmsType"`                 // 1=分类 2=数据权限
	OrderNum int    `gorm:"not null;"json:"orderNum" form:"orderNum"`               // 排序
	Intro    string `gorm:"not null;"json:"intro" form:"intro"`                     // 说明
	Ctime    int64  `gorm:"not null;"json:"ctime" form:"ctime"`                     // 创建时间
	Utime    int64  `gorm:"not null;"json:"utime" form:"utime"`                     // 更新时间

}

func (t *Pms) TableName() string {
	return "pms"
}

// AddPms insert a new Pms into database and returns
// last inserted Id on success.
func PmsCreate(db *gorm.DB, data *Pms) (err error) {
	data.Ctime = time.Now().Unix()
	data.Utime = time.Now().Unix()

	if err = db.Create(data).Error; err != nil {
		mus.Logger.Error("create pms error", zap.Error(err))
		return
	}
	return
}

func PmsUpdate(db *gorm.DB, paramId int, ups Ups) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}

	if err = db.Table("pms").Where(sql, binds...).Updates(ups).Error; err != nil {
		mus.Logger.Error("pms update error", zap.Error(err))
		return
	}
	return
}

// UpdateX Update的扩展方法，根据Cond更新一条或多条记录
func PmsUpdateX(db *gorm.DB, conds Conds, ups Ups) (err error) {
	sql, binds := BuildQuery(conds)
	if err = db.Table("pms").Where(sql, binds...).Updates(ups).Error; err != nil {
		mus.Logger.Error("pms update error", zap.Error(err))
		return
	}
	return
}

// Delete 根据主键删除一条记录。如果有delete_time则软删除，否则硬删除。
func PmsDelete(db *gorm.DB, paramId int) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}

	if err = db.Table("pms").Where(sql, binds...).Delete(&Pms{}).Error; err != nil {
		mus.Logger.Error("pms delete error", zap.Error(err))
		return
	}

	return
}

// DeleteX Delete的扩展方法，根据Cond删除一条或多条记录。如果有delete_time则软删除，否则硬删除。
func PmsDeleteX(db *gorm.DB, conds Conds) (err error) {
	sql, binds := BuildQuery(conds)

	if err = db.Table("pms").Where(sql, binds...).Delete(&Pms{}).Error; err != nil {
		mus.Logger.Error("pms delete error", zap.Error(err))
		return
	}

	return
}

// Info 根据PRI查询单条记录
func PmsInfo(db *gorm.DB, paramId int) (resp Pms, err error) {

	var sql = "`id`= ?"

	var binds = []interface{}{paramId}

	if err = db.Table("pms").Where(sql, binds...).First(&resp).Error; err != nil {
		mus.Logger.Error("pms info error", zap.Error(err))
		return
	}
	return
}

// InfoX Info的扩展方法，根据Cond查询单条记录
func PmsInfoX(db *gorm.DB, conds Conds) (resp Pms, err error) {

	sql, binds := BuildQuery(conds)

	if err = db.Table("pms").Where(sql, binds...).First(&resp).Error; err != nil {
		mus.Logger.Error("pms info error", zap.Error(err))
		return
	}
	return
}

// List 查询list，extra[0]为sorts
func PmsList(conds Conds, extra ...string) (resp []*Pms, err error) {

	sql, binds := BuildQuery(conds)

	sorts := ""
	if len(extra) >= 1 {
		sorts = extra[0]
	}
	if err = mus.Db.Table("pms").Where(sql, binds...).Order(sorts).Find(&resp).Error; err != nil {
		mus.Logger.Error("pms info error", zap.Error(err))
		return
	}
	return
}

// ListMap 查询map，map遍历的时候是无序的，所以指定sorts参数没有意义
func PmsListMap(conds Conds) (resp map[int]*Pms, err error) {

	sql, binds := BuildQuery(conds)

	mysqlSlice := make([]*Pms, 0)
	resp = make(map[int]*Pms, 0)
	if err = mus.Db.Table("pms").Where(sql, binds...).Find(&mysqlSlice).Error; err != nil {
		mus.Logger.Error("pms info error", zap.Error(err))
		return
	}
	for _, value := range mysqlSlice {
		resp[value.Id] = value
	}
	return
}

// ListPage 根据分页条件查询list
func PmsListPage(conds Conds, reqList *trans.ReqPage) (total int, respList []*Pms) {

	if reqList.PageSize == 0 {
		reqList.PageSize = 10
	}
	if reqList.Current == 0 {
		reqList.Current = 1
	}
	sql, binds := BuildQuery(conds)

	db := mus.Db.Table("pms").Where(sql, binds...)
	respList = make([]*Pms, 0)
	db.Count(&total)
	db.Order(reqList.Sort).Offset((reqList.Current - 1) * reqList.PageSize).Limit(reqList.PageSize).Find(&respList)
	return
}

type GetByRoleIdData struct {
	Id      int    `json:"id"`
	RoleId  int    `json:"roleId"`
	Name    string `json:"name"`
	PmsCode string `json:"pmsCode"`
}

// get by role_id
func PmsByRoleId(roleId int) ([]GetByRoleIdData, int64) {
	var rdps []GetByRoleIdData
	var total int64
	db := mus.Db
	fields := "role_pms.role_id,pms.name,pms.pms_code,pms.id"
	query := db.Table("role_pms").
		Select(fields).
		Joins("left join pms on role_pms.pms_id=pms.id").
		Where("role_pms.role_id=?", roleId)

	rows, _ := query.Rows()
	query.Count(&total)
	defer rows.Close()
	for rows.Next() {
		var data GetByRoleIdData
		_ = db.ScanRows(rows, &data)
		rdps = append(rdps, data)
	}

	return rdps, total
}
