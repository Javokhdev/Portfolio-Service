package postgres

import (
	pb "github.com/Javokhdev/Portfolio-Service/genprotos"
)

type InitRoot interface {
	Skill() Skill
	Experience() Experience
	Education()	Education
	Project() Project

}
type Skill interface {
	CreateSkill(Skill *pb.Skill) (*pb.Void, error)
	GetByIdSkill(id *pb.ById) (*pb.Skill, error)
	GetAllSkill(_ *pb.Skill) (*pb.GetAllSkills, error)
	UpdateSkill(Skill *pb.Skill) (*pb.Void, error)
	DeleteSkill(id *pb.ById) (*pb.Void, error)
}

type Experience interface {
	CreateExperience(Experience *pb.Experience) (*pb.Void, error)
	GetByIdExperience(id *pb.ById) (*pb.Experience, error)
	GetAllExperience(_ *pb.Experience) (*pb.GetAllExperiences, error)
	UpdateExperience(Experience *pb.Experience) (*pb.Void, error)
	DeleteExperience(id *pb.ById) (*pb.Void, error)
}

type Education interface {
	CreateEducation(Education *pb.Education) (*pb.Void, error)
	GetByIdEducation(id *pb.ById) (*pb.Education, error)
	GetAllEducation(_ *pb.Education) (*pb.GetAllEducations, error)
	UpdateEducation(Education *pb.Education) (*pb.Void, error)
	DeleteEducation(id *pb.ById) (*pb.Void, error)
}

type Project interface {
	CreateProject(Project *pb.Project) (*pb.Void, error)
	GetByIdProject(id *pb.ById) (*pb.Project, error)
	GetAllProject(_ *pb.Project) (*pb.GetAllProjects, error)
	UpdateProject(Project *pb.Project) (*pb.Void, error)
	DeleteProject(id *pb.ById) (*pb.Void, error)
}



