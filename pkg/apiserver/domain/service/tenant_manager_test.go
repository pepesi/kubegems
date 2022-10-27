package service

import (
	"context"
	"reflect"
	"testing"

	"kubegems.io/kubegems/pkg/apiserver/domain/model"
	"kubegems.io/kubegems/pkg/apiserver/domain/repository"
	"kubegems.io/kubegems/pkg/apiserver/options"
)

func Test_tenantManager_CreateTenant(t *testing.T) {
	type fields struct {
		repo repository.GenericRepo[*model.Tenant]
	}
	type args struct {
		ctx    context.Context
		tenant *model.Tenant
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Tenant
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mt := &tenantManager{
				repo: tt.fields.repo,
			}
			got, err := mt.CreateTenant(tt.args.ctx, tt.args.tenant)
			if (err != nil) != tt.wantErr {
				t.Errorf("tenantManager.CreateTenant() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tenantManager.CreateTenant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tenantManager_GetTenant(t *testing.T) {
	type fields struct {
		repo repository.GenericRepo[*model.Tenant]
	}
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Tenant
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mt := &tenantManager{
				repo: tt.fields.repo,
			}
			got, err := mt.GetTenant(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("tenantManager.GetTenant() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tenantManager.GetTenant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tenantManager_ListTenant(t *testing.T) {
	type fields struct {
		repo repository.GenericRepo[*model.Tenant]
	}
	type args struct {
		ctx  context.Context
		opts []options.Option
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.Tenant
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mt := &tenantManager{
				repo: tt.fields.repo,
			}
			got, err := mt.ListTenant(tt.args.ctx, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("tenantManager.ListTenant() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tenantManager.ListTenant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tenantManager_DeleteTenant(t *testing.T) {
	type fields struct {
		repo repository.GenericRepo[*model.Tenant]
	}
	type args struct {
		ctx  context.Context
		opts []options.Option
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mt := &tenantManager{
				repo: tt.fields.repo,
			}
			if err := mt.DeleteTenant(tt.args.ctx, tt.args.opts...); (err != nil) != tt.wantErr {
				t.Errorf("tenantManager.DeleteTenant() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_tenantManager_ModifyTenant(t *testing.T) {
	type fields struct {
		repo repository.GenericRepo[*model.Tenant]
	}
	type args struct {
		ctx    context.Context
		name   string
		tenant *model.Tenant
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mt := &tenantManager{
				repo: tt.fields.repo,
			}
			if err := mt.ModifyTenant(tt.args.ctx, tt.args.name, tt.args.tenant); (err != nil) != tt.wantErr {
				t.Errorf("tenantManager.ModifyTenant() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewTenantManager(t *testing.T) {
	type args struct {
		tenantRepo repository.GenericRepo[*model.Tenant]
	}
	tests := []struct {
		name string
		args args
		want *tenantManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTenantManager(tt.args.tenantRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTenantManager() = %v, want %v", got, tt.want)
			}
		})
	}
}
