package roles

import (
	"net/http"

	"genops-master/internal/logic/roles"
	"genops-master/internal/svc"
	"genops-master/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetRoleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := roles.NewGetRoleLogic(r.Context(), svcCtx)
		resp, err := l.GetRole(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
