package router

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/aportela/gotask/internal/fileserver"
	"github.com/aportela/gotask/internal/handlers"
)

func NewRouter() http.Handler {
	baseRouter := chi.NewRouter()

	baseRouter.Use(middleware.Logger)

	apiRouter := chi.NewRouter()
	apiRouter.Get("/hello", handlers.DefaultHandler)
	apiRouter.Get("/projects", handlers.SearchProjectsHandler)
	apiRouter.Get("/tasks", handlers.SearchTasksHandler)
	baseRouter.Mount("/api", apiRouter)

	workDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	publicWebPath := http.Dir(filepath.Join(workDir, "public"))
	fileserver.FileServer(baseRouter, "/static", publicWebPath)

	return baseRouter
}
