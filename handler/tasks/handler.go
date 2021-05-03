package tasks

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
	var tasks []models.Task
	err := models.ListTasks(h.db, &tasks)
	if err != nil {
		render.Status(request, http.StatusInternalServerError)
		render.JSON(writer, request, map[string]string{"message": err.Error()})
		return
	}
	render.JSON(writer, request, tasks)
}

func (h *handler) Find(writer http.ResponseWriter, request *http.Request) {
	param := chi.URLParam(request, "id")
	id, err := strconv.Atoi(param)
	if err != nil {
		render.Status(request, http.StatusBadRequest)
		render.JSON(writer, request, map[string]string{"message": "Id not found"})
		return
	}
	var task models.Task
	err = models.GetTask(h.db, &task, id)
	if err != nil {
		render.Status(request, http.StatusNotFound)
		render.JSON(writer, request, map[string]string{"message": "Task not found"})
		return
	}
	render.JSON(writer, request, task)
}

func (h *handler) Create(writer http.ResponseWriter, request *http.Request) {
	var task models.Task
	body := request.Body
	err := json.NewDecoder(body).Decode(&task)
	if err != nil {
		render.Status(request, http.StatusBadRequest)
		render.JSON(writer, request, map[string]string{"message": "Incorrect body: " + err.Error()})
		return
	}
	err = models.CreateTask(h.db, &task)
	if err != nil {
		render.Status(request, http.StatusBadRequest)
		render.JSON(writer, request, map[string]string{"message": "Error on create Task: " + err.Error()})
		return
	}
	render.JSON(writer, request, task)
}

func (h *handler) Update(writer http.ResponseWriter, request *http.Request) {
	panic("implement me")
}

func (h *handler) Delete(writer http.ResponseWriter, request *http.Request) {
	panic("implement me")
}