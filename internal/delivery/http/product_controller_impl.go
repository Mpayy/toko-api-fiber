package http

import (
	"errors"
	"strconv"
	"toko-api-fiber/internal/exception"
	"toko-api-fiber/internal/model"
	"toko-api-fiber/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
)

type ProductControllerImpl struct {
	Validate       *validator.Validate
	Log            *logrus.Logger
	ProductUsecase usecase.ProductUsecase
}

func NewProductController(productUsecase usecase.ProductUsecase, log *logrus.Logger, validate *validator.Validate) ProductController {
	return &ProductControllerImpl{
		ProductUsecase: productUsecase,
		Log:            log,
		Validate:       validate,
	}
}

func (c *ProductControllerImpl) Create(ctx fiber.Ctx) error {
	request := new(model.CreateProductRequest)

	err := ctx.Bind().JSON(request)
	if err != nil {
		c.Log.WithError(err).Warn("Failed to parse request")
		return fiber.ErrBadRequest
	}

	err = c.Validate.Struct(request)
	if err != nil {
		fieldErrors := exception.ExtractValidationErrors(err)
		c.Log.WithFields(logrus.Fields{
			"errors": fieldErrors,
		}).Warn("Validation failed for create product")
		return &exception.ValidationErrorWithFields{
			Errors: fieldErrors,
		}
	}

	response, err := c.ProductUsecase.Create(ctx.Context(), request)
	if err != nil {
		return err
	}

	return ctx.JSON(model.WebResponse[*model.ProductResponse]{Data: response})
}

func (c *ProductControllerImpl) Update(ctx fiber.Ctx) error {
	request := new(model.UpdateProductRequest)

	err := ctx.Bind().JSON(request)
	if err != nil {
		c.Log.WithError(err).Warn("Failed to parse request")
		return fiber.ErrBadRequest
	}

	idParam := ctx.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.Log.WithError(err).Warn("Failed to parse ID")
		return fiber.ErrBadRequest
	}

	request.ID = int64(id)

	err = c.Validate.Struct(request)
	if err != nil {
		fieldErrors := exception.ExtractValidationErrors(err)
		c.Log.WithFields(logrus.Fields{
			"errors": fieldErrors,
		}).Warn("Validation failed for update product")
		return &exception.ValidationErrorWithFields{
			Errors: fieldErrors,
		}
	}

	response, err := c.ProductUsecase.Update(ctx.Context(), request)
	if err != nil {
		if errors.Is(err, exception.ErrNotFound) {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse[*model.ProductResponse]{Data: response})
}

func (c *ProductControllerImpl) Delete(ctx fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.Log.WithError(err).Warn("Failed to parse ID")
		return fiber.ErrBadRequest
	}

	err = c.ProductUsecase.Delete(ctx.Context(), int64(id))
	if err != nil {
		if errors.Is(err, exception.ErrNotFound) {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse[bool]{Data: true})
}

func (c *ProductControllerImpl) GetAll(ctx fiber.Ctx) error {
	page := fiber.Query(ctx, "page", 1)
	size := fiber.Query(ctx, "size", 10)

	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 1
	}

	request := &model.PaginationRequest{
		Page: &page,
		Size: &size,
	}

	response, totalItems, totalPage, err := c.ProductUsecase.GetAll(ctx.Context(), request)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse[[]*model.ProductResponse]{
		Data: response,
		Paging: &model.PaginationResponse{
			Page:       *request.Page,
			TotalPages: totalPage,
			TotalItems: totalItems,
		},
	})
}

func (c *ProductControllerImpl) GetByID(ctx fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.Log.WithError(err).Warn("Failed to parse ID")
		return fiber.ErrBadRequest
	}

	response, err := c.ProductUsecase.GetByID(ctx.Context(), int64(id))
	if err != nil {
		if errors.Is(err, exception.ErrNotFound) {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse[*model.ProductResponse]{Data: response})
}

func (c *ProductControllerImpl) Patch(ctx fiber.Ctx) error {
	request := new(model.PatchProductRequest)

	err := ctx.Bind().JSON(request)
	if err != nil {
		c.Log.WithError(err).Warn("Failed to parse request")
		return fiber.ErrBadRequest
	}

	idParam := ctx.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.Log.WithError(err).Warn("Failed to parse ID")
		return fiber.ErrBadRequest
	}

	request.ID = int64(id)

	err = c.Validate.Struct(request)
	if err != nil {
		fieldErrors := exception.ExtractValidationErrors(err)
		c.Log.WithFields(logrus.Fields{
			"errors": fieldErrors,
		}).Warn("Validation failed for patch product")
		return &exception.ValidationErrorWithFields{
			Errors: fieldErrors,
		}
	}

	response, err := c.ProductUsecase.Patch(ctx.Context(), request)
	if err != nil {
		if errors.Is(err, exception.ErrNotFound) {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse[*model.ProductResponse]{Data: response})
}
