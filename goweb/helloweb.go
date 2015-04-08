package goweb

import (
	"net/http"
)

type SongHandler struct {
}

func HttpHandler() {
	// set router
	http.Handle("/", &SongHandler{})
}

func (self *SongHandler) ServeHTTP(resWr http.ResponseWriter, req *http.Request) {
}
