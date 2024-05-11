package global

import (
	"TTMS_self/model/config"
	"TTMS_self/pkg/goroutine/work"
	"TTMS_self/pkg/logger"
	"TTMS_self/pkg/token"
	ut "github.com/go-playground/universal-translator"
)

var (
	RootDir  string
	Settings config.AllConfig
	Logger   *logger.Log
	Trans    ut.Translator
	Worker   *work.Worker
	Maker    token.Maker
)
