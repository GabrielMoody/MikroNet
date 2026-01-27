package handler_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/GabrielMoody/MikroNet/services/authentication/internal/controller"
	"github.com/GabrielMoody/MikroNet/services/authentication/internal/dto"
	"github.com/GabrielMoody/MikroNet/services/authentication/internal/handler"
	"github.com/GabrielMoody/MikroNet/services/authentication/internal/repository"
	"github.com/GabrielMoody/MikroNet/services/authentication/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/suite"
	gormpostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AuthHandlerTestSuite struct {
	suite.Suite
	app      *fiber.App
	db       *gorm.DB
	pgxPool  *pgxpool.Pool
	httpPort string
}

func (suite *AuthHandlerTestSuite) SetupSuite() {
	// Set environment variables for the application
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "123")
	os.Setenv("DB_NAME", "mikronet")
	os.Setenv("JWT_SECRET", "test-secret")
	os.Setenv("JWT_ISS", "test-iss")

	connStr := "host=localhost port=5432 user=test password=test dbname=test_db sslmode=disable"

	// Initialize gorm
	gormDB, err := gorm.Open(gormpostgres.Open(connStr), &gorm.Config{})
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.db = gormDB

	// Run migrations
	sqlFile, err := ioutil.ReadFile("../../../../init.sql")
	if err != nil {
		suite.T().Fatal(err)
	}

	// Split the SQL file into individual statements
	sqlStatements := bytes.Split(sqlFile, []byte(";"))

	for _, statement := range sqlStatements {
		if len(bytes.TrimSpace(statement)) > 0 {
			if err := suite.db.Exec(string(statement)).Error; err != nil {
				suite.T().Fatal(err)
			}
		}
	}

	// Initialize fiber app
	app := fiber.New()
	repo := repository.NewAuthRepo(suite.db)
	authService := service.NewAuthService(repo)
	authController := controller.NewAuthController(authService)
	api := app.Group("/")
	handler.AuthHandler(api, authController)
	suite.app = app
}

func (suite *AuthHandlerTestSuite) TearDownTest() {
	// Clean up the database after each test
	suite.db.Exec("TRUNCATE TABLE authentications, users, drivers CASCADE")
}

func TestAuthHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(AuthHandlerTestSuite))
}

func (suite *AuthHandlerTestSuite) TestHealthCheck() {
	req := httptest.NewRequest("GET", "/healthcheck", nil)
	resp, err := suite.app.Test(req)
	suite.NoError(err)
	suite.Equal(http.StatusOK, resp.StatusCode)
}

func (suite *AuthHandlerTestSuite) TestCreateUser_Success() {
	user := dto.UserRegistrationsReq{
		Email:                "test@example.com",
		Password:             "password",
		PasswordConfirmation: "password",
		Name:                 "Test User",
		PhoneNumber:          "123456789",
	}
	body, _ := json.Marshal(user)
	req := httptest.NewRequest("POST", "/register/user", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := suite.app.Test(req)
	suite.NoError(err)
	suite.Equal(http.StatusCreated, resp.StatusCode)
}

func (suite *AuthHandlerTestSuite) TestCreateUser_Duplicate() {
	user := dto.UserRegistrationsReq{
		Email:                "test@example.com",
		Password:             "password",
		PasswordConfirmation: "password",
		Name:                 "Test User",
		PhoneNumber:          "123456789",
	}
	body, _ := json.Marshal(user)

	// First request
	req1 := httptest.NewRequest("POST", "/register/user", bytes.NewReader(body))
	req1.Header.Set("Content-Type", "application/json")
	resp1, err1 := suite.app.Test(req1)
	suite.NoError(err1)
	suite.Equal(http.StatusCreated, resp1.StatusCode)

	// Second request with the same email
	req2 := httptest.NewRequest("POST", "/register/user", bytes.NewReader(body))
	req2.Header.Set("Content-Type", "application/json")
	resp2, err2 := suite.app.Test(req2)
	suite.NoError(err2)
	suite.Equal(http.StatusConflict, resp2.StatusCode)
}

func (suite *AuthHandlerTestSuite) TestCreateDriver_Success() {
	driver := dto.DriverRegistrationsReq{
		Email:                "driver@example.com",
		Password:             "password",
		PasswordConfirmation: "password",
		Name:                 "Test Driver",
		PlateNumber:          "B 1234 CD",
		PhoneNumber:          "123456789",
	}
	body, _ := json.Marshal(driver)
	req := httptest.NewRequest("POST", "/register/driver", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := suite.app.Test(req)
	suite.NoError(err)
	suite.Equal(http.StatusCreated, resp.StatusCode)
}

func (suite *AuthHandlerTestSuite) TestLoginUser_Success() {
	// Create user first
	user := dto.UserRegistrationsReq{
		Email:                "test@example.com",
		Password:             "password",
		PasswordConfirmation: "password",
		Name:                 "Test User",
		PhoneNumber:          "123456789",
	}
	body, _ := json.Marshal(user)
	req := httptest.NewRequest("POST", "/register/user", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	suite.app.Test(req)

	// Login
	loginReq := dto.UserLoginReq{
		Email:    "test@example.com",
		Password: "password",
	}
	loginBody, _ := json.Marshal(loginReq)
	loginReq_ := httptest.NewRequest("POST", "/login", bytes.NewReader(loginBody))
	loginReq_.Header.Set("Content-Type", "application/json")

	resp, err := suite.app.Test(loginReq_)
	suite.NoError(err)
	suite.Equal(http.StatusOK, resp.StatusCode)

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	suite.Equal("success", result["status"])
	suite.NotNil(result["data"])
}

func (suite *AuthHandlerTestSuite) TestLoginUser_NotFound() {
	loginReq := dto.UserLoginReq{
		Email:    "notfound@example.com",
		Password: "password",
	}
	loginBody, _ := json.Marshal(loginReq)
	req := httptest.NewRequest("POST", "/login", bytes.NewReader(loginBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := suite.app.Test(req)
	suite.NoError(err)
	suite.Equal(http.StatusNotFound, resp.StatusCode)
}

func (suite *AuthHandlerTestSuite) TestLoginUser_IncorrectPassword() {
	// Create user first
	user := dto.UserRegistrationsReq{
		Email:                "test@example.com",
		Password:             "password",
		PasswordConfirmation: "password",
		Name:                 "Test User",
		PhoneNumber:          "123456789",
	}
	body, _ := json.Marshal(user)
	req := httptest.NewRequest("POST", "/register/user", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	suite.app.Test(req)

	// Login with incorrect password
	loginReq := dto.UserLoginReq{
		Email:    "test@example.com",
		Password: "wrongpassword",
	}
	loginBody, _ := json.Marshal(loginReq)
	loginReq_ := httptest.NewRequest("POST", "/login", bytes.NewReader(loginBody))
	loginReq_.Header.Set("Content-Type", "application/json")

	resp, err := suite.app.Test(loginReq_)
	suite.NoError(err)
	suite.Equal(http.StatusUnauthorized, resp.StatusCode)
}
