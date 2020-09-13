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

type Menu struct {
	Id                  int            `orm:"auto"json:"id" form:"id"`          // id
	Pid                 int            `json:"pid" form:"pid"`                  // 上级菜单id
	AppId               int            `json:"appId" form:"appId"`              // 应用id
	Name                string         `orm:"size(255)"json:"name" form:"name"` // 菜单名称
	Path                string         `orm:"size(255)"json:"path" form:"url"`  // 路由url
	Key                 string         `gorm:"-"json:"key"`
	PmsCode             string         `orm:"size(255)"json:"pmsCode" form:"pmsCode"` // 标识
	MenuPmsKey          MenuPmsKeyJson `gorm:"-"json:"menuPmsKey"`                    // 扩展信息
	MenuType            int            `json:"menuType" form:"menuType"`              // 类型 1=菜单 2=按钮
	Icon                string         `orm:"size(255)"json:"icon" form:"icon"`       // 图标
	OrderNum            int            `json:"orderNum" form:"orderNum"`              // 排序
	State               int            `json:"state"`
	Ctime               int64          `json:"ctime" form:"ctime"` // 创建时间
	Utime               int64          `json:"utime" form:"utime"` // 更新时间
	Actions             *MenuTrees     `json:"actions"`
	Children            *MenuTrees     `gorm:"-"json:"children"`
	ProLayoutParentKeys []string       `gorm:"-"json:"pro_layout_parentKeys"` // ant design pro
}

func (t *Menu) TableName() string {
	return "menu"
}

type MenuPmsKeyJson []string

func (c MenuPmsKeyJson) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *MenuPmsKeyJson) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

type MenuCheck struct {
	Id       int    `orm:"auto"json:"menuId" form:"id"`      // id
	Pid      int    `json:"pid" form:"pid"`                  // 上级菜单id
	AppId    int    `json:"appId" form:"appId"`              // 应用id
	MenuType int    `json:"menuType" form:"menuType"`        // 类型 1=菜单 2=按钮
	Name     string `orm:"size(255)"json:"name" form:"name"` // 菜单名称
	Url      string `orm:"size(255)"json:"url" form:"url"`   // 路由url
	Actions  []int  `json:"actions"`
}

func (t *MenuCheck) TableName() string {
	return "menu"
}

// MenuTrees 菜单树列表
type MenuTrees []*Menu

type Menus []*Menu

func (m Menus) ToTree() MenuTrees {
	mTreeMap := make(map[int]*Menu)
	for _, item := range m {
		mTreeMap[item.Id] = item
	}

	list := make(MenuTrees, 0)
	for _, item := range m {
		item.Key = item.Path
		if item.Pid == 0 {
			//if item.ProLayoutParentKeys == nil {
			//	item.ProLayoutParentKeys = make([]string,0)
			//}
			//item.ProLayoutParentKeys = append(item.ProLayoutParentKeys,item.Path)
			list = append(list, item)
			continue
		}

		if pItem, ok := mTreeMap[item.Pid]; ok {
			if item.MenuType == 2 {
				if pItem.Actions == nil {
					children := MenuTrees{item}
					pItem.Actions = &children
					continue
				}
				*pItem.Actions = append(*pItem.Actions, item)
			} else {
				if item.ProLayoutParentKeys == nil {
					item.ProLayoutParentKeys = make([]string, 0)
				}
				item.ProLayoutParentKeys = append(item.ProLayoutParentKeys, pItem.Path)

				if pItem.Children == nil {
					children := MenuTrees{item}
					pItem.Children = &children
					continue
				}

				*pItem.Children = append(*pItem.Children, item)
			}
		}
	}
	return list
}

type MenuChecks []*MenuCheck

func (m MenuChecks) ToCheckArr() MenuChecks {
	mTreeMap := make(map[int]*MenuCheck)
	for _, item := range m {
		mTreeMap[item.Id] = item
	}

	list := make(MenuChecks, 0)
	for _, item := range m {
		if item.MenuType == 1 {
			list = append(list, item)
			continue
		}

		if pItem, ok := mTreeMap[item.Pid]; ok {
			if item.MenuType == 2 {
				if pItem.Actions == nil {
					pItem.Actions = make([]int, 0)
				}
				pItem.Actions = append(pItem.Actions, item.Id)
			}
		}
	}
	return list
}

func MenuAllListByIds(ids []int) MenuChecks {
	respList := make(MenuChecks, 0)
	mus.Db.Where("id in (?)", ids).Find(&respList)
	return respList
}

func MenuListByIds(ids []int) Menus {
	menus := make(Menus, 0)
	mus.Db.Where("id in (?) and menu_type=1", ids).Find(&menus)
	return menus
}

