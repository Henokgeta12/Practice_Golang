package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type handler struct{
	store *types.Userstore
}

func newhandler(store *types.Userstore) *handler{
	return &handlers{store : store}
}
func (h *handler) RegisterRoutes(router *max.router){

	router.HandleFunc("/login", h.handlelogin).Methods("POST")
	router.HandleFunc("/register", h.handleregister).Methods("POST")
}

func (h *handler) handlelogin(w http.ResponseWritter, r *http.Request){

	
}

func (h *handler) handleRegister(w http.ResponseWriter,r *http.Request){

	var payload types.registeruserpayload;

	if err := utils.parseJSON(r,&payload); err != nil{
		utils.WriteError(w,http.StatusBadRequest,err)
		return
	}

}