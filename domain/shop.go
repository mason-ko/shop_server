package domain

import "github.com/graphql-go/graphql"

type Shop struct {
	Id              string `json:"id"`
	PlaceName       string `json:"place_name"`
	AddressName     string `json:"address_name"`
	CategoryName    string `json:"category_name"`
	Address         string `json:"address"`
	RoadAddressName string `json:"road_address_name"`

	Phone    string `json:"phone"`
	PlaceUrl string `json:"place_url"`

	X string `json:"x"`
	Y string `json:"y"`
}

var ShopType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Shop",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"place_name": &graphql.Field{
			Type: graphql.String,
		},
		"address_name": &graphql.Field{
			Type: graphql.String,
		},
		"category_name": &graphql.Field{
			Type: graphql.String,
		},
		"address": &graphql.Field{
			Type: graphql.String,
		},
		"road_address_name": &graphql.Field{
			Type: graphql.String,
		},
		"phone": &graphql.Field{
			Type: graphql.String,
		},
		"place_url": &graphql.Field{
			Type: graphql.String,
		},
		"x": &graphql.Field{
			Type: graphql.String,
		},
		"y": &graphql.Field{
			Type: graphql.String,
		},
	},
})

type ShopServiceSearchResponse struct {
	Shops      []Shop `json:"shops"`
	TotalCount int    `json:"total_count"`
}

type ShopService interface {
	Get(id string) (Shop, error)
	Search(x, y string, page, size int) (ShopServiceSearchResponse, error)
}

type ShopRepository interface {
	Get(id string) (Shop, error)
	Search(x, y string, page, size int) (ShopServiceSearchResponse, error)
}
