package httpserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/natanaelrusli/segokuning-be/internal/config"
	"github.com/natanaelrusli/segokuning-be/internal/handler"
	"github.com/natanaelrusli/segokuning-be/internal/pkg/database"
	"github.com/natanaelrusli/segokuning-be/internal/pkg/encryptutils"
	"github.com/natanaelrusli/segokuning-be/internal/pkg/jwtutils"
	"github.com/natanaelrusli/segokuning-be/internal/repository"
	"github.com/natanaelrusli/segokuning-be/internal/usecase"
)

func InitGinServer(cfg *config.Config) {
	db, err := database.InitPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}

	passwordEncryptor := encryptutils.NewBcryptPasswordEncryptor(cfg.App.BCryptCost)
	jwtUtil := jwtutils.NewJWTUtil(cfg.Jwt)

	r := gin.New()

	userRepository := repository.NewUserRepository(db)
	authUsecase := usecase.NewAuthUsecaseImpl(userRepository, passwordEncryptor, jwtUtil)
	authHandler := handler.NewAuthHandler(authUsecase)

	r.POST("/v1/user/register", authHandler.Register)
	r.POST("/v1/user/login", authHandler.Login)

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.HttpServer.Host, cfg.HttpServer.Port),
		Handler: r,
	}

	srv.ListenAndServe()
}
