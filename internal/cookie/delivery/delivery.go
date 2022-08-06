package cookieDelivery

import (
	cr "b2b/m/internal/cookie/repository"
	cu "b2b/m/internal/cookie/usecase"
	ent "b2b/m/internal/entities"
	cnst "b2b/m/pkg/constants"
	chttp "b2b/m/pkg/customhttp"
	"b2b/m/pkg/domain"
	"encoding/json"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
)

type CookieHandler interface {
	SetCookieAndToken(ctx *fasthttp.RequestCtx, cookie string, id int)
	setCookie(ctx *fasthttp.RequestCtx, cookie string, id int)
	DeleteCookie(ctx *fasthttp.RequestCtx, cookie string)
}

type cookieHandler struct {
	CookieUseCase domain.CookieUseCase
}

func NewCookieHandler(CookieUseCase domain.CookieUseCase) CookieHandler {
	return &cookieHandler{
		CookieUseCase: CookieUseCase,
	}
}

func CreateDelivery(db *pgxpool.Pool) CookieHandler {
	cookieLayer := NewCookieHandler(cu.NewCookieUseCase(cr.NewCookieStorage(db)))
	return cookieLayer
}

func (s *cookieHandler) SetCookieAndToken(ctx *fasthttp.RequestCtx, cookie string, id int) {
	s.setCookie(ctx, cookie, id)
	setToken(ctx, cookie)
}

func (s *cookieHandler) setCookie(ctx *fasthttp.RequestCtx, cookie string, id int) {
	var c fasthttp.Cookie
	c.SetKey(cnst.CookieName)
	c.SetValue(cookie)
	c.SetMaxAge(int(time.Hour))
	c.SetHTTPOnly(true)
	c.SetSameSite(fasthttp.CookieSameSiteStrictMode)
	ctx.Response.Header.SetCookie(&c)

	s.CookieUseCase.Add(cookie, id)
}

func (s *cookieHandler) DeleteCookie(ctx *fasthttp.RequestCtx, cookie string) {
	var c fasthttp.Cookie
	c.SetKey(cnst.CookieName)
	c.SetValue("")
	c.SetMaxAge(0)
	c.SetExpire(time.Now().Add(-1))
	c.SetSameSite(fasthttp.CookieSameSiteStrictMode)
	ctx.Response.Header.SetCookie(&c)

	s.CookieUseCase.Delete(cookie)
}

func setToken(ctx *fasthttp.RequestCtx, hash string) {
	t := ent.Token{Token: hash}
	r := chttp.Resp{Data: t, Msg: "OK"}
	bytes, err := json.Marshal(r)
	if err != nil {
		r.Msg = "error while marshalling JSON"
		log.Printf("error while marshalling JSON: %s", err)
	}
	ctx.Write(bytes)
}
