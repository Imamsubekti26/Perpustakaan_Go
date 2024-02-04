package controllers

import (
	"fmt"

	"github.com/Imamsubekti26/Perpustakaan_Go/models"
	jwt "github.com/Imamsubekti26/Perpustakaan_Go/utils/JWT"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		DB: db,
	}
}

func (u *UserController) Login(c *fiber.Ctx) error {
	var user models.Users

	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "" || password == "" {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Username or Password is empty",
		})
	}

	if err := u.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Username or Password not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Username or Password not found",
		})
	}

	token, err := jwt.GenerateToken(user.Username, user.IsAdmin)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to generate token: %s", err),
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"token": token,
	})
}

func (u *UserController) Register(c *fiber.Ctx) error {
	var user models.Users

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	if user.Username == "" || user.Password == "" || user.Surename == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Please fill all input form",
		})
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	if err := u.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(&user)
}
