package apiServer

import (
	"net/http"
	"strings"

	apiModels "github.com/grandminingpool/pool-api/api/generated"
	serverHandelrs "github.com/grandminingpool/pool-api/internal/common/server/handlers"
)

type Server struct {
	srv *apiModels.Server
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", strings.Join([]string{
		http.MethodGet,
		http.MethodOptions,
	}, ", "))

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)

		return
	}

	s.srv.ServeHTTP(w, r)
}

func CreateServer(h apiModels.Handler) (*Server, error) {
	srv, err := apiModels.NewServer(
		h,
		apiModels.WithErrorHandler(serverHandelrs.ErrorHandler()),
		apiModels.WithMethodNotAllowed(serverHandelrs.MethodNotAllowedHandler()),
		apiModels.WithNotFound(serverHandelrs.NotFoundHandler()),
	)
	if err != nil {
		return nil, err
	}

	return &Server{
		srv: srv,
	}, nil
}
