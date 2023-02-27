package delivery

import (
	"b2b/m/internal/gateway/productsCategories/usecase"
	"b2b/m/internal/models"
	cnst "b2b/m/pkg/constants"
	chttp "b2b/m/pkg/customhttp"
	"b2b/m/pkg/error_adapter"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/valyala/fasthttp"
)

type ProductsCategoriesDelivery interface {
	GetCategoryById(ctx *fasthttp.RequestCtx)
	SearchCategories(ctx *fasthttp.RequestCtx)
	SearchProducts(ctx *fasthttp.RequestCtx)
	GetProductsList(ctx *fasthttp.RequestCtx)
}

type productsCategoriesDelivery struct {
	errorAdapter error_adapter.HttpAdapter
	manager      usecase.ProductsCategoriesUseCase
}

func (u *productsCategoriesDelivery) GetCategoryById(ctx *fasthttp.RequestCtx) {
	var request = &models.GetCategoryByIdRequest{}
	if err := json.Unmarshal(ctx.Request.Body(), request); err != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.SetBody([]byte(cnst.WrongRequestBody))
		return
	}

	response, err := u.manager.GetCategoryById(ctx, request)
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

func (u *productsCategoriesDelivery) GetProductsList(ctx *fasthttp.RequestCtx) {
	params, err := chttp.GetQueryParams(ctx)
	if err != nil {
		fmt.Println("PIZDEC V GetProductsList", err)
		httpError := u.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}

	response, err := u.manager.GetProductsList(ctx, params)
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

func (u *productsCategoriesDelivery) SearchProducts(ctx *fasthttp.RequestCtx) {
	params, err := chttp.GetQueryParams(ctx)
	if err != nil {
		httpError := u.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}
	var request = &chttp.SearchItemName{}
	if err := json.Unmarshal(ctx.Request.Body(), request); err != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.SetBody([]byte(cnst.WrongRequestBody))
		return
	}
	var searchBody = &chttp.SearchItemNameWithSkipLimit{
		Name:  request.Name,
		Skip:  params.Skip,
		Limit: params.Limit,
	}
	response, err := u.manager.SearchProducts(ctx, searchBody)
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

func (u *productsCategoriesDelivery) SearchCategories(ctx *fasthttp.RequestCtx) {
	var request = &chttp.SearchItemName{}
	if err := json.Unmarshal(ctx.Request.Body(), request); err != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.SetBody([]byte(cnst.WrongRequestBody))
		return
	}

	response, err := u.manager.SearchCategories(ctx, request)
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
