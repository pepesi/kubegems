package endpoint

import (
	"context"

	agg_services "kubegems.io/kubegems/pkg/apiserver/aggregate"
	v1 "kubegems.io/kubegems/pkg/apiserver/interfaces/apis/v1"
)

func NewTenantServer(business agg_services.TenantService) v1.TenantServiceServer {
	return &tenantImpl{
		tenantSvc: business,
	}
}

type tenantImpl struct {
	v1.UnimplementedTenantServiceServer
	tenantSvc agg_services.TenantService
}

func (s *tenantImpl) CreateTenant(ctx context.Context, req *v1.CreateTenantRequest) (*v1.CreateTenantResponse, error) {
	err := req.ValidateAll()
	if err != nil {
		return nil, err
	}
	resp := &v1.CreateTenantResponse{
		Succeed: false,
		Message: "xxxx",
	}
	_, err = s.tenantSvc.CreateTenant(ctx, req)
	if err != nil {
		resp.Message = err.Error()
		resp.Succeed = false
	}
	resp.Message = "ok"
	resp.Succeed = true
	return resp, nil
}

func (s *tenantImpl) DeleteTenant(ctx context.Context, req *v1.DeleteTenantRequest) (*v1.DeleteTenantResponse, error) {
	resp := &v1.DeleteTenantResponse{
		Message: "ok",
		Succeed: true,
	}
	if err := s.tenantSvc.DeleteTenant(ctx, req.Name); err != nil {
		resp.Message = err.Error()
		resp.Succeed = false
	}
	return resp, nil
}

func (s *tenantImpl) CreateTenantClusterResourceQuota(ctx context.Context, req *v1.CreateTenantClusterResourceQuotaRequest) (*v1.CreateTenantClusterResourceQuotaResponse, error) {
	resp := &v1.CreateTenantClusterResourceQuotaResponse{
		Message: "ok",
		Code:    uint32(v1.StatusCode_STATUS_CODE_BAD_REQUEST),
	}
	return resp, nil
}
