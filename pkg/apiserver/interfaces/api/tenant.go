package api

import (
	"github.com/emicklei/go-restful/v3"
	agg_services "kubegems.io/kubegems/pkg/apiserver/aggregate"
	apis "kubegems.io/kubegems/pkg/apiserver/interfaces/dto/v1"
)

type TenantHTTPService struct {
	iface *TenantHTTPInterface
}

func NewTenantAPIService(iface *TenantHTTPInterface) *TenantHTTPService {
	return &TenantHTTPService{
		iface: iface,
	}
}

func (svc *TenantHTTPService) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/v1/tenant").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON).Doc("tenant api")
	println(ws.RootPath())

	ws.Route(ws.GET("/{name}").
		To(svc.iface.getTenant).
		Doc("retrieve tenant detail data").
		Param(restful.PathParameter("name", "tenant name").Required(true)).
		Writes(apis.RetrieveTenantResp{}).
		Returns(200, "ok", apis.RetrieveTenantResp{}).
		Returns(404, "not found", apis.NotFoundResp{}))
	ws.Route(ws.GET("/").
		To(svc.iface.listTenants).
		Doc("list tenants").
		Param(restful.QueryParameter("page", "page").DefaultValue("1")).
		Param(restful.QueryParameter("size", "size").DefaultValue("10")).
		Writes(apis.ListTenantResp{}).
		Returns(200, "ok", apis.ListTenantResp{}))
	ws.Route(ws.POST("/").
		To(svc.iface.createTeant).
		Doc("create a tenant").
		Param(restful.BodyParameter("name", "tenant name").Required(true)).
		Returns(201, "ok", apis.CreateUpdateTenantResp{}).
		Returns(400, "bad request", apis.BadRequest{}))
	ws.Route(ws.DELETE("/{name}").
		To(svc.iface.deleteTenant).
		Doc("delete tenant").
		Param(restful.PathParameter("name", "tenant name")).
		Returns(204, "ok", nil))
	ws.Route(ws.PUT("/{name}").
		To(svc.iface.modifyTenant).
		Doc("modify the tenant").
		Param(restful.PathParameter("name", "tenant name")).
		Returns(200, "ok", apis.CreateUpdateTenantResp{}).
		Returns(400, "bad request", apis.BadRequest{}))

	ws.Route(ws.POST("/{name}/clusterquota").To(svc.iface.createTenantClusterQuota))
	ws.Route(ws.PUT("/{name}/clusterquota/{cluster}").To(svc.iface.modifyTenantClusterQuota))
	ws.Route(ws.DELETE("/{name}/clusterquota/{cluster}").To(svc.iface.deleteTenantClusterQuota))
	ws.Route(ws.GET("/{name}/clusterquota/{cluster}").To(svc.iface.getTenantClusterQuota))
	ws.Route(ws.GET("/{name}/clusterquota/").To(svc.iface.listTenantClusterQuota))

	ws.Route(ws.POST("/{name}/member").To(svc.iface.addMember))
	ws.Route(ws.PUT("/{name}/member/{member_name}").To(svc.iface.modifyMemberRole))
	ws.Route(ws.DELETE("/{name}/member/{member_name}").To(svc.iface.deleteMember))
	ws.Route(ws.GET("/{name}/member").To(svc.iface.listMembers))
	return ws
}

func NewTenantHTTPInterface(tenantSvc agg_services.TenantService) *TenantHTTPInterface {
	return &TenantHTTPInterface{
		tenantSvc: tenantSvc,
	}
}

type TenantHTTPInterface struct {
	tenantSvc agg_services.TenantService
}

func (iface *TenantHTTPInterface) getTenant(req *restful.Request, resp *restful.Response) {
	tenantName := req.PathParameter("name")
	tenant, err := iface.tenantSvc.GetTenant(req.Request.Context(), tenantName)
	if err != nil {
		resp.WriteError(400, err)
		return
	}
	resp.WriteEntity(tenant)
}

func (iface *TenantHTTPInterface) createTeant(req *restful.Request, resp *restful.Response) {
	reqdata := &apis.CreateUpdateTenantReq{}
	if err := req.ReadEntity(reqdata); err != nil {
		resp.WriteError(400, err)
		return
	}
	if errs := apis.Validate(req.Request.Context(), reqdata); errs != nil {
		resp.WriteHeaderAndEntity(400, errs)
		return
	}
	respdata, err := iface.tenantSvc.CreateTenant(req.Request.Context(), nil)
	if err != nil {
		resp.WriteError(400, err)
		return
	}
	resp.WriteAsJson(respdata)
}

func (iface *TenantHTTPInterface) deleteTenant(req *restful.Request, resp *restful.Response) {
	tenantName := req.PathParameter("name")
	err := iface.tenantSvc.DeleteTenant(req.Request.Context(), tenantName)
	if err != nil {
		resp.WriteError(400, err)
		return
	}
	resp.WriteHeaderAndEntity(204, "no content")
}

func (iface *TenantHTTPInterface) modifyTenant(req *restful.Request, resp *restful.Response) {
	tenantName := req.PathParameter("name")
	reqdata := &apis.CreateUpdateTenantReq{}
	if err := req.ReadEntity(reqdata); err != nil {
		resp.WriteError(400, err)
		return
	}
	if err := iface.tenantSvc.UpdateTenant(req.Request.Context(), tenantName, reqdata); err != nil {
		resp.WriteError(400, err)
		return
	}
}

func (iface *TenantHTTPInterface) listTenants(req *restful.Request, resp *restful.Response) {
	req.QueryParameter("page")
	req.QueryParameter("size")
	respdata, err := iface.tenantSvc.ListTenants(req.Request.Context())
	if err != nil {
		resp.WriteError(400, err)
		return
	}
	resp.WriteAsJson(respdata)
}

func (iface *TenantHTTPInterface) createTenantClusterQuota(req *restful.Request, resp *restful.Response) {
	tenantName := req.PathParameter("tenant")
	clusterName := req.PathParameter("cluster")
	reqdata := &apis.CreateTenantClusterQuotaReq{}
	if err := req.ReadEntity(reqdata); err != nil {
		resp.WriteError(400, err)
	}
	respdata, err := iface.tenantSvc.CreateTenantClusterQuota(tenantName, clusterName, reqdata)
	if err != nil {
		resp.WriteError(400, err)
		return
	}
	resp.WriteAsJson(respdata)
}

func (iface *TenantHTTPInterface) modifyTenantClusterQuota(req *restful.Request, resp *restful.Response) {
}

func (iface *TenantHTTPInterface) deleteTenantClusterQuota(req *restful.Request, resp *restful.Response) {
}

func (iface *TenantHTTPInterface) getTenantClusterQuota(req *restful.Request, resp *restful.Response) {
}

func (iface *TenantHTTPInterface) listTenantClusterQuota(req *restful.Request, resp *restful.Response) {
}

func (iface *TenantHTTPInterface) addMember(req *restful.Request, resp *restful.Response) {
}

func (iface *TenantHTTPInterface) modifyMemberRole(req *restful.Request, resp *restful.Response) {
}

func (iface *TenantHTTPInterface) deleteMember(req *restful.Request, resp *restful.Response) {
}

func (iface *TenantHTTPInterface) listMembers(req *restful.Request, resp *restful.Response) {
	tenantName := req.PathParameter("tenant")
	iface.tenantSvc.ListMembers(tenantName)
}
