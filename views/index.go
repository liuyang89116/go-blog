package views

import (
	"go-blog/common"
	"go-blog/service"
	"log"
	"net/http"
)

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	homeResponse, err := service.GetAllIndexInfo()
	if err != nil {
		log.Println("Get index response error:", err)
		panic(err)
	}

	index.WriteData(w, homeResponse)
}
