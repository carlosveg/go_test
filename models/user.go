package models

// Forma con la que se manejara la info del usuario
type User struct {
	Username string `json:"usuario"`
	Correo   string `json:"correo"`
	Telefono string `json:"telefono"`
	Password string `json:"contrasena"`
}
