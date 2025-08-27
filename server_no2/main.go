package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	RealName string `json:"realname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	app := fiber.New()
	InitRedis()

	app.Post("/login", loginHandler)

	app.Listen(":3000")
}

func loginHandler(c *fiber.Ctx) error {
	type LoginInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}

	key := "login_" + input.Username
	val, err := Rdb.Get(ctx, key).Result()
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "user not found"})
	}

	var user User
	if err := json.Unmarshal([]byte(val), &user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to parse user data"})
	}

	// Hash password input
	h := sha1.New()
	h.Write([]byte(input.Password))
	hashedInput := hex.EncodeToString(h.Sum(nil))

	if user.Password != hashedInput {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid password"})
	}

	return c.JSON(fiber.Map{
		"message":  "login success",
		"realname": user.RealName,
		"email":    user.Email,
	})
}
