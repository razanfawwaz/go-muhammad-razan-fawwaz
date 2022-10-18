package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http/httptest"
	"part3-orm/config"
	"part3-orm/models"
	"regexp"
	"testing"
)

func TestGetUserController(t *testing.T) {
	t.Log()

	dbGormTest, mock, err := sqlmock.New()

	assert.NoError(t, err)

	dbGorm, _ := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      dbGormTest,
	}))

	config.DB = dbGorm

	row := sqlmock.NewRows([]string{
		"name", "email"}).AddRow("Razan Fawwaz", "razan@mail.com")

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`id` = ? AND `users`.`deleted_at` IS NULL")).
		WithArgs(1).
		WillReturnRows(row)
	mock.ExpectCommit()

	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Body               models.User
		HasReturnBody      bool
		ExpectedBody       models.User
	}{
		{
			"succes",
			200,
			"GET",
			models.User{
				Name: "Razan",
			},
			true,
			models.User{
				Name: "Razan Fawwaz",
			},
		},
	}

	for _, v := range testCase {
		t.Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")

			err := GetUserController(ctx)
			assert.NoError(t, err)

			assert.Equal(t, v.ExpectedStatusCode, w.Result().StatusCode)
			t.Log(w.Result())

			if v.HasReturnBody {
				var resp map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&resp)
				assert.NoError(t, err)
				assert.Equal(t, v.ExpectedBody.Name, resp["user"])
				t.Log()
			}
		})
	}
}

func TestGetAllUserController(t *testing.T) {
	t.Log()

	dbGormTest, mock, err := sqlmock.New()

	assert.NoError(t, err)

	dbGorm, _ := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      dbGormTest,
	}))

	config.DB = dbGorm

	row := sqlmock.NewRows([]string{
		"name", "email"}).
		AddRow("Razan Fawwaz", "razan@mail.com").
		AddRow("Razan Muhammad", "razan@ymail.com")

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` ")).
		WillReturnRows(row)
	mock.ExpectCommit()

	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Body               models.User
		HasReturnBody      bool
		ExpectedBody       []interface{}
	}{
		{
			"succes",
			200,
			"GET",
			models.User{
				Name: "Razan",
			},
			true, []interface{}{"Razan Fawwaz", "Razan Muhammad"},
		},
	}

	for _, v := range testCase {
		t.Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/users")

			err := GetUsersControllers(ctx)
			assert.NoError(t, err)

			assert.Equal(t, v.ExpectedStatusCode, w.Result().StatusCode)
			t.Log(w.Result())

			if v.HasReturnBody {
				var resp map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&resp)
				assert.NoError(t, err)
				t.Log(resp)
				assert.Equal(t, v.ExpectedBody, resp["user"])
				t.Log()
			}
		})
	}
}

func TestCreateUserController(t *testing.T) {
	t.Log()

	dbGormTest, mock, err := sqlmock.New()

	assert.NoError(t, err)

	dbGorm, _ := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      dbGormTest,
	}))

	config.DB = dbGorm

	// mock sql insert
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`name`,`email`,`password`,`token`) VALUES (?,?,?,?,?,?,?)")).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, "Razan Fawwaz", "razan@mail.com", "123456", "").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		// body make json
		Body          models.User
		HasReturnBody bool
		ExpectedBody  string
	}{
		{
			"succes",
			200,
			"POST",
			models.User{
				Name:     "Razan Fawwaz",
				Email:    "razan@mail.com",
				Password: "123456",
				Token:    "",
			},
			true,
			"success create user",
		},
	}

	for _, v := range testCase {
		t.Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/users")
			ctx.Request().Header.Set("content-type", "application/json")
			err := CreateUserController(ctx)
			assert.NoError(t, err)

			assert.Equal(t, v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&resp)
				t.Log(resp)
				assert.NoError(t, err)
				assert.Equal(t, v.ExpectedBody, resp["message"])
				t.Log()
			}
		})
	}
}

// TestUpdateUserController func

func TestUpdateUserController(t *testing.T) {
	t.Log()

	dbGormTest, mock, err := sqlmock.New()

	assert.NoError(t, err)

	dbGorm, _ := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      dbGormTest,
	}))

	config.DB = dbGorm

	// mock sql insert
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("UPDATE `users` SET `updated_at`=?,`name`=? WHERE id = ? AND `users`.`deleted_at` IS NULL")).
		WithArgs(sqlmock.AnyArg(), "Razan Fawwaz", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		// body make json
		Body          models.User
		HasReturnBody bool
		ExpectedBody  string
	}{
		{
			"succes",
			200,
			"PUT",
			models.User{
				Name: "Razan Fawwaz",
			},
			true,
			"success update user",
		},
	}

	for _, v := range testCase {
		t.Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/users/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")
			ctx.Request().Header.Set("content-type", "application/json")
			err := UpdateUserController(ctx)
			assert.NoError(t, err)

			assert.Equal(t, v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&resp)
				t.Log(resp)
				assert.NoError(t, err)
				assert.Equal(t, v.ExpectedBody, resp["message"])
				t.Log()
			}
		})
	}
}

func TestDeleteUserController(t *testing.T) {
	t.Log()

	dbGormTest, mock, err := sqlmock.New()

	assert.NoError(t, err)

	dbGorm, _ := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      dbGormTest,
	}))

	config.DB = dbGorm

	// mock sql insert
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("UPDATE `users` SET `deleted_at`=? WHERE `users`.`id` = ? AND `users`.`deleted_at` IS NULL")).
		WithArgs(sqlmock.AnyArg(), 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		// body make json
		Body          models.User
		HasReturnBody bool
		ExpectedBody  string
	}{
		{
			"succes",
			200,
			"DELETE",
			models.User{
				Name: "Razan Fawwaz",
			},
			true,
			"success delete user",
		},
	}

	for _, v := range testCase {
		t.Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/users/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")
			ctx.Request().Header.Set("content-type", "application/json")
			err := DeleteUserController(ctx)
			assert.NoError(t, err)

			assert.Equal(t, v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&resp)
				t.Log(resp)
				assert.NoError(t, err)
				assert.Equal(t, v.ExpectedBody, resp["message"])
				t.Log()
			}
		})
	}
}
