package http

import (
	"strconv"
	"toko-api-fiber/internal/model"
	"toko-api-fiber/internal/usecase"

	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
)

type ProductControllerImpl struct {
	Log            *logrus.Logger
	ProductUseCase usecase.ProductUseCase
}

func NewProductController(productUseCase usecase.ProductUseCase, log *logrus.Logger) ProductController {
	return &ProductControllerImpl{
		ProductUseCase: productUseCase,
		Log:            log,
	}
}

func (c *ProductControllerImpl) Create(ctx fiber.Ctx) error {
	request := new(model.CreateProductRequest)

	err := ctx.Bind().JSON(request)
	if err != nil {
		c.Log.WithError(err).Error("Failed to parse request")
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(model.WebResponse{
			Code:   fiber.ErrBadRequest.Code,
			Status: fiber.ErrBadRequest.Message,
			Data:   err.Error(),
		})
	}

	response, err := c.ProductUseCase.Create(ctx.Context(), request)
	if err != nil {
		c.Log.WithError(err).Error("Failed to create product")
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(model.WebResponse{
			Code:   fiber.ErrBadRequest.Code,
			Status: fiber.ErrBadRequest.Message,
			Data:   err.Error(),
		})
	}

	return ctx.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (c *ProductControllerImpl) Update(ctx fiber.Ctx) error {
	request := new(model.UpdateProductRequest)

	err := ctx.Bind().JSON(request)
	if err != nil {
		c.Log.WithError(err).Error("Failed to parse request")
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(model.WebResponse{
			Code:   fiber.ErrBadRequest.Code,
			Status: fiber.ErrBadRequest.Message,
			Data:   err.Error(),
		})
	}

	idParam := ctx.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.Log.WithError(err).Error("Failed to parse ID")
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(model.WebResponse{
			Code:   fiber.ErrBadRequest.Code,
			Status: fiber.ErrBadRequest.Message,
			Data:   err.Error(),
		})
	}

	request.ID = int64(id)

	response, err := c.ProductUseCase.Update(ctx.Context(), request)
	if err != nil {
		c.Log.WithError(err).Error("Failed to update product")
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(model.WebResponse{
			Code:   fiber.ErrBadRequest.Code,
			Status: fiber.ErrBadRequest.Message,
			Data:   err.Error(),
		})
	}

	return ctx.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (c *ProductControllerImpl) Delete(ctx fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.Log.WithError(err).Error("Failed to parse ID")
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(model.WebResponse{
			Code:   fiber.ErrBadRequest.Code,
			Status: fiber.ErrBadRequest.Message,
			Data:   err.Error(),
		})
	}

	err = c.ProductUseCase.Delete(ctx.Context(), int64(id))
	if err != nil {
		c.Log.WithError(err).Error("Failed to delete product")
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(model.WebResponse{
			Code:   fiber.ErrBadRequest.Code,
			Status: fiber.ErrBadRequest.Message,
			Data:   err.Error(),
		})
	}

	return ctx.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	})
}

func (c *ProductControllerImpl) GetAll(ctx fiber.Ctx) error {
	response, err := c.ProductUseCase.GetAll(ctx.Context())
	if err != nil {
		c.Log.WithError(err).Error("Failed to get all products")
		return ctx.Status(fiber.ErrInternalServerError.Code).JSON(model.WebResponse{
			Code:   fiber.ErrInternalServerError.Code,
			Status: fiber.ErrInternalServerError.Message,
			Data:   err.Error(),
		})
	}

	return ctx.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (c *ProductControllerImpl) GetByID(ctx fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.Log.WithError(err).Error("Failed to parse ID")
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(model.WebResponse{
			Code:   fiber.ErrBadRequest.Code,
			Status: fiber.ErrBadRequest.Message,
			Data:   err.Error(),
		})
	}

	response, err := c.ProductUseCase.GetByID(ctx.Context(), int64(id))
	if err != nil {
		c.Log.WithError(err).Error("Failed to get product by ID")
		return ctx.Status(fiber.ErrInternalServerError.Code).JSON(model.WebResponse{
			Code:   fiber.ErrInternalServerError.Code,
			Status: fiber.ErrInternalServerError.Message,
			Data:   err.Error(),
		})
	}

	return ctx.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}
