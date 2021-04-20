package handler

import (
	"github.com/gguibittencourt/go-restapi/handler/users"
	"go.uber.org/fx"
)

var Module = fx.Options(
	users.Module,
)
