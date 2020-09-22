package api

import (
	dingtalk "github.com/bullteam/go-dingtalk/src"
	"github.com/jinzhu/gorm"
	"oapms/pkg/model"
	"oapms/pkg/mus"
	"oapms/pkg/router/core"
	"oapms/pkg/service"
	"strings"
)

func Sync(c *core.Context) {
	departments, err := service.DingTalk.GetDepartments()
	if err != nil {
		c.JSONErrTips(err.Error(), err)
		return
	}
	tx := mus.Db.Begin()
	err = syncDepartment(tx, departments)
	if err != nil {
		tx.Rollback()
		c.JSONErrTips(err.Error(), err)
		return
	}
	tx.Commit()
	c.JSONOK()
}

func syncDepartment(tx *gorm.DB, departments []dingtalk.Department) (err error) {
	for _, department := range departments {
		var info model.Department
		info, err = model.DepartmentInfo(tx, department.Id)
		// 如果是系统错误
		if err != nil && !gorm.IsRecordNotFoundError(err) {
			return
		}
		// 如果不存在创建
		if gorm.IsRecordNotFoundError(err) {
			err = model.DepartmentCreate(tx, &model.Department{
				Id:   department.Id,
				Name: department.Name,
				Pid:  department.ParentId,
			})
			if err != nil {
				return
			}

			err = syncUser(tx, department.Id)
			if err != nil {
				return
			}
			continue
		}

		// 如果存在更新
		err = model.DepartmentUpdate(tx, info.Id, model.Ups{
			"name": department.Name,
			"pid":  department.ParentId,
		})
		if err != nil {
			return
		}

		err = syncUser(tx, department.Id)
		if err != nil {
			return
		}
	}
	return
}

func syncUser(tx *gorm.DB, id int) (err error) {
	users, err := service.DingTalk.GetUsers(id)
	if err != nil {
		return
	}
	for _, user := range users {
		state := 0
		if !user.Active {
			state = 2
		}
		userName := user.Name
		if user.Email != "" {
			userName = strings.Split(user.Email, "@")[0]
		}
		var info model.User
		info, err = model.UserInfoX(tx, model.Conds{
			"nickname": user.Name,
		})
		// 如果是系统错误
		if err != nil && !gorm.IsRecordNotFoundError(err) {
			return
		}

		// 如果不存在创建
		if gorm.IsRecordNotFoundError(err) {
			err = model.UserCreate(tx, &model.User{
				Nickname:      user.Name,
				Username:      userName,
				Email:         user.Email,
				Avatar:        user.Avatar,
				DepartmentIds: user.Department,
				State:         state,
			})
			if err != nil {
				return
			}
			continue
		}

		// 看是否为已经激活
		if user.Active {
			state = info.State
		}

		// 如果存在更新
		err = model.UserUpdate(tx, info.Uid, model.Ups{
			"username":       userName,
			"email":          user.Email,
			"avatar":         user.Avatar,
			"department_ids": model.DepartmentIds(user.Department),
			"state":          state,
		})
		if err != nil {
			return
		}
	}

	return
}
