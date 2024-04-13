package data

import "github.com/carlosveg/go_test/models"

var Users = map[string]models.User{
	"test1@test.com": {
		Username: "usuario1",
		Correo:   "test1@test.com",
		Telefono: "1234567890",
		Password: "Abc@1234",
	},
	"test2@test.com": {
		Username: "usuario2",
		Correo:   "test2@test.com",
		Telefono: "9876543210",
		Password: "Abc@123",
	},
	"test3@test.com": {
		Username: "usuario3",
		Correo:   "test3@test.com",
		Telefono: "9876543210",
		Password: "Abc@123",
	},
}
