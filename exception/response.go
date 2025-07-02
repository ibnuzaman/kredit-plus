package exception

import (
	"fmt"
	"kredit-plus/constant"
	"kredit-plus/internal/model"
	"kredit-plus/logger"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type exception struct {
	log *zerolog.Logger
}

type Exception interface {
	Error(err error)
}

func NewException() Exception {
	return &exception{
		log: logger.GetWithoutCaller("exception"),
	}
}

func (e *exception) getCaller(skips ...int) string {
	skip := 2
	if len(skips) > 0 {
		skip = skips[0]
	}

	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return ""
	}

	return fmt.Sprintf("%s:%d", file, line)
}

func (e *exception) baseError(err error) {
	e.log.Error().Str("caller", e.getCaller(3)).Err(err).Msg("SERVER ERROR")
	panic(model.NewErrorMessage(fiber.StatusInternalServerError, constant.MsgInternalServerError, nil))
}

func (e *exception) Error(err error) {
	if err != nil {
		e.baseError(err)
	}
}
