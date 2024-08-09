package delivery

import (
	"b2b/m/internal/gateway/company/usecase"
	chttp "b2b/m/pkg/customhttp"
	"b2b/m/pkg/error_adapter"
	"net/http"
	"strconv"

	"github.com/valyala/fasthttp"
)

type CompanyDelivery interface {
	GetCompanyById(ctx *fasthttp.RequestCtx)
	GetCompanyByItnFromDaData(ctx *fasthttp.RequestCtx)
}

type companyDelivery struct {
	errorAdapter error_adapter.HttpAdapter
	manager      usecase.CompanyUseCase
}

func (c *companyDelivery) GetCompanyByItnFromDaData(ctx *fasthttp.RequestCtx) {
	param, _ := ctx.UserValue("itn").(string)
	response, err := c.manager.GetCompanyByItnFromDaData(ctx, param)
	if err != nil {
		httpError := c.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}
	b, err := chttp.ApiResp(response, err)
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(b)
}

func (c *companyDelivery) GetCompanyById(ctx *fasthttp.RequestCtx) {
	param, _ := strconv.Atoi(ctx.UserValue("id").(string))
	response, err := c.manager.GetCompanyById(ctx, int64(param))
	if err != nil {
		httpError := c.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}
	b, err := chttp.ApiResp(response, err)
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(b)
}

// func (c *companyDelivery) GetCompanyByUserCookie(ctx *fasthttp.RequestCtx) {
// 	response, err := c.manager.GetCompanyByUserCookie(ctx, string(ctx.Request.Header.Cookie(cnst.CookieName)))
// 	if err != nil {
// 		httpError := c.errorAdapter.AdaptError(err)
// 		ctx.SetStatusCode(httpError.Code)
// 		ctx.SetBody([]byte(httpError.MSG))
// 		return
// 	}
// 	b, err := chttp.ApiResp(response, err)
// 	ctx.SetStatusCode(http.StatusOK)
// 	ctx.SetBody(b)
// }

func NewCompanyDelivery(
	errorAdapter error_adapter.HttpAdapter,
	manager usecase.CompanyUseCase,
) CompanyDelivery {
	return &companyDelivery{
		errorAdapter: errorAdapter,
		manager:      manager,
	}
}
