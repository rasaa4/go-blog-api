package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"blog-api/config"
	"blog-api/db"
	"blog-api/internal/handler"
	"blog-api/internal/middleware"
	"blog-api/internal/repository"
	"blog-api/internal/usecase"
)

func main() {
	cfg := config.LoadConfig()

	database := db.ConnectDB(cfg.DBUser, cfg.DBPass, cfg.DBName)

	userRepo := repository.NewMySQLUserRepository(database)
	userUsecase := usecase.NewUserUsecase(userRepo, cfg.JWTSecret)
	userHandler := handler.NewUserHandler(userUsecase)

	postRepo := repository.NewMySQLPostRepository(database)
	postUsecase := usecase.NewPostUsecase(postRepo)
	postHandler := handler.NewPostHandler(postUsecase)

	r := chi.NewRouter()
	r.Post("register", userHandler.Register)
	r.Post("/login", userHandler.Login)
	r.Group(func(r chi.Router) {
		r.Use(middleware.JWTAuth("secret"))
		r.Post("/posts", postHandler.Create)
		r.Get("/posts", postHandler.GetPosts)
		r.Get("/posts/{id}", postHandler.GetPostByID)
		r.Delete("/posts/{id}", postHandler.Delete)
		r.Put("/posts/{id}", postHandler.Update)

	})
	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)

}
