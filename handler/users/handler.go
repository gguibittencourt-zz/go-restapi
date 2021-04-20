package users

import (
	"github.com/gguibittencourt/go-restapi/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Provide(New)

type Handler interface {
	List(http.ResponseWriter, *http.Request)
	Find(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

type Params struct {
	fx.In

	Logger *zap.Logger
	DB     *gorm.DB
}

type handler struct {
	logger *zap.Logger
	db     *gorm.DB
}

func New(p Params) (Handler, error) {
	return &handler{
		logger: p.Logger,
		db:     p.DB,
	}, nil
}

func (h *handler) List(writer http.ResponseWriter, request *http.Request) {
	var users []models.User
	h.db.Find(&users)
	render.JSON(writer, request, users)
}

func (h *handler) Find(writer http.ResponseWriter, request *http.Request) {
	param := chi.URLParam(request, "id")
	id, err := strconv.Atoi(param)
	if err != nil {
		render.Status(request, http.StatusBadRequest)
		render.JSON(writer, request, map[string]string{"message": "Id not found"})
		return
	}
	var user models.User
	h.db.Where("id = ?", id).First(&user)
	if user.Id == 0 {
		render.Status(request, http.StatusNotFound)
		render.JSON(writer, request, map[string]string{"message": "User not found"})
		return
	}
	render.JSON(writer, request, user)
}

func (h *handler) Create(writer http.ResponseWriter, request *http.Request) {
	panic("implement me")
}

func (h *handler) Update(writer http.ResponseWriter, request *http.Request) {
	panic("implement me")
}

func (h *handler) Delete(writer http.ResponseWriter, request *http.Request) {
	panic("implement me")
}

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}