package service

import (
	"fmt"
	"go.uber.org/zap"
	"oapms/pkg/model"
	"oapms/pkg/mus"
	"oapms/pkg/service/pmsplugin"
	"strconv"
	"strings"
)

func (*pms) AssignMenuPms(roleId int, menuIds []int, dataPermCount int) (err error) {
	roleData, err := model.RoleInfo(mus.Db, roleId)
	if err != nil {
		return
	}

	menus := model.MenuApiListByIds(menuIds)

	if len(menus) > 0 {
		var policies [][]string
		for _, m := range menus {
			fmt.Println("minfo", m.Path, m.PmsCode)
			if m.Path == "" && m.PmsCode != "" {
				//Do not allow comma which would cause panic error with casbin rules
				m.PmsCode = strings.Replace(m.PmsCode, ",", "|", -1)
				policies = append(policies, []string{roleData.Name, m.PmsCode, "*", strconv.Itoa(roleData.AppId)})
			}
		}
		mus.Logger.Debug("policies info", zap.Any("policies", policies))

		OverwritePerm(roleData.Name, strconv.Itoa(roleData.AppId), policies)
	} else {
		// if we have data permission sets , should not remove entire role relative records
		// fixed issue #23
		if dataPermCount > 0 {
			DeletePermPolicy(roleData.Name)
		} else {
			DeletePerm(roleData.Name)
		}
	}
	return
}

// assign data permission
func (*pms) AssignDataPms(roleId int, dataPermIds []int) error {
	var (
		oldDataPermIds []int
		dataPmsArr     []model.DataPms
		err            error
	)

	// Get the old data permission list of the current role
	oldRoleDataPerms, _ := model.PmsByRoleId(roleId)

	tx := mus.Db.Begin()

	// Delete all old data permissions for this role and insert new ones
	if len(oldRoleDataPerms) > 0 {
		for _, v := range oldRoleDataPerms {
			oldDataPermIds = append(oldDataPermIds, v.Id)
		}
		err = model.RolePmsDeleteMulti(tx, roleId, oldDataPermIds)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	// Insert new data permissions
	if len(dataPermIds) > 0 {
		for _, v := range dataPermIds {
			oneDataPms := model.DataPms{}
			oneDataPms.RoleId = roleId
			oneDataPms.DataPmsId = v

			dataPmsArr = append(dataPmsArr, oneDataPms)
		}
		err = model.RolePmsInsertMulti(tx, dataPmsArr)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

// OverwritePerm : overwrite permissions
// remove or create policy
func OverwritePerm(roleName, appId string, polices [][]string) {
	currentPerms := pmsplugin.GetAllPermsByRoleDomain(roleName, appId)
	for k1, newPerm := range polices {
		for k2, currentPerm := range currentPerms {
			if newPerm[0] == currentPerm[0] &&
				newPerm[1] == currentPerm[1] &&
				newPerm[2] == currentPerm[2] &&
				newPerm[3] == currentPerm[3] {
				// 如果相等
				polices[k1] = []string{"-skip"}
				currentPerms[k2] = []string{"-skip"}
			}
		}
	}
	for _, newPerm := range polices {
		if newPerm[0] == "-skip" {
			continue
		}
		pmsplugin.AddPerm(newPerm)
	}
	for _, remPerm := range currentPerms {
		if remPerm[0] == "-skip" {
			continue
		}
		pmsplugin.DelPerm(remPerm)
	}
}

// DeletePerm : delete role in casbin policies
func DeletePerm(roleName string) {
	pmsplugin.DelRole(roleName)
}

// DeletePermPolicy : delete role in casbin policies
func DeletePermPolicy(roleName string) {
	pmsplugin.DeleteRolePolicy(roleName)
}
