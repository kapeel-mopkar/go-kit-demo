package crypto

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kapeel-mopkar/go-kit-demo/account/dto"
)

func DecodeUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req dto.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeEmailRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req dto.GetUserRequest
	vars := mux.Vars(r)

	req = dto.GetUserRequest{
		Id: vars["id"],
	}

	return req, nil
}

func EncodeReponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
