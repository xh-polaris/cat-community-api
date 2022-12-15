package like

import (
	"net/http"

	"github.com/xh-polaris/meowchat-bff/internal/logic/like"
	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserLikedHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserLikedReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := like.NewGetUserLikedLogic(r.Context(), svcCtx)
		resp, err := l.GetUserLiked(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
