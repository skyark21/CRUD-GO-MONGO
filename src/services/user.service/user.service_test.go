package user_service_test

import (
	"testing"
	"time"

	m "../../models"
	userService "../user.service"
)

func TestCreate(t *testing.T) {
	user := m.User{
		Name:      "Luis",
		Email:     "luis.sanchez@gmail.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := userService.Create(user)
	if err != nil {
		t.Error("La prueba de creación de datos de usuario ha fallado...")
		t.Fail()
	} else {
		t.Log("La prueba de cración de datos de usaurio finalizó con éxito...")
	}
}
func TestRead(t *testing.T) {

}
func TestUpdate(t *testing.T) {

}
func TestDelete(t *testing.T) {

}
