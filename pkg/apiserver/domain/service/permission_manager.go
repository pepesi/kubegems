package service

import "kubegems.io/kubegems/pkg/apiserver/domain/model"

type PermissionManager interface {
	CopyUserPermFrom(from, to *model.User)
}

type permissionManager struct{}

func (mgr *permissionManager) CopyUserPermFrom(from, to *model.User) {}

var _ PermissionManager = &permissionManager{}
