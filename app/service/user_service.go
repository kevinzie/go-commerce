package service

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"kevinzie/go-commerce/app/models"
	repository "kevinzie/go-commerce/app/repository"
	"kevinzie/go-commerce/pkg/config"
	"kevinzie/go-commerce/pkg/security"
	"kevinzie/go-commerce/pkg/utils"
	"kevinzie/go-commerce/platform/cache"
)

func GetUsers(c *fiber.Ctx) models.BaseResponseModel {
	ctx := context.Background()
	//now := time.Now().Unix()
	//var (
	//	entity models.Users
	//)
	//redisData := cache.RedisData("get", "user:data", entity)
	//if redisData != nil {
	//	return utils.StatusOK(redisData)
	//}

	modelQuery := config.Database.Preload("Profiles")
	repo := repository.NewRepository[models.Users](modelQuery)
	users, err := repo.GetAll(ctx, c)
	if err != nil {
		fmt.Println(err)
		return utils.StatusFail("Error occurred while getting users")
	}
	cache.RedisData("set", "user:data", users)

	return utils.StatusOK(users)

}

func GetUserById(c *fiber.Ctx) models.BaseResponseModel {
	ctx := context.Background()
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.StatusFail("Invalid request")
	}
	repo := repository.NewRepository[models.Users](config.Database)
	user, err := repo.GetById(id, ctx)
	if user.ID == 0 {
		return utils.StatusNotFound("User not found")
	}
	if err != nil {
		return utils.StatusFail("Error occurred while getting user")
	}
	return utils.StatusOK(user)
}

func CreateUser(c *fiber.Ctx) models.BaseResponseModel {
	ctx := context.Background()
	user := models.Users{}
	//input := models.Users{}

	if err := c.BodyParser(&user); err != nil {
		return utils.StatusFail("Invalid request")
	}
	modelQuery := config.Database
	user.Uuid = uuid.New()
	hasPassword, _ := security.HashPassword(user.Password)
	user.Password = hasPassword
	repo := repository.NewRepository[models.Users](modelQuery)
	err := repo.Add(&user, ctx)
	if err != nil {
		return utils.StatusFail("Error occurred while adding user")
	}

	return utils.StatusOK("User added successfully")
}

func UpdateUser(c *fiber.Ctx) models.BaseResponseModel {
	ctx := context.Background()
	user := &models.Users{}
	id := c.Params("id")
	if err := c.BodyParser(&user); err != nil {
		fmt.Println(err)
		return utils.StatusFail("Invalid request")
	}
	repo := repository.NewRepository[models.Users](config.Database)
	err := repo.Update(id, user, ctx)
	if err != nil {
		return utils.StatusFail("Error occurred while adding user")
	}
	return utils.StatusOK("User added successfully")
}

func DeleteUser(c *fiber.Ctx) models.BaseResponseModel {
	ctx := context.Background()
	id, err := c.ParamsInt("id")
	if err != nil {
		return utils.StatusFail("Invalid request")
	}
	repo := repository.NewRepository[models.Users](config.Database)
	err = repo.Delete(id, ctx)
	if err != nil {
		return utils.StatusFail("Error occurred while deleting user")
	}

	return utils.StatusOK("User deleted successfully")
}
