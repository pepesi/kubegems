package agg_services

import (
	"kubegems.io/kubegems/pkg/apiserver/domain/model"
)

type ApplicationService struct{}

// 应用聚合上下文
// 应用-环境管理
// 应用部署
// 应用成员
// 应用其他运维动作(exec, logs)

// JoinEnvironment 加入环境
func (s *ApplicationService) JoinEnvironment(target model.Environment) {}

// LeaveEnvironment 离开环境
func (s *ApplicationService) LeaveEnvironment(target model.Environment) {}

// Deploy 部署到目标环境
func (s *ApplicationService) Deploy(target model.Environment, diff string) {}

// Rollback 回滚到目标revision
func (s *ApplicationService) Rollback(target model.Environment, revision string) {}

// ExportManifest 导出manifests
func (s *ApplicationService) ExportManifest(target model.Environment, revision string) {}

// AddMember 添加应用成员
func (s *ApplicationService) AddMember(app string, user model.User, role string) {}

// ModifyMember 修改应用成员
func (s *ApplicationService) ModifyMember(app string, user model.User, role string) {}

// DeleteMember 删除成员
func (s *ApplicationService) DeleteMember(app string, user model.User) {}

// Logs 查询应用日志
func (s *ApplicationService) Logs(target model.Environment, selector map[string]string) {}

// RelatedWorkloads 查看应用关联的workloads(拓扑数据)
func (s *ApplicationService) RelatedWorkloads(target model.Environment, selector map[string]string) {}
