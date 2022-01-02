package graphql

import (
	"atos.com/domain"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	handler "github.com/graphql-go/graphql-go-handler"
)

type ShopHandler struct {
	shopServcie domain.ShopService
}

func RegisterGraphQlShopHandler(engine *gin.Engine, svc domain.ShopService) {
	h := ShopHandler{
		shopServcie: svc,
	}

	var queryType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"shops": &graphql.Field{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"x": &graphql.ArgumentConfig{
						Type:         graphql.String,
						DefaultValue: nil,
						Description:  "",
					},
					"y": &graphql.ArgumentConfig{
						Type:         graphql.String,
						DefaultValue: nil,
						Description:  "",
					},
					"page": &graphql.ArgumentConfig{
						Type:         graphql.Int,
						DefaultValue: 0,
						Description:  "",
					},
					"size": &graphql.ArgumentConfig{
						Type:         graphql.Int,
						DefaultValue: 20,
						Description:  "",
					},
				},
				Resolve: h.Search,
			},
			"shop": &graphql.Field{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type:         graphql.String,
						DefaultValue: nil,
						Description:  "",
					},
				},
				Resolve: h.GetShop,
			},
		},
	})

	var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		// ...
		Query: queryType,
	})

	gqlHandler := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	engine.POST("/graphql", func(c *gin.Context) {
		gqlHandler.ServeHTTP(c.Writer, c.Request)
	})
}

func (h *ShopHandler) Search(p graphql.ResolveParams) (interface{}, error) {
	x, o := p.Args["x"].(string)
	if !o {
		return nil, domain.ErrBadParamInput
	}
	y, o := p.Args["y"].(string)
	if !o {
		return nil, domain.ErrBadParamInput
	}
	page, o := p.Args["page"].(int)
	if !o {
		return nil, domain.ErrBadParamInput
	}
	size, o := p.Args["size"].(int)
	if !o {
		return nil, domain.ErrBadParamInput
	}

	shops, err := h.shopServcie.Search(x, y, page, size)
	if err != nil {
		return nil, err
	}
	return shops, nil
}

func (h *ShopHandler) GetShop(p graphql.ResolveParams) (interface{}, error) {
	//id := c.Param("id")
	//if id == "" {
	//	c.JSON(http.StatusBadRequest, domain.ErrBadParamInput)
	//	return
	//}
	//
	//shop, err := h.shopServcie.Get(id)
	//if err != nil {
	//	c.JSON(http.StatusNotFound, domain.ErrNotFound)
	//	return
	//}
	//
	//c.JSON(http.StatusOK, shop)
	return nil, nil
}
