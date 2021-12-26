package domain

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
