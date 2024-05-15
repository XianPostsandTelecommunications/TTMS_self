package global

import (
	ut "github.com/go-playground/universal-translator"
	"mognolia/internal/model/config"
	"mognolia/internal/pkg/goroutine/work"
	"mognolia/internal/pkg/logger"
	"mognolia/internal/pkg/token"
)

var (
	RootDir  string
	Settings config.AllConfig
	Logger   *logger.Log
	Trans    ut.Translator
	Worker   *work.Worker
	Maker    token.Maker
)
