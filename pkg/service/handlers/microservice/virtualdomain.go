// Copyright 2022 The kubegems.io Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package microservice

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"kubegems.io/kubegems/pkg/i18n"
	"kubegems.io/kubegems/pkg/service/handlers"
	"kubegems.io/kubegems/pkg/service/handlers/base"
	microservice "kubegems.io/kubegems/pkg/service/handlers/microservice/options"
	"kubegems.io/kubegems/pkg/service/models"
)

type VirtualDomainHandler struct {
	base.BaseHandler
	MicroserviceOptions *microservice.MicroserviceOptions
}

// ListVirtualDomain 列表 VirtualDomain
// @Tags        VirtualDomain
// @Summary     VirtualDomain列表
// @Description VirtualDomain列表
// @Accept      json
// @Produce     json
// @Param       VirtualDomainName query    string                                                                       false "VirtualDomainName"
// @Param       VirtualDomainID   query    string                                                                       false "VirtualDomainID"
// @Param       page              query    int                                                                          false "page"
// @Param       size              query    int                                                                          false "page"
// @Param       search            query    string                                                                       false "search in (VirtualDomainName)"
// @Success     200               {object} handlers.ResponseStruct{Data=handlers.PageData{List=[]models.VirtualDomain}} "VirtualDomain"
// @Router      /v1/virtualdomain [get]
// @Security    JWT
func (h *VirtualDomainHandler) ListVirtualDomain(c *gin.Context) {
	var list []models.VirtualDomain
	query, err := handlers.GetQuery(c, nil)
	if err != nil {
		handlers.NotOK(c, err)
		return
	}
	cond := &handlers.PageQueryCond{
		Model:        "VirtualDomain",
		SearchFields: []string{"VirtualDomainName"},
		// Join:         handlers.Args("left join virtual_spaces on virtual_spaces.virtual_domain_id = virtual_domains.id"),
		// Select:       handlers.Args("virtual_domains.*, if(virtual_spaces.virtual_domain_id is null, false, true) as is_using"),
	}
	total, page, size, err := query.PageList(h.GetDB(), cond, &list)
	if err != nil {
		handlers.NotOK(c, err)
		return
	}
	handlers.OK(c, handlers.Page(total, list, page, size))
}

// GetVirtualDomain VirtualDomain详情
// @Tags        VirtualDomain
// @Summary     VirtualDomain详情
// @Description get VirtualDomain详情
// @Accept      json
// @Produce     json
// @Param       virtualdomain_id path     uint                                               true "virtualdomain_id"
// @Success     200              {object} handlers.ResponseStruct{Data=models.VirtualDomain} "VirtualDomain"
// @Router      /v1/virtualdomain/{virtualdomain_id} [get]
// @Security    JWT
func (h *VirtualDomainHandler) GetVirtualDomain(c *gin.Context) {
	// get vd
	vd := models.VirtualDomain{}
	if err := h.GetDB().First(&vd, c.Param("virtualdomain_id")).Error; err != nil {
		handlers.NotOK(c, err)
		return
	}
	handlers.OK(c, vd)
}

// PostVirtualDomain 创建VirtualDomain
// @Tags        VirtualDomain
// @Summary     创建VirtualDomain
// @Description 创建VirtualDomain
// @Accept      json
// @Produce     json
// @Param       param body     models.VirtualDomain                               true "表单"
// @Success     200   {object} handlers.ResponseStruct{Data=models.VirtualDomain} "VirtualDomain"
// @Router      /v1/virtualdomain [post]
// @Security    JWT
func (h *VirtualDomainHandler) PostVirtualDomain(c *gin.Context) {
	var vd models.VirtualDomain
	if err := c.BindJSON(&vd); err != nil {
		handlers.NotOK(c, err)
		return
	}

	action := i18n.Sprintf(context.TODO(), "create")
	module := i18n.Sprintf(context.TODO(), "virtual domain")
	h.SetAuditData(c, action, module, vd.VirtualDomainName)
	h.SetExtraAuditData(c, models.ResVirtualDomain, vd.ID)

	u, _ := h.GetContextUser(c)
	vd.CreatedBy = u.GetUsername()
	vd.IsActive = true

	if err := h.GetDB().Save(&vd).Error; err != nil {
		handlers.NotOK(c, err)
		return
	}

	handlers.OK(c, vd)
}

