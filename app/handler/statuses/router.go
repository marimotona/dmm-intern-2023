package statuses

import (
	"net/http"
	"yatter-backend-go/app/domain/repository"
	"yatter-backend-go/app/handler/auth"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	sr repository.Status
	ar repository.Account
}

func NewRouter(sr repository.Status, ar repository.Account) http.Handler {
	r := chi.NewRouter()

	h := &handler{sr, ar}

	r.With(auth.Middleware(ar)).Post("/", h.PostStatus)

	return r
}
