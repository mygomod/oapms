package service

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"oapms/pkg/model"
	"oapms/pkg/model/typex"
	"oapms/pkg/mus"
)

type pmsMenu struct{}

func InitPmsMenu() *pmsMenu {
	return &pmsMenu{}
}

// todo menu not delete
// todo transaction
func (p *pmsMenu) AutoMenu(appId int, list []typex.Url) {

	menuCodeMap := p.MenuCodeMap(list)
	codeNameMap := p.CodeNameMap(list)

	var err error
	for _, value := range list {
		info := model.MenuPms{}
		err = mus.Db.Select("id").Where("`key` = ? and `app_id` = ?", value.Key, appId).Find(&info).Error
		if err != nil && !gorm.IsRecordNotFoundError(err) {
			mus.Logger.Error("get auto menu error", zap.Any("pms extend", value))
			continue
		}

		create := model.MenuPms{
			PmsCode: value.Code,
			Key:     value.Key,
			AppId:   appId,
		}

		if gorm.IsRecordNotFoundError(err) {
			err = model.MenuPmsCreate(mus.Db, &create)
			if err != nil {
				mus.Logger.Error("create auto menu error", zap.Any("pms extend", info))
				continue
			}
		} else {
			err = model.MenuPmsUpdate(mus.Db, info.Id, model.Ups{
				"pms_code": value.Code,
			})
			if err != nil {
				mus.Logger.Error("update auto menu error", zap.Any("pms extend", value))
				continue
			}
		}
	}

	// 自动添加到菜单
	for menu, codeArr := range menuCodeMap {
		info := model.Menu{}
		// 肯定找得到的,没找到error
		err = mus.Db.Select("id").Where("`path` = ? and `app_id` = ? and `menu_type` = ?", menu, appId, 1).Find(&info).Error
		if err != nil {
			mus.Logger.Error("get auto menu error", zap.Any("info", info), zap.Error(err))
			continue
		}

		for _, code := range codeArr {
			apiInfo := model.Menu{}
			err = mus.Db.Select("id").Where("`pms_code` = ? and `app_id` = ? and `menu_type` = ?", code, appId, 2).Find(&apiInfo).Error
			if err != nil && !gorm.IsRecordNotFoundError(err) {
				mus.Logger.Error("get auto menu error", zap.Any("info", info), zap.Error(err))
				continue
			}

			create := model.Menu{
				PmsCode:  code,
				Name:     codeNameMap[code],
				AppId:    appId,
				MenuType: 2,
				State:    1,
				Pid:      info.Id,
			}

			if gorm.IsRecordNotFoundError(err) {
				err = model.MenuCreate(mus.Db, &create)
				if err != nil {
					mus.Logger.Error("create auto menu error", zap.Any("pms extend", info))
					continue
				}
			} else {
				err = model.MenuUpdate(mus.Db, apiInfo.Id, model.Ups{
					"name": codeNameMap[code],
					"pid":  info.Id,
				})
				if err != nil {
					mus.Logger.Error("update auto menu error", zap.Any("pms extend", apiInfo))
					continue
				}
			}

		}
	}
}

func (*pmsMenu) MenuCodeMap(list []typex.Url) (output map[string][]string) {
	output = make(map[string][]string, 0)
	for _, value := range list {
		for _, menu := range value.Menus {
			codeArr, ok := output[menu]
			if !ok {
				codeArr = make([]string, 0)
			}
			codeArr = append(codeArr, value.Code)
			output[menu] = codeArr
		}
	}
	return
}

func (*pmsMenu) CodeNameMap(list []typex.Url) (output map[string]string) {
	output = make(map[string]string, 0)
	for _, value := range list {
		output[value.Code] = value.Name
	}
	return
}
