package main

import (
	"atos.com/shop/controller/http"
	"atos.com/shop/repository"
	"atos.com/shop/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	shopRepo := repository.NewShopRepository()
	shopSvc := service.NewShopService(shopRepo)

	http.RegisterShopHandler(r, shopSvc)

	r.Run()
}
