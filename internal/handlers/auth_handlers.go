package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/saurabhraut1212/nextjs_golang_project/internal/auth"
	"github.com/saurabhraut1212/nextjs_golang_project/internal/config"
	"github.com/saurabhraut1212/nextjs_golang_project/internal/models"
	"github.com/saurabhraut1212/nextjs_golang_project/internal/repo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	cfg *config.Config
	ur  *repo.UserRepo
}

func NewAuthHandler(cfg *config.Config, ur *repo.UserRepo) *AuthHandler {
	return &AuthHandler{cfg: cfg, ur: ur}
}

type registerReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var body registerReq
	if err := c.BodyParser(&body); err != nil || body.Email == "" || body.Password == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if existing, _ := h.ur.ByEmail(ctx, body.Email); existing != nil {
		return c.Status(http.StatusConflict).JSON(fiber.Map{"error": "email already exists"})
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	u := &models.User{
		Name: body.Name, Email: body.Email, PasswordHash: string(hash),
	}
	if err := h.ur.Create(ctx, u); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	access, refresh, aExp, rExp, err := auth.GenerateTokens(h.cfg.JWTSecret, u.ID.Hex(), h.cfg.AccessTTL, h.cfg.RefreshTTL)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "token error"})
	}
	auth.SetAuthCookies(c, access, aExp, refresh, rExp, h.cfg.Env == "prod")
	return c.Status(http.StatusCreated).JSON(fiber.Map{"_id": u.ID, "name": u.Name, "email": u.Email})
}

type loginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var body loginReq
	if err := c.BodyParser(&body); err != nil || body.Email == "" || body.Password == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	u, _ := h.ur.ByEmail(ctx, body.Email)
	if u == nil || bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(body.Password)) != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "invalid credentials"})
	}
	access, refresh, aExp, rExp, err := auth.GenerateTokens(h.cfg.JWTSecret, u.ID.Hex(), h.cfg.AccessTTL, h.cfg.RefreshTTL)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "token error"})
	}
	auth.SetAuthCookies(c, access, aExp, refresh, rExp, h.cfg.Env == "prod")
	return c.JSON(fiber.Map{"_id": u.ID, "name": u.Name, "email": u.Email})
}

func (h *AuthHandler) Me(c *fiber.Ctx) error {
	uidHex, _ := c.Locals("userId").(string)
	if uidHex == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}
	id, _ := primitive.ObjectIDFromHex(uidHex)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	u, _ := h.ur.ByID(ctx, id)
	if u == nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "not found"})
	}
	return c.JSON(fiber.Map{"_id": u.ID, "name": u.Name, "email": u.Email})
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	auth.ClearAuthCookies(c)
	return c.JSON(fiber.Map{"ok": true})
}
