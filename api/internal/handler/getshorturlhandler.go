package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"shorturl-service/api/internal/logic"
	"shorturl-service/api/internal/svc"
	"shorturl-service/api/internal/types"
)

func GetShortUrlHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetShortUrlRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetShortUrlLogic(r.Context(), svcCtx)
		resp, err := l.GetShortUrl(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
