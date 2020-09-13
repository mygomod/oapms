package service

import (
	"go.uber.org/zap"
	"oapms/pkg/model"
	"oapms/pkg/mus"
	"oapms/pkg/service/pmsplugin"
	"strconv"
	"strings"
)

var (
	/* Dao layer */
	RootAppId = 0
)

type pms struct{}

func InitPms() *pms {
	pmsplugin.Invoker(true)
	return &pms{}
}

const (
	UserStatusNotPub = iota
	UserStatusNormal
	UserStatusLock
)

// AssignRole - assign roles to specific user
// 这个方法同时作用与用户角色，用户用户组
func (p *pms) AssignRole(userId int, roleIds []int) {
	var roles [][]string
	for _, roleId := range roleIds {
		if userId == 0 || roleId == 0 {
			continue
		}

		roles = append(roles, []string{strconv.Itoa(userId), strconv.Itoa(roleId)})
	}
	p.OverwriteRoles(strconv.Itoa(userId), roles)
}

// OverwriteRoles : assign roles to specific user
func (*pms) OverwriteRoles(userId string, newRoles [][]string) {
	currentRoles := pmsplugin.GetGroupsByUser(userId)
	for k1, newRole := range newRoles {
		for k2, currentRole := range currentRoles {
			if newRole[0] == currentRole[0] && newRole[1] == currentRole[1] {
				newRoles[k1] = []string{"-skip"}
				currentRoles[k2] = []string{"-skip"}
			}
		}
	}
	for _, newRole := range newRoles {
		if newRole[0] == "-skip" {
			continue
		}
		pmsplugin.AddGroup(newRole)
	}
	for _, rmRole := range currentRoles {
		if rmRole[0] == "-skip" {
			continue
		}
		pmsplugin.DelGroup(rmRole)
	}
}

// GetRelatedDomains - get related app
func (p *pms) GetRelatedApp(uid string, skipRoot bool) []model.App {
	apps := make([]model.App, 0)
	var single = map[int]bool{}
	//1.get roles by user
	roleIds := p.GetAllRoleIds(uid)
	mus.Logger.Debug("get group user roles", zap.String("uid", uid), zap.Any("roleIds", roleIds))
	//2.get domains by roles
	for _, roleId := range roleIds {
		// role name
		// 数据结构 2 数据中台管理员
		role, err := model.RoleInfo(mus.Db, roleId)
		if err != nil {
			mus.Logger.Error("get role info err", zap.Error(err), zap.Int("roleId", roleId))
			continue
		}

		// 排除oa系统
		if skipRoot {
			if role.AppId == RootAppId {
				continue
			}
		}
		if _, ok := single[role.AppId]; !ok {
			appInfo := model.App{}
			err = mus.Db.Select("aid,name,url").Where("aid = ?", role.AppId).Find(&appInfo).Error
			if err != nil {
				mus.Logger.Error("get app info error", zap.Error(err))
				continue
			}
			single[role.AppId] = true
			apps = append(apps, appInfo)
		}
	}
	return apps
}

// GetAllRoles would return all roles of a user
func (*pms) GetAllRoles(uid string) []string {
	groups := pmsplugin.GetGroupsByUser(uid)
	roles := []string{}
	for _, group := range groups {
		roles = append(roles, group[1])
	}
	return roles
}

// GetAllRoles would return all roles of a user
func (*pms) GetAllRoleIds(uid string) []int {
	groups := pmsplugin.GetGroupsByUser(uid)
	roles := []int{}
	for _, group := range groups {
		roleId, err := strconv.Atoi(group[1])
		if err != nil {
			continue
		}
		roles = append(roles, roleId)
	}
	return roles
}

// GetAllPermissions - get all permission by specific user
func (*pms) GetAllPermissions(uid string) []string {
	perms := []string{}
	var path = map[string]bool{}

	// 数据结构
	// pms[0]   pms[1]
	// 系统设置  /system/menu:show
	for _, perm := range pmsplugin.GetAllPmsByUser(uid) {
		prefix := strings.Split(perm[1], ":")
		seg := strings.Split(prefix[0], "/")
		ss := ""
		for _, s := range seg[1:] {
			ss += "/" + s
			if ok := path[ss]; !ok {
				path[ss] = true
				perms = append(perms, ss)
			}
		}
		perms = append(perms, perm[1])
	}
	return perms
}

//GetPermissionsOfDomain - Get pure permission list  in specific domain(another backend system)
func (*pms) GetPermissionsOfApp(uid string, app string) []string {
	perms := pmsplugin.GetAllPmsByUser(uid)
	var polices []string
	for _, p := range perms {
		if p[3] == app {
			polices = append(polices, p[1])
		}
	}
	return polices
}

//GetDataPermissionsOfDomain - Get data permission list  in specific domain(another backend system)
func (us *pms) GetDataPermissionsOfDomain(uid, appId string) []map[string]string {
	gs := pmsplugin.GetGroupsByUser(uid)
	var (
		polices []map[string]string
		roles   []string
	)
	for _, p := range gs {
		roles = append(roles, p[1])
	}
	dmHash := map[int]bool{}
	for _, dm := range us.GetRelatedApp(uid, false) {
		if strconv.Itoa(dm.Aid) != appId {
			continue
		}
		dmHash[dm.Aid] = true
	}
	for _, r := range model.RolesByRoleIds(roles) {
		for _, dp := range r.Pms {
			// 2 为数据权限
			if dp.PmsType == 2 {
				if _, ok := dmHash[dp.AppId]; ok {
					polices = append(polices, map[string]string{
						"pms":    dp.PmsCode,
						"rule":   dp.PmsRule,
						"weight": strconv.Itoa(dp.OrderNum),
					})
				}
			}
		}
	}
	return polices
}

//GetMenusOfDomain - get menus in specific domain
func (*pms) GetMenusOfDomain(uid string, appId string) []*model.Role {
	roles := pmsplugin.GetGroupsByUser(uid)
	var roleNames []string
	for _, r := range roles {
		roleNames = append(roleNames, r[1])
	}
	return model.RolesByRoleIds(roleNames)
}

// GetAppMenu -  get specific user's menus of specific app
// 根据uid和appId，获取该用户在某个应用下的菜单
func (p *pms) GetAppMenu(uid string, appId string) model.Menus {
	roles := p.GetAllRoles(uid)
	mus.Logger.Debug("get roles", zap.String("mod", "pms_get_app_menu"), zap.String("uid", uid), zap.Strings("roles", roles))
	mids := []int{}
	for _, r := range model.RolesByRoleIds(roles) {
		if strconv.Itoa(r.AppId) == appId {
			mids = append(mids, r.MenuIds...)
		}
	}

	return model.MenuListByIds(mids)
}

//CheckPermission - check user's permission in specific domain with specific policy
func (p *pms) CheckPermission(uid string, appId int, policy string) bool {
	info, err := model.MenuPmsInfoX(mus.Db, model.Conds{
		"app_id": appId,
		"key":    policy,
	})
	if err != nil {
		return false
	}
	// 权限必须是str，否则会判断有问题
	appIdStr := strconv.Itoa(appId)

	policy = info.PmsCode
	roles := p.GetAllRoles(uid)
	for _, role := range roles {
		flag := pmsplugin.Enforce(role, policy, "*", appIdStr)
		mus.Logger.Debug("check permission", zap.String("appid", appIdStr), zap.String("role", role), zap.String("policy", policy), zap.Bool("flag", flag))
		if flag {
			return true
		}
	}
	return false
}
