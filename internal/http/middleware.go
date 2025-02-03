package http

import (
	"github.com/freitzzz/gameboy-db-api/internal/logging"
	"github.com/freitzzz/gameboy-db-api/internal/service"
	"github.com/labstack/echo/v4"
)

const serviceContainerKey = "service.container"

// Container for all dependencies required in service context.
type serviceContainer struct {
	Games *service.GamesService
}

func ServiceContainer(
	games *service.GamesService,
) serviceContainer {
	return serviceContainer{
		Games: games,
	}
}

func registerMiddlewares(e *echo.Echo, sc serviceContainer) {
	e.Use(
		func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(ectx echo.Context) error {
				req := ectx.Request()

				logging.Info("Host: %s | Method: %s | Path: %s | Client IP: %s", req.Host, req.Method, req.URL.RequestURI(), ectx.RealIP())
				return next(ectx)
			}
		},
		func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(ectx echo.Context) error {
				ectx.Set(serviceContainerKey, sc)

				return next(ectx)
			}
		},
	)
}

func withServiceContainer(ectx echo.Context, with func(serviceContainer) error) error {
	container, ok := ectx.Get(serviceContainerKey).(serviceContainer)
	if !ok {
		return echo.ErrFailedDependency
	}

	return with(container)
}

func callAndReply[T any](ectx echo.Context, cb func() (T, error)) error {
	v, err := cb()
	if err != nil {
		return err
	}

	return ectx.JSON(200, v)
}
