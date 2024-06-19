package service

import (
	"context"
	"log"

	pb "github.com/Javokhdev/Portfolio-Service/genprotos"
	s "github.com/Javokhdev/Portfolio-Service/storage"
)

type SkillService struct {
	stg s.InitRoot
	pb.UnimplementedSkillServiceServer
}

func NewSkillService(stg s.InitRoot) *SkillService {
	return &SkillService{stg: stg}
}

func (c *SkillService) CreateSkill(ctx context.Context, Skill *pb.Skill) (*pb.Void, error) {
	pb, err := c.stg.Skill().CreateSkill(Skill)
	if err != nil {
		log.Print(err)
	}
	return pb, err
}

func (c *SkillService) GetAllSkill(ctx context.Context, pb *pb.Skill) (*pb.GetAllSkills, error) {
	Skills, err := c.stg.Skill().GetAllSkill(pb)
	if err != nil {
		log.Print(err)
	}

	return Skills, err
}

func (c *SkillService) GetByIdSkill(ctx context.Context, id *pb.ById) (*pb.Skill, error) {
	prod, err := c.stg.Skill().GetByIdSkill(id)
	if err != nil {
		log.Print(err)
	}

	return prod, err
}

func (c *SkillService) UpdateSkill(ctx context.Context, Skill *pb.Skill) (*pb.Void, error) {
	
	pb, err := c.stg.Skill().UpdateSkill(Skill)
	if err != nil {
		log.Print(err)
	}

	return pb, err
}

func (c *SkillService) DeleteSkill(ctx context.Context, id *pb.ById) (*pb.Void, error) {
	pb, err := c.stg.Skill().DeleteSkill(id)
	if err != nil {
		log.Print(err)
	}

	return pb, err
}
