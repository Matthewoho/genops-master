package permissions

import (
	"net/http"

	"genops-master/internal/logic/permissions"
	"genops-master/internal/svc"
	"genops-master/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListPermissionsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListPermissionsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := permissions.NewListPermissionsLogic(r.Context(), svcCtx)
		resp, err := l.ListPermissions(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
