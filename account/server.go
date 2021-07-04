package account

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/kapeel-mopkar/go-kit-demo/account/crypto"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("POST").Path("/users").Handler(httptransport.NewServer(
		endpoints.CreateUser,
		crypto.DecodeUserRequest,
		crypto.EncodeReponse,
	))

	r.Methods("GET").Path("/users/{id}").Handler(httptransport.NewServer(
		endpoints.GetUser,
		crypto.DecodeEmailRequest,
		crypto.EncodeReponse,
	))

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
