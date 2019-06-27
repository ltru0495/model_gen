package routers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func SetAppRoutes(router *mux.Router) *mux.Router {
	appRouter := mux.NewRouter()

	appRouter.HandleFunc("/", controllers.IndexGET)

	router.PathPrefix("/").Handler(appRouter)
	return router
}
