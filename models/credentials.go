package models

// Struct para Validar las credenciales introducidas
type Credentials struct {
	UsuarioCorreo string `json:"usuario_correo"`
	Password      string `json:"contrasena"`
}
