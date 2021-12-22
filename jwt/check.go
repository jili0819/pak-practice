package jwt

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jili/pkg-practice/errcode"
	"net/http"
	"strconv"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func responseError(w http.ResponseWriter, codeErr *errcode.Error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	result := Result{Code: codeErr.Code(), Msg: codeErr.Msg()}
	jsonData, _ := json.Marshal(result)
	w.Write(jsonData)
}

func JWTMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenStr := r.Header.Get("Authorization")
			if tokenStr == "" {
				responseError(w, errcode.New400Error("error"))
				return
			}
			claim, err := ParseToken(tokenStr)
			if err != nil || claim == nil {
				responseError(w, errcode.New400Error("error"))
				return
			}
			// 设置参数到header中
			r.Header.Set("user_id", strconv.FormatInt(claim.UserID, 10))
			next.ServeHTTP(w, r)
		})
	}
}
