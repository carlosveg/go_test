package services

import (
	"strings"

	"github.com/carlosveg/go_test/assets"
	"github.com/carlosveg/go_test/data"
	"github.com/carlosveg/go_test/models"
	"github.com/gofiber/fiber/v2"
)

// Función que agrega un usuario validando cada campo
// En caso de error muestra el respectivo mensaje
// En caso de éxito registra al usuario y lo retorna con un códido de estado 200
func AddUser(c *fiber.Ctx) error {
	var response models.Response
	user := models.User{}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	isValid, emptyFields := assets.ValidateEmptyFields(user)

	// Si existe 1 o más campos vacios se retorna un mensaje indicando el/los campos faltantes
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

	// Se valida si el telefono o correo ingresados ya se encuentran registrados
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

// Función de prueba para validar que se guarden correctamente los usuarios
func ListUsers(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(data.Users)
}

// Función Login, valida las credenciales del usuario
// Un usuario puede loggearse con su correo o su usuario
// En caso de error o campos faltantes muestra el respectivo mensaje
// En caso de éxito registra al usuario y lo retorna con un códido de estado 200
func Login(c *fiber.Ctx) error {
	var response models.Response
	var user models.User
	var credentials models.Credentials

	if err := c.BodyParser(&credentials); err != nil {
		return err
	}

	isValid, emptyFields := assets.ValidateLogin(credentials)

	if !isValid {
		message := "Faltan campos: "

		for _, field := range emptyFields {
			message += field + ", "
		}

		response = models.Response{Message: message}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	// Verificamos si introdujo un correo o un usuario
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

	// Validamos credenciales
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
