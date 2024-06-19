package postgres_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Javokhdev/Portfolio-Service/genprotos"
	"github.com/Javokhdev/Portfolio-Service/storage/postgres"
	"github.com/stretchr/testify/assert"
)

func TestCreateExperience(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	storage := postgres.NewExperienceStorage(db)

	mock.ExpectExec("INSERT INTO experiences").WithArgs(sqlmock.AnyArg(), "user1", "title1", "company1", "description1", "start_date1", "end_date1").WillReturnResult(sqlmock.NewResult(1, 1))

	exp := &genprotos.Experience{
		UserId:      "user1",
		Title:       "title1",
		Company:     "company1",
		Description: "description1",
		StartDate:   "start_date1",
		EndDate:     "end_date1",
	}

	_, err = storage.CreateExperience(exp)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetByIdExperience(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	storage := postgres.NewExperienceStorage(db)

	rows := sqlmock.NewRows([]string{"user_id", "title", "company", "description", "start_date", "end_date"}).
		AddRow("user1", "title1", "company1", "description1", "start_date1", "end_date1")

	mock.ExpectQuery("SELECT user_id, title, company, description, start_date, end_date from experiences where id =").
		WithArgs("1").
		WillReturnRows(rows)

	id := &genprotos.ById{Id: "1"}
	exp, err := storage.GetByIdExperience(id)
	assert.NoError(t, err)
	assert.NotNil(t, exp)
	assert.Equal(t, "user1", exp.UserId)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAllExperience(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	storage := postgres.NewExperienceStorage(db)

	rows := sqlmock.NewRows([]string{"user_id", "title", "company", "description", "start_date", "end_date"}).
		AddRow("user1", "title1", "company1", "description1", "start_date1", "end_date1").
		AddRow("user2", "title2", "company2", "description2", "start_date2", "end_date2")

	mock.ExpectQuery("SELECT user_id, title, company, description, start_date, end_date from experiences where deleted_at=0").
		WillReturnRows(rows)

	exp, err := storage.GetAllExperience(&genprotos.Experience{})
	assert.NoError(t, err)
	assert.NotNil(t, exp)
	assert.Len(t, exp.Experiences, 2)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateExperience(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	storage := postgres.NewExperienceStorage(db)

	mock.ExpectExec("UPDATE experiences SET user_id =").WithArgs("user1", "title1", "company1", "description1", "start_date1", "end_date1", "1").WillReturnResult(sqlmock.NewResult(1, 1))

	exp := &genprotos.Experience{
		Id:          "1",
		UserId:      "user1",
		Title:       "title1",
		Company:     "company1",
		Description: "description1",
		StartDate:   "start_date1",
		EndDate:     "end_date1",
	}

	_, err = storage.UpdateExperience(exp)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteExperience(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	storage := postgres.NewExperienceStorage(db)

	mock.ExpectExec("update experiences set deleted_at=").WithArgs(sqlmock.AnyArg(), "1").WillReturnResult(sqlmock.NewResult(1, 1))

	id := &genprotos.ById{Id: "1"}

	_, err = storage.DeleteExperience(id)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
