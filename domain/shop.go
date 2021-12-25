package domain

type Shop struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	Category   []string `json:"category"`
	Address    string   `json:"address"`
	Display    string   `json:"display"`
	TelDisplay string   `json:"telDisplay"`
	ThumUrl    string   `json:"thumUrl"`

	X string `json:"x"`
	Y string `json:"y"`
}

type ShopService interface {
	Get(id string) (Shop, error)
	Search(x, y string, limit int) ([]Shop, error)
}

type ShopRepository interface {
	Get(id string) (Shop, error)
	Search(x, y string, limit int) ([]Shop, error)
}
