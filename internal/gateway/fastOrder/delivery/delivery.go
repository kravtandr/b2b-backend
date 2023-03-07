package delivery

import (
	"b2b/m/internal/gateway/fastOrder/usecase"
	"b2b/m/internal/models"
	cnst "b2b/m/pkg/constants"
	"b2b/m/pkg/error_adapter"
	"encoding/json"
	"net/http"

	"github.com/valyala/fasthttp"
)

type FastOrderDelivery interface {
	FastOrder(ctx *fasthttp.RequestCtx)
}

type fastOrderDelivery struct {
	errorAdapter error_adapter.HttpAdapter
	manager      usecase.FastOrderUseCase
}

func (u *fastOrderDelivery) FastOrder(ctx *fasthttp.RequestCtx) {
	var request = &models.FastOrderRequest{}
	if err := json.Unmarshal(ctx.Request.Body(), request); err != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.SetBody([]byte(cnst.WrongRequestBody))
		return
	}

	err := u.manager.FastOrder(ctx, request)
	if err != nil {
		httpError := u.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}

	//b, err := chttp.ApiResp({}, err)
	ctx.SetStatusCode(http.StatusOK)
	//ctx.SetBody(b)
}

func NewFastOrderDelivery(
	errorAdapter error_adapter.HttpAdapter,
	manager usecase.FastOrderUseCase,
) FastOrderDelivery {
	return &fastOrderDelivery{
		errorAdapter: errorAdapter,
		manager:      manager,
	}
}
