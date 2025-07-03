package middleware

import (
	"errors"
	"kredit-plus/constant"
	"kredit-plus/internal/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func (m *middleware) Auth() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authTokens, isFound := ctx.GetReqHeaders()[fiber.HeaderAuthorization]
		if !isFound || len(authTokens) == 0 {
			return m.errUnauthorized(ctx)
		}

		authToken := authTokens[0]
		if len(authToken) < 7 || authToken[:7] != "Bearer " {
			return m.errUnauthorized(ctx)
		}

		authToken = authToken[7:]
		payload, isValid := m.validateJwtToken(authToken)
		if !isValid {
			return m.errUnauthorized(ctx)
		}

		customerId, err := strconv.Atoi(payload.ID)
		if err != nil {
			return m.errUnauthorized(ctx)
		}

		user, err := m.authRepo.FindById(ctx.Context(), customerId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return m.errForbidden(ctx)
			}
			return m.errUnauthorized(ctx)
		}

		ctx.Locals(constant.KeyLocalsAuthUser, user.ToAuthMe())

		return ctx.Next()
	}
}

func (m *middleware) errUnauthorized(ctx *fiber.Ctx) error {
	code := fiber.StatusUnauthorized
	return ctx.Status(code).JSON(model.BaseResponse{
		Code:    code,
		Message: http.StatusText(code),
		Data:    nil,
	})
}

func (m *middleware) errForbidden(ctx *fiber.Ctx) error {
	code := fiber.StatusForbidden
	return ctx.Status(code).JSON(model.BaseResponse{
		Code:    code,
		Message: http.StatusText(code),
		Data:    nil,
	})
}

func (m *middleware) validateJwtToken(token string) (*jwt.RegisteredClaims, bool) {
	if len(token) == 0 {
		return nil, false
	}

	var claims jwt.RegisteredClaims
	secret := m.conf.Auth.SecretKey
	value, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	}, jwt.WithLeeway(5*time.Second))

	if err != nil || !value.Valid {
		return nil, false
	}

	if _, ok := value.Claims.(*jwt.RegisteredClaims); !ok {
		return nil, false
	}

	return &claims, true
}
