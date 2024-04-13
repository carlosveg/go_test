package services

import (
	"strings"

	"github.com/carlosveg/go_test/assets"
	"github.com/carlosveg/go_test/data"
	"github.com/carlosveg/go_test/models"
	"github.com/gofiber/fiber/v2"
)

func AddUser(c *fiber.Ctx) error {
	var response models.Response
	user := models.User{}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	isValid, emptyFields := assets.ValidateEmptyFields(user)

	if !isValid {
		message := "Faltan campos: "

		for _, field := range emptyFields {
			message += field + ", "
		}

		response = models.Response{Message: message}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if !assets.ValidateEmail(user.Correo) {
		response = models.Response{Message: "Formato de correo inválido"}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if !assets.ValidatePhone(user.Telefono) {
		response = models.Response{Message: "Formato de teléfono inválido"}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if assets.TelefonoRegistrado(user.Telefono, data.Users) || assets.CorreoRegistrado(user.Correo, data.Users) {
		response = models.Response{Message: "El correo/telefono ya se encuentra registrado"}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if !assets.ValidatePassword(user.Password) {
		response = models.Response{Message: "Formato de contraseña inválido"}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	data.Users[user.Correo] = user

	return c.Status(fiber.StatusOK).JSON(user)
}

func ListUsers(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(data.Users)
}

func Login(c *fiber.Ctx) error {
	var response models.Response
	var user models.User
	var credentials models.Credentials

	if err := c.BodyParser(&credentials); err != nil {
		return err
	}

	if strings.Contains(credentials.UsuarioCorreo, "@") {
		user = data.Users[credentials.UsuarioCorreo]

	} else {
		for _, u := range data.Users {
			if u.Username == credentials.UsuarioCorreo {
				user = u
				break
			}
		}
	}
	if user.Username == "" || user.Password != credentials.Password {
		response = models.Response{Message: "usuario / contraseña incorrectos"}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	tokenString, err := assets.GenerarTokenJWT(credentials.UsuarioCorreo)

	if err != nil {
		response = models.Response{Message: "Error al generar token JWT"}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(models.TokenResponse{Token: tokenString})

}
