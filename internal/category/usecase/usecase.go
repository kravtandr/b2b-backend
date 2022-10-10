package categoryUseCase

import (
	chttp "b2b/m/pkg/customhttp"
	"b2b/m/pkg/domain"
	"log"
)

func NewCategoryUseCase(categoryStorage domain.CategoryStorage) domain.CategoryUseCase {
	return categoryUseCase{categoryStorage: categoryStorage}
}

type categoryUseCase struct {
	categoryStorage domain.CategoryStorage
}

// func (c categoryUseCase) GetCategoryById(key string) (value []byte, err error) {
// 	category, err := c.categoryStorage.GetCategoryById(key)
// 	if err != nil {
// 		return []byte{}, err
// 	}
// 	bytes, err := chttp.ApiResp(category, err)
// 	if err != nil {
// 		log.Printf("error while marshalling JSON: %s", err)
// 	}
// 	return bytes, err
// }

// func (c categoryUseCase) GetCategoriesInIndustry(key string) (value []byte, err error) {
// 	category, err := c.categoryStorage.GetCategoriesInIndustry(key)
// 	if err != nil {
// 		return []byte{}, err
// 	}
// 	bytes, err := chttp.ApiResp(category, err)
// 	if err != nil {
// 		log.Printf("error while marshalling JSON: %s", err)
// 	}
// 	return bytes, err
// }

func (c categoryUseCase) GetAllCategories() (value []byte, err error) {
	categories, err := c.categoryStorage.GetAllCategories()
	if err != nil {
		return []byte{}, err
	}
	bytes, err := chttp.ApiResp(categories, err)
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
	}
	return bytes, err
}

func (c categoryUseCase) SearchCategories(param string) (value []byte, err error) {
	categories, err := c.categoryStorage.SearchCategories(param)
	if err != nil {
		return []byte{}, err
	}
	bytes, err := chttp.ApiResp(categories, err)
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
	}
	return bytes, err
}
