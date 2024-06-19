package postgres_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Javokhdev/Portfolio-Service/genprotos"
	"github.com/Javokhdev/Portfolio-Service/storage/postgres"
	"github.com/stretchr/testify/assert"
)

func TestCreateProject(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	storage := postgres.NewProjectsStorage(db)

	mock.ExpectExec("INSERT INTO projects").WithArgs(sqlmock.AnyArg(), "user1", "title1", "description1", "url1").WillReturnResult(sqlmock.NewResult(1, 1))

	project := &genprotos.Project{
		UserId:      "user1",
		Title:       "title1",
		Description: "description1",
		Url:         "url1",
	}

	_, err = storage.CreateProject(project)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetByIdProject(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	storage := postgres.NewProjectsStorage(db)

	rows := sqlmock.NewRows([]string{"user_id", "title", "description", "url"}).
		AddRow("user1", "title1", "description1", "url1")

	mock.ExpectQuery("SELECT user_id, title, description, url from projects where id =").
		WithArgs("1").
		WillReturnRows(rows)

	id := &genprotos.ById{Id: "1"}
	project, err := storage.GetByIdProject(id)
	assert.NoError(t, err)
	assert.NotNil(t, project)
	assert.Equal(t, "user1", project.UserId)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAllProject(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	storage := postgres.NewProjectsStorage(db)

	rows := sqlmock.NewRows([]string{"user_id", "title", "description", "url"}).
		AddRow("user1", "title1", "description1", "url1").
		AddRow("user2", "title2", "description2", "url2")

	mock.ExpectQuery("SELECT user_id, title, description, url from projects where deleted_at=0").
		WillReturnRows(rows)

	project, err := storage.GetAllProject(&genprotos.Project{})
	assert.NoError(t, err)
	assert.NotNil(t, project)
	assert.Len(t, project.Projects, 2)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateProject(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	storage := postgres.NewProjectsStorage(db)

	mock.ExpectExec("UPDATE projects SET user_id =").WithArgs("user1", "title1", "description1", "url1", "1").WillReturnResult(sqlmock.NewResult(1, 1))

	project := &genprotos.Project{
		Id:          "1",
		UserId:      "user1",
		Title:       "title1",
		Description: "description1",
		Url:         "url1",
	}

	_, err = storage.UpdateProject(project)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteProject(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	storage := postgres.NewProjectsStorage(db)

	mock.ExpectExec("update projects set deleted_at=").WithArgs(sqlmock.AnyArg(), "1").WillReturnResult(sqlmock.NewResult(1, 1))

	id := &genprotos.ById{Id: "1"}

	_, err = storage.DeleteProject(id)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
