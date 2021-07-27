package web

import (
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
	fs := http.FileServer(http.Dir("static/"))
	ss.server.Handler = cors(fs)
	err := ss.server.ListenAndServe()
	if err != nil {
		panic("error running server!")
	}
}

func cors(fs http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		fs.ServeHTTP(writer, request)
	}
}
