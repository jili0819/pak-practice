package httpRouter

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jili/pkg-practice/errcode"
	"github.com/jili/pkg-practice/jwt"
	"net/http"
	"strconv"
	"time"
)

func StartRouter() {
	router := mux.NewRouter()
	v1 := router.PathPrefix("/api/v1").Subrouter()
	v1.NotFoundHandler = http.NotFoundHandler()

	v1.MethodNotAllowedHandler = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	})
	// auth router
	authRouter := v1.PathPrefix("/auth").Subrouter()
	{
		authRouter.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
			token, err := jwt.GenToken(1, time.Now().Unix()+86400)
			if err != nil {
				ResponseError(writer, errcode.New400Error(err.Error()))
				return
			}
			ResponseSuccess(writer, token)
			return
		}).Methods(http.MethodPost, http.MethodOptions)
	}
	// index router
	indexRouter := v1.PathPrefix("/index").Subrouter()
	indexRouter.Use(jwt.JWTMiddleware())
	{
		indexRouter.HandleFunc("", func(writer http.ResponseWriter, request *http.Request) {
			userIdStr := request.Header.Get("user_id")
			userId, _ := strconv.ParseInt(userIdStr, 10, 64)
			ResponseSuccess(writer, UserInfo{UserID: userId})
			return
		}).Methods(http.MethodGet, http.MethodOptions)
	}

	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		return nil
	})

	http.ListenAndServe("localhost:8080", router)
}

type UserInfo struct {
	UserID int64 `json:"user_id"`
}

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func response(w http.ResponseWriter, statusCode int, data Result) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	jsonData, _ := json.Marshal(data)
	w.Write(jsonData)
}

func GetMemberID(r *http.Request) (int64, error) {
	userIDStr := r.Header.Get("user_id")
	if userIDStr == "" {
		return 0, errcode.New400Error("user_id获取失败")
	}
	return strconv.ParseInt(userIDStr, 10, 64)
}

func ResponseError(w http.ResponseWriter, err error) {
	var errnoCode *errcode.Error

	result := Result{
		Code: errnoCode.Code(),
		Msg:  errnoCode.Msg(),
	}
	response(w, http.StatusOK, result)
}

func ResponseSuccess(w http.ResponseWriter, data interface{}) {
	result := Result{
		Code: 100,
		Msg:  "success",
		Data: data,
	}
	response(w, http.StatusOK, result)
}

func Response(w http.ResponseWriter, code int, message string, data interface{}) {
	result := Result{
		Code: code,
		Msg:  message,
		Data: data,
	}
	response(w, http.StatusOK, result)
}