// PutVirtualDomain 更新VirtualDomain
// @Tags        VirtualDomain
// @Summary     更新VirtualDomain
// @Description 更新VirtualDomain
// @Accept      json
// @Produce     json
// @Param       virtualdomain_id path     uint                                               true "virtualdomain_id"
// @Param       param            body     models.VirtualDomain                               true "表单"
// @Success     200              {object} handlers.ResponseStruct{Data=models.VirtualDomain} "VirtualDomain"
// @Router      /v1/virtualdomain/{virtualdomain_id} [put]
// @Security    JWT
func (h *VirtualDomainHandler) PutVirtualDomain(c *gin.Context) {
	var obj models.VirtualDomain
	if err := h.GetDB().First(&obj, c.Param("virtualdomain_id")).Error; err != nil {
		handlers.NotOK(c, err)
		return
	}

	action := i18n.Sprintf(context.TODO(), "update")
	module := i18n.Sprintf(context.TODO(), "virtual domain")
	h.SetAuditData(c, action, module, obj.VirtualDomainName)
	h.SetExtraAuditData(c, models.ResTenant, obj.ID)

	if err := c.BindJSON(&obj); err != nil {
		handlers.NotOK(c, err)
		return
	}
	if strconv.Itoa(int(obj.ID)) != c.Param("virtualdomain_id") {
		handlers.NotOK(c, i18n.Errorf(c, "URL parameter mismatched with body"))
		return
	}
	if err := h.GetDB().Save(&obj).Error; err != nil {
		handlers.NotOK(c, err)
		return
	}
	handlers.OK(c, obj)
}

// DeleteVirtualDomain 删除 VirtualDomain
// @Tags        VirtualDomain
// @Summary     删除 VirtualDomain
// @Description 删除 VirtualDomain
// @Accept      json
// @Produce     json
// @Param       virtualdomain_id path     uint                    true "virtualdomain_id"
// @Success     200              {object} handlers.ResponseStruct "resp"
// @Router      /v1/virtualdomain/{virtualdomain_id} [delete]
// @Security    JWT
func (h *VirtualDomainHandler) DeleteVirtualDomain(c *gin.Context) {
	// get vd
	vd := models.VirtualDomain{}
	if err := h.GetDB().First(&vd, c.Param("virtualdomain_id")).Error; err != nil {
		handlers.NotOK(c, err)
		return
	}

	action := i18n.Sprintf(context.TODO(), "delete")
	module := i18n.Sprintf(context.TODO(), "virtual domain")
	h.SetAuditData(c, action, module, vd.VirtualDomainName)
	h.SetExtraAuditData(c, models.ResVirtualDomain, vd.ID)

	if err := h.GetDB().Delete(&vd).Error; err != nil {
		handlers.NotOK(c, err)
		return
	}

	handlers.OK(c, "")
}

// InjectVirtualDomain 为 service 设置 serviceentry
// TODO:
func (h *VirtualDomainHandler) InjectVirtualDomain(c *gin.Context) {}

// UnInjectVirtualDomain 为 service 取消设置 serviceentry
// TODO:
func (h *VirtualDomainHandler) UnInjectVirtualDomain(c *gin.Context) {}

func (h *VirtualDomainHandler) RegistRouter(rg *gin.RouterGroup) {
	rg.GET("/virtualdomain", h.ListVirtualDomain)
	rg.POST("/virtualdomain", h.PostVirtualDomain)
	rg.GET("/virtualdomain/:virtualdomain_id", h.GetVirtualDomain)
	rg.PUT("/virtualdomain/:virtualdomain_id", h.PutVirtualDomain)
	rg.DELETE("/virtualdomain/:virtualdomain_id", h.DeleteVirtualDomain)
	rg.PUT("/virtualdomain/:virtualdomain_id/actions/inject", h.InjectVirtualDomain)
	rg.PUT("/virtualdomain/:virtualdomain_id/actions/uninject", h.InjectVirtualDomain)
}
