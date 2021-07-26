package web

import (
	"github.com/gorilla/mux"
	"net/http"
	"weather-api-parser/web/controllers"
)

type WeatherServer interface {
	HandleRoutes(router *mux.Router)
	Run()
}

func NewWeatherServer() WeatherServer {
	return &weatherServer{
		server: &http.Server{
			Addr: ":8080",
		},
	}
}

type weatherServer struct {
	server *http.Server
}

func (ws weatherServer) Run() {
	router := mux.NewRouter()
	ws.HandleRoutes(router)

	err := ws.server.ListenAndServe()
	if err != nil {
		panic("error running server!")
	}
}

func (ws *weatherServer) HandleRoutes(router *mux.Router) {
	router.HandleFunc("/weather/{city}", controllers.HandleCurrentWeatherByCityName).Methods(http.MethodGet)

	ws.server.Handler = router
}
