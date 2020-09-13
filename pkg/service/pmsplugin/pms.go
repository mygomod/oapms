package pmsplugin

import (
	rediswatcher "github.com/billcobbler/casbin-redis-watcher"
	"github.com/casbin/casbin"
	"github.com/spf13/viper"
	"oapms/pkg/mus"
	"oapms/pkg/service/pmsplugin/adapter"
	"sync"
)

var (
	enforcer     *casbin.Enforcer
	enforcerLock = &sync.Mutex{}
)

// SetUp permission handler
func Invoker(enable bool) {
	enforcer = casbin.NewEnforcer(viper.GetString("casbin.rule.path"), adapter.NewMysqlAdapter())
	if enable {
		//Distributed watcher
		w, _ := rediswatcher.NewWatcher(viper.GetString("casbin.redisHost"), rediswatcher.Password(viper.GetString("casbin.redisPwd")))
		enforcer.SetWatcher(w)
		// @Overwrite
		// See if policy changed and do distributed notification
		_ = w.SetUpdateCallback(func(s string) {
			mus.Logger.Info("Casbin policies changed")
			enforcerLock.Lock()
			_ = enforcer.LoadPolicy()
			enforcerLock.Unlock()
		})
	}
}

// SetUpForTest : for unit tests
func SetUpForTest(dir string) {
	enforcer = casbin.NewEnforcer(dir+"/rbac_model.conf", dir+"/perm_test.csv")
}

// AddGroup : method of group policy adding
//first : user
//second : group
func AddGroup(params ...interface{}) bool {
	enforcerLock.Lock()
	defer enforcerLock.Unlock()
	return enforcer.AddGroupingPolicy(params...)
}

// AddGroupRole : assign role to a user group
func AddGroupRole(params ...interface{}) bool {
	enforcerLock.Lock()
	defer enforcerLock.Unlock()
	return enforcer.AddNamedGroupingPolicy("g2", params...)
}

// DelGroup : method of group policy deleting
func DelGroup(params ...interface{}) bool {
	enforcerLock.Lock()
	defer enforcerLock.Unlock()
	return enforcer.RemoveGroupingPolicy(params...)
}

// DelGroupPerm : delete user group - role connection
func DelGroupPerm(params ...interface{}) bool {
	enforcerLock.Lock()
	defer enforcerLock.Unlock()
	return enforcer.RemoveNamedGroupingPolicy("g2", params...)
}

// GetGroupsByUser : get groups by specific user
// roles
// role 0 uid
// role 1 role name
func GetGroupsByUser(userId string) [][]string {
	return enforcer.GetFilteredGroupingPolicy(0, userId)
}

// AddPerm : method for permission policy adding
//sub,obj,act,domain
func AddPerm(params ...interface{}) bool {
	enforcerLock.Lock()
	defer enforcerLock.Unlock()
	return enforcer.AddPolicy(params...)
}

// DelPerm : delete permission policy
func DelPerm(params ...interface{}) bool {
	enforcerLock.Lock()
	defer enforcerLock.Unlock()
	return enforcer.RemovePolicy(params...)
}

// DeleteFilteredPerm
func DelFilteredPerm(fieldIndex int, params ...string) bool {
	enforcerLock.Lock()
	defer enforcerLock.Unlock()
	return enforcer.RemoveFilteredPolicy(fieldIndex, params...)
}

// Enforce : check permission
func Enforce(params ...interface{}) bool {
	enforcerLock.Lock()
	defer enforcerLock.Unlock()
	return enforcer.Enforce(params...)
}

// DelRoleByName : delete all specific role policy of domain
func DelRoleByDomain(role string, domain string) {
	DelFilteredPerm(0, role, "", "", domain)
}

// DelRole : delete specific role
func DelRole(role string) {
	DelFilteredPerm(0, role)
	enforcer.RemoveFilteredGroupingPolicy(1, role)
}

// DeleteRolePolicy : delete policy row
func DeleteRolePolicy(role string) {
	DelFilteredPerm(0, role)
}

// GetAllPermsByRoleDomain : get policies by role and appId
func GetAllPermsByRoleDomain(role string, appId string) [][]string {
	perms := enforcer.GetFilteredNamedPolicy("p", 0, role, "", "", appId)
	return perms
}

// GetAllPmsByRole : get all permission across domains
func GetAllPmsByRole(role string) [][]string {
	perms := enforcer.GetFilteredNamedPolicy("p", 0, role, "", "", "")
	return perms
}

// GetAllPermsByUser : get all permission across app
func GetAllPmsByUser(uid string) [][]string {
	pms := enforcer.GetFilteredNamedGroupingPolicy("g", 0, uid, "", "", "")
	var policies [][]string
	for _, policy := range pms {
		rp := GetAllPmsByRole(policy[1])
		policies = append(policies, rp...)
	}
	return policies
}

//dangerous! do not call until you really need it
func CommitChange() {
	_ = enforcer.SavePolicy()
}
