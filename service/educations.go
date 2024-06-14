package service

import (
	"context"
	"log"

	pb "github.com/Javokhdev/Portfolio-Service/genprotos"
	s "github.com/Javokhdev/Portfolio-Service/storage"
)

type EducationService struct {
	stg s.InitRoot
	pb.UnimplementedEducationServiceServer
}

func NewEducationService(stg s.InitRoot) *EducationService {
	return &EducationService{stg: stg}
}

func (c *EducationService) CreateEducation(ctx context.Context, Education *pb.Education) (*pb.Void, error) {
	pb, err := c.stg.Education().CreateEducation(Education)
	if err != nil {
		log.Print(err)
	}
	return pb, err
}

func (c *EducationService) GetAllEducation(ctx context.Context, pb *pb.Education) (*pb.GetAllEducations, error) {
	Educations, err := c.stg.Education().GetAllEducation(pb)
	if err != nil {
		log.Print(err)
	}

	return Educations, err
}

func (c *EducationService) GetByIdEducation(ctx context.Context, id *pb.ById) (*pb.Education, error) {
	prod, err := c.stg.Education().GetByIdEducation(id)
	if err != nil {
		log.Print(err)
	}

	return prod, err
}

func (c *EducationService) UpdateEducation(ctx context.Context, Education *pb.Education) (*pb.Void, error) {
	pb, err := c.stg.Education().UpdateEducation(Education)
	if err != nil {
		log.Print(err)
	}

	return pb, err
}

func (c *EducationService) DeleteEducation(ctx context.Context, id *pb.ById) (*pb.Void, error) {
	pb, err := c.stg.Education().DeleteEducation(id)
	if err != nil {
		log.Print(err)
	}

	return pb, err
}
