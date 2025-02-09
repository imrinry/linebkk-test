package auth_test

import (
	"line-bk-api/internal/auth"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRegisterAuthRoutes(t *testing.T) {
	app := fiber.New()
	mockAuthHandler := &mockAuthHandler{}

	auth.RegisterAuthRoutes(app, mockAuthHandler)

	// Test login with pin route
	req := httptest.NewRequest("POST", "/api/v1/auth/login/pin", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	// Test login with password route
	req = httptest.NewRequest("POST", "/api/v1/auth/login/password", nil)
	resp, err = app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

type mockAuthHandler struct {
}

func (m *mockAuthHandler) LoginWithPinCode(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func (m *mockAuthHandler) LoginWithPassword(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
