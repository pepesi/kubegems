package apis

import "time"

type TenantBase struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Remark    string    `json:"remark"`
	Enabled   bool      `json:"enabled"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type RetrieveTenantResp struct {
	TenantBase
}

type ListTenantResp struct {
	Tenants []TenantBase `json:"tenants"`
}

type CreateUpdateTenantReq struct {
	Name    string `json:"name" validate:"required"`
	Remark  string `json:"remark" validate:"required"`
	Enabled bool   `json:"enabled" validate:"required"`
}

type CreateUpdateTenantResp struct {
	TenantBase
}

type QuotaBase struct {
	ID        uint   `json:"id"`
	ClusterID uint   `json:"clusterId"`
	RelKind   string `json:"relKind"`
	RelName   string `json:"relName"`
	Pod       int    `json:"pod"`
	CPU       int    `json:"cpu"`
	Memory    int    `json:"memory"`
}

type CreateTenantClusterQuotaReq struct {
	Datas map[string]interface{} `json:"datas"`
}

type CreateTenantClusterQuotaResp struct {
	Datas map[string]interface{} `json:"datas"`
}
