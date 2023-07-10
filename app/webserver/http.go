package webserver

import "net/http"

type Server struct {
	Port string
}

func (s Server) SumHandler(w http.ResponseWriter, r *http.Request) {

}
