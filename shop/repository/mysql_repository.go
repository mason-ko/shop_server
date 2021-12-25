package repository

import (
	"encoding/json"
	"strconv"
	"strings"

	"atos.com/domain"
	"atos.com/static"
)

type ShopRepository struct {
	shopsById map[string]domain.Shop
	shops     map[string][]domain.Shop
}

func NewShopRepository() domain.ShopRepository {
	repo := &ShopRepository{
		shopsById: map[string]domain.Shop{},
		shops:     map[string][]domain.Shop{},
	}

	shops1 := loadShops(static.Shops1)
	shops2 := loadShops(static.Shops2)
	shops3 := loadShops(static.Shops3)

	for _, v := range shops1 {
		repo.shopsById[v.Id] = v
		key := getPath(v.X, 6) + getPath(v.Y, 5)
		if _, o := repo.shops[key]; !o {
			repo.shops[key] = []domain.Shop{v}
		} else {
			repo.shops[key] = append(repo.shops[key], v)
		}
	}
	for _, v := range shops2 {
		repo.shopsById[v.Id] = v
		key := getPath(v.X, 6) + getPath(v.Y, 5)
		if _, o := repo.shops[key]; !o {
			repo.shops[key] = []domain.Shop{v}
		} else {
			repo.shops[key] = append(repo.shops[key], v)
		}
	}
	for _, v := range shops3 {
		repo.shopsById[v.Id] = v
		key := getPath(v.X, 6) + getPath(v.Y, 5)
		if _, o := repo.shops[key]; !o {
			repo.shops[key] = []domain.Shop{v}
		} else {
			repo.shops[key] = append(repo.shops[key], v)
		}
	}

	return repo
}

func loadShops(bytes string) []domain.Shop {
	shops := []domain.Shop{}
	json.Unmarshal([]byte(bytes), &shops)
	return shops
}

func getPath(x string, d int) string {
	x = strings.ReplaceAll(x, ".", "")
	return x[:d]
}

func (s *ShopRepository) Get(id string) (domain.Shop, error) {
	v, o := s.shopsById[id]
	if !o {
		return domain.Shop{}, domain.ErrNotFound
	}
	return v, nil
}

func (s *ShopRepository) Search(x, y string, limit int) ([]domain.Shop, error) {
	xKey, _ := strconv.Atoi(getPath(x, 6))
	yKey, _ := strconv.Atoi(getPath(y, 5))
	shops := s.bfsSearch(xKey, yKey, limit)
	return shops, nil
}

func (s *ShopRepository) bfsSearch(xKey, yKey int, limit int) []domain.Shop {
	type tuple struct {
		x int
		y int
	}

	q := []tuple{}
	q = append(q, tuple{
		x: xKey,
		y: yKey,
	})

	oldPath := map[string]bool{}
	shops := []domain.Shop{}

	for len(q) > 0 && len(shops) < limit {
		k := q[0]
		q = q[1:]

		key := strconv.Itoa(k.x) + strconv.Itoa(k.y)
		if _, o := oldPath[key]; o {
			continue
		}
		oldPath[key] = true
		items := s.shops[key]
		if len(items) > 0 {
			if len(items)+len(shops) <= limit {
				shops = append(shops, items...)
			} else {
				cnt := limit - len(shops)
				shops = append(shops, items[:cnt]...)
			}
		}

		q = append(q, tuple{x: k.x, y: k.y + 1})
		q = append(q, tuple{x: k.x + 1, y: k.y})
		q = append(q, tuple{x: k.x + 1, y: k.y + 1})
		q = append(q, tuple{x: k.x, y: k.y - 1})
		q = append(q, tuple{x: k.x - 1, y: k.y})
		q = append(q, tuple{x: k.x - 1, y: k.y - 1})
	}

	return shops
}
