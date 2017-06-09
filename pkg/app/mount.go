package app

import (
	"net/http"

	"encoding/json"

	"github.com/acoshift/go-apiserver-example/pkg/api"
)

func responseHandler(w http.ResponseWriter, r *http.Request, v interface{}) {
	// TODO: choose response type from request
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func bindRequest(r *http.Request, req interface{}) error {
	// TODO: choose decoder from request content type
	return json.NewDecoder(r.Body).Decode(req)
}

// MountUserController mounts user controller to http mux
func MountUserController(mux *http.ServeMux, ctrl api.UserController) {
	// this problem solve in https://github.com/acoshift/mount

	mux.Handle("/user.get", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := api.IDRequest{}
		err := bindRequest(r, &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		resp, err := ctrl.Get(r.Context(), &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		responseHandler(w, r, resp)
	}))

	mux.Handle("/user.list", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := api.ListRequest{}
		err := bindRequest(r, &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		resp, err := ctrl.List(r.Context(), &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		responseHandler(w, r, resp)
	}))

	mux.Handle("/user.delete", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := api.IDRequest{}
		err := bindRequest(r, &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		resp, err := ctrl.Delete(r.Context(), &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		responseHandler(w, r, resp)
	}))
}
