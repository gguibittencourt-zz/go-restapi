package users

import (
	"encoding/json"
	"github.com/gguibittencourt/go-restapi/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"gorm.io/gorm"
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

func New(p Params) Handler {
	return &handler{
		logger: p.Logger,
		db:     p.DB,
	}
}

func (h *handler) List(writer http.ResponseWriter, request *http.Request) {
	var users []models.User
	err := models.ListUsers(h.db, &users)
	if err != nil {
		render.Status(request, http.StatusInternalServerError)
		render.JSON(writer, request, map[string]string{"message": err.Error()})
		return
	}
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
	err = models.GetUser(h.db, &user, id)
	if err != nil {
		render.Status(request, http.StatusNotFound)
		render.JSON(writer, request, map[string]string{"message": "User not found"})
		return
	}
	render.JSON(writer, request, user)
}

func (h *handler) Create(writer http.ResponseWriter, request *http.Request) {
	var user models.User
	body, _ := request.GetBody()
	err := json.NewDecoder(body).Decode(&user)
	if err != nil {
		render.Status(request, http.StatusBadRequest)
		render.JSON(writer, request, map[string]string{"message": "Incorrect body: " + err.Error()})
		return
	}
	err = models.CreateUser(h.db, &user)
	if err != nil {
		render.Status(request, http.StatusBadRequest)
		render.JSON(writer, request, map[string]string{"message": "Error on create user: " + err.Error()})
		return
	}
	render.JSON(writer, request, user)
}

func (h *handler) Update(writer http.ResponseWriter, request *http.Request) {
	panic("implement me")
}

func (h *handler) Delete(writer http.ResponseWriter, request *http.Request) {
	panic("implement me")
}