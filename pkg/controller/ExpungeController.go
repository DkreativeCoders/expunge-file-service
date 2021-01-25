package controller

import (
	"encoding/json"
	"github.com/DkreativeCoders/expunge-file-service/pkg/service/Expunge"
	"github.com/gorilla/mux"
	"github.com/kpango/glg"
	"net/http"
)

func NewExpungeController(router *mux.Router, fileExpunge Expunge.IFileExpunge) {
	handler := &expungeController{
		fileExpunge: fileExpunge,
	}

	router.HandleFunc("/api/v1/flush", handler.flush).Methods("GET")

	//return userControllerHandler{userService}
}

type expungeController struct {
	fileExpunge Expunge.IFileExpunge
}


func (e expungeController) flush(w http.ResponseWriter, r *http.Request) {
	glg.Log("flush")
	e.fileExpunge.ExecuteDeleteTask()
	glg.Log("End flush")
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Expunge task completed")
}
