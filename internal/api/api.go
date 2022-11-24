package api

import (
	"context"
	"net/http"
	"strconv"

	"github.com/alesr/apierrordemo/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go.uber.org/zap"
)

type userService interface {
	FetchUser(ctx context.Context, id string) (*service.User, error)
}

type API struct {
	logger  *zap.Logger
	service userService
}

func New(logger *zap.Logger, svc userService) *API {
	return &API{
		logger:  logger,
		service: svc,
	}
}

func (api *API) FetchUser(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")

	if err := validateUserID(id); err != nil {
		api.logger.Info("could not validate user id", zap.Error(err))
		render.Status(req, http.StatusBadRequest)
		render.JSON(w, req, err)
		return
	}

	user, err := api.service.FetchUser(context.TODO(), id)
	if err != nil {
		api.logger.Info("could not fetch user", zap.Error(err))

		e := transportError(err)

		render.Status(req, e.StatusCode)
		render.JSON(w, req, e)
		return
	}

	render.JSON(w, req, User{
		Name:  user.Name,
		Email: user.Email,
	})
}

// Syntax check
func validateUserID(id string) error {
	if _, err := strconv.Atoi(id); err != nil {
		return ErrInvalidUserID
	}
	return nil
}
