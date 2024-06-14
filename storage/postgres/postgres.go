package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Javokhdev/Portfolio-Service/config"
	st "github.com/Javokhdev/Portfolio-Service/storage"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db              *sql.DB
	Skills         	st.Skill
	Experiences     st.Experience
	Educations 		st.Education
	Projects        st.Project
}

func NewPostgresStorage() (st.InitRoot, error) {
	config := config.Load()
	con := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.PostgresUser, config.PostgresPassword,
		config.PostgresHost, config.PostgresPort,
		config.PostgresDatabase)
	db, err := sql.Open("postgres", con)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Storage{Db: db, Skills: &SkillsStorage{db}, Experiences: &ExperienceStorage{db}, Educations: &EducationsStorage{db}, Projects: &ProjectsStorage{db}}, nil

}

func (s *Storage) Skill() st.Skill {
	if s.Skills == nil {
		s.Skills = &SkillsStorage{s.Db}
	}
	return s.Skills
}

func (s *Storage) Experience() st.Experience {
	if s.Experiences == nil {
		s.Experiences = &ExperienceStorage{s.Db}
	}
	return s.Experiences
}

func (s *Storage) Education() st.Education {
	if s.Educations == nil {
		s.Educations = &EducationsStorage{s.Db}
	}
	return s.Educations
}

func (s *Storage) Project() st.Project {
	if s.Projects == nil {
		s.Projects = &ProjectsStorage{s.Db}
	}
	return s.Projects
}
