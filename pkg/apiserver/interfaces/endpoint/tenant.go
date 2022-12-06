package endpoint

import (
	"context"
	"errors"

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
	r := &v1.CreateTenantResponse{
		Succeed: false,
		Message: "xxxx",
	}
	return r, nil
}

func (s *tenantImpl) DeleteTenant(ctx context.Context, req *v1.DeleteTenantRequest) (*v1.DeleteTenantResponse, error) {
	return nil, errors.New("not implemented")
}
