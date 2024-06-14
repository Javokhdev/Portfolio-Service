package service

import (
	"context"
	"log"

	pb "github.com/Javokhdev/Portfolio-Service/genprotos"
	s "github.com/Javokhdev/Portfolio-Service/storage"
)

type ProjectService struct {
	stg s.InitRoot
	pb.UnimplementedProjectServiceServer
}

func NewProjectService(stg s.InitRoot) *ProjectService {
	return &ProjectService{stg: stg}
}

func (c *ProjectService) CreateProject(ctx context.Context, Project *pb.Project) (*pb.Void, error) {
	pb, err := c.stg.Project().CreateProject(Project)
	if err != nil {
		log.Print(err)
	}
	return pb, err
}

func (c *ProjectService) GetAllProject(ctx context.Context, pb *pb.Project) (*pb.GetAllProjects, error) {
	project, err := c.stg.Project().GetAllProject(pb)
	if err != nil {
		log.Print(err)
	}

	return project, err
}

func (c *ProjectService) GetByIdProject(ctx context.Context, id *pb.ById) (*pb.Project, error) {
	prod, err := c.stg.Project().GetByIdProject(id)
	if err != nil {
		log.Print(err)
	}

	return prod, err
}

func (c *ProjectService) UpdateProject(ctx context.Context, Project *pb.Project) (*pb.Void, error) {
	pb, err := c.stg.Project().UpdateProject(Project)
	if err != nil {
		log.Print(err)
	}

	return pb, err
}

func (c *ProjectService) DeleteProject(ctx context.Context, id *pb.ById) (*pb.Void, error) {
	pb, err := c.stg.Project().DeleteProject(id)
	if err != nil {
		log.Print(err)
	}

	return pb, err
}
