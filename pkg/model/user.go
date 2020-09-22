// @BeeOverwrite NO
// @BeeGenerateTime 20200821_113539
package model

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"oapms/pkg/mus"
	"oapms/pkg/trans"

	"time"
)

type User struct {
	Uid           int           `gorm:"not null;primary_key;AUTO_INCREMENT"json:"uid" form:"uid"` // uid
	Nickname      string        `gorm:"not null"json:"nickname" form:"nickname"`                  // nickname
	Username      string        `gorm:"not null"json:"username" form:"username"`                  // nickname
	DepartmentIds DepartmentIds `gorm:"not null;type:json"json:"departmentIds"`                   // DepartmentIds
	Email         string        `gorm:"not null"json:"email" form:"email"`                        // email
	Avatar        string        `gorm:"not null"json:"avatar" form:"avatar"`                      // avatar
	Password      string        `gorm:"not null"json:"-" form:"password"`                         // password
	State         int           `gorm:"not null"json:"state"form:"state"`                         // 状态
	Gender        int64         `gorm:"not null"json:"gender" form:"gender"`                      // gender
	Birthday      int64         `gorm:"not null"json:"birthday" form:"birthday"`                  // birthday
	Ctime         int64         `gorm:"not null"json:"ctime" form:"ctime"`                        // 创建时间
	Utime         int64         `gorm:"not null"json:"utime" form:"utime"`                        // 更新时间
	LastLoginIp   string        `gorm:"not null"json:"lastLoginIp" form:"lastLoginIp"`            // last_login_ip
	LastLoginTime int64         `gorm:"not null"json:"lastLoginTime" form:"lastLoginTime"`        // last_login_time

}

func (t *User) TableName() string {
	return "user"
}

type DepartmentIds []int

func (c DepartmentIds) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *DepartmentIds) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

// AddUser insert a new User into database and returns
// last inserted Id on success.
func UserCreate(db *gorm.DB, data *User) (err error) {
	data.Ctime = time.Now().Unix()
	data.Utime = time.Now().Unix()

	if err = db.Create(data).Error; err != nil {
		mus.Logger.Error("create user error", zap.Error(err))
		return
	}
	return
}

func UserUpdate(db *gorm.DB, paramId int, ups Ups) (err error) {
	var sql = "`uid`=?"
	var binds = []interface{}{paramId}

	if err = db.Table("user").Where(sql, binds...).Updates(ups).Error; err != nil {
		mus.Logger.Error("user update error", zap.Error(err))
		return
	}
	return
}

// UpdateX Update的扩展方法，根据Cond更新一条或多条记录
func UserUpdateX(db *gorm.DB, conds Conds, ups Ups) (err error) {
	sql, binds := BuildQuery(conds)
	if err = db.Table("user").Where(sql, binds...).Updates(ups).Error; err != nil {
		mus.Logger.Error("user update error", zap.Error(err))
		return
	}
	return
}

// Delete 根据主键删除一条记录。如果有delete_time则软删除，否则硬删除。
func UserDelete(db *gorm.DB, paramId int) (err error) {
	var sql = "`uid`=?"
	var binds = []interface{}{paramId}

	if err = db.Table("user").Where(sql, binds...).Delete(&User{}).Error; err != nil {
		mus.Logger.Error("user delete error", zap.Error(err))
		return
	}

	return
}

// DeleteX Delete的扩展方法，根据Cond删除一条或多条记录。如果有delete_time则软删除，否则硬删除。
func UserDeleteX(db *gorm.DB, conds Conds) (err error) {
	sql, binds := BuildQuery(conds)

	if err = db.Table("user").Where(sql, binds...).Delete(&User{}).Error; err != nil {
		mus.Logger.Error("user delete error", zap.Error(err))
		return
	}

	return
}

// Info 根据PRI查询单条记录
func UserInfo(db *gorm.DB, paramId int) (resp User, err error) {

	var sql = "`uid`= ?"

	var binds = []interface{}{paramId}

	if err = db.Table("user").Where(sql, binds...).First(&resp).Error; err != nil {
		mus.Logger.Error("user info error", zap.Error(err))
		return
	}
	return
}

// InfoX Info的扩展方法，根据Cond查询单条记录
func UserInfoX(db *gorm.DB, conds Conds) (resp User, err error) {

	sql, binds := BuildQuery(conds)

	if err = db.Table("user").Where(sql, binds...).First(&resp).Error; err != nil {
		mus.Logger.Error("user info error", zap.Error(err))
		return
	}
	return
}

// List 查询list，extra[0]为sorts
func UserList(conds Conds, extra ...string) (resp []*User, err error) {

	sql, binds := BuildQuery(conds)

	sorts := ""
	if len(extra) >= 1 {
		sorts = extra[0]
	}
	if err = mus.Db.Table("user").Where(sql, binds...).Order(sorts).Find(&resp).Error; err != nil {
		mus.Logger.Error("user info error", zap.Error(err))
		return
	}
	return
}

// ListMap 查询map，map遍历的时候是无序的，所以指定sorts参数没有意义
func UserListMap(conds Conds) (resp map[int]*User, err error) {
	sql, binds := BuildQuery(conds)

	mysqlSlice := make([]*User, 0)
	resp = make(map[int]*User, 0)
	if err = mus.Db.Table("user").Where(sql, binds...).Find(&mysqlSlice).Error; err != nil {
		mus.Logger.Error("user info error", zap.Error(err))
		return
	}
	for _, value := range mysqlSlice {
		resp[value.Uid] = value
	}
	return
}

// ListPage 根据分页条件查询list
func UserListPage(conds Conds, reqList *trans.ReqPage) (total int, respList []*User) {

	if reqList.PageSize == 0 {
		reqList.PageSize = 10
	}
	if reqList.Current == 0 {
		reqList.Current = 1
	}
	sql, binds := BuildQuery(conds)

	db := mus.Db.Table("user").Where(sql, binds...)
	respList = make([]*User, 0)
	db.Count(&total)
	db.Order(reqList.Sort).Offset((reqList.Current - 1) * reqList.PageSize).Limit(reqList.PageSize).Find(&respList)
	return
}
