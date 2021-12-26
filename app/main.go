package main

import (
	"fmt"
	"path/filepath"
	"runtime"

	"atos.com/shop/controller/http"
	"atos.com/shop/repository"
	"atos.com/shop/service"
	"github.com/gin-gonic/gin"
)

func main() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	fmt.Println("load basepath " + basepath)

	r := gin.Default()

	shopRepo := repository.NewKaKaoShopRepository()
	shopSvc := service.NewShopService(shopRepo)

	http.RegisterShopHandler(r, shopSvc)

	r.Run()
}
