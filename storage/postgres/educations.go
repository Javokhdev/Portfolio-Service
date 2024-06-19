package postgres

import (
	"database/sql"
	"fmt"
	"time"

	pb "github.com/Javokhdev/Portfolio-Service/genprotos"

	"github.com/google/uuid"
)

type EducationsStorage struct {
	db *sql.DB
}

func NewEducationsStorage(db *sql.DB) *EducationsStorage {
	return &EducationsStorage{db: db}
}

func (p *EducationsStorage) CreateEducation(edu *pb.Education) (*pb.Void, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO educations (id, user_id, institution, degree, field_of_study, start_date, end_date)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := p.db.Exec(query, id, edu.UserId, edu.Institution, edu.Degree, edu.FieldOfStudy, edu.StartDate, edu.EndDate)
	return nil, err
}

func (p *EducationsStorage) GetByIdEducation(id *pb.ById) (*pb.Education, error) {
	query := `
			SELECT user_id, institution, degree, field_of_study, start_date, end_date from educations 
			where id =$1 and deleted_at=0 
		`
	row := p.db.QueryRow(query, id.Id)

	var edu pb.Education

	err := row.Scan(&edu.UserId, &edu.Institution, &edu.Degree, &edu.FieldOfStudy, &edu.StartDate, &edu.EndDate)
	if err != nil {
		return nil, err
	}

	return &edu, nil
}

func (p *EducationsStorage) GetAllEducation(rest *pb.Education) (*pb.GetAllEducations, error) {
	edu := &pb.GetAllEducations{}
	var query string
	query = ` SELECT user_id, institution, degree, field_of_study, start_date, end_date from educations 
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

	if len(rest.Institution) > 0 {
		query += fmt.Sprintf(" and institution=$%d", count)
		count++
		arr = append(arr, rest.Institution)
	}

	if len(rest.Degree) > 0 {
		query += fmt.Sprintf(" and degree=$%d", count)
		count++
		arr = append(arr, rest.Degree)
	}

	if len(rest.FieldOfStudy) > 0 {
		query += fmt.Sprintf(" and field_of_study=$%d", count)
		count++
		arr = append(arr, rest.FieldOfStudy)
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
		var r pb.Education
		err = row.Scan(&r.UserId, &r.Institution, &r.Degree, &r.FieldOfStudy, &r.StartDate, &r.EndDate)
		if err != nil {
			return nil, err
		}
		edu.Educations = append(edu.Educations, &r)
	}
	return edu, nil
}

func (p *EducationsStorage) UpdateEducation(Edu *pb.Education) (*pb.Void, error) {
	query := `
		UPDATE educations
		SET user_id = $1, institution = $2, degree = $3, field_of_study = $4, start_date = $5, end_date = $6
		WHERE id = $7
	`
	_, err := p.db.Exec(query, Edu.UserId, Edu.Institution, Edu.Degree, Edu.FieldOfStudy, Edu.StartDate, Edu.EndDate, Edu.Id)
	return nil, err
}

func (p *EducationsStorage) DeleteEducation(id *pb.ById) (*pb.Void, error) {
	query := `
		update educations set deleted_at=$1
		where id = $2
	`
	_, err := p.db.Exec(query, time.Now().Unix(), id.Id)
	return nil, err
}
