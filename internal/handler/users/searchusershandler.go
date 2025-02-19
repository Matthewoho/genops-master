package users

import (
	"net/http"

	"genops-master/internal/logic/users"
	"genops-master/internal/svc"
	"genops-master/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SearchUsersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchUsersReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := users.NewSearchUsersLogic(r.Context(), svcCtx)
		resp, err := l.SearchUsers(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
