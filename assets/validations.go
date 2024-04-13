package assets

import (
	"regexp"

	"github.com/carlosveg/go_test/models"
)

func ValidatePassword(password string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9@$&]{6,12}$`)
	return regex.MatchString(password)
}

func ValidatePhone(telefono string) bool {
	regex := regexp.MustCompile(`^\d{10}$`)
	return regex.MatchString(telefono)
}

func ValidateEmail(correo string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return regex.MatchString(correo)
}

func ValidateEmptyFields(user models.User) (bool, []string) {
	var emptyFields []string

	if user.Username == "" {
		emptyFields = append(emptyFields, "usuario")
	}
	if user.Correo == "" {
		emptyFields = append(emptyFields, "correo")
	}
	if user.Telefono == "" {
		emptyFields = append(emptyFields, "telefono")
	}
	if user.Password == "" {
		emptyFields = append(emptyFields, "contrasena")
	}

	if len(emptyFields) > 0 {
		return false, emptyFields
	}

	return true, nil
}

func CorreoRegistrado(email string, users map[string]models.User) bool {
	_, registrado := users[email]
	return registrado
}

func TelefonoRegistrado(phone string, users map[string]models.User) bool {
	for _, user := range users {
		if user.Telefono == phone {
			return true
		}
	}
	return false
}

func ValidateLogin(credentials models.Credentials) (bool, []string) {
	var emptyFields []string

	if credentials.UsuarioCorreo == "" {
		emptyFields = append(emptyFields, "usuario")
	}

	if credentials.Password == "" {
		emptyFields = append(emptyFields, "usuario")
	}

	if len(emptyFields) > 0 {
		return false, emptyFields
	}

	return true, nil
}
