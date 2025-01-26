package models

import (
	"database/sql"
	"sqli/initializers"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestSecureLogin_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	initializers.DB = db

	mock.ExpectPrepare("SELECT \\* FROM credentials where username=\\? AND password=\\?").
		ExpectQuery().
		WithArgs("testuser", "password123").
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password", "isadmin"}).
			AddRow(1, "testuser", "password123", true))

	user, err := SecureLogin("testuser", "password123")

	assert.NoError(t, err)
	assert.Equal(t, "testuser", user.Username)
	assert.Equal(t, true, user.IsAdmin)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSecureLogin_InvalidCredentials(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	initializers.DB = db

	mock.ExpectPrepare("SELECT \\* FROM credentials where username=\\? AND password=\\?").
		ExpectQuery().
		WithArgs("wronguser", "wrongpass").
		WillReturnError(sql.ErrNoRows)

	user, err := SecureLogin("wronguser", "wrongpass")

	assert.Error(t, err)
	assert.Equal(t, InvalidCredentialsErr, err)
	assert.Equal(t, User{}, user)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSecureChangePassword_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	initializers.DB = db

	mock.ExpectPrepare("UPDATE credentials SET password=\\? WHERE username=\\? AND password=\\?").
		ExpectExec().
		WithArgs("newpass123", "testuser", "oldpass123").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = SecureChangePassword("testuser", "oldpass123", "newpass123")

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSecureChangePassword_InvalidCredentials(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	initializers.DB = db

	mock.ExpectPrepare("UPDATE credentials SET password=\\? WHERE username=\\? AND password=\\?").
		ExpectExec().
		WithArgs("newpass123", "testuser", "wrongpass").
		WillReturnResult(sqlmock.NewResult(0, 0))

	err = SecureChangePassword("testuser", "wrongpass", "newpass123")

	assert.Error(t, err)
	assert.Equal(t, InvalidCredentialsErr, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSecureForgotPassword_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	initializers.DB = db

	mock.ExpectPrepare("SELECT \\* FROM credentials where username=\\?").
		ExpectQuery().
		WithArgs("testuser").
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password", "isadmin"}).
			AddRow(1, "testuser", "password123", false))

	err = SecureForgotPassword("testuser")

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSecureForgotPassword_UserNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	initializers.DB = db

	mock.ExpectPrepare("SELECT \\* FROM credentials where username=\\?").
		ExpectQuery().
		WithArgs("nonexistent").
		WillReturnError(sql.ErrNoRows)

	err = SecureForgotPassword("nonexistent")

	assert.Error(t, err)
	assert.Equal(t, InvalidCredentialsErr, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
