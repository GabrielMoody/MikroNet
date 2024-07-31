package test

import (
	"encoding/json"
	"fmt"
	"github.com/GabrielMoody/MikroNet/authentication/internal/handler"
	"github.com/GabrielMoody/MikroNet/authentication/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var app *fiber.App

func init() {
	app = fiber.New()

	db := models.DatabaseTestInit()

	api := app.Group("/api")

	handler.ProfileHandler(api, db)

	err := app.Listen("localhost:8000")
	if err != nil {
		return
	}
}

func TestCreateUserHandling(t *testing.T) {
	reqBody := strings.NewReader(`{
    "nama_lengkap":"Gabriel Moody Waworundeng",
    "email":"gabriel@gmail.com",
    "nomor_telepon":"0822991283201",
    "kata_sandi":"galjdwieoenfiow",
    "konfirmasi_kata_sandi": "galjdwieoenfiow"
	}`)

	req := httptest.NewRequest(http.MethodPost, "/api/profile", reqBody)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	resp, _ := app.Test(req)
	body, _ := io.ReadAll(resp.Body)
	var respBody map[string]interface{}
	err := json.Unmarshal(body, &respBody)

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 201, resp.StatusCode, "Status code must be 201")
	assert.Equal(t, fmt.Sprintf(`{"data":"%s","status":"success"}`, respBody["data"]), string(body))
}

func TestCreateUserAlreadyExist(t *testing.T) {
	reqBody := strings.NewReader(`{
    "nama_lengkap":"Gabriel Moody Waworundeng",
    "email":"gabriel@gmail.com",
    "nomor_telepon":"0822991283201",
    "kata_sandi":"galjdwieoenfiow",
    "konfirmasi_kata_sandi": "galjdwieoenfiow"
	}`)

	req := httptest.NewRequest(http.MethodPost, "/api/profile", reqBody)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	app.Test(req)
	resp, _ := app.Test(req)
	body, _ := io.ReadAll(resp.Body)
	var respBody map[string]interface{}
	err := json.Unmarshal(body, &respBody)

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 500, resp.StatusCode, "Status code must be 500")
	assert.Equal(t, `{"error":"email dan/atau nomor telepon telah terdaftar","status":"error"}`, string(body))
}