func MenuApiListByIds(ids []int) Menus {
	menus := make(Menus, 0)
	mus.Db.Where("id in (?) and menu_type=2", ids).Find(&menus)
	return menus
}

// AddMenu insert a new Menu into database and returns
// last inserted Id on success.
func MenuCreate(db *gorm.DB, data *Menu) (err error) {
	data.Ctime = time.Now().Unix()
	data.Utime = time.Now().Unix()

	if err = db.Create(data).Error; err != nil {
		mus.Logger.Error("create menu error", zap.Error(err))
		return
	}
	return
}

func MenuUpdate(db *gorm.DB, paramId int, ups Ups) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}

	if err = db.Table("menu").Where(sql, binds...).Updates(ups).Error; err != nil {
		mus.Logger.Error("menu update error", zap.Error(err))
		return
	}
	return
}

// UpdateX Update的扩展方法，根据Cond更新一条或多条记录
func MenuUpdateX(db *gorm.DB, conds Conds, ups Ups) (err error) {
	sql, binds := BuildQuery(conds)
	if err = db.Table("menu").Where(sql, binds...).Updates(ups).Error; err != nil {
		mus.Logger.Error("menu update error", zap.Error(err))
		return
	}
	return
}

// Delete 根据主键删除一条记录。如果有delete_time则软删除，否则硬删除。
func MenuDelete(db *gorm.DB, paramId int) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}

	if err = db.Table("menu").Where(sql, binds...).Delete(&Menu{}).Error; err != nil {
		mus.Logger.Error("menu delete error", zap.Error(err))
		return
	}

	return
}

// DeleteX Delete的扩展方法，根据Cond删除一条或多条记录。如果有delete_time则软删除，否则硬删除。
func MenuDeleteX(db *gorm.DB, conds Conds) (err error) {
	sql, binds := BuildQuery(conds)

	if err = db.Table("menu").Where(sql, binds...).Delete(&Menu{}).Error; err != nil {
		mus.Logger.Error("menu delete error", zap.Error(err))
		return
	}

	return
}

// Info 根据PRI查询单条记录
func MenuInfo(db *gorm.DB, paramId int) (resp Menu, err error) {

	var sql = "`id`= ?"

	var binds = []interface{}{paramId}

	if err = db.Table("menu").Where(sql, binds...).First(&resp).Error; err != nil {
		mus.Logger.Error("menu info error", zap.Error(err))
		return
	}
	return
}

// InfoX Info的扩展方法，根据Cond查询单条记录
func MenuInfoX(db *gorm.DB, conds Conds) (resp Menu, err error) {

	sql, binds := BuildQuery(conds)

	if err = db.Table("menu").Where(sql, binds...).First(&resp).Error; err != nil {
		mus.Logger.Error("menu info error", zap.Error(err))
		return
	}
	return
}

// List 查询list，extra[0]为sorts
func MenuList(conds Conds, extra ...string) (resp []*Menu, err error) {

	sql, binds := BuildQuery(conds)

	sorts := ""
	if len(extra) >= 1 {
		sorts = extra[0]
	}
	if err = mus.Db.Table("menu").Where(sql, binds...).Order(sorts).Find(&resp).Error; err != nil {
		mus.Logger.Error("menu info error", zap.Error(err))
		return
	}
	return
}

// ListMap 查询map，map遍历的时候是无序的，所以指定sorts参数没有意义
func MenuListMap(conds Conds) (resp map[int]*Menu, err error) {

	sql, binds := BuildQuery(conds)

	mysqlSlice := make([]*Menu, 0)
	resp = make(map[int]*Menu, 0)
	if err = mus.Db.Table("menu").Where(sql, binds...).Find(&mysqlSlice).Error; err != nil {
		mus.Logger.Error("menu info error", zap.Error(err))
		return
	}
	for _, value := range mysqlSlice {
		resp[value.Id] = value
	}
	return
}

// ListPage 根据分页条件查询list
func MenuListPage(conds Conds, reqList *trans.ReqPage) (total int, respList []*Menu) {

	if reqList.PageSize == 0 {
		reqList.PageSize = 10
	}
	if reqList.Current == 0 {
		reqList.Current = 1
	}
	sql, binds := BuildQuery(conds)

	db := mus.Db.Table("menu").Where(sql, binds...)
	respList = make([]*Menu, 0)
	db.Count(&total)
	db.Order(reqList.Sort).Offset((reqList.Current - 1) * reqList.PageSize).Limit(reqList.PageSize).Find(&respList)
	return
}
