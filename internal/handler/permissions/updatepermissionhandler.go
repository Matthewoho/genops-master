package permissions

import (
	"net/http"

	"genops-master/internal/logic/permissions"
	"genops-master/internal/svc"
	"genops-master/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdatePermissionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdatePermissionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := permissions.NewUpdatePermissionLogic(r.Context(), svcCtx)
		resp, err := l.UpdatePermission(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
