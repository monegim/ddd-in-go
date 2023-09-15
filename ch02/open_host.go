package ch02

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserHandler interface {
	IsUserPaymentActive(ctx context.Context, userID string) bool
}

type UserActiveResponse struct {
	IsActive bool
}

func router(u UserHandler) {
	m := chi.NewRouter()
	m.Get("/user/{userID}/payment/active", func(writer http.ResponseWriter, request *http.Request) {
		// check auth, etc
		uID := chi.URLParam(request, "userID")
		if uID == "" {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		isActive := u.IsUserPaymentActive(request.Context(), uID)
		b, err := json.Marshal(UserActiveResponse{IsActive: isActive})
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, _ = writer.Write(b)
	})
}
