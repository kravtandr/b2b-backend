package customhttp

import (
	ent "b2b/m/internal/entities"
	cnst "b2b/m/pkg/constants"
	"encoding/json"
	"log"
	"time"

	"github.com/valyala/fasthttp"
)

func SetCookieAndToken(ctx *fasthttp.RequestCtx, cookie string, id int) {
	setCookie(ctx, cookie, id)
	setToken(ctx, cookie)
}

func setCookie(ctx *fasthttp.RequestCtx, cookie string, id int) {
	var c fasthttp.Cookie
	c.SetKey(cnst.CookieName)
	c.SetValue(cookie)
	c.SetMaxAge(int(time.Hour))
	c.SetHTTPOnly(true)
	c.SetSameSite(fasthttp.CookieSameSiteStrictMode)
	ctx.Response.Header.SetCookie(&c)
	ent.CookieDB[cookie] = id
}

func DeleteCookie(ctx *fasthttp.RequestCtx, cookie string) {
	var c fasthttp.Cookie
	c.SetKey(cnst.CookieName)
	c.SetValue("")
	c.SetMaxAge(0)
	c.SetExpire(time.Now().Add(-1))
	c.SetSameSite(fasthttp.CookieSameSiteStrictMode)
	ctx.Response.Header.SetCookie(&c)

	delete(ent.CookieDB, cookie)
}

func setToken(ctx *fasthttp.RequestCtx, hash string) {
	t := ent.Token{Token: hash}
	bytes, err := json.Marshal(t)

	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		ctx.Write([]byte("{}"))
		return
	}

	ctx.Write(bytes)
}

func CheckCookie(ctx *fasthttp.RequestCtx) bool {
	cookieHash := string(ctx.Request.Header.Cookie(cnst.CookieName))
	if _, found := ent.CookieDB[cookieHash]; !found {
		return false
	}
	return true
}
