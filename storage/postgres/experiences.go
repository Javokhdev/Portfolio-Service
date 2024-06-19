package postgres

import (
	"database/sql"
	"fmt"
	"time"

	pb "github.com/Javokhdev/Portfolio-Service/genprotos"

	"github.com/google/uuid"
)

type ExperienceStorage struct {
	db *sql.DB
}

func NewExperienceStorage(db *sql.DB) *ExperienceStorage {
	return &ExperienceStorage{db: db}
}

func (p *ExperienceStorage) CreateExperience(exp *pb.Experience) (*pb.Void, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO experiences (id, user_id, title, company, description, start_date, end_date)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := p.db.Exec(query, id, exp.UserId, exp.Title, exp.Company, exp.Description, exp.StartDate, exp.EndDate)
	return nil, err
}

func (p *ExperienceStorage) GetByIdExperience(id *pb.ById) (*pb.Experience, error) {
	query := `
			SELECT user_id, title, company, description, start_date, end_date from experiences 
			where id =$1 and deleted_at=0 
		`
	row := p.db.QueryRow(query, id.Id)

	var exp pb.Experience

	err := row.Scan(&exp.UserId, &exp.Title, &exp.Company, &exp.Description, &exp.StartDate, &exp.EndDate)
	if err != nil {
		return nil, err
	}

	return &exp, nil
}

func (p *ExperienceStorage) GetAllExperience(rest *pb.Experience) (*pb.GetAllExperiences, error) {
	exp := &pb.GetAllExperiences{}
	var query string
	query = ` SELECT user_id, title, company, description, start_date, end_date from experiences 
			where deleted_at=0`
	var arr []interface{}
	count := 1

	if len(rest.Id) > 0 {
		query += fmt.Sprintf(" and id=$%d", count)
		count++
		arr = append(arr, rest.Id)
	}
	
	if len(rest.UserId) > 0 {
		query += fmt.Sprintf(" and user_id=$%d", count)
		count++
		arr = append(arr, rest.UserId)
	}

	if len(rest.Title) > 0 {
		query += fmt.Sprintf(" and title=$%d", count)
		count++
		arr = append(arr, rest.Title)
	}

	if len(rest.Company) > 0 {
		query += fmt.Sprintf(" and company=$%d", count)
		count++
		arr = append(arr, rest.Company)
	}

	if len(rest.Description) > 0 {
		query += fmt.Sprintf(" and description=$%d", count)
		count++
		arr = append(arr, rest.Description)

	}

	if len(rest.StartDate) > 0 {
		query += fmt.Sprintf(" and start_date=$%d", count)
		count++
		arr = append(arr, rest.StartDate)

	}

	if len(rest.EndDate) > 0 {
		query += fmt.Sprintf(" and end_date=$%d", count)
		count++
		arr = append(arr, rest.EndDate)
	}

	row, err := p.db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var r pb.Experience
		err = row.Scan(&r.UserId, &r.Title, &r.Company, &r.Description, &r.StartDate, &r.EndDate)
		if err != nil {
			return nil, err
		}
		exp.Experiences = append(exp.Experiences, &r)
	}
	return exp, nil
}

func (p *ExperienceStorage) UpdateExperience(exp *pb.Experience) (*pb.Void, error) {
	query := `
		UPDATE experiences
		SET user_id = $1, title = $2, company = $3, description = $4, start_date = $5, end_date = $6
		WHERE id = $7
	`
	_, err := p.db.Exec(query, exp.UserId, exp.Title, exp.Company, exp.Description, exp.StartDate, exp.EndDate, exp.Id)
	return nil, err
}

func (p *ExperienceStorage) DeleteExperience(id *pb.ById) (*pb.Void, error) {
	query := `
		update experiences set deleted_at=$1
		where id = $2
	`
	_, err := p.db.Exec(query, time.Now().Unix(), id.Id)
	return nil, err
}
