package web

import (
	"github.com/gorilla/mux"
	"net/http"
)

type StaticServer interface {
	Run()
}

func NewStaticServer() StaticServer {
	return &staticServer{
		server: &http.Server{
			Addr: ":8000",
		},
	}
}

type staticServer struct {
	server *http.Server
}

func (ss *staticServer) Run() {
	fs := http.FileServer(http.Dir("static"))
	router := mux.NewRouter()
	router.Handle("/static/", http.StripPrefix("/static", fs))
	ss.server.Handler = router

	err := ss.server.ListenAndServe()
	if err != nil {
		panic("error running server!")
	}
}
