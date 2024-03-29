package routers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func SetAppRoutes(router *mux.Router) *mux.Router {
	appRouter := mux.NewRouter()

	appRouter.HandleFunc("/", controllers.IndexGET)

	appRouter.HandleFunc("/user/create", controllers.UserCreate).Methods("GET", "POST")
	appRouter.HandleFunc("/users", controllers.UserTable).Methods("GET")
	appRouter.HandleFunc("/user/create", controllers.UserCreate).Methods("GET", "POST")
	appRouter.HandleFunc("/users", controllers.UserTable).Methods("GET")
	appRouter.HandleFunc("/user/create", controllers.UserCreate).Methods("GET", "POST")
	appRouter.HandleFunc("/users", controllers.UserTable).Methods("GET")
	router.PathPrefix("/").Handler(appRouter)
	return router
}
