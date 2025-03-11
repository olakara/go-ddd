package chapter2

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type UserHandler interface {
	IsUserSubscriptionActive(ctx context.Context, userID string) bool
}

type UserActiveResponse struct {
	IsActive bool
}

func router(u UserHandler) {
	m := mux.NewRouter()

	m.HandleFunc("/user/{id}/subscription/active", func(w http.ResponseWriter, r *http.Request) {
		// get user id from request
		uID := mux.Vars(r)["id"]
		if uID == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// call u.IsUserSubscriptionActive
		isActive := u.IsUserSubscriptionActive(r.Context(), uID)
		data, err := json.Marshal(UserActiveResponse{isActive})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// write response
		w.Write(data)
	}).Methods(http.MethodGet)
}
