package http

import (
	"net/http"

	"atos.com/domain"
	"github.com/gin-gonic/gin"
)

type ShopHandler struct {
	shopService domain.ShopService
}

func RegisterShopHandler(engine *gin.Engine, svc domain.ShopService) {
	h := ShopHandler{
		shopService: svc,
	}

	engine.GET("/search", h.Search)
	engine.GET("/shop/:id", h.GetShop)
}

func (h *ShopHandler) Search(c *gin.Context) {
	var p struct {
		X    string `form:"x"`
		Y    string `form:"y"`
		Size int    `form:"size"`
		Page int    `form:"page"`
	}

	if err := c.ShouldBindQuery(&p); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrBadParamInput)
		return
	}

	if p.Size == 0 {
		p.Size = 20
	}
	if p.Page == 0 {
		p.Page = 1
	}

	shops, err := h.shopService.Search(p.X, p.Y, p.Page, p.Size)
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

	shop, err := h.shopService.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrNotFound)
		return
	}

	c.JSON(http.StatusOK, shop)
}
