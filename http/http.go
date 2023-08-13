package http

import (
	"encoding/json"
	"fmt"
	"github.com/charmbracelet/log"
	"net/http"
	"time"
)

// EncodeResponse encodes the response to JSON and writes it to the response writer
func EncodeResponse(w http.ResponseWriter, response interface{}) {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.Write([]byte("Error encoding response"))
	}
}

// EncodeError encodes the error to JSON and writes it to the response writer
func EncodeError(w http.ResponseWriter, response error) {
	w.WriteHeader(http.StatusInternalServerError)
	EncodeResponse(w, response.Error())
}

// MoshaHttpService is an interface for any Mosha HTTP service.
type MoshaHttpService interface {
	MakeHandler() http.Handler
	GetPort() string
	GetName() string
}

// StartHttpService starts a Mosha HTTP service.
func StartHttpService(mhs MoshaHttpService) error {
	log.Infof("Starting %s http on %s", mhs.GetName(), mhs.GetPort())

	server := &http.Server{
		Addr:              fmt.Sprintf(":%s", mhs.GetPort()),
		Handler:           mhs.MakeHandler(),
		ReadHeaderTimeout: 3 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		return fmt.Errorf("unable to start service %q: %s", mhs.GetName(), err)
	}
	return nil
}
