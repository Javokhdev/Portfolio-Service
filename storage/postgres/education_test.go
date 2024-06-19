package postgres_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Javokhdev/Portfolio-Service/genprotos"
	"github.com/Javokhdev/Portfolio-Service/storage/postgres"
	"github.com/stretchr/testify/assert"
)

func TestCreateEducation(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	storage := postgres.NewEducationsStorage(db)

	mock.ExpectExec("INSERT INTO educations").WithArgs(sqlmock.AnyArg(), "user1", "institution1", "degree1", "field1", "start_date1", "end_date1").WillReturnResult(sqlmock.NewResult(1, 1))

	edu := &genprotos.Education{
		UserId:       "user1",
		Institution:  "institution1",
		Degree:       "degree1",
		FieldOfStudy: "field1",
		StartDate:    "start_date1",
		EndDate:      "end_date1",
	}

	_, err = storage.CreateEducation(edu)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetByIdEducation(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	storage := postgres.NewEducationsStorage(db)

	rows := sqlmock.NewRows([]string{"user_id", "institution", "degree", "field_of_study", "start_date", "end_date"}).
		AddRow("user1", "institution1", "degree1", "field1", "start_date1", "end_date1")

	mock.ExpectQuery("SELECT user_id, institution, degree, field_of_study, start_date, end_date from educations where id =").
		WithArgs("1").
		WillReturnRows(rows)

	id := &genprotos.ById{Id: "1"}
	edu, err := storage.GetByIdEducation(id)
	assert.NoError(t, err)
	assert.NotNil(t, edu)
	assert.Equal(t, "user1", edu.UserId)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAllEducation(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	storage := postgres.NewEducationsStorage(db)

	rows := sqlmock.NewRows([]string{"user_id", "institution", "degree", "field_of_study", "start_date", "end_date"}).
		AddRow("user1", "institution1", "degree1", "field1", "start_date1", "end_date1").
		AddRow("user2", "institution2", "degree2", "field2", "start_date2", "end_date2")

	mock.ExpectQuery("SELECT user_id, institution, degree, field_of_study, start_date, end_date from educations where deleted_at=0").
		WillReturnRows(rows)

	edu, err := storage.GetAllEducation(&genprotos.Education{})
	assert.NoError(t, err)
	assert.NotNil(t, edu)
	assert.Len(t, edu.Educations, 2)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateEducation(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	storage := postgres.NewEducationsStorage(db)

	mock.ExpectExec("UPDATE educations SET user_id =").WithArgs("user1", "institution1", "degree1", "field1", "start_date1", "end_date1", "1").WillReturnResult(sqlmock.NewResult(1, 1))

	edu := &genprotos.Education{
		Id:           "1",
		UserId:       "user1",
		Institution:  "institution1",
		Degree:       "degree1",
		FieldOfStudy: "field1",
		StartDate:    "start_date1",
		EndDate:      "end_date1",
	}

	_, err = storage.UpdateEducation(edu)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteEducation(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	storage := postgres.NewEducationsStorage(db)

	mock.ExpectExec("update educations set deleted_at=").WithArgs(sqlmock.AnyArg(), "1").WillReturnResult(sqlmock.NewResult(1, 1))

	id := &genprotos.ById{Id: "1"}

	_, err = storage.DeleteEducation(id)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
