package utils

import (
	"github.com/gofiber/fiber/v2"
	"kevinzie/go-commerce/app/models"
)

func ResponseSuccess(data any) fiber.Map {
	return fiber.Map{
		"status": "success",
		"data":   data,
	}
}

func ResponseError(data any) fiber.Map {
	return fiber.Map{
		"status": "success",
		"data":   data,
	}
}

// custom response
func NewResponse(status int, data interface{}, message string, success bool) models.BaseResponseModel {
	return models.BaseResponseModel{
		Status:  status,
		Data:    data,
		Message: message,
		Success: success,
	}
}

// returns http 200 OK
func StatusOK(data interface{}, paginate ...any) models.BaseResponseModel {
	if paginate == nil {
		return models.BaseResponseModel{
			Status:  fiber.StatusOK,
			Success: true,
			Message: "OK",
			Data:    data,
		}
	}
	return models.BaseResponseModel{
		Status:    fiber.StatusOK,
		Success:   true,
		Total:     10,
		TotalPage: 0,
		Message:   "OK",
		Data:      data,
	}
}

// returns http 400
func StatusFail(message string) models.BaseResponseModel {
	return models.BaseResponseModel{
		Status:  fiber.StatusBadRequest,
		Message: message,
		Success: false,
	}
}

// retuns http 401
func StatusUnauthorized(message string) models.BaseResponseModel {
	return models.BaseResponseModel{
		Status:  fiber.StatusUnauthorized,
		Message: message,
		Success: false,
	}
}

// returns http 500
func UnhandledError() models.BaseResponseModel {
	return models.BaseResponseModel{
		Status:  fiber.StatusInternalServerError,
		Message: "Unhandled error occurred. Please try again later",
		Success: false,
	}
}

// returns http 404
func StatusNotFound(message string) models.BaseResponseModel {
	return models.BaseResponseModel{
		Status:  fiber.StatusNotFound,
		Message: message,
		Success: false,
	}
}
