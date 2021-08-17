package handler

import (
	"net/http"

	"zero_examples/common/response/errorx"
	"zero_examples/service/user/api/internal/logic"
	"zero_examples/service/user/api/internal/svc"
	"zero_examples/service/user/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func searchHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// validate the request
		var req types.SearchReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorx.NewDefaultError(err.Error()))
			return
		}

		l := logic.NewSearchLogic(r.Context(), ctx)
		resp, err := l.Search(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
