package postgres

import (
	"database/sql"
	"fmt"
	"time"

	pb "github.com/Javokhdev/Portfolio-Service/genprotos"

	"github.com/google/uuid"
)

type ProjectsStorage struct {
	db *sql.DB
}

func NewProjectsStorage(db *sql.DB) *ProjectsStorage {
	return &ProjectsStorage{db: db}
}

func (p *ProjectsStorage) CreateProject(project *pb.Project) (*pb.Void, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO skills (id, user_id, title, description, url)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := p.db.Exec(query, id, project.UserId, project.Title, project.Description, project.Url)
	return nil, err
}

func (p *ProjectsStorage) GetByIdProject(id *pb.ById) (*pb.Project, error) {
	query := `
			SELECT user_id, title, description, url from projects 
			where id =$1 and delated_at=0 
		`
	row := p.db.QueryRow(query, id.Id)

	var Project pb.Project

	err := row.Scan(&Project.UserId, &Project.Title, &Project.Description, &Project.Url)
	if err != nil {
		return nil, err
	}

	return &Project, nil
}

func (p *ProjectsStorage) GetAllProject(rest *pb.Project) (*pb.GetAllProjects, error) {
	project := &pb.GetAllProjects{}
	var query string
	query = ` SELECT user_id, title, description, url from projects 
			where delated_at=0`
	var arr []interface{}
	count := 1
	if len(rest.UserId) > 0 {
		query += fmt.Sprintf(" and user_id=$%d", count)
		count++
		arr = append(arr, rest.UserId)
	}

	row, err := p.db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var r pb.Project
		err = row.Scan(&r.UserId, &r.Title, &r.Description, &r.Url)
		if err != nil {
			return nil, err
		}
		project.Projects = append(project.Projects, &r)
	}
	return project, nil
}

func (p *ProjectsStorage) UpdateProject(Project *pb.Project) (*pb.Void, error) {
	query := `
		UPDATE projects
		SET user_id = $1, title = $2, description = $3, url = $4
		WHERE id = $5
	`
	_, err := p.db.Exec(query, Project.UserId, Project.Title, Project.Description, Project.Url, Project.Id)
	return nil, err
}

func (p *ProjectsStorage) DeleteProject(id *pb.ById) (*pb.Void, error) {
	query := `
		update from projects set delated_at=$1
		where id = $2
	`
	_, err := p.db.Exec(query, time.Now().Unix(), id.Id)
	return nil, err
}
