package service

import (
	dingtalk "github.com/bullteam/go-dingtalk/src"
	"oapms/pkg/mus"
)

type dingTalk struct {
}

// GetDepartment - get departments of dingding
func (d *dingTalk) GetDepartments() ([]dingtalk.Department, error) {
	_ = mus.DingTalkClient.RefreshCompanyAccessToken()
	list, err := mus.DingTalkClient.DepartmentList(1, "zh_CN")
	if err != nil {
		return nil, err
	}
	return list.Department, nil
}

// GetUsers - get users of dingding
func (d *dingTalk) GetUsers(departmentId int) ([]dingtalk.UDetailedList, error) {
	list, err := mus.DingTalkClient.UserList(departmentId)
	if err != nil {
		return nil, err
	}
	return list.UserList, nil
}
