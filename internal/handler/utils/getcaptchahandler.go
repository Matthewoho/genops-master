package utils

import (
	"net/http"

	"genops-master/internal/biz"
	"genops-master/internal/logic/utils"
	"genops-master/internal/svc"
	"genops-master/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCaptchaReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := utils.NewGetCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.GetCaptcha(svcCtx, &req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, biz.GenerateCaptchaError)
		} else {
			httpx.OkJsonCtx(r.Context(), w, biz.Success(resp))
		}
	}
}
