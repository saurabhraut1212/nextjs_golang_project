package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/saurabhraut1212/nextjs_golang_project/internal/repo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoHandler struct{ tr *repo.TodoRepo }

func NewTodoHandler(tr *repo.TodoRepo) *TodoHandler { return &TodoHandler{tr: tr} }

func userIDFromCtx(c *fiber.Ctx) (primitive.ObjectID, error) {
	uidHex, _ := c.Locals("userId").(string)
	return primitive.ObjectIDFromHex(uidHex)
}

func (h *TodoHandler) List(c *fiber.Ctx) error {
	uid, err := userIDFromCtx(c)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	items, err := h.tr.ListByUser(ctx, uid)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(items)
}

type createReq struct {
	Title string `json:"title"`
}

func (h *TodoHandler) Create(c *fiber.Ctx) error {
	var body createReq
	if err := c.BodyParser(&body); err != nil || body.Title == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "title is required"})
	}
	uid, err := userIDFromCtx(c)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	item, err := h.tr.Create(ctx, uid, body.Title)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(item)
}

type updateReq struct {
	Title     *string `json:"title"`
	Completed *bool   `json:"completed"`
}

func (h *TodoHandler) Update(c *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	var body updateReq
	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	set := bson.M{}
	if body.Title != nil {
		set["title"] = *body.Title
	}
	if body.Completed != nil {
		set["completed"] = *body.Completed
	}
	if len(set) == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "no fields to update"})
	}

	uid, err := userIDFromCtx(c)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	item, err := h.tr.Update(ctx, uid, id, set)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(item)
}

func (h *TodoHandler) Delete(c *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	uid, err := userIDFromCtx(c)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := h.tr.Delete(ctx, uid, id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"ok": true})
}
