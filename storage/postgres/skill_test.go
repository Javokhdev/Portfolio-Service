package postgres_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Javokhdev/Portfolio-Service/genprotos"
	"github.com/Javokhdev/Portfolio-Service/storage/postgres"
	"github.com/stretchr/testify/assert"
)

func TestCreateSkill(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	storage := postgres.NewSkillsStorage(db)

	mock.ExpectExec("INSERT INTO skills").WithArgs(sqlmock.AnyArg(), "user1", "skill1", "level1").WillReturnResult(sqlmock.NewResult(1, 1))

	skill := &genprotos.Skill{
		UserId: "user1",
		Name:   "skill1",
		Level:  "level1",
	}

	_, err = storage.CreateSkill(skill)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetByIdSkill(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	storage := postgres.NewSkillsStorage(db)

	rows := sqlmock.NewRows([]string{"user_id", "name", "level"}).
		AddRow("user1", "skill1", "level1")

	mock.ExpectQuery("SELECT user_id, name, level from skills where id =").
		WithArgs("1").
		WillReturnRows(rows)

	id := &genprotos.ById{Id: "1"}
	skill, err := storage.GetByIdSkill(id)
	assert.NoError(t, err)
	assert.NotNil(t, skill)
	assert.Equal(t, "user1", skill.UserId)
	assert.Equal(t, "skill1", skill.Name)
	assert.Equal(t, "level1", skill.Level)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAllSkill(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	storage := postgres.NewSkillsStorage(db)

	rows := sqlmock.NewRows([]string{"user_id", "name", "level"}).
		AddRow("user1", "skill1", "level1").
		AddRow("user2", "skill2", "level2")

	mock.ExpectQuery("SELECT user_id, name, level from skills where deleted_at=0").
		WillReturnRows(rows)

	skill, err := storage.GetAllSkill(&genprotos.Skill{})
	assert.NoError(t, err)
	assert.NotNil(t, skill)
	assert.Len(t, skill.Skills, 2)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateSkill(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	storage := postgres.NewSkillsStorage(db)

	mock.ExpectExec("UPDATE skills SET user_id =").WithArgs("user1", "skill1", "level1", "1").WillReturnResult(sqlmock.NewResult(1, 1))

	skill := &genprotos.Skill{
		Id:     "1",
		UserId: "user1",
		Name:   "skill1",
		Level:  "level1",
	}

	_, err = storage.UpdateSkill(skill)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteSkill(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	storage := postgres.NewSkillsStorage(db)

	mock.ExpectExec("update skills set deleted_at=").WithArgs(sqlmock.AnyArg(), "1").WillReturnResult(sqlmock.NewResult(1, 1))

	id := &genprotos.ById{Id: "1"}

	_, err = storage.DeleteSkill(id)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
