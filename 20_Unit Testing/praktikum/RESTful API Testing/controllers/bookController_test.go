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

func TestGetBookController(t *testing.T) {
	t.Log()

	dbGormTest, mock, err := sqlmock.New()

	assert.NoError(t, err)

	dbGorm, _ := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      dbGormTest,
	}))

	config.DB = dbGorm

	row := sqlmock.NewRows([]string{
		"book_name", "desc", "author"}).AddRow("Pemrograman Dasar", "belajar pemrograman", "razan")

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `books` WHERE `books`.`id` = ?")).
		WithArgs(1).
		WillReturnRows(row)
	mock.ExpectCommit()

	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Body               models.Book
		HasReturnBody      bool
		ExpectedBody       models.Book
	}{
		{
			"succes",
			200,
			"GET",
			models.Book{
				BookName: "Pemrograman Dasar",
			},
			true,
			models.Book{
				BookName: "Pemrograman Dasar",
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

			err := GetBookController(ctx)
			assert.NoError(t, err)

			assert.Equal(t, v.ExpectedStatusCode, w.Result().StatusCode)
			t.Log(w.Result())

			if v.HasReturnBody {
				var resp map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&resp)
				t.Log(resp)
				assert.NoError(t, err)
				assert.Equal(t, v.ExpectedBody.BookName, resp["book"])
				t.Log()
			}
		})
	}
}

func TestGetAllBookController(t *testing.T) {
	t.Log()

	dbGormTest, mock, err := sqlmock.New()

	assert.NoError(t, err)

	dbGorm, _ := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      dbGormTest,
	}))

	config.DB = dbGorm

	row := sqlmock.NewRows([]string{
		"book_name", "desc", "author"}).
		AddRow("Java", "belajar Java", "Razan Fawwaz").
		AddRow("Golang", "belajar Golang", "Razan Fawwaz")

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `books` ")).
		WillReturnRows(row)
	mock.ExpectCommit()

	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Body               models.Book
		HasReturnBody      bool
		ExpectedBody       []interface{}
	}{
		{
			"succes",
			200,
			"GET",
			models.Book{
				BookName: "Golang",
			},
			true, []interface{}{"Java", "Golang"},
		},
	}

	for _, v := range testCase {
		t.Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/books")

			err := GetBooksControllers(ctx)
			assert.NoError(t, err)

			assert.Equal(t, v.ExpectedStatusCode, w.Result().StatusCode)
			t.Log(w.Result())

			if v.HasReturnBody {
				var resp map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&resp)
				assert.NoError(t, err)
				t.Log(resp)
				assert.Equal(t, v.ExpectedBody, resp["book"])
				t.Log()
			}
		})
	}
}

func TestCreateBookController(t *testing.T) {
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
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `books` (`created_at`,`updated_at`,`deleted_at`,`book_name`,`desc`,`author`) VALUES (?,?,?,?,?,?)")).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, "Java Dasar", "Belajar Java", "Razan").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		// body make json
		Body          models.Book
		HasReturnBody bool
		ExpectedBody  string
	}{
		{
			"succes",
			200,
			"POST",
			models.Book{
				BookName: "Java Dasar",
				Desc:     "Belajar Java",
				Author:   "Razan",
			},
			true,
			"success create book",
		},
	}

	for _, v := range testCase {
		t.Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/books")
			ctx.Request().Header.Set("content-type", "application/json")
			err := CreateBookController(ctx)
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

func TestUpdateBookController(t *testing.T) {
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
	mock.ExpectExec(regexp.QuoteMeta("UPDATE `books` SET `updated_at`=?,`book_name`=? WHERE id = ?")).
		WithArgs(sqlmock.AnyArg(), "Java Expert", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		// body make json
		Body          models.Book
		HasReturnBody bool
		ExpectedBody  string
	}{
		{
			"succes",
			200,
			"PUT",
			models.Book{
				BookName: "Java Expert",
			},
			true,
			"success update book",
		},
	}

	for _, v := range testCase {
		t.Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/books/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")
			ctx.Request().Header.Set("content-type", "application/json")
			err := UpdateBookController(ctx)
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

func TestDeleteBookController(t *testing.T) {
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
	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `books` WHERE `books`.`id` = ?")).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		// body make json
		Body          models.Book
		HasReturnBody bool
		ExpectedBody  string
	}{
		{
			"succes",
			200,
			"DELETE",
			models.Book{
				BookName: "Belajar Java",
			},
			true,
			"success delete book",
		},
	}

	for _, v := range testCase {
		t.Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/books/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")
			ctx.Request().Header.Set("content-type", "application/json")
			err := DeleteBookController(ctx)
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
