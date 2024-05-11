package logger

import (
	"go.uber.org/zap"
)

type Log struct {
	*zap.Logger
}
