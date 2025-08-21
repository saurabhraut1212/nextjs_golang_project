package router

import (
	//"context"
	//"time"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	// "github.com/saurabhraut1212/internal/config"
	// "github.com/saurabhraut1212/internal/middleware"
	// "github.com/saurabhraut1212/internal/db"
	// "github.com/saurabhraut1212/internal/handlers"
	// "github.com/saurabhraut1212/internal/repo"
)

func New() *fiber.App {
	//cfg := config.Load()
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		//AllowOrigins: cfg.CORSOrigin,
		AllowMethods: "GET,POST,PATCH,DELETE,OPTIONS",
		AllowHeaders: "*",
	}))

	// Mongo client
	// client, err := db.New(cfg.MongoURI)
	// if err != nil { panic(err) }
	// database := client.Database(cfg.MongoDB)

	// Repos & indexes
	//ur := repo.NewUserRepo(database)
	//_ = ur.EnsureIndexes(context.Background())
	// todoRepo := repo.NewTodoRepo(database) // next step

	// Handlers
	// ah := handlers.NewAuthHandler(cfg, ur)
	// th := handlers.NewTodoHandler() // placeholder

	// Routes
	app.Get("/health", func(c *fiber.Ctx) error { return c.SendString("OK") })

	//api := app.Group("/api")
	//auth := api.Group("/auth")
	// auth.Post("/register", ah.Register)
	// auth.Post("/login", ah.Login)

	// protected := api.Group("/", middleware.RequireAuth(cfg))
	// protected.Get("/auth/me", ah.Me)
	// protected.Post("/auth/logout", ah.Logout)

	// todos := protected.Group("/todos")
	// todos.Get("/", th.List)
	// todos.Post("/", th.Create)
	// todos.Patch("/:id", th.Update)
	// todos.Delete("/:id", th.Delete)

	// Graceful shutdown: disconnect Mongo
	// app.Hooks().OnShutdown(func() error {
	// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// 	defer cancel()
	// 	//return client.Disconnect(ctx)
	// })

	return app
}
