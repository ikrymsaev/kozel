package controllers

import (
	"fmt"
	pb_repositories "go-kozel/internal/infrastructure/pb/repositories"
	"go-kozel/internal/services"
	"log"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

type AuthController struct {
	app         *pocketbase.PocketBase
	authService *services.AuthService
	sessionRepo *pb_repositories.SessionsRepository
}

func NewAuthController(app *pocketbase.PocketBase) *AuthController {
	return &AuthController{
		app:         app,
		authService: services.NewAuthService(app),
		sessionRepo: pb_repositories.NewSessionsRepository(app),
	}
}

func (m *AuthController) Register() {
	m.app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.POST("/api/auth/signIn", m.signIn, apis.ActivityLogger(m.app))
		e.Router.POST("/api/auth/signUp", m.signUp, apis.ActivityLogger(m.app))
		return nil
	})
}

type SignBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (m *AuthController) signUp(c echo.Context) error {
	var body SignBody
	if err := c.Bind(&body); err != nil {
		log.Println("Failed to parse request body!:", err)
		err := fmt.Errorf("request body is invalid")
		return apis.NewBadRequestError(err.Error(), nil)
	}
	if body.Username == "" || body.Password == "" {
		log.Println("Request body is invalid!")
		err := fmt.Errorf("request body is invalid")
		return apis.NewBadRequestError(err.Error(), nil)
	}

	log.Println("Creating the user record...")
	usersCollection, err := m.app.Dao().FindCollectionByNameOrId("users")

	newUserRecord := models.NewRecord(usersCollection)
	if err != nil {
		log.Println("Failed to create user record!:", err)
		err := fmt.Errorf("failed to create user record")
		return apis.NewApiError(500, err.Error(), nil)
	}
	newUserRecord.Set("username", body.Username)
	newUserRecord.SetEmail(fmt.Sprintf("%s@localhost", body.Username))
	newUserRecord.SetPassword(body.Password)
	err = m.app.Dao().SaveRecord(newUserRecord)
	if err != nil {
		log.Println("Failed to save user record!:", err)
		err := fmt.Errorf("failed to save user record")
		return apis.NewApiError(500, err.Error(), nil)
	}

	m.sessionRepo.AddSessionLog(newUserRecord, c.RealIP())

	log.Println("Returning auth token...")
	return m.authService.GenUserAuthResult(c, newUserRecord)
}

func (m *AuthController) signIn(c echo.Context) error {
	var body SignBody
	if err := c.Bind(&body); err != nil {
		log.Println("Failed to parse request body!:", err)
		err := fmt.Errorf("request body is invalid")
		return apis.NewBadRequestError(err.Error(), nil)
	}
	if body.Username == "" || body.Password == "" {
		log.Println("Request body is invalid!")
		err := fmt.Errorf("request body is invalid")
		return apis.NewBadRequestError(err.Error(), nil)
	}

	log.Println("Getting the user record...")
	userRecord, err := m.app.Dao().FindFirstRecordByData("users", "username", body.Username)
	if err != nil {
		log.Println("Failed to get user!:", err)
		err := fmt.Errorf("failed to get user")
		return apis.NewApiError(500, err.Error(), nil)
	}

	m.sessionRepo.AddSessionLog(userRecord, c.RealIP())

	log.Println("Returning auth token...")
	return m.authService.GenUserAuthResult(c, userRecord)
}
