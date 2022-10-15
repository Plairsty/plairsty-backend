package model

import (
	projectPb "awesomeProject/internal/proto/project"
	"context"
)

func (app *Application) AddProject(_ context.Context, req *projectPb.AddProjectRequest) (*projectPb.AddProjectResponse, error) {
	userId := req.GetUserId()
	project := req.GetProject()
	err := app.persistence.Project.Insert(userId, project)
	if err != nil {
		return nil, err
	}
	return &projectPb.AddProjectResponse{
		Status: projectPb.STATUS_STATUS_ADDED,
	}, nil
}

func (app *Application) GetProject(_ context.Context, req *projectPb.GetProjectRequest) (*projectPb.GetProjectResponse, error) {
	userId := req.GetUserId()
	projectId := req.GetProjectId()
	project, err := app.persistence.Project.Get(userId, projectId)
	if err != nil {
		return nil, err
	}
	return &projectPb.GetProjectResponse{
		Project: project,
	}, nil
}

func (app *Application) GetProjects(_ context.Context, req *projectPb.GetProjectsRequest) (*projectPb.GetProjectsResponse, error) {
	userId := req.GetUserId()
	projects, err := app.persistence.Project.GetAll(userId)
	if err != nil {
		return nil, err
	}
	return &projectPb.GetProjectsResponse{
		Projects: projects,
	}, nil
}

func (app *Application) UpdateProject(_ context.Context, req *projectPb.UpdateProjectRequest) (*projectPb.UpdateProjectResponse, error) {
	userId := req.GetUserId()
	project := req.GetProject()
	err := app.persistence.Project.Update(userId, project)
	if err != nil {
		return nil, err
	}
	return &projectPb.UpdateProjectResponse{
		Status: projectPb.STATUS_STATUS_MODIFIED,
	}, nil
}

func (app *Application) DeleteProject(_ context.Context, req *projectPb.DeleteProjectRequest) (*projectPb.DeleteProjectResponse, error) {
	userId := req.GetUserId()
	projectId := req.GetProjectId()
	err := app.persistence.Project.Delete(userId, projectId)
	if err != nil {
		return nil, err
	}
	return &projectPb.DeleteProjectResponse{
		Status: projectPb.STATUS_STATUS_DELETED,
	}, nil
}

func (app *Application) GetProjectsBySemester(_ context.Context, req *projectPb.GetProjectsBySemesterRequest) (*projectPb.GetProjectsBySemesterResponse, error) {
	userId := req.GetUserId()
	semester := req.GetSemester()
	projects, err := app.persistence.Project.GetProjectsBySemester(userId, semester)
	if err != nil {
		return nil, err
	}
	return &projectPb.GetProjectsBySemesterResponse{
		Projects: projects,
	}, nil
}
