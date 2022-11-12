package delivery

import (
	"b2b/m/internal/gateway/productsCategories/usecase"
	"b2b/m/internal/models"
	cnst "b2b/m/pkg/constants"
	chttp "b2b/m/pkg/customhttp"
	"b2b/m/pkg/error_adapter"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"net/http"
)

type ProductsCategoriesDelivery interface {
	GetAllcategories(ctx *fasthttp.RequestCtx)
}

type productsCategoriesDelivery struct {
	errorAdapter error_adapter.HttpAdapter
	manager      usecase.ProductsCategoriesUseCase
}

func (u *productsCategoriesDelivery) GetAllcategories(ctx *fasthttp.RequestCtx) {
	fmt.Println("kfljdblkqsdjbfhlkdjqsbflkdsefblkdsafbljasdbfh")
	var request = &models.GetCategorieByIdRequest{}
	if err := json.Unmarshal(ctx.Request.Body(), request); err != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.SetBody([]byte(cnst.WrongRequestBody))
		return
	}

	response, err := u.manager.GetAllcategories(ctx, request)
	if err != nil {
		httpError := u.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}

	b, err := chttp.ApiResp(response, err)
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(b)
}

func NewProductsCategoriesDelivery(
	errorAdapter error_adapter.HttpAdapter,
	manager usecase.ProductsCategoriesUseCase,
) ProductsCategoriesDelivery {
	return &productsCategoriesDelivery{
		errorAdapter: errorAdapter,
		manager:      manager,
	}
}
