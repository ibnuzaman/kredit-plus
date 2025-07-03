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
	BadRequest(messages ...string)
	BadRequestErr(err error, messages ...string)
	Unauthorized(messages ...string)
	UnauthorizedErr(err error, messages ...string)
	UnauthorizedBool(isError bool, messages ...string)
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

func (e *exception) getMessage(defaultMessage string, messages ...string) string {
	if len(messages) > 0 {
		return messages[0]
	}

	return defaultMessage
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

func (e *exception) baseBadRequest(messages ...string) {
	e.log.Warn().Str("caller", e.getCaller(4)).Msg("BAD_REQUEST")
	panic(model.NewErrorMessage(fiber.StatusBadRequest, e.getMessage(constant.MsgBadRequest, messages...), nil))
}

func (e *exception) BadRequest(messages ...string) {
	e.baseBadRequest(messages...)
}

func (e *exception) BadRequestErr(err error, messages ...string) {
	if err != nil {
		e.baseBadRequest(messages...)
	}
}

func (e *exception) baseUnauthorized(messages ...string) {
	e.log.Warn().Str("caller", e.getCaller(3)).Msg("UNAUTHORIZED")
	panic(model.NewErrorMessage(fiber.StatusUnauthorized, e.getMessage(constant.MsgUnauthorized, messages...), nil))
}

func (e *exception) Unauthorized(messages ...string) {
	e.baseUnauthorized(messages...)
}

func (e *exception) UnauthorizedErr(err error, messages ...string) {
	if err != nil {
		e.baseUnauthorized(messages...)
	}
}

func (e *exception) UnauthorizedBool(isError bool, messages ...string) {
	if isError {
		e.baseUnauthorized(messages...)
	}
}
