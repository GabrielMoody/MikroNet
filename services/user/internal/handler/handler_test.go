package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/GabrielMoody/MikroNet/services/user/internal/handler"
	"github.com/GabrielMoody/MikroNet/services/user/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/suite"
	gormsqlite "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserHandlerTestSuite struct {
	suite.Suite
	app *fiber.App
	db  *gorm.DB
}

func (suite *UserHandlerTestSuite) SetupSuite() {
	os.Setenv("JWT_SECRET", "test-secret")
	os.Setenv("JWT_ISS", "test-iss")

	gormDB, err := gorm.Open(gormsqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.db = gormDB

	sqlFile, err := ioutil.ReadFile("../../../../init.sql")
	if err != nil {
		suite.T().Fatal(err)
	}

	sqlStatements := bytes.Split(sqlFile, []byte(";"))

	for _, statement := range sqlStatements {
		if len(bytes.TrimSpace(statement)) > 0 {
			if err := suite.db.Exec(string(statement)).Error; err != nil {
				// We can ignore errors here if they are due to syntax differences between PG and SQLite
			}
		}
	}

	app := fiber.New()
	api := app.Group("/")
	handler.UserHandler(api, suite.db, nil)
	suite.app = app
}

func (suite *UserHandlerTestSuite) TearDownTest() {
	suite.db.Exec("DELETE FROM users")
	suite.db.Exec("DELETE FROM orders")
	suite.db.Exec("DELETE FROM authentications")
}

func TestUserHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}

func (suite *UserHandlerTestSuite) TestHealthCheck() {
	req := httptest.NewRequest("GET", "/healthcheck", nil)
	resp, err := suite.app.Test(req)
	var result map[string]string
	json.NewDecoder(resp.Body).Decode(&result)
	suite.NoError(err)
	suite.Equal(http.StatusOK, resp.StatusCode)
	suite.Equal(map[string]string{"status": "pass"}, result)
}

func (suite *UserHandlerTestSuite) generateTestJWTWithStringID(userID int64, role string) string {
	claims := jwt.MapClaims{
		"id":    fmt.Sprintf("%d", userID),
		"email": fmt.Sprintf("testuser%d@example.com", userID),
		"role":  role,
		"kid":   role + "-kid",
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
		"iss":   os.Getenv("JWT_ISS"),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return t
}

func (suite *UserHandlerTestSuite) generateTestJWTWithNumericID(userID int64, role string) string {
	claims := jwt.MapClaims{
		"id":    userID,
		"email": fmt.Sprintf("testuser%d@example.com", userID),
		"role":  role,
		"kid":   role + "-kid",
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
		"iss":   os.Getenv("JWT_ISS"),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return t
}

func (suite *UserHandlerTestSuite) TestGetUser_Success() {
	// Setup: Create a user in the database
	testUser := model.User{
		ID:       1,
		Username: "testuser@example.com",
		Fullname: "Test User",
	}
	suite.db.Create(&testUser)

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer "+suite.generateTestJWTWithStringID(1, "user"))

	resp, err := suite.app.Test(req)
	suite.NoError(err)

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	suite.Equal("success", result["status"])
}
