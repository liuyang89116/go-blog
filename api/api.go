package api

import "net/http"

var API = &Api{}

type Api struct {
}

func (*Api) SaveAndUpdatePost(writer http.ResponseWriter, request *http.Request) {

}
