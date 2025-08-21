package router

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/saurabhraut1212/nextjs_golang_project/internal/config"
	"github.com/saurabhraut1212/nextjs_golang_project/internal/db"
	"github.com/saurabhraut1212/nextjs_golang_project/internal/handlers"
	"github.com/saurabhraut1212/nextjs_golang_project/internal/middleware"
	"github.com/saurabhraut1212/nextjs_golang_project/internal/repo"
)

func New() *fiber.App {
	cfg := config.Load()
	app := fiber.New()

	// Middlewares
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: cfg.CORSOrigin,
		AllowMethods: "GET,POST,PATCH,DELETE,OPTIONS",
		AllowHeaders: "*",
	}))

	// Mongo client
	client, err := db.New(cfg.MongoURI)
	if err != nil {
		panic(err)
	}
	database := client.Database(cfg.MongoDB)

	// Repositories & indexes
	userRepo := repo.NewUserRepo(database)
	_ = userRepo.EnsureIndexes(context.Background())

	todoRepo := repo.NewTodoRepo(database)
	_ = todoRepo.EnsureIndexes(context.Background()) // ✅ ensure index for todos

	// Handlers
	authHandler := handlers.NewAuthHandler(cfg, userRepo)
	todoHandler := handlers.NewTodoHandler(todoRepo) // ✅ inject todoRepo

	// Healthcheck
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// Public routes
	api := app.Group("/api")
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	// Protected routes
	protected := api.Group("/", middleware.RequireAuth(cfg))
	protected.Get("/auth/me", authHandler.Me)
	protected.Post("/auth/logout", authHandler.Logout)

	// Todo routes
	todos := protected.Group("/todos")
	todos.Get("/", todoHandler.List)
	todos.Post("/", todoHandler.Create)
	todos.Patch("/:id", todoHandler.Update)
	todos.Delete("/:id", todoHandler.Delete)

	// Graceful shutdown: disconnect Mongo
	app.Hooks().OnShutdown(func() error {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return client.Disconnect(ctx)
	})

	return app
}
