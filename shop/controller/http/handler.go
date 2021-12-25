package http

import (
	"net/http"

	"atos.com/domain"
	"github.com/gin-gonic/gin"
)

type ShopHandler struct {
	shopServcie domain.ShopService
}

func RegisterShopHandler(engine *gin.Engine, svc domain.ShopService) {
	h := ShopHandler{
		shopServcie: svc,
	}

	engine.GET("/search", h.Search)
	engine.GET("/shop/:id", h.GetShop)
}

func (h *ShopHandler) Search(c *gin.Context) {
	x := c.Param("x")
	y := c.Param("y")
	if x == "" || y == "" {
		c.JSON(http.StatusBadRequest, domain.ErrBadParamInput)
		return
	}

	shops, err := h.shopServcie.Search(x, y)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, shops)
}

func (h *ShopHandler) GetShop(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, domain.ErrBadParamInput)
		return
	}

	shop, err := h.shopServcie.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrNotFound)
		return
	}

	c.JSON(http.StatusOK, shop)
}
