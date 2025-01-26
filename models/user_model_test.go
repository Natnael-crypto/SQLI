package models

import (
	"sqli/initializers"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestVulnLogin(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	initializers.DB = db

	rows := sqlmock.NewRows([]string{"id", "username", "password", "is_admin"}).
		AddRow(1, "testuser", "password123", false)

	mock.ExpectQuery("SELECT \\* FROM credentials WHERE username='testuser' AND PASSWORD='password123'").
		WillReturnRows(rows)

	user, err := VulnLogin("testuser", "password123")

	assert.NoError(t, err)
	assert.Equal(t, "testuser", user.Username)
	assert.Equal(t, "password123", user.Password)
	assert.False(t, user.IsAdmin)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestVulnLogin_ErrorInQuery(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	initializers.DB = db

	mock.ExpectQuery("SELECT \\* FROM credentials WHERE username='testuser' AND PASSWORD='password123'").
		WillReturnError(InvalidCredentialsErr)

	user, err := VulnLogin("testuser", "password123")

	assert.Error(t, err)
	assert.Equal(t, User{}, user)
	assert.Equal(t, InvalidCredentialsErr, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSecureLogin(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	initializers.DB = db

	rows := sqlmock.NewRows([]string{"id", "username", "password", "is_admin"}).
		AddRow(1, "testuser", "password123", false)

	mock.ExpectPrepare("SELECT \\* FROM credentials where username=\\? AND password=\\?").
		ExpectQuery().
		WithArgs("testuser", "password123").
		WillReturnRows(rows)

	user, err := SecureLogin("testuser", "password123")

	assert.NoError(t, err)
	assert.Equal(t, "testuser", user.Username)
	assert.Equal(t, "password123", user.Password)
	assert.False(t, user.IsAdmin)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSecureLogin_ErrorInPrepare(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	initializers.DB = db

	mock.ExpectPrepare("SELECT \\* FROM credentials where username=\\? AND password=\\?").
		WillReturnError(SomethingWentWrongErr)

	user, err := SecureLogin("testuser", "password123")

	assert.Error(t, err)
	assert.Equal(t, User{}, user)
	assert.Equal(t, SomethingWentWrongErr, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSecureLogin_ErrorInQuery(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	initializers.DB = db

	mock.ExpectPrepare("SELECT \\* FROM credentials where username=\\? AND password=\\?").
		ExpectQuery().
		WithArgs("testuser", "password123").
		WillReturnError(InvalidCredentialsErr)

	user, err := SecureLogin("testuser", "password123")

	assert.Error(t, err)
	assert.Equal(t, User{}, user)
	assert.Equal(t, InvalidCredentialsErr, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestVulnChangePassword(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	initializers.DB = db

	mock.ExpectExec("UPDATE credentials SET password='newpassword' WHERE username='testuser' AND password='oldpassword'").
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = VulnChangePassword("testuser", "oldpassword", "newpassword")

	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestVulnChangePassword_ErrorInExec(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	initializers.DB = db

	mock.ExpectExec("UPDATE credentials SET password='newpassword' WHERE username='testuser' AND password='oldpassword'").
		WillReturnError(InvalidCredentialsErr)

	err = VulnChangePassword("testuser", "oldpassword", "newpassword")

	assert.Error(t, err)
	assert.Equal(t, InvalidCredentialsErr, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSecureChangePassword(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	initializers.DB = db

	mock.ExpectPrepare("UPDATE credentials SET password=\\? WHERE username=\\? AND password=\\?").
		ExpectExec().
		WithArgs("newpassword", "testuser", "oldpassword").
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = SecureChangePassword("testuser", "oldpassword", "newpassword")

	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSecureChangePassword_ErrorInPrepare(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	initializers.DB = db

	mock.ExpectPrepare("UPDATE credentials SET password=\\? WHERE username=\\? AND password=\\?").
		WillReturnError(SomethingWentWrongErr)

	err = SecureChangePassword("testuser", "oldpassword", "newpassword")

	assert.Error(t, err)
	assert.Equal(t, SomethingWentWrongErr, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSecureChangePassword_ErrorInExec(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	initializers.DB = db

	mock.ExpectPrepare("UPDATE credentials SET password=\\? WHERE username=\\? AND password=\\?").
		ExpectExec().
		WithArgs("newpassword", "testuser", "oldpassword").
		WillReturnError(InvalidCredentialsErr)

	err = SecureChangePassword("testuser", "oldpassword", "newpassword")

	assert.Error(t, err)
	assert.Equal(t, InvalidCredentialsErr, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestVulnForgotPassword(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	initializers.DB = db

	rows := sqlmock.NewRows([]string{"id", "username", "password", "is_admin"}).
		AddRow(1, "testuser", "password123", false)

	mock.ExpectQuery("SELECT \\* FROM credentials where username='testuser'").
		WillReturnRows(rows)

	err = VulnForgotPassword("testuser")

	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestVulnForgotPassword_ErrorInQuery(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	initializers.DB = db

	mock.ExpectQuery("SELECT \\* FROM credentials where username='testuser'").
		WillReturnError(InvalidCredentialsErr)

	err = VulnForgotPassword("testuser")

	assert.Error(t, err)
	assert.Equal(t, InvalidCredentialsErr, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSecureForgotPassword(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	initializers.DB = db

	rows := sqlmock.NewRows([]string{"id", "username", "password", "is_admin"}).
		AddRow(1, "testuser", "password123", false)

	mock.ExpectPrepare("SELECT \\* FROM credentials where username=\\?").
		ExpectQuery().
		WithArgs("testuser").
		WillReturnRows(rows)

	err = SecureForgotPassword("testuser")

	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSecureForgotPassword_ErrorInPrepare(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	initializers.DB = db

	mock.ExpectPrepare("SELECT \\* FROM credentials where username=\\?").
		WillReturnError(SomethingWentWrongErr)

	err = SecureForgotPassword("testuser")

	assert.Error(t, err)
	assert.Equal(t, SomethingWentWrongErr, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSecureForgotPassword_ErrorInQuery(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	initializers.DB = db

	mock.ExpectPrepare("SELECT \\* FROM credentials where username=\\?").
		ExpectQuery().
		WithArgs("testuser").
		WillReturnError(InvalidCredentialsErr)

	err = SecureForgotPassword("testuser")

	assert.Error(t, err)
	assert.Equal(t, InvalidCredentialsErr, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}