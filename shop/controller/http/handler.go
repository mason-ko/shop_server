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
	var p struct {
		X     string `form:"x"`
		Y     string `form:"y"`
		Limit int    `form:"limit"`
	}

	if err := c.ShouldBindQuery(&p); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrBadParamInput)
		return
	}

	if p.Limit == 0 {
		p.Limit = 20
	}

	shops, err := h.shopServcie.Search(p.X, p.Y, p.Limit)
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
