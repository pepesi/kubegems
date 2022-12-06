package interfaces

import (
	"gorm.io/gorm"
)

/*
	func globalLogging(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
		log.Printf("[LOG] %s,%s\n", req.Request.Method, req.Request.URL)
		chain.ProcessFilter(req, resp)
	}

	func setlang(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
		if l := req.Request.Header.Get("Language"); l != "" {
			req.Request = req.Request.WithContext(context.WithValue(req.Request.Context(), options.LanguageKey, l))
		}
		chain.ProcessFilter(req, resp)
	}
*/
func InitHTTPServer(db *gorm.DB) error {
	// TODO: WIRE

	/*
		infraOpt := infrastructure.NewInfraOption(db)
		tenantRepo := repository.RepoFor(&model.Tenant{}, infraOpt)
		tenantRelRepo := repository.RepoFor(&model.TenantUserRel{}, infraOpt)
		userRepo := repository.RepoFor(&model.User{}, infraOpt)
		clusterRepo := repository.RepoFor(&model.Cluster{}, infraOpt)
		quotaRepo := repository.RepoFor(&model.Quota{}, infraOpt)
		tenantMgr := service.NewTenantManager(tenantRepo)
		tenant_agg_serivice := agg_services.NewTenantService(tenantMgr)
		tenant_member_agg_serivice := agg_services.NewTenantMemberService(tenantRepo, tenantRelRepo, userRepo)
		tenant_quota_agg_service := agg_services.NewTenantResourceQuotaService(tenantRepo, clusterRepo, quotaRepo)
		tenantHTTPInterface := api.NewTenantHTTPInterface(tenant_agg_serivice, tenant_member_agg_serivice, tenant_quota_agg_service)
		tenantAPI := api.NewTenantAPIService(tenantHTTPInterface)
		container := restful.NewContainer()
		container.Filter(globalLogging)
		container.Filter(setlang)
		container.Add(tenantAPI.WebService())

		for _, webs := range container.RegisteredWebServices() {
			for _, l := range webs.Routes() {
				log.Printf("[routers] %s  %s\n", l.Method, l.Path)
			}
		}
		return http.ListenAndServe(":8080", container)
	*/
	return nil
}
