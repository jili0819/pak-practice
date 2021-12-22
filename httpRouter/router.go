package httpRouter

import (
	"github.com/gorilla/mux"
	"github.com/jili/pkg-practice/jwt"
	"net/http"
)

func init() {
	router := mux.NewRouter()
	v1 := router.PathPrefix("/api/v1").Subrouter()
	// auth router
	authRouter := v1.PathPrefix("/auth").Subrouter()
	{
		authRouter.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {

		}).Methods(http.MethodPost, http.MethodOptions)
	}
	// index router
	indexRouter := v1.PathPrefix("/index").Subrouter()
	indexRouter.Use(jwt.JWTMiddleware())
	{
		indexRouter.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		}).Methods(http.MethodPost, http.MethodOptions)
	}

	http.ListenAndServe("localhost", router)
}
