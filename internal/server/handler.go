package server

import (
	"net/http"
	"strings"

	"github.com/DLzer/go-echo-boilerplate/pkg/metric"
	apiMiddlewares "github.com/DLzer/go-echo-boilerplate/pkg/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	usersHttp "github.com/DLzer/go-echo-boilerplate/internal/users/http"
	usersRepository "github.com/DLzer/go-echo-boilerplate/internal/users/repository"
	usersService "github.com/DLzer/go-echo-boilerplate/internal/users/service"
)

// Map Server Handlers
func (s *Server) MapHandlers(e *echo.Echo) error {
	metrics, err := metric.CreateMetrics(s.cfg.Metrics.URL, s.cfg.Metrics.ServiceName)
	if err != nil {
		s.logger.Errorf("CreateMetrics Error: %s", err)
	}
	s.logger.Infof(
		"Metrics available URL: %s, ServiceName: %s",
		s.cfg.Metrics.URL,
		s.cfg.Metrics.ServiceName,
	)

	usersPgRepo := usersRepository.NewUsersPgRepo(s.db)
	usersService := usersService.NewUsersService(s.cfg, usersPgRepo, s.logger)
	usersHandler := usersHttp.NewUsersHandler(s.cfg, usersService, s.logger)

	// Swagger Setup
	// docs.SwaggerInfo.Title = "Echo Rest API"
	// e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Middlewares
	mw := apiMiddlewares.NewMiddlewareManager(s.cfg, []string{"*"}, s.logger)

	e.Use(mw.RequestLoggerMiddleware)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXRequestID, "x-api-key"},
	}))
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         1 << 10, // 1 KB
		DisablePrintStack: true,
		DisableStackAll:   true,
	}))
	e.Use(middleware.RequestID())
	e.Use(mw.MetricsMiddleware(metrics))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
		// Skipper: func(c echo.Context) bool {
		// 	return strings.Contains(c.Request().URL.Path, "swagger")
		// },
	}))
	e.Use(middleware.Secure())
	e.Use(middleware.BodyLimit("2M"))
	if s.cfg.Server.Debug {
		e.Use(mw.DebugMiddleware)
	}

	if s.cfg.Server.Mode != "Development" {
		e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
			Skipper: func(c echo.Context) bool {
				return strings.Contains(c.Request().URL.Path, "/health")
			},
			KeyLookup: "header:x-api-key",
			Validator: func(key string, c echo.Context) (bool, error) {
				return key == s.cfg.Server.ApiKey, nil
			},
		}))
	}

	// Define Route Group
	v1 := e.Group("/v1")
	usersGroup := v1.Group("/users")
	health := v1.Group("/health")

	// Map groups to handlers
	usersHttp.UserRoutes(usersGroup, usersHandler, *mw)

	// Health route function
	health.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})

	return nil
}
