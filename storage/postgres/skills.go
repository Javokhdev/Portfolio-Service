package postgres

import (
	"database/sql"
	"fmt"
	"time"

	pb "github.com/Javokhdev/Portfolio-Service/genprotos"

	"github.com/google/uuid"
)

type SkillsStorage struct {
	db *sql.DB
}

func NewSkillsStorage(db *sql.DB) *SkillsStorage {
	return &SkillsStorage{db: db}
}

func (p *SkillsStorage) CreateSkill(skill *pb.Skill) (*pb.Void, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO skills (id, user_id, name, level)
		VALUES ($1, $2, $3, $4)
	`
	_, err := p.db.Exec(query, id, skill.UserId, skill.Name, skill.Level)
	return nil, err
}

func (p *SkillsStorage) GetByIdSkill(id *pb.ById) (*pb.Skill, error) {
	query := `
			SELECT user_id, name, level from skills 
			where id =$1 and delated_at=0 
		`
	row := p.db.QueryRow(query, id.Id)

	var Skill pb.Skill

	err := row.Scan(&Skill.UserId, &Skill.Name, &Skill.Level)
	if err != nil {
		return nil, err
	}

	return &Skill, nil
}

func (p *SkillsStorage) GetAllSkill(rest *pb.Skill) (*pb.GetAllSkills, error) {
	Skill := &pb.GetAllSkills{}
	var query string
	query = ` SELECT user_id, name, level from skills 
			where delated_at=0`
	var arr []interface{}
	count := 1
	if len(rest.Name) > 0 {
		query += fmt.Sprintf(" and name=$%d", count)
		count++
		arr = append(arr, rest.Name)
	}
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
		var r pb.Skill
		err = row.Scan(&r.UserId, &r.Name, &r.Level)
		if err != nil {
			return nil, err
		}
		Skill.Skills = append(Skill.Skills, &r)
	}
	return Skill, nil
}

func (p *SkillsStorage) UpdateSkill(Skill *pb.Skill) (*pb.Void, error) {
	query := `
		UPDATE skills
		SET user_id = $1, name = $2, level = $3
		WHERE id = $4
	`
	_, err := p.db.Exec(query, Skill.UserId, Skill.Name, Skill.Level, Skill.Id)
	return nil, err
}

func (p *SkillsStorage) DeleteSkill(id *pb.ById) (*pb.Void, error) {
	query := `
		update from skills set delated_at=$1
		where id = $2
	`
	_, err := p.db.Exec(query, time.Now().Unix(), id.Id)
	return nil, err
}
