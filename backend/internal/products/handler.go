package products

import (
	"net/http"
	"shopsocial-backend/pkg/constants"
	"shopsocial-backend/pkg/logger"
	"shopsocial-backend/pkg/responses"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
)

type ProductHandler struct {
	Service *ProductService
}

func NewProductHandler(service *ProductService) *ProductHandler {
	return &ProductHandler{Service: service}
}

// CREATE
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		responses.SendError(c, http.StatusBadRequest, constants.ErrInvalidRequest, err)
		return
	}

	createdProduct, err := h.Service.CreateProduct(&product)
	if err != nil {
		responses.SendError(c, http.StatusInternalServerError, constants.FormatMessage(constants.ErrCreationFailed, constants.EntityProduct), err)
		return
	}

	logger.Log.Info("Product created", zap.String("id", createdProduct.ID.Hex()))
	responses.SendCreated(c, constants.FormatMessage(constants.SuccessCreated, constants.EntityProduct), createdProduct)
}

// READ
func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	logger.Log.Info("Fetching product", zap.String("id", id))

	product, err := h.Service.GetProductByID(id)
	if err != nil {
		responses.SendError(c, http.StatusNotFound, constants.FormatMessage(constants.ErrNotFound, constants.EntityProduct), err)
		return
	}

	responses.SendSuccess(c, constants.FormatMessage(constants.SuccessFetched, constants.EntityProduct), product)
}

// UPDATE
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var updateData bson.M
	if err := c.ShouldBindJSON(&updateData); err != nil {
		responses.SendError(c, http.StatusBadRequest, constants.ErrInvalidRequest, err)
		return
	}

	updatedProduct, err := h.Service.UpdateProduct(id, updateData)
	if err != nil {
		responses.SendError(c, http.StatusInternalServerError, constants.FormatMessage(constants.ErrUpdateFailed, constants.EntityProduct), err)
		return
	}

	logger.Log.Info("Product updated", zap.String("id", id))
	responses.SendSuccess(c, constants.FormatMessage(constants.SuccessUpdated, constants.EntityProduct), updatedProduct)
}

// DELETE
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	logger.Log.Info("Deleting product", zap.String("id", id))

	err := h.Service.DeleteProduct(id)
	if err != nil {
		responses.SendError(c, http.StatusInternalServerError, constants.FormatMessage(constants.ErrDeletionFailed, constants.EntityProduct), err)
		return
	}

	logger.Log.Info("Product deleted", zap.String("id", id))
	responses.SendDeleted(c, constants.FormatMessage(constants.SuccessDeleted, constants.EntityProduct))
}
