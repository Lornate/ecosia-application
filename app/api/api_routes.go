package api

import (
	"ecosia-application/app/api/controller"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	getFavoriteTree = "GetFavoriteTree"
)

func GetRoutes() *mux.Router {

	v1Router := mux.NewRouter()

	//GET
	v1Router.HandleFunc("/", controller.GetFavouriteTree).Methods(http.MethodGet).Name(getFavoriteTree)

	return v1Router
}
