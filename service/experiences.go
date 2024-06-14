package service

import (
	"context"
	"log"

	pb "github.com/Javokhdev/Portfolio-Service/genprotos"
	s "github.com/Javokhdev/Portfolio-Service/storage"
)

type ExperienceService struct {
	stg s.InitRoot
	pb.UnimplementedExperienceServiceServer
}

func NewExperienceService(stg s.InitRoot) *ExperienceService {
	return &ExperienceService{stg: stg}
}

func (c *ExperienceService) CreateExperience(ctx context.Context, Experience *pb.Experience) (*pb.Void, error) {
	pb, err := c.stg.Experience().CreateExperience(Experience)
	if err != nil {
		log.Print(err)
	}
	return pb, err
}

func (c *ExperienceService) GetAllExperience(ctx context.Context, pb *pb.Experience) (*pb.GetAllExperiences, error) {
	Experiences, err := c.stg.Experience().GetAllExperience(pb)
	if err != nil {
		log.Print(err)
	}

	return Experiences, err
}

func (c *ExperienceService) GetByIdExperience(ctx context.Context, id *pb.ById) (*pb.Experience, error) {
	prod, err := c.stg.Experience().GetByIdExperience(id)
	if err != nil {
		log.Print(err)
	}

	return prod, err
}

func (c *ExperienceService) UpdateExperience(ctx context.Context, Experience *pb.Experience) (*pb.Void, error) {
	pb, err := c.stg.Experience().UpdateExperience(Experience)
	if err != nil {
		log.Print(err)
	}

	return pb, err
}

func (c *ExperienceService) DeleteExperience(ctx context.Context, id *pb.ById) (*pb.Void, error) {
	pb, err := c.stg.Experience().DeleteExperience(id)
	if err != nil {
		log.Print(err)
	}

	return pb, err
}
