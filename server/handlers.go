package server

import (
	"echelon/exec"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type request struct {
	Cmd   string `json:"cmd"`
	Os    string `json:"os"`
	Stdin string `json:"stdin"`
}
type response struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
}

func initMux() *chi.Mux {
	mux := chi.NewMux()
	mux.Route("/api/v1", func(r chi.Router) {
		r.Mount("/remote-execution", remoteExecution())
	})
	return mux
}
func remoteExecution() http.Handler {
	r := chi.NewRouter()
	r.Post("/", func(rw http.ResponseWriter, r *http.Request) {
		req := []request{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Println(err)
			http.Error(rw, "Bad Request", http.StatusBadRequest)
		}

		var stdout, stderr, tempOut, tempErr string
		for _, v := range req {
			tempOut, tempErr = exec.RunCommand(r.Context(), v.Cmd, v.Os, v.Stdin)
			stdout += tempOut
			stderr += tempErr
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		err := json.NewEncoder(rw).Encode(response{Stdout: stdout, Stderr: stderr})
		if err != nil {
			log.Println(err)
		}
	})
	return r
}
