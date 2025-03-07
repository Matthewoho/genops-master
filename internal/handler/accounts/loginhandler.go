package accounts

import (
	"fmt"
	"net/http"

	"genops-master/internal/biz"
	"genops-master/internal/logic/accounts"
	"genops-master/internal/svc"
	"genops-master/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			fmt.Println(err)
			return
		}

		l := accounts.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, biz.Success(resp))
		}
	}
}
