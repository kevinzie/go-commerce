package controllers

import (
	"github.com/gofiber/fiber/v2"
	"kevinzie/go-commerce/app/service"
)

// GetUser func gets users by given ID or 404 error.
// @Description Get users by given ID.
// @Summary get users by given ID
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.Users
// @Router /v1/user/{id} [get]
func GetUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(service.GetUserById(c))
}

// GetUsers func get all users.
// @Description Get all User.
// @Summary get users
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} models.Users
// @Router /v1/users [get]
func GetUsers(c *fiber.Ctx) error {
	return c.Status(200).JSON(service.GetUsers(c))
}

// CreateUser func Create user.
// @Description Create user.
// @Summary Creat new user
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} models.Users
// @Router /v1/user [post]
func CreateUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(service.CreateUser(c))
}

// DeleteUser func Delete user.
// @Description Delete user.
// @Summary Delete new user
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} models.Users
// @Router /v1/user [delete]
func DeleteUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(service.DeleteUser(c))
}

// UpdateUser func Update user.
// @Description Update user.
// @Summary Update new user
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} models.Users
// @Router /v1/user [put]
func UpdateUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(service.UpdateUser(c))
}
