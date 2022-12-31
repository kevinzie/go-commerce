package service

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"kevinzie/go-commerce/app/models"
	"kevinzie/go-commerce/app/repository"
	"kevinzie/go-commerce/pkg/config"
	"kevinzie/go-commerce/pkg/security"
	"kevinzie/go-commerce/pkg/utils"
	"log"
	"strconv"
)

//type SigninData struct {
//	Email    string `json:"email" xml:"email" form:"email"`
//	Password string `json:"password" xml:"password" form:"password"`
//}

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

func Login(c *fiber.Ctx) models.BaseResponseModel {

	ctx := context.Background()
	input := models.Users{}
	if fuckerr := c.BodyParser(&input); fuckerr != nil {
		return utils.StatusFail(fuckerr.Error())
	}
	repo := repository.NewRepository[models.Users](config.Database)
	input.Email = utils.NormalizeEmail(input.Email)
	user, err := repo.GetByEmail(input.Email, ctx)
	if err != nil {
		return utils.StatusFail(utils.ErrInvalidEmail.Error())
	}

	if user.ID == 0 {
		return utils.StatusFail("Account not found or not activated")
	}

	fmt.Println("tester:", user)

	verifyPassword := security.CheckPasswordHash(input.Password, user.Password)
	if verifyPassword == false {
		//log.Printf(err.Error())
		return utils.StatusFail(utils.ErrInvalidCredentials.Error())
	}

	token, err := security.NewToken(strconv.Itoa(int(user.ID)))
	if err != nil {
		log.Printf("%s signin failed: %v\n", input.Email, err.Error())
		return utils.StatusUnauthorized(err.Error())
	}

	return utils.StatusOK(fiber.Map{
		"user":  user,
		"token": token,
	})
	//if err != nil {
	//	log.Printf("%s signin failed: %v\n", input.Email, err.Error())
	//	return ctx.
	//		Status(http.StatusUnauthorized).
	//		JSON(util.NewJError(util.ErrInvalidCredentials))
	//}
	//if err != nil {
	//	log.Printf("%s signin failed: %v\n", input.Email, err.Error())
	//	return ctx.
	//		Status(http.StatusUnauthorized).
	//		JSON(util.NewJError(util.ErrInvalidCredentials))
	//}
	//return ctx.
	//	Status(http.StatusOK).
	//	JSON(fiber.Map{
	//		"user":  user,
	//		"token": fmt.Sprintf("Bearer %s", token),
	//	})
}
