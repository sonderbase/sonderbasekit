package middlewarex

import (
	"net/http"
	"strings"

	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/samber/lo"
)

type CORSOption struct {
	AllowOrigins []string
	AllowMethods []string
	AllowHeaders []string
}

func CORS(opt *CORSOption) khttp.ServerOption {
	isAllOriginAllowed := lo.Contains(opt.AllowOrigins, "*")
	allowedMethods := strings.Join(opt.AllowMethods, ",")
	allowedHeaders := strings.Join(opt.AllowHeaders, ",")

	return khttp.Filter(
		func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				origin := r.Header.Get("origin")
				if isAllOriginAllowed || lo.Contains(opt.AllowOrigins, origin) {
					w.Header().Set("Access-Control-Allow-Origin", origin)
				}
				w.Header().Set("Access-Control-Allow-Methods", allowedMethods)
				w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
				if r.Method == http.MethodOptions {
					return
				}
				next.ServeHTTP(w, r)
			})
		},
	)
}
