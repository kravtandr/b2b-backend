package delivery

import (
	"encoding/json"
	"net/http"
	"path/filepath"

	"b2b/m/internal/gateway/media/usecase"
	"b2b/m/internal/models"
	"b2b/m/pkg/constants"
	"b2b/m/pkg/error_adapter"

	"github.com/valyala/fasthttp"
)

type MediaDelivery interface {
	UploadFile(ctx *fasthttp.RequestCtx)
}

type mediaDelivery struct {
	manager      usecase.MediaUsecase
	errorAdapter error_adapter.HttpAdapter
}

func (m *mediaDelivery) UploadFile(ctx *fasthttp.RequestCtx) {
	form, err := ctx.FormFile(constants.FileKey)
	if err != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		return
	}
	file, err := form.Open()
	if err != nil {
		ctx.SetStatusCode(http.StatusConflict)
		return
	}
	defer file.Close()

	filename, err := m.manager.UploadFile(ctx, file, filepath.Ext(form.Filename))
	if err != nil {
		ctx.SetStatusCode(http.StatusTeapot)
		return
	}

	b, _ := json.Marshal(models.UploadFileResponse{Filename: filename})
	ctx.Response.SetBody(b)
	ctx.SetStatusCode(http.StatusOK)
}

func NewMediaDelivery(
	manager usecase.MediaUsecase,
	errorAdapter error_adapter.HttpAdapter,
) MediaDelivery {
	return &mediaDelivery{
		manager:      manager,
		errorAdapter: errorAdapter,
	}
}
