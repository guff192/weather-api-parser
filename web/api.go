package web

import (
	"github.com/gorilla/mux"
	"net/http"
	"weather-api-parser/web/controllers"
)

type WeatherAPIServer interface {
	HandleRoutes(router *mux.Router)
	Run()
}

func NewWeatherAPIServer() WeatherAPIServer {
	return &weatherAPIServer{
		server: &http.Server{
			Addr: ":8080",
		},
	}
}

type weatherAPIServer struct {
	server *http.Server
}

func (ws weatherAPIServer) Run() {
	router := mux.NewRouter()
	ws.HandleRoutes(router)

	err := ws.server.ListenAndServe()
	if err != nil {
		panic("error running server!")
	}
}

func (ws *weatherAPIServer) HandleRoutes(router *mux.Router) {
	router.HandleFunc("/weather/{city}", controllers.HandleCurrentWeatherByCityName).Methods(http.MethodGet)

	ws.server.Handler = router
}
