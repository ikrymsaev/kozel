package services

import (
	"go-kozel/internal/domain"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tokens"
)

type AuthService struct {
	pb *pocketbase.PocketBase
}

func NewAuthService(pb *pocketbase.PocketBase) *AuthService {
	return &AuthService{
		pb: pb,
	}
}

func (s *AuthService) GetUserFromToken(token string) (domain.User, error) {
	userRecord, err := s.pb.Dao().FindAuthRecordByToken(
		token,
		s.pb.Settings().RecordAuthToken.Secret,
	)
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{ID: userRecord.Id, Username: userRecord.Username()}, nil
}

type SignedUser struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

// Overrided method from pocketbase
func (s *AuthService) GenUserAuthResult(c echo.Context, authRecord *models.Record) error {
	token, tokenErr := tokens.NewRecordAuthToken(s.pb, authRecord)
	if tokenErr != nil {
		return apis.NewBadRequestError("Failed to create auth token.", tokenErr)
	}

	event := new(core.RecordAuthEvent)
	event.HttpContext = c
	event.Collection = authRecord.Collection()
	event.Record = authRecord
	event.Token = token

	return s.pb.OnRecordAuthRequest().Trigger(event, func(e *core.RecordAuthEvent) error {
		user := SignedUser{Id: e.Record.Id, Username: e.Record.Username()}
		result := map[string]any{
			"token": e.Token,
			"user":  user,
		}
		return e.HttpContext.JSON(http.StatusOK, result)
	})
}
