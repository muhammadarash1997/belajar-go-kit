package transports

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/muhammadarash1997/go-kit-http/endpoints"
	"github.com/muhammadarash1997/go-kit-http/models"
)

func NewHTTPServer(ctx context.Context, endpoints endpoints.Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("POST").Path("/user").Handler(httptransport.NewServer(
		endpoints.CreateUser, // Akan dijalankan setelah decodeUserRequest dijalankan
		decodeUserRequest,    // Akan dijalankan pertama kali dan bertugas mengekstrak payload bertipe JSON.Request menjadi payload bertipe interface{}
		encodeResponse,       // Akan dijalankan terakhir yang mana bertugas mengekstrak data bertipe interface{} menjadi data bertipe JSON.Response dan bertugas untuk mengembalikan response ke client
	))

	r.Methods("GET").Path("/user/{id}").Handler(httptransport.NewServer(
		endpoints.GetUser,
		decodeEmailRequest,
		encodeResponse,
	))

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/")
		next.ServeHTTP(w, r)
	})
}

// decdecodeUserRequest untuk mengubah payload bertipe json (yang diterima dari client) ke struct model yang selanjutnya akan dipakai dan diproses oleh endpoints
func decodeUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req models.CreateUserRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// decdecodeEmailRequest untuk mengubah payload bertipe json (yang diterima dari client) ke struct model yang selanjutnya akan dipakai dan diproses oleh endpoints
func decodeEmailRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req models.GetUserRequest

	vars := mux.Vars(r)
	req = models.GetUserRequest{
		Id: vars["id"],
	}

	return req, nil
}

// encodeResponse untuk mengubah data yang telah diproses oleh endpoints (yang bertipe struct) ke json dan akan langsung dikirim ke client
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
