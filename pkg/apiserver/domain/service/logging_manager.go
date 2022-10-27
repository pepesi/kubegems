package service

import (
	"context"

	"kubegems.io/kubegems/pkg/apiserver/options"
)

// 日志模块
// 日志插件分采集插件和查询插件
// 查询插件 需要由插件管理器实现 Install, Uninstall，在集群上真实安装和卸载
// api 侧需要实现 LogQueryer 的接口，适配 对应安装的 查询插件
// api 侧需要实现 LogCollector 的接口，适配 对应安装的 采集插件

// LogQueryer 日志查询
type LogQuryer interface {
	Labels(ctx context.Context, opts ...options.Option) []string
	Logs(ctx context.Context, opts ...options.Option) []string
}

// LogCollector 负责日志采集器的管理
type LogCollector interface {
	// SetUpCollector 安装一个日志采集器
	SetUpCollector(ctx context.Context, target, name, path string) error
	// UpdateCollector 更新一个日志采集器
	UpdateCollector(ctx context.Context, target, name, path string) error
	// DeleteCollector 删除采集器
	DeleteCollector(ctx context.Context, trget, name string) error
}
