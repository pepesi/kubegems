package service

// 监控和告警

// 监控插件
// 服务端插件，agent插件
// 监控采集规则管理器
// 监控查询器

// 告警插件
// 告警依赖监控插件, 服务端插件
// 告警规则管理器, 告警通道管理器

type MonitorManager interface {
	AddMonitorUnits(target, template, values interface{})
	DeleteMonitorUnits(target, template string)
}

type AlertManager interface {
	AddRule(ruleName, ruleContent string) error
	ModifyRule(ruleName, ruleContent string) error
	DeleteRule(ruleName string) error
}

type MonitorUnitTemplateManager interface {
	ListTemplate()
	CreateTemplate()
	DeleteTemplate()
	ModifyTemplate()
}
