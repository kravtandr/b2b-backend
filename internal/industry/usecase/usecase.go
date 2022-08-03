package industryUseCase

import (
	chttp "b2b/m/pkg/customhttp"
	"b2b/m/pkg/domain"
	"log"
)

func NewIndustryUseCase(industryStorage domain.IndustryStorage) domain.IndustryUseCase {
	return industryUseCase{industryStorage: industryStorage}
}

type industryUseCase struct {
	industryStorage domain.IndustryStorage
}

func (c industryUseCase) GetAllIndustries() (value []byte, err error) {
	category, err := c.industryStorage.GetAllIndustries()
	if err != nil {
		return []byte{}, err
	}
	bytes, err := chttp.ApiResp(category)
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
	}
	return bytes, err
}
