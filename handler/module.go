package handler

import (
	"github.com/gguibittencourt/go-restapi/handler/tasks"
	"go.uber.org/fx"
)

var Module = fx.Options(
	tasks.Module,
)
