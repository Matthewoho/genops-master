package permissions

import (
	"net/http"

	"genops-master/internal/logic/permissions"
	"genops-master/internal/svc"
	"genops-master/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetPermissionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPermissionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := permissions.NewGetPermissionLogic(r.Context(), svcCtx)
		resp, err := l.GetPermission(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
