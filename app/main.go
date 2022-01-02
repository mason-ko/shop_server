package main

import (
	"atos.com/shop/controller/graphql"
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

	graphql.RegisterGraphQlShopHandler(r, shopSvc)
	http.RegisterShopHandler(r, shopSvc)

	r.Run()
}
