package handler

import (
	"net/http"

	"github.com/xh-polaris/cat-community-svc/api/internal/logic"
	"github.com/xh-polaris/cat-community-svc/api/internal/svc"
	"github.com/xh-polaris/cat-community-svc/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateUserDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateUserDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUpdateUserDetailLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUserDetail(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}