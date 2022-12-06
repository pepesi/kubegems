package endpoint

import (
	"context"

	agg_services "kubegems.io/kubegems/pkg/apiserver/aggregate"
	v1 "kubegems.io/kubegems/pkg/apiserver/interfaces/apis/v1"
)

func NewProjectServer(projectSvc agg_services.TenantProjectService) v1.ProjectServiceServer {
	return &projectImpl{
		projectSvc: projectSvc,
	}
}

type projectImpl struct {
	v1.UnimplementedProjectServiceServer
	projectSvc agg_services.TenantProjectService
}

func (s *projectImpl) CreateProject(ctx context.Context, req *v1.CreateProjectRequest) (*v1.CreateProjectResponse, error) {
	resp := &v1.CreateProjectResponse{}
	project, err := s.projectSvc.AddProject(ctx, req.Tenant, req.Project)
	if err != nil {
		return nil, err
	}
	resp.Code = 0
	resp.Message = "ok"
	resp.Project = &v1.Project{
		Id:    uint32(project.ID),
		Name:  project.Name,
		Intro: project.Remark,
	}
	return resp, nil
}
