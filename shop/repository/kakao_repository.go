package repository

import (
	"atos.com/domain"
	"atos.com/postgreSQL"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type kakaoShopRepository struct {
	key string
}

func NewKaKaoShopRepository() domain.ShopRepository {
	repo := &kakaoShopRepository{
		key: "",
	}

	key, err := postgreSQL.GetKey()
	if err != nil {
		fmt.Println(err)
	}

	repo.key = key

	return repo
}

func (s *kakaoShopRepository) Get(id string) (domain.Shop, error) {
	return domain.Shop{}, domain.ErrNotFound
}

func (s *kakaoShopRepository) Search(x, y string, page, size int) (domain.ShopServiceSearchResponse, error) {
	return s.call(x, y, page, size)
}

func (s *kakaoShopRepository) call(x, y string, page, size int) (domain.ShopServiceSearchResponse, error) {
	// Request 객체 생성
	params := url.Values{}
	params.Add("x", x)
	params.Add("y", y)
	params.Add("page", strconv.Itoa(page))
	params.Add("size", strconv.Itoa(size))
	params.Add("query", "애견미용")

	req, err := http.NewRequest("GET", fmt.Sprintf("https://dapi.kakao.com/v2/local/search/keyword.json?%s", params.Encode()), nil)
	if err != nil {
		return domain.ShopServiceSearchResponse{}, err
	}

	//필요시 헤더 추가 가능
	req.Header.Add("Authorization", "KakaoAK "+s.key)

	// Client객체에서 Request 실행
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 결과 출력
	bytes, _ := ioutil.ReadAll(resp.Body)

	var r struct {
		Documents []domain.Shop          `json:"documents"`
		Meta      map[string]interface{} `json:"meta"`
	}
	fmt.Println("=========== resp ", string(bytes))

	err = json.Unmarshal(bytes, &r)
	if err != nil {
		return domain.ShopServiceSearchResponse{}, err
	}

	totalCount := 0
	t, o := r.Meta["total_count"]
	if o {
		if iv, oo := t.(float64); oo {
			totalCount = int(iv)
		}
	}

	return domain.ShopServiceSearchResponse{
		Shops:      r.Documents,
		TotalCount: totalCount,
	}, nil
}
