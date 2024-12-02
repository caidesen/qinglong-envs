package middleware

import (
	"context"
	"net/http"
	"qinglong-envs/pkg/httpapi"
)

func BindParamsID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := httpapi.GetIntInPath(r, "id")
		if err != nil {
			httpapi.ErrorHandler(w, err)
			return
		}
		r = r.WithContext(context.WithValue(r.Context(), "id", id))
		next.ServeHTTP(w, r)
	})
}

func GetIDInCtx(ctx context.Context) int {
	id := ctx.Value("id")
	if id == nil {
		return 0
	}
	return id.(int)
}
