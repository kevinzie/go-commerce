package service

import (
	"github.com/gofiber/fiber/v2"
	"kevinzie/go-commerce/app/models"
	"kevinzie/go-commerce/pkg/utils"
)

func ChangePassword(c *fiber.Ctx) models.BaseResponseModel {
	//ctx := context.Background()
	//user := &models.Users{}
	//if err := c.BodyParser(&user); err != nil {
	//	fmt.Println(err)
	//	return utils.StatusFail("Invalid request")
	//}
	//repo := repository.NewRepository[models.Users](config.Database)
	//params := fiber.Map{
	//	"UserName": "He"
	//}
	//err := repo.Where( , ctx)
	//if err != nil {
	//	return utils.StatusFail("Error occurred while adding user")
	//}
	//return utils.StatusOK("User added successfully")
	return utils.StatusFail("Invalid request")
}
